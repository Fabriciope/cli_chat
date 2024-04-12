package sender

import (
	"encoding/json"
	"net"

	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

type requestSender struct {
	connection *net.TCPConn
}

func NewRequestSender(conn *net.TCPConn) *requestSender {
	return &requestSender{connection: conn}
}

func (sender *requestSender) SendRequest(request dto.Request) error {
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
