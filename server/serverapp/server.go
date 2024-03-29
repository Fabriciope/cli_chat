package serverapp

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"
	"sync"

	"github.com/Fabriciope/cli_chat/shared"
)

const (
	ip   = "localhost"
	port = 5000
)

type Server struct {
	mutex               *sync.Mutex
	listener            *net.TCPListener
	handlersForRequests handlersMap
	clients             map[string]*Client
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
		clients:             make(map[string]*Client),
	}, nil
}

func (server *Server) InitServer() {
	handlers := newRequestHandlers(server)
	server.setHandlerForEachRequest(handlers)
	server.run()
}

func (server *Server) setHandlerForEachRequest(handlers *RequestHandlers) {
	server.handlersForRequests = handlersMap{
		shared.LoginActionName:       (*handlers).loginHandler,
		shared.SendMessageActionName: (*handlers).sendMessageInChat,
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
			return
		}

		var request shared.Request

		if err = json.Unmarshal(buf[:bufSize], &request); err != nil {
			sender.sendMessage(conn, shared.Response{
				Name:    "unknown",
				Err:     true,
				Payload: "Invalid request",
			})
			return
		}

		var response *shared.Response = server.handleRequest(ctx, request)
		err = sender.sendMessage(conn, *response)
		if err != nil {
			return
		}
	}
}

func (server *Server) handleRequest(ctx context.Context, request shared.Request) *shared.Response {
	for actionName, handler := range server.handlersForRequests {
		if actionName == request.Name {
			// TODO: serealizar os dados de entrada
			return handler(ctx, request)
		}
	}

	return &shared.Response{Name: "unknown", Err: true, Payload: "Action name unknown"}
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
    client := newClient(conn, username)

    server.lock()
    server.clients[conn.RemoteAddr().String()] = client
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

func (server *Server) userByRemoteAddr(remoteAddr string) (*Client, error) {
    client, ok := server.clients[remoteAddr]
    if !ok {
        return nil, errors.New("client does not exist")
    }

    return client, nil
}
