package clientapp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Fabriciope/cli_chat/client/clientapp/cui"
	"github.com/Fabriciope/cli_chat/shared"
)

// TODO: colocar como variaveis globais no container
const (
	remoteIp   = "localhost"
	remotePort = 5000

	localIp   = "localhost"
	localPort = 3000
)

type Client struct {
	connection          *net.TCPConn
	requestSender       sender
	responsesFromServer chan shared.Response
	loggedIn            bool

	cui *cui.CUI
}

func NewClient() (*Client, error) {
	remoteAddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", remoteIp, remotePort))
	//localAddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", localIp, localPort))
	conn, err := net.DialTCP("tcp", nil, remoteAddr)
	if err != nil {
		return nil, err
	}

	cui, err := cui.NewCUI()
	if err != nil {
		return nil, err
	}

	return &Client{
		connection:          conn,
		requestSender:       newRequestSender(conn),
		responsesFromServer: make(chan shared.Response),
		loggedIn:            false,
		cui:                 cui,
	}, nil
}

// TODO: retornar error para tratar no main
func (client *Client) InitChat() {
	defer client.connection.Close()

	//	err := client.cui.DrawLoading(100, cui.Magenta)
	//	if err != nil {
	//		log.Panicln(err.Error())
	//	}

	controller := newController(client)
	inputScanner := bufio.NewScanner(os.Stdin)

	client.login(inputScanner, controller.loginHandler())

    client.cui.DrawConsoleUserInterface()

	go client.listenToServer()
	go controller.catchResponsesAndHandle()

	client.listenToInput(inputScanner, controller)
}

func (client *Client) login(inputScanner *bufio.Scanner, handler func(string) error) {
	client.cui.DrawLoginInterface()

	for inputScanner.Scan() {
		username := strings.Trim(inputScanner.Text(), " ")
		if username == "" {
			client.cui.DrawLoginError("invalid username!")
			continue
		}

		err := handler(username)
		if err != nil {
			client.cui.DrawLoginError(err.Error() + ", try again.")
			continue
		}

		return
	}
}

func (client *Client) listenToServer() {
	for {
		var buf = make([]byte, 1024)
		n, err := client.connection.Read(buf)
		if err != nil {
			return
		}

		var responseFromServer shared.Response
		if err = json.Unmarshal(buf[:n], &responseFromServer); err != nil {
			return
		}

		client.responsesFromServer <- responseFromServer
	}
}

func (client *Client) listenToInput(inputScanner *bufio.Scanner, controller *controller) {
	for {
		if !inputScanner.Scan() || inputScanner.Err() == nil {
			return
		}

		controller.handleInput(inputScanner.Text())
	}
}

func (client *Client) awaitResponseFromServer() (responseFromServer shared.Response, err error) {
	var buf = make([]byte, 1024)
	n, err := client.connection.Read(buf)
	if err != nil {
		return
	}

	if err = json.Unmarshal(buf[:n], &responseFromServer); err != nil {
		return
	}

	return
}
