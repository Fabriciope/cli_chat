package serverapp

import (
	"context"
	"encoding/json"
	"net"
	"sync"

	"github.com/Fabriciope/cli_chat/shared"
)

// TODO: renomear struct para responseSender e criar uma interface Sender

type Sender struct {
	server *Server
}

func newSender(server *Server) *Sender {
	return &Sender{server: server}
}

func (sender *Sender) propagateMessage(ctx context.Context, response shared.Response) {
	conn := ctx.Value("connection").(*net.TCPConn)
	clients := sender.server.clients
	var wg *sync.WaitGroup
	wg.Add(len(clients) - 1)

	var send = func(receiver *net.TCPConn, response shared.Response, wg *sync.WaitGroup) {
		defer wg.Done()

		responseJson, _ := json.Marshal(response)
		receiver.Write([]byte(responseJson))
	}

	sender.server.lock()
	for addr := range clients {
		if addr.String() != conn.RemoteAddr().String() {
			go send(clients[addr].connection, response, wg)
		}
	}

	wg.Wait()
	sender.server.unlock()
}

func (sender *Sender) sendMessage(receiver *net.TCPConn, response shared.Response) (err error) {
	responseJson, _ := json.Marshal(response)

	sender.server.lock()
	_, err = receiver.Write(responseJson)
	sender.server.unlock()

	return
}
