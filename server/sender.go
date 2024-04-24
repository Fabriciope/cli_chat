package server

import (
	"encoding/json"
	"net"
	"sync"

	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

type responseSender struct {
	server *Server
}

func newResponseSender(server *Server) *responseSender {
	return &responseSender{server: server}
}

func (sender *responseSender) propagateMessage(exceptConn *net.TCPConn, response dto.Response) error {
	clients := sender.server.clients

	responseJson, _ := json.Marshal(response)
	wg := new(sync.WaitGroup)
	sendMsgErr := make(chan error, 1)
	send := func(addr string) {
		defer wg.Done()

		_, err := clients[addr].connection.Write([]byte(responseJson))
		sendMsgErr <- err
	}

	sender.server.mutex.Lock()
	for addr := range clients {
		if exceptConn != nil {
			if addr != exceptConn.RemoteAddr().String() {
				wg.Add(1)
				go send(addr)
			}
			continue
		}

		wg.Add(1)
		go send(addr)
	}

	go func() {
		wg.Wait()
		close(sendMsgErr)
		sender.server.mutex.Unlock()
	}()

	for err := range sendMsgErr {
		if err != nil {
			return err
		}
	}

	return nil
}

func (sender *responseSender) sendMessage(receiver *net.TCPConn, response dto.Response) (err error) {
	responseJson, _ := json.Marshal(response)

	sender.server.mutex.Lock()
	_, err = receiver.Write(responseJson)
	sender.server.mutex.Unlock()

	return
}
