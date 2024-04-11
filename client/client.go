package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Fabriciope/cli_chat/client/controller"
	"github.com/Fabriciope/cli_chat/client/controller/handler"
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared"
)

// TODO: colocar como variaveis globais no container
const (
	remoteIp   = "localhost"
	remotePort = 5000

	localIp   = "localhost"
	localPort = 3000
)

type User struct {
	connection   *net.TCPConn
	inputScanner *bufio.Scanner
	controller   *controller.Controller
	cui          *cui.CUI
	loggedIn     bool
}

// TODO: recever cui como parametro como interface
func NewUser() (*User, error) {
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

	myUser := &User{
		connection:   conn,
		inputScanner: bufio.NewScanner(os.Stdin),
		cui:          cui,
		loggedIn:     false,
	}
	myUser.controller = controller.NewController(handler.NewHandler(myUser))

	return myUser, nil
}

func (user *User) CloseConnection() error {
	return user.connection.Close()
}

func (user *User) InitChat() {
	defer user.CloseConnection()

	go user.CUI().InitApp()

	user.login()

	go user.listenToServer()
	user.listenToInput()
}

func (user *User) login() {
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

func (user *User) listenToServer() {
	for {
		var buf = make([]byte, 1024)
		n, err := user.connection.Read(buf)
		if err != nil {
			// TODO: fazer algo quando a conexao com o servidor e perdida
			// TODO: tirar o parametro chatLine como um ponteiro no cui
			// TODO: tira o Info field do chatline e deixar somento o timestamp
			user.cui.DrawNewLineInChat(&cui.ChatLine{
				Info:      "[insert time]",
				InfoColor: escapecode.Red,
				Text:      "connection to the server was lost",
			})
			user.CloseConnection()
			os.Exit(1)
			return
		}

		var responseFromServer shared.Response
		if err = json.Unmarshal(buf[:n], &responseFromServer); err != nil {
			return
		}

		user.controller.HandleResponse(responseFromServer)
	}
}

func (user *User) listenToInput() {
	for user.inputScanner.Scan() {
		if user.inputScanner.Err() != nil {
			return
		}

		input := strings.Trim(user.inputScanner.Text(), " ")
		if input == "" {
			user.CUI().RedrawTypingBox()
			continue
		}

		user.controller.HandleInput(input)
		user.CUI().RedrawTypingBox()
	}
}

func (user *User) AwaitResponseFromServer() (responseFromServer shared.Response, err error) {
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

func (user *User) Conn() *net.TCPConn {
	return user.connection
}

func (user *User) CUI() *cui.CUI {
	return user.cui
}

func (user *User) LoggedIn() bool {
	return user.loggedIn
}

func (user *User) SetLoggedInAs(logged bool) {
	user.loggedIn = logged
	user.CUI().SetLoggedAs(true)
}
