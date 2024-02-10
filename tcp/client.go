package tcp

import (
	"bufio"
	"io"
	"log/slog"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/zusi/zusi-go/tcp/message"
)

type Client struct {
	conn           net.Conn
	zusiVersion    string
	messageChan    chan message.Message
	errorChan      chan error
	abort          chan struct{}
	messageHandler func(*message.Message) (keepMessage bool, err error)
}

func New() *Client {
	client := &Client{}
	client.abort = make(chan struct{})

	return client
}

func NewWithHandler(messageHandler func(*message.Message) (keepMessage bool, err error)) *Client {
	client := New()
	client.messageHandler = messageHandler

	return client
}

func (c *Client) Receive() (*message.Message, error) {
	select {
	case msg, ok := <-c.messageChan:
		if !ok {
			return nil, io.EOF
		}
		return &msg, nil
	case err := <-c.errorChan:
		return nil, err
	}
}

func (c *Client) ConnectWithType(address string, typ message.ClientTyp) error {
	if c.conn != nil {
		return errors.New("already connected to a server. first close old connection")
	}

	conn, err := net.DialTimeout("tcp", address, 10*time.Second)
	if err != nil {
		return errors.Wrap(err, "error connecting to zusi host")
	}

	c.conn = conn

	bufReader := bufio.NewReader(c.conn)
	err = c.handshake(bufReader, typ)
	if err != nil {
		return errors.Wrap(err, "error during handshake")
	}

	c.messageChan = make(chan message.Message)
	c.errorChan = make(chan error)
	go c.read(bufReader)

	return err
}

func (c *Client) handshake(reader io.Reader, typ message.ClientTyp) error {
	handshakeMsg := message.Message{
		Verbindungsaufbau: &message.VerbindungsaufbauMessage{
			Hello: &message.HelloMessage{
				ProtokollVersion: Uint16(2),
				ClientTyp:        Uint16(uint16(typ)),
				Name:             String("Client TCP-Test"),
				Version:          String("1.0-alpha"),
			},
		},
	}

	err := c.Send(handshakeMsg)
	if err != nil {
		return err
	}

	msg, err := Read(reader)
	if err != nil {
		slog.With("err", err).Error("Error unmarshalling initial message")
		return errors.Wrap(err, "unmarshaling initial message failed")
	}
	if msg.Verbindungsaufbau == nil || msg.Verbindungsaufbau.AckHello == nil {
		slog.With("message", msg).Error("Initial message is not of type Verbindungsaufbau.AckHello")
		return errors.New("Initial message is not of type Verbindungsaufbau.AckHello")
	}
	if *msg.Verbindungsaufbau.AckHello.ErrorCode != 0 {
		slog.With("message", msg, "errorCode", msg.Verbindungsaufbau.AckHello.ErrorCode).Error("ErrorCode of initial message != 0")
		return errors.New("ErrorCode of initial message != 0")
	}

	c.zusiVersion = *msg.Verbindungsaufbau.AckHello.ZusiVersion

	return nil
}

func (c *Client) read(reader io.Reader) {
	for {
		msg, err := Read(reader)
		if err != nil {
			if errors.Cause(err) == io.EOF {
				c.errorChan <- errors.Wrap(err, "remote connection died")
				close(c.messageChan)
				return
			}

			c.errorChan <- errors.Wrap(err, "error unmarshalling message. one message is lost but keep receiving")
			continue
		}

		if c.messageHandler != nil {
			keep, err := c.messageHandler(msg)
			if err != nil {
				c.errorChan <- errors.Wrap(err, "error handling message internally but keep receiving")
			}
			if !keep {
				continue
			}
		}

		c.messageChan <- *msg
	}
}

func (c *Client) Send(message message.Message) error {
	bts, err := MarshalMessage(message)
	if err != nil {
		slog.With("err", err, "message", message).Warn("Error marshalling message")
		return err
	}

	_, err = c.conn.Write(bts)
	if err != nil {
		slog.With("err", err, "message", message).Warn("Error writing message to stream")
		return err
	}

	return nil
}

func (c *Client) Close() error {
	err := c.conn.Close()
	c.conn = nil

	if c.abort != nil {
		c.abort <- struct{}{}
	}

	return err
}
