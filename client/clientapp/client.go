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
	connection   *net.TCPConn
	inputScanner *bufio.Scanner

	cui *cui.CUI

	loggedIn bool
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
		connection:   conn,
		inputScanner: bufio.NewScanner(os.Stdin),
		cui:          cui,
		loggedIn:     false,
	}, nil
}

// TODO: retornar error para tratar no main
func (client *Client) InitChat() {
	defer client.connection.Close()

	go client.cui.InitApp()

	controller := newController(client)

	client.login(controller.loginHandler())

	go client.listenToServer(controller)
	client.listenToInput(controller)
}

func (client *Client) login(handler func(string) error) {
	for client.inputScanner.Scan() {
		username := strings.Trim(client.inputScanner.Text(), " ")
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

func (client *Client) listenToServer(controller *controller) {
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

		controller.handleResponse(responseFromServer)
	}
}

func (client *Client) listenToInput(controller *controller) {
	for {
		if !client.inputScanner.Scan() || client.inputScanner.Err() == nil {
			return
		}

		controller.handleInput(client.inputScanner.Text())
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
