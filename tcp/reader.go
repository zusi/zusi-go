package tcp

import (
	"errors"
	"fmt"
	"io"

	"github.com/zusi/zusi-go/tcp/message"
)

func Read(reader io.Reader) (*message.Message, error) {
	length, attr, err := ReadHeader(reader)
	if length != 0x0 {
		return nil, errors.New("No new struct received")
	}

	msg := message.Message{}

	if err != nil && errors.Is(err, io.EOF) {
		return &msg, err
	}

	switch attr {
	case 0x0001:
		msg.Verbindungsaufbau, err = readMessageVerbindungsaufbauMessage(reader)
	case 0x0002:
		msg.Fahrpult, err = readFahrpultFahrpultMessage(reader)
	default:
		fmt.Printf("unknown field %v", attr)
		_, err = ReadString(reader, length)
	}

	if err != nil && errors.Is(err, io.EOF) {
		return &msg, err
	}

	return &msg, nil
}
