package serverapp

import (
	"context"
	"encoding/json"
	"net"
	"sync"

	"github.com/Fabriciope/cli_chat/shared"
)

type responseSender struct {
	server *Server
}

func newResponseSender(server *Server) *responseSender {
	return &responseSender{server: server}
}

// TODO:   verificar se a conexao esta aberta antes de escrever
func (sender *responseSender) propagateMessage(ctx context.Context, response shared.Response) error {
	conn := ctx.Value("connection").(*net.TCPConn)
	responseJson, _ := json.Marshal(response)

	clients := sender.server.clients

	wg := new(sync.WaitGroup)
	wg.Add(len(clients) - 1)

	sendMsgErr := make(chan error, len(clients)-1)

	sender.server.lock()
	for addr := range clients {
		if addr != conn.RemoteAddr().String() {
			go func() {
				defer wg.Done()

				_, err := clients[addr].connection.Write([]byte(responseJson))
				sendMsgErr <- err
			}()
		}
	}

	go func() {
		wg.Wait()
		close(sendMsgErr)
		sender.server.unlock()
	}()

	for err := range sendMsgErr {
		if err != nil {
			return err
		}
	}

	return nil
}

func (sender *responseSender) sendMessage(receiver *net.TCPConn, response shared.Response) (err error) {
	responseJson, _ := json.Marshal(response)

	sender.server.lock()
	_, err = receiver.Write(responseJson)
	sender.server.unlock()

	return
}
