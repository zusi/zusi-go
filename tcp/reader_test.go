package tcp

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zusi/zusi-go/tcp/message"
)

var result *message.Message

func BenchmarkBeispiel1Generated(b *testing.B) {
	var msg *message.Message

	for i := 0; i < b.N; i++ {
		byteReader := bytes.NewReader(beispiel1bytes)
		msg, _ = Read(byteReader)
	}

	result = msg
}

func BenchmarkBeispiel5Generated(b *testing.B) {
	var msg *message.Message

	for i := 0; i < b.N; i++ {
		byteReader := bytes.NewReader(beispiel5bytes)
		msg, _ = Read(byteReader)
	}

	result = msg
}

func TestReadBeispiel(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		byteReader := bytes.NewReader(beispiel1bytes)

		res, err := Read(byteReader)

		assert.NoError(t, err)
		assert.Equal(t, beispiel1msg.Verbindungsaufbau.Hello.Version, res.Verbindungsaufbau.Hello.Version)
		assert.Equal(t, beispiel1msg.Verbindungsaufbau.Hello.Name, res.Verbindungsaufbau.Hello.Name)
		assert.Equal(t, beispiel1msg.Verbindungsaufbau.Hello.ProtokollVersion, res.Verbindungsaufbau.Hello.ProtokollVersion)
		assert.Equal(t, beispiel1msg.Verbindungsaufbau.Hello.ClientTyp, res.Verbindungsaufbau.Hello.ClientTyp)
		assert.Nil(t, res.Fahrpult)
		assert.Nil(t, res.Verbindungsaufbau.AckHello)

		_, err = byteReader.ReadByte()
		assert.EqualError(t, err, "EOF")
	})

	t.Run("2", func(t *testing.T) {
		byteReader := bytes.NewReader(beispiel2bytes)

		res, err := Read(byteReader)

		assert.NoError(t, err)
		assert.Equal(t, beispiel2msg.Verbindungsaufbau.AckHello.ErrorCode, res.Verbindungsaufbau.AckHello.ErrorCode)
		assert.Equal(t, beispiel2msg.Verbindungsaufbau.AckHello.ZusiVerbindungsinfo, res.Verbindungsaufbau.AckHello.ZusiVerbindungsinfo)
		assert.Equal(t, beispiel2msg.Verbindungsaufbau.AckHello.ZusiVersion, res.Verbindungsaufbau.AckHello.ZusiVersion)
		assert.Nil(t, res.Fahrpult)
		assert.Nil(t, res.Verbindungsaufbau.Hello)

		_, err = byteReader.ReadByte()
		assert.EqualError(t, err, "EOF")
	})

	t.Run("3", func(t *testing.T) {
		byteReader := bytes.NewReader(beispiel3bytes)

		res, err := Read(byteReader)

		assert.NoError(t, err)
		assert.NotNil(t, res.Fahrpult)
		assert.NotNil(t, res.Fahrpult.NeededData)
		assert.NotNil(t, res.Fahrpult.NeededData.Fuehrerstandsanzeigen)
		assert.Equal(t, beispiel3msg.Fahrpult.NeededData.Fuehrerstandsanzeigen.Anzeigen, res.Fahrpult.NeededData.Fuehrerstandsanzeigen.Anzeigen)
		assert.Nil(t, res.Verbindungsaufbau)
		assert.Nil(t, res.Fahrpult.Control)
		assert.Nil(t, res.Fahrpult.AckNeededData)
		assert.Nil(t, res.Fahrpult.DataFtd)

		_, err = byteReader.ReadByte()
		assert.EqualError(t, err, "EOF")
	})

	t.Run("4", func(t *testing.T) {
		byteReader := bytes.NewReader(beispiel4bytes)

		res, err := Read(byteReader)

		assert.NoError(t, err)
		assert.NotNil(t, res.Fahrpult)
		assert.NotNil(t, res.Fahrpult.AckNeededData)
		assert.Equal(t, beispiel4msg.Fahrpult.AckNeededData.ErrorCode, res.Fahrpult.AckNeededData.ErrorCode)
		assert.Nil(t, res.Verbindungsaufbau)
		assert.Nil(t, res.Fahrpult.Control)
		assert.Nil(t, res.Fahrpult.NeededData)
		assert.Nil(t, res.Fahrpult.DataFtd)

		_, err = byteReader.ReadByte()
		assert.EqualError(t, err, "EOF")
	})

	t.Run("5", func(t *testing.T) {
		byteReader := bytes.NewReader(beispiel5bytes)

		res, err := Read(byteReader)

		assert.NoError(t, err)
		assert.NotNil(t, res.Fahrpult)
		assert.NotNil(t, res.Fahrpult.DataFtd)
		assert.Equal(t, beispiel5msg.Fahrpult.DataFtd.Geschwindigkeit, res.Fahrpult.DataFtd.Geschwindigkeit)
		assert.Equal(t, beispiel5msg.Fahrpult.DataFtd.LmSchleudern, res.Fahrpult.DataFtd.LmSchleudern)
		assert.Nil(t, res.Verbindungsaufbau)
		assert.Nil(t, res.Fahrpult.Control)
		assert.Nil(t, res.Fahrpult.NeededData)
		assert.Nil(t, res.Fahrpult.AckNeededData)

		_, err = byteReader.ReadByte()
		assert.EqualError(t, err, "EOF")
	})
}
