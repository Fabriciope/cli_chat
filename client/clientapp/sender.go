package clientapp

import (
	"encoding/json"
	"net"

	"github.com/Fabriciope/cli_chat/shared"
)

type sender interface {
	sendRequest(shared.Request) error
}

type requestSender struct {
	connection *net.TCPConn
}

func newRequestSender(conn *net.TCPConn) *requestSender {
	return &requestSender{connection: conn}
}

func (sender *requestSender) sendRequest(request shared.Request) error {
	requestJson, err := json.Marshal(request)
	if err != nil {
		return err
	}

	_, err = sender.connection.Write(requestJson)
	if err != nil {
		return err
	}

	return nil
}
