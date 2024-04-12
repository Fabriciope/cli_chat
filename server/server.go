package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

const (
	ip   = "localhost"
	port = 5000
)

type Server struct {
	mutex               *sync.Mutex
	listener            *net.TCPListener
	handlersForRequests handlersMap
	clients             map[string]*client
}

func NewServer() (*Server, error) {
	addr := &net.TCPAddr{IP: net.ParseIP(ip), Port: port}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		mutex:               &sync.Mutex{},
		listener:            listener,
		handlersForRequests: make(handlersMap),
		clients:             make(map[string]*client),
	}, nil
}

func (server *Server) InitServer() {
	handlers := newRequestHandlers(server)
	server.setHandlerForEachRequest(handlers)
	server.run()
}

func (server *Server) setHandlerForEachRequest(handlers *RequestHandlers) {
	server.handlersForRequests = handlersMap{
		dto.LoginActionName:       (*handlers).loginHandler,
		dto.SendMessageActionName: (*handlers).sendMessageInChat,
	}
}

func (server *Server) run() {
	log.Printf("started server at :%d\n", port)
	for {
		conn, err := server.listener.AcceptTCP()
		if err != nil {
			return
		}
		context := context.WithValue(context.Background(), "connection", conn)
		go server.clientHandler(context)
	}
}

func (server *Server) clientHandler(ctx context.Context) {
	conn := ctx.Value("connection").(*net.TCPConn)
	log.Printf("new client from %s\n", conn.RemoteAddr().String())
	defer conn.Close()

	sender := newResponseSender(server)
	for {
		var buf [1024]byte
		bufSize, err := conn.Read(buf[0:])
		if err != nil {
			client, _ := server.userByRemoteAddr(conn.RemoteAddr().String())
			username := client.username
			server.removeClient(conn.RemoteAddr().String())
			sender.sendMessage(conn, dto.Response{
				Name:    dto.LogoutActionName,
				Err:     false,
				Payload: fmt.Sprintf("%s disconnected from the chat", username),
			})
			return
		}

		var request dto.Request

		if err = json.Unmarshal(buf[:bufSize], &request); err != nil {
			sender.sendMessage(conn, dto.Response{
				Name:    "unknown",
				Err:     true,
				Payload: "Invalid request",
			})
			return
		}

		var response *dto.Response = server.handleRequest(ctx, request)
		err = sender.sendMessage(conn, *response)
		if err != nil {
			return
		}
	}
}

func (server *Server) handleRequest(ctx context.Context, request dto.Request) *dto.Response {
	for actionName, handler := range server.handlersForRequests {
		if actionName == request.Name {
			return handler(ctx, request)
		}
	}

	return &dto.Response{Name: "unknown", Err: true, Payload: "Action name unknown"}
}

func (server *Server) lock() {
	server.mutex.Lock()
}

func (server *Server) unlock() {
	server.mutex.Unlock()
}

func (server *Server) addClient(ctx context.Context, username string) error {
	if server.hasClient(username) {
		return errors.New("user already exists")
	}

	conn := ctx.Value("connection").(*net.TCPConn)

	var chosenColor escapecode.ColorCode
loop:
	for _, color := range availableColors {
		if !server.colorIsAlreadyInUse(color) {
			chosenColor = color
			break loop
		}
	}
	client := newClient(conn, username, chosenColor)

	server.lock()
	server.clients[conn.RemoteAddr().String()] = client
	server.unlock()

	return nil
}

func (server *Server) removeClient(remoteAddr string) error {
	client, err := server.userByRemoteAddr(remoteAddr)
	if err != nil {
		return err
	}

	server.lock()
	delete(server.clients, client.RemoteAddr())
	server.unlock()

	return nil
}

func (server *Server) hasClient(username string) bool {
	for i := range server.clients {
		if server.clients[i].username == username {
			return true
		}
	}

	return false
}

func (server *Server) colorIsAlreadyInUse(color escapecode.ColorCode) bool {
	for i := range server.clients {
		if server.clients[i].color == color {
			return true
		}
	}

	return false
}

func (server *Server) userByRemoteAddr(remoteAddr string) (*client, error) {
	client, ok := server.clients[remoteAddr]
	if !ok {
		return nil, errors.New("client does not exist")
	}

	return client, nil
}
