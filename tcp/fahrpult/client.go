package fahrpult

import (
	"github.com/zusi/zusi-go/tcp"
	"github.com/zusi/zusi-go/tcp/message"
	"github.com/zusi/zusi-go/tcp/message/fahrpult"
)

type Client struct {
	*tcp.Client
}

func New() *Client {
	tcpClient := tcp.New()
	c := &Client{tcpClient}

	return c
}

func (c *Client) Connect(address string) error {
	err := c.ConnectWithType(address, message.Fahrpult)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendFp(msg fahrpult.FahrpultMessage) error {
	return c.Send(message.Message{Fahrpult: &msg})
}
