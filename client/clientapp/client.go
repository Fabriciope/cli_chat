package clientapp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Fabriciope/cli_chat/client/clientapp/controller"
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

type MyUser struct {
	connection   *net.TCPConn
	inputScanner *bufio.Scanner

	cui *cui.CUI

	loggedIn bool
}

func NewUser() (*MyUser, error) {
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

	return &MyUser{
		connection:   conn,
		inputScanner: bufio.NewScanner(os.Stdin),
		cui:          cui,
		loggedIn:     false,
	}, nil
}

// TODO: retornar error para tratar no main
func (user *MyUser) InitChat() {
	defer user.connection.Close()

	go user.CUI().InitApp()

	controller := controller.NewController(user)

	user.login(controller.LoginHandler())

	go user.listenToServer(controller)
	user.listenToInput(controller)
}

func (client *MyUser) login(handler func(string) error) {
	for client.inputScanner.Scan() {
		username := strings.Trim(client.inputScanner.Text(), " ")
		if username == "" {
			client.CUI().DrawLoginError("invalid username!")
			client.CUI().DrawLoginError("invalid username!")
		}

		err := handler(username)
		if err != nil {
			client.CUI().DrawLoginError(err.Error() + ", try again.")
			continue
		}

		return
	}
}

func (client *MyUser) listenToServer(controller *controller.Controller) {
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

		controller.HandleResponse(responseFromServer)
	}
}

func (client *MyUser) listenToInput(controller *controller.Controller) {
	for client.inputScanner.Scan() {
		if client.inputScanner.Err() != nil {
			return
		}

		input := strings.Trim(client.inputScanner.Text(), " ")
		if input == "" {
			client.CUI().RedrawTypingBox()
			continue
		}

		controller.HandleInput(client.inputScanner.Text())
		client.CUI().RedrawTypingBox()
	}
}

func (client *MyUser) AwaitResponseFromServer() (responseFromServer shared.Response, err error) {
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

func (client *MyUser) Conn() *net.TCPConn {
	return client.connection
}

func (user *MyUser) CUI() *cui.CUI {
	return user.cui
}

func (user *MyUser) LoggedIn() bool {
	return user.loggedIn
}

func (user *MyUser) SetLoggedInAs(logged bool) {
	user.loggedIn = logged
}
