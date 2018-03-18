package message

import (
	"encoding/json"
	"github.com/zusi/zusi-go/tcp/message/fahrpult"
)

type Message struct {
	Verbindungsaufbau *VerbindungsaufbauMessage `zusi:"0001" json:"verbindungsaufbau,omitempty"`
	Fahrpult          *fahrpult.FahrpultMessage `zusi:"0002" json:"fahrpult,omitempty"`
}

func (m *Message) String() string {
	res, err := json.Marshal(m)
	if err != nil {
		return "Could not convert Message to JSON"
	}

	return string(res)
}

// 5.3.2 Verbindungsaufbau
type VerbindungsaufbauMessage struct {
	Hello    *HelloMessage    `zusi:"0001" json:"hello,omitempty"`
	AckHello *AckHelloMessage `zusi:"0002" json:"ack_hello,omitempty"`
}

func (VerbindungsaufbauMessage) IsToplevel() bool {
	return true
}

// 5.3.2.1 Befehl 00 01 - HELLO (Client → Zusi)
type HelloMessage struct {
	ProtokollVersion *uint16 `zusi:"0001" json:"protokoll_version,omitempty"`
	ClientTyp        *uint16 `zusi:"0002" json:"client_typ,omitempty"` // 01: Zusi, 02: Fahrpult
	Name             *string `zusi:"0003" json:"name,omitempty"`
	Version          *string `zusi:"0004" json:"version,omitempty"`
}

// 5.3.2.2 Befehl 00 02 - ACK_HELLO (Zusi → Client)
type AckHelloMessage struct {
	ZusiVersion         *string `zusi:"0001" json:"zusi_version,omitempty"`
	ZusiVerbindungsinfo *string `zusi:"0002" json:"zusi_verbindungsinfo,omitempty"`
	ErrorCode           *byte   `zusi:"0003" json:"error_code,omitempty"` // Der Client wurde akzeptiert, wenn das Byte auf 00 steht. Wird der Client nicht akzeptiert, wird stattdes- sen ein anderes Byte gesendet. Der Server bricht dar- aufhin die Verbindung ab.
}

type IsToplevel interface {
	IsToplevel() bool
}
