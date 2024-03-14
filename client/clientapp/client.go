package clientapp

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/Fabriciope/cli_chat/shared"
)

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
}

func NewClient() *Client {
	var remoteAddr, _ = net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", remoteIp, remotePort))
	var localAddr, _ = net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", localIp, localPort))
	var conn, err = net.DialTCP("tcp", localAddr, remoteAddr)
	if err != nil {
		log.Panicln(err)
		return nil
	}

	return &Client{
		connection:          conn,
		requestSender:       newRequestSender(conn),
		responsesFromServer: make(chan shared.Response),
		loggedIn:            false,
	}
}

func (client *Client) InitChat() {
	defer client.connection.Close()

	var controller = newController(client)

	go client.listenToServer()
	go controller.catchResponsesAndHandle()

	var inputScanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	var loginHandler inputHandler = controller.inputHandler(shared.LoginActionName)
	if err := client.login(inputScanner, loginHandler); err != nil {
		log.Fatal(err.Error())
	}

	client.listenToInput(inputScanner, controller)
}

func (client *Client) listenToServer() {
	for {
		var buf = make([]byte, 1024)
		var n, err = client.connection.Read(buf)
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

func (client *Client) login(inputScanner *bufio.Scanner, handler inputHandler) error {
	fmt.Print("Enter with your username: ")
	for inputScanner.Scan() {
		var username string = strings.Trim(inputScanner.Text(), " ")
		if username == "" {
			fmt.Println("Err: invalid username ↑")
			fmt.Print("Enter with your username again: ")
			continue
		}

		handler(username)
		return nil
	}

	return errors.New("cannot get the username")
}

func (client *Client) listenToInput(inputScanner *bufio.Scanner, controller *controller) {
	for {
		fmt.Print("--> ")
		if !inputScanner.Scan() || inputScanner.Err() == nil {
			return
		}

		var err error = controller.handleInput(inputScanner.Text())
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
