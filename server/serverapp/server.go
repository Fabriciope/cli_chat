package serverapp

import (
	"context"
	"encoding/json"
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
	clients             map[net.Addr]*Client
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
		clients:             make(map[net.Addr]*Client),
	}, nil
}

func (server *Server) InitServer() {
	handlers := newRequestHandlers(server)
	server.setHandlerForEachRequest(handlers)
	server.run()
}

func (server *Server) setHandlerForEachRequest(handlers *RequestHandlers) {
	server.handlersForRequests = handlersMap{
		shared.LoginActionName: (*handlers).loginHandler,
	}
}

func (server *Server) run() {
	log.Printf("started server at :%d\n", port)
	for {
		conn, err := server.listener.AcceptTCP()
		if err != nil {
			return
		}

		type myKey string
		var key myKey = "connection"
		context := context.WithValue(context.Background(), key, conn)
		go server.clientHandler(context)
	}
}

func (server *Server) clientHandler(ctx context.Context) {
	conn := ctx.Value("connection").(*net.TCPConn)
	log.Printf("new client from %s\n", conn.RemoteAddr().String())
	defer conn.Close()

	sender := newSender(server)
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
