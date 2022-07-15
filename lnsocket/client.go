package lnsocket

import (
	"encoding/json"

	lnsocket "github.com/jb55/lnsocket/go"
)

type Client struct {
	socket  lnsocket.LNSocket
	nodeID  string
	address string
}

func New(nodeId string, address string) *Client {
	client := Client{
		socket:  lnsocket.LNSocket{},
		nodeID:  nodeId,
		address: address,
	}
	client.socket.GenKey()
	return &client
}

func (instance *Client) Connect() error {
	return instance.socket.ConnectAndInit(instance.address, instance.nodeID)
}

func (instance *Client) Disconnect() {
	instance.socket.Disconnect()
}

func (instance *Client) Call(method string, params map[string]any, runes string) (map[string]any, error) {
	jsonParm, _ := json.Marshal(params)
	bodyStr, err := instance.socket.Rpc(runes, method, string(jsonParm))
	var res map[string]any
	if err := json.Unmarshal([]byte(bodyStr), &res); err != nil {
		return nil, err
	}
	return res, err
}
