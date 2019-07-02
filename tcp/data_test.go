package tcp

import (
	"github.com/zusi/zusi-go/tcp/message"
	"github.com/zusi/zusi-go/tcp/message/fahrpult"
)

var helloTestMessageBytes = []byte{
	0x04, 0x00, 0x00, 0x00, // Länge 4 Bytes → es folgt ein Attribut, Länge 4 bytes
	0x01, 0x00, // ID x0001: Protokoll-Version
	0x02, 0x00, // Protokoll-Version „2“ (Word)
	0x04, 0x00, 0x00, 0x00, // Länge 4 Bytes → es folgt ein Attribut, Länge 4 bytes
	0x02, 0x00, // ID x0002: Client-Typ
	0x02, 0x00, // Client-Typ „Fahrpult“ (Word)
	0x0A, 0x00, 0x00, 0x00, // Länge 10 Bytes → es folgt ein Attribut
	0x03, 0x00, // ID x0003: Klartextstring
	0x46, 0x61, 0x68, 0x72, 0x70, 0x75, 0x6C, 0x74, // String „Fahrpult“ (8 Zeichen, da 2 bytes für die ID)
	0x05, 0x00, 0x00, 0x00, // Länge 5 Bytes → es folgt ein Attribut
	0x04, 0x00, // ID x0004: Version
	0x32, 0x2E, 0x30, // String „2.0“
}

type helloTestMessage struct {
	ProtokollVersion uint16 `zusi:"0001"`
	ClientTyp        uint16 `zusi:"0002"` // 01: Zusi, 02: Fahrpult
	Name             string `zusi:"0003"`
	Version          string `zusi:"0004"`
}

func (helloTestMessage) IsTopMost() bool {
	return true
}

var beispiel1bytes = []byte{
	0x00, 0x00, 0x00, 0x00, // Länge 0 Bytes → es beginnt ein Knoten
	0x01, 0x00, // ID 1: Verbindungsaufbau
	0x00, 0x00, 0x00, 0x00, // Länge 0 Bytes → es beginnt ein Knoten
	0x01, 0x00, // ID 1: HELLO-Befehl
	0x04, 0x00, 0x00, 0x00, // Länge 4 Bytes → es folgt ein Attribut, Länge 4 bytes
	0x01, 0x00, // ID x0001: Protokoll-Version
	0x02, 0x00, // Protokoll-Version „2“ (Word)
	0x04, 0x00, 0x00, 0x00, // Länge 4 Bytes → es folgt ein Attribut, Länge 4 bytes
	0x02, 0x00, // ID x0002: Client-Typ
	0x02, 0x00, // Client-Typ „Fahrpult“ (Word)
	0x0A, 0x00, 0x00, 0x00, // Länge 10 Bytes → es folgt ein Attribut
	0x03, 0x00, // ID x0003: Klartextstring
	0x46, 0x61, 0x68, 0x72, 0x70, 0x75, 0x6C, 0x74, // String „Fahrpult“ (8 Zeichen, da 2 bytes für die ID)
	0x05, 0x00, 0x00, 0x00, // Länge 5 Bytes → es folgt ein Attribut
	0x04, 0x00, // ID x0004: Version
	0x32, 0x2E, 0x30, // String „2.0“
	0xFF, 0xFF, 0xFF, 0xFF, // Ende Knoten
	0xFF, 0xFF, 0xFF, 0xFF, // Ende Knoten
}

var beispiel1msg = message.Message{Verbindungsaufbau: &message.VerbindungsaufbauMessage{
	Hello: &message.HelloMessage{
		ProtokollVersion: Uint16(2),
		ClientTyp:        Uint16(2),
		Name:             String("Fahrpult"),
		Version:          String("2.0"),
	},
}}

var beispiel2bytes = []byte{
	0x00, 0x00, 0x00, 0x00, // Länge 0 Bytes → es beginnt ein Knoten
	0x01, 0x00, // ID 1: Verbindungsaufbau
	0x00, 0x00, 0x00, 0x00, // Länge 0 Bytes → es beginnt ein Knoten
	0x02, 0x00, // ID 2: ACK_HELLO-Befehl
	0x09, 0x00, 0x00, 0x00, // Länge 9 Bytes → es folgt ein Attribut
	0x01, 0x00, // ID x0001: Zusi-Version
	0x33, 0x2E, 0x30, 0x2E, 0x31, 0x2E, 0x30, // String „3.0.1.0“
	0x03, 0x00, 0x00, 0x00, // Länge 3 Bytes → es folgt ein Attribut
	0x02, 0x00, // ID x0002: Zusi-Verbindungsinfo
	0x30,                   // String „0“
	0x03, 0x00, 0x00, 0x00, // Länge 3 Bytes → es folgt ein Attribut
	0x03, 0x00, // ID x0003: Ergebnis
	0x00,                   // 0 (Byte) -> Verbindung akzeptiert
	0x0A, 0x00, 0x00, 0x00, // Länge 10 Bytes → es folgt ein Attribut
	0x04, 0x00, // ID x0004: Fahrplananfangszeit
	0x00, 0x00, 0x00, 0x00, 0xD0, 0x35, 0xE4, 0x40, // 41390,5 (double) -> Anfangszeit in Tagen
	0xFF, 0xFF, 0xFF, 0xFF, // Ende Knoten
	0xFF, 0xFF, 0xFF, 0xFF, // Ende Knoten
}

var beispiel2msg = message.Message{Verbindungsaufbau: &message.VerbindungsaufbauMessage{
	AckHello: &message.AckHelloMessage{
		ZusiVersion:         String("3.0.1.0"),
		ZusiVerbindungsinfo: String("0"),
		ErrorCode:           Byte(0),
		FahrplanStartZeit:   Float64(41390.5),
	},
}}

var beispiel3bytes = []byte{
	0x00, 0x00, 0x00, 0x00, // Länge 0 Bytes → es beginnt ein Knoten
	0x02, 0x00, // ID 0x0002: Client-Anwendung Typ 2 (Fahrpult)
	0x00, 0x00, 0x00, 0x00, // Länge 0 Bytes → es beginnt ein Knoten
	0x03, 0x00, // ID 0x0003: NEEDED_DATA-Befehl
	0x00, 0x00, 0x00, 0x00, // Länge 0 Bytes → es beginnt ein Knoten
	0x0A, 0x00, // ID x000A: Führerstandsanzeigen
	0x04, 0x00, 0x00, 0x00, // Länge 4 Bytes → es folgt ein Attribut
	0x01, 0x00, // ID 0x0001: Führerstands ID
	0x01, 0x00, // Nr. 1: Geschwindigkeit (Word)
	0x04, 0x00, 0x00, 0x00, // Länge 4 Bytes → es folgt ein Attribut
	0x01, 0x00, // ID 0x0001: Führerstands ID
	0x1B, 0x00, // Nr. 1B: Schleudern (Word)
	0xFF, 0xFF, 0xFF, 0xFF, // Ende Knoten
	0xFF, 0xFF, 0xFF, 0xFF, // Ende Knoten
	0xFF, 0xFF, 0xFF, 0xFF, // Ende Knoten
}

var beispiel3msg = message.Message{Fahrpult: &fahrpult.FahrpultMessage{
	NeededData: &fahrpult.NeededDataMessage{
		Fuehrerstandsanzeigen: &fahrpult.Fuehrerstandsanzeigen{
			Anzeigen: []uint16{0x0001, 0x001B},
		},
	},
}}

var beispiel4bytes = []byte{
	0x00, 0x00, 0x00, 0x00,
	0x02, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x04, 0x00,
	0x03, 0x00, 0x00, 0x00,
	0x01, 0x00,
	0x00,
	0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF,
}

var beispiel4msg = message.Message{Fahrpult: &fahrpult.FahrpultMessage{
	AckNeededData: &fahrpult.AckNeededData{
		ErrorCode: Byte(0),
	},
}}

var beispiel5bytes = []byte{
	0x00, 0x00, 0x00, 0x00,
	0x02, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x0A, 0x00,
	0x06, 0x00, 0x00, 0x00,
	0x01, 0x00,
	0xAE, 0x47, 0x3D, 0x41,
	0x06, 0x00, 0x00, 0x00,
	0x1B, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF,
}

var beispiel5msg = message.Message{Fahrpult: &fahrpult.FahrpultMessage{
	DataFtd: &fahrpult.DataFtd{
		Geschwindigkeit: Float32(11.83),
		LmSchleudern:    Float32(0),
	},
}}

type StructWithSlice struct {
	Sub []Substruct `zusi:"0001"`
}

type Substruct struct {
	Name string `zusi:"0001"`
}

var structWithSliceExample = StructWithSlice{Sub: []Substruct{
	{Name: "A"},
	{Name: "B"},
}}

var structWithSliceBytes = []byte{
	0x00, 0x00, 0x00, 0x00,
	0x01, 0x00,
	0x03, 0x00, 0x00, 0x00,
	0x01, 0x00,
	0x41,
	0xff, 0xff, 0xff, 0xff,
	0x00, 0x00, 0x00, 0x00,
	0x01, 0x00,
	0x03, 0x00, 0x00, 0x00,
	0x01, 0x00,
	0x42,
	0xff, 0xff, 0xff, 0xff,
}
