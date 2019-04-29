package fpn

import "encoding/xml"

type Fpn struct {
	XMLName xml.Name `xml:"Zusi"`
	Text    string   `xml:",chardata"`
	Info    struct {
		Text         string `xml:",chardata"`
		DateiTyp     string `xml:"DateiTyp,attr"`
		Version      string `xml:"Version,attr"`
		MinVersion   string `xml:"MinVersion,attr"`
		AutorEintrag struct {
			Text      string `xml:",chardata"`
			AutorID   string `xml:"AutorID,attr"`
			AutorName string `xml:"AutorName,attr"`
		} `xml:"AutorEintrag"`
		Datei []struct {
			Text      string `xml:",chardata"`
			Dateiname string `xml:"Dateiname,attr"`
		} `xml:"Datei"`
	} `xml:"Info"`
	Fahrplan struct {
		Text                 string `xml:",chardata"`
		AnfangsZeit          string `xml:"AnfangsZeit,attr"`
		BefehlsKonfiguration struct {
			Text      string `xml:",chardata"`
			Dateiname string `xml:"Dateiname,attr"`
		} `xml:"BefehlsKonfiguration"`
		Begruessungsdatei struct {
			Text      string `xml:",chardata"`
			Dateiname string `xml:"Dateiname,attr"`
		} `xml:"Begruessungsdatei"`
		Zug []struct {
			Text  string `xml:",chardata"`
			Datei struct {
				Text      string `xml:",chardata"`
				Dateiname string `xml:"Dateiname,attr"`
			} `xml:"Datei"`
		} `xml:"Zug"`
		StrModul []struct {
			Text  string `xml:",chardata"`
			Datei struct {
				Text      string `xml:",chardata"`
				Dateiname string `xml:"Dateiname,attr"`
			} `xml:"Datei"`
			P   string `xml:"p"`
			Phi string `xml:"phi"`
		} `xml:"StrModul"`
		UTM struct {
			Text     string `xml:",chardata"`
			UTMWE    string `xml:"UTM_WE,attr"`
			UTMNS    string `xml:"UTM_NS,attr"`
			UTMZone  string `xml:"UTM_Zone,attr"`
			UTMZone2 string `xml:"UTM_Zone2,attr"`
		} `xml:"UTM"`
	} `xml:"Fahrplan"`
}
