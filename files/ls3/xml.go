package ls3

import "encoding/xml"

type Ls3 struct {
	XMLName xml.Name `xml:"Zusi"`
	Text    string   `xml:",chardata"`
	Info    struct {
		Text         string `xml:",chardata"`
		DateiTyp     string `xml:"DateiTyp,attr"`
		Version      string `xml:"Version,attr"`
		MinVersion   string `xml:"MinVersion,attr"`
		AutorEintrag struct {
			Text       string `xml:",chardata"`
			AutorID    string `xml:"AutorID,attr"`
			AutorName  string `xml:"AutorName,attr"`
			AutorEmail string `xml:"AutorEmail,attr"`
		} `xml:"AutorEintrag"`
	} `xml:"Info"`
	Landschaft struct {
		Text        string `xml:",chardata"`
		Verknuepfte []struct {
			Text        string `xml:",chardata"`
			Flags       string `xml:"Flags,attr"`
			BoundingR   string `xml:"BoundingR,attr"`
			SichtbarBis string `xml:"SichtbarBis,attr"`
			Vorlade     string `xml:"Vorlade,attr"`
			LODbit      string `xml:"LODbit,attr"`
			Datei       struct {
				Text      string `xml:",chardata"`
				Dateiname string `xml:"Dateiname,attr"`
			} `xml:"Datei"`
			P struct {
				Text string `xml:",chardata"`
				X    string `xml:"X,attr"`
				Y    string `xml:"Y,attr"`
				Z    string `xml:"Z,attr"`
			} `xml:"p"`
			Phi string `xml:"phi"`
			Sk  string `xml:"sk"`
		} `xml:"Verknuepfte"`
		Ankerpunkt []struct {
			Text         string `xml:",chardata"`
			AnkerKat     string `xml:"AnkerKat,attr"`
			Beschreibung string `xml:"Beschreibung,attr"`
			P            struct {
				Text string `xml:",chardata"`
				X    string `xml:"X,attr"`
				Y    string `xml:"Y,attr"`
				Z    string `xml:"Z,attr"`
			} `xml:"p"`
			Phi struct {
				Text string `xml:",chardata"`
				Z    string `xml:"Z,attr"`
			} `xml:"phi"`
		} `xml:"Ankerpunkt"`
		Animation struct {
			Text            string `xml:",chardata"`
			AniBeschreibung string `xml:"AniBeschreibung,attr"`
			AniLoopen       string `xml:"AniLoopen,attr"`
			AniNrs          struct {
				Text  string `xml:",chardata"`
				AniNr string `xml:"AniNr,attr"`
			} `xml:"AniNrs"`
		} `xml:"Animation"`
		VerknAnimation struct {
			Text      string `xml:",chardata"`
			AniIndex  string `xml:"AniIndex,attr"`
			AniNr     string `xml:"AniNr,attr"`
			AniGeschw string `xml:"AniGeschw,attr"`
			AniPunkt  []struct {
				Text       string `xml:",chardata"`
				AniDimmung string `xml:"AniDimmung,attr"`
				AniZeit    string `xml:"AniZeit,attr"`
				P          string `xml:"p"`
				Q          struct {
					Text string `xml:",chardata"`
					W    string `xml:"W,attr"`
				} `xml:"q"`
			} `xml:"AniPunkt"`
		} `xml:"VerknAnimation"`
	} `xml:"Landschaft"`
}
