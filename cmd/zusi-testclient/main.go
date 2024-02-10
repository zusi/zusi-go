package main

import (
	"flag"
	"io"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
	"github.com/zusi/zusi-go/tcp"
	"github.com/zusi/zusi-go/tcp/fahrpult"
	msg "github.com/zusi/zusi-go/tcp/message/fahrpult"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

var zusiUri = flag.String("zusi", "localhost:1436", "uri of zusi server")

func main() {
	flag.Parse()

	fp := fahrpult.New()
	err := fp.Connect(*zusiUri)
	if err != nil {
		log.Fatal(err)
		return
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
		log.Warn(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func receiveLogger(client *tcp.Client) {
	for {
		ms, err := client.Receive()
		if err != nil {
			if err == io.EOF {
				log.WithError(err).Fatal()
			}
			log.WithError(err).Warn()
		} else {
			log.WithField("message", ms).Info()
		}
	}
}
