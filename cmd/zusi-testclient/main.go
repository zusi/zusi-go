package main

import (
	"flag"
	"io"
	"log/slog"
	"os"
	"os/signal"

	"github.com/zusi/zusi-go/tcp"
	"github.com/zusi/zusi-go/tcp/fahrpult"
	msg "github.com/zusi/zusi-go/tcp/message/fahrpult"
)

var zusiUri = flag.String("zusi", "localhost:1436", "uri of zusi server")

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	flag.Parse()

	if err := mainErr(); err != nil {
		slog.With("err", err).Error("running client failed")
		os.Exit(1)
	}
}

func mainErr() error {
	fp := fahrpult.New()
	err := fp.Connect(*zusiUri)
	if err != nil {
		return err
	}
	defer fp.Close()

	go receiveLogger(fp.Client)

	err = fp.SendFp(msg.FahrpultMessage{
		NeededData: &msg.NeededDataMessage{
			Fuehrerstandsanzeigen: &msg.Fuehrerstandsanzeigen{
				Anzeigen: []uint16{0x0001},
			},
		},
	})
	if err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	return nil
}

func receiveLogger(client *tcp.Client) {
	for {
		ms, err := client.Receive()
		if err != nil {
			if err == io.EOF {
				slog.With("err", err).Error("connection closed")
				return
			}
			slog.With("err", err).Warn("error while receiving message")
		} else {
			slog.With("message", ms).Info("received message")
		}
	}
}
