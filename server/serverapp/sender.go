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

func (sender *Sender) propagateMessage(ctx context.Context, response shared.Response) error {
	conn := ctx.Value("connection").(*net.TCPConn)
	clients := sender.server.clients
	wg := new(sync.WaitGroup)
	wg.Add(len(clients) - 1)

	sendMsgErr := make(chan error)

	sender.server.lock()

	for addr := range clients {
		if addr.String() != conn.RemoteAddr().String() {
			go sender.sendMessageWithGoroutine(clients[addr].connection, response, sendMsgErr, wg)
		}
	}

    waitAndCloseChannel := func(wg *sync.WaitGroup, ch chan error) {
        wg.Wait()
        close(ch)
    }
    go waitAndCloseChannel(wg, sendMsgErr)

    sender.server.unlock()

	for err := range sendMsgErr {
		if err != nil {
			close(sendMsgErr)
			return err
		}
    }

	return nil
}

func (sender *Sender) sendMessageWithGoroutine(receiver *net.TCPConn, response shared.Response, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	responseJson, _ := json.Marshal(response)

	sender.server.lock()
	_, err := receiver.Write([]byte(responseJson))
	sender.server.unlock()

	errCh <- err
}

func (sender *Sender) sendMessage(receiver *net.TCPConn, response shared.Response) (err error) {
	responseJson, _ := json.Marshal(response)

	sender.server.lock()
	_, err = receiver.Write(responseJson)
	sender.server.unlock()

	return
}
