package bus

import (
	"encoding/json"
	"fmt"
	"net"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Push(message Message) error {
	//send message on tcp connection
	conn, err := net.Dial("tcp", PUSH_PORT)
	if err != nil {
		return fmt.Errorf("failed to connect to push port: %v", err)
	}

	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshall message: %v", err)
	}
	payload = append(payload, '\n')

	_, err = conn.Write(payload)

	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	defer conn.Close()
	return nil
}
func (c *Client) StopTask(taskId string, disable bool, delete bool) error {
	var msg Message
	if disable {
		msg = newMessage("disable", taskId)
	} else if delete {
		msg = newMessage("delete", taskId)
	} else {
		msg = newMessage("stop", taskId)
	}
	return c.Push(msg)
}
