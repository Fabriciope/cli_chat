package clientapp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Fabriciope/cli_chat/client/clientapp/controller"
	"github.com/Fabriciope/cli_chat/client/clientapp/controller/handlers"
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
    controller *controller.Controller
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


    myUser := &MyUser{
		connection:   conn,
		inputScanner: bufio.NewScanner(os.Stdin),
		cui:          cui,
		loggedIn:     false,
	}
    myUser.controller = controller.NewController(handlers.NewHandler(myUser))

    return myUser, nil
}

// TODO: fazer algo quando a conexao com o servidor e perdida
func (user *MyUser) InitChat() {
	defer user.connection.Close()

	go user.CUI().InitApp()

	user.login()

	go user.listenToServer()
	user.listenToInput()
}

func (user *MyUser) login() {
	for user.inputScanner.Scan() {
		username := strings.Trim(user.inputScanner.Text(), " ")
		if username == "" {
			user.CUI().DrawLoginError("invalid username!")
            continue
		}

        err := user.controller.LoginHandler()(username)
		if err != nil {
			user.CUI().DrawLoginError(err.Error() + ", try again.")
			continue
		}

		return
	}
}

func (user *MyUser) listenToServer() {
	for {
		var buf = make([]byte, 1024)
		n, err := user.connection.Read(buf)
		if err != nil {
			return
		}

		var responseFromServer shared.Response
		if err = json.Unmarshal(buf[:n], &responseFromServer); err != nil {
			return
		}

		user.controller.HandleResponse(responseFromServer)
	}
}

func (user *MyUser) listenToInput() {
	for user.inputScanner.Scan() {
		if user.inputScanner.Err() != nil {
			return
		}

		input := strings.Trim(user.inputScanner.Text(), " ")
		if input == "" {
			user.CUI().RedrawTypingBox()
			continue
		}

		user.controller.HandleInput(user.inputScanner.Text())
		user.CUI().RedrawTypingBox()
	}
}

func (user *MyUser) AwaitResponseFromServer() (responseFromServer shared.Response, err error) {
	var buf = make([]byte, 1024)
	n, err := user.connection.Read(buf)
	if err != nil {
		return
	}

	if err = json.Unmarshal(buf[:n], &responseFromServer); err != nil {
		return
	}

	return
}

func (user *MyUser) Conn() *net.TCPConn {
	return user.connection
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
