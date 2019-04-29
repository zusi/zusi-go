package trn

import "encoding/xml"

// .trn
type Trn struct {
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
	Zug struct {
		Text             string `xml:",chardata"`
		Gattung          string `xml:"Gattung,attr"`
		Nummer           string `xml:"Nummer,attr"`
		Zuglauf          string `xml:"Zuglauf,attr"`
		Prio             string `xml:"Prio,attr"`
		BremsstellungZug string `xml:"BremsstellungZug,attr"`
		SpAnfang         string `xml:"spAnfang,attr"`
		MBrh             string `xml:"MBrh,attr"`
		ReisendenDichte  string `xml:"ReisendenDichte,attr"`
		FahrplanGruppe   string `xml:"FahrplanGruppe,attr"`
		Rekursionstiefe  string `xml:"Rekursionstiefe,attr"`
		FahrstrName      string `xml:"FahrstrName,attr"`
		Zugtyp           string `xml:"Zugtyp,attr"`
		Buchfahrplandll  string `xml:"Buchfahrplandll,attr"`
		Datei            struct {
			Text      string `xml:",chardata"`
			Dateiname string `xml:"Dateiname,attr"`
			NurInfo   string `xml:"NurInfo,attr"`
		} `xml:"Datei"`
		BuchfahrplanRohDatei struct {
			Text      string `xml:",chardata"`
			Dateiname string `xml:"Dateiname,attr"`
		} `xml:"BuchfahrplanRohDatei"`
		FahrplanEintrag []struct {
			Text                         string `xml:",chardata"`
			Ank                          string `xml:"Ank,attr"`
			Abf                          string `xml:"Abf,attr"`
			Betrst                       string `xml:"Betrst,attr"`
			FplEintrag                   string `xml:"FplEintrag,attr"`
			FzgVerbandAktion             string `xml:"FzgVerbandAktion,attr"`
			FzgVerbandWendeSignalabstand string `xml:"FzgVerbandWendeSignalabstand,attr"`
			FahrplanSignalEintrag        []struct {
				Text           string `xml:",chardata"`
				FahrplanSignal string `xml:"FahrplanSignal,attr"`
			} `xml:"FahrplanSignalEintrag"`
		} `xml:"FahrplanEintrag"`
		FahrzeugVarianten struct {
			Text         string `xml:",chardata"`
			Bezeichnung  string `xml:"Bezeichnung,attr"`
			ZufallsWert  string `xml:"ZufallsWert,attr"`
			FahrzeugInfo []struct {
				Text        string `xml:",chardata"`
				IDHaupt     string `xml:"IDHaupt,attr"`
				IDNeben     string `xml:"IDNeben,attr"`
				DotraModus  string `xml:"DotraModus,attr"`
				SASchaltung string `xml:"SASchaltung,attr"`
				Gedreht     string `xml:"Gedreht,attr"`
				Datei       struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
				} `xml:"Datei"`
			} `xml:"FahrzeugInfo"`
		} `xml:"FahrzeugVarianten"`
	} `xml:"Zug"`
}

// .timetable.xml
type Timetable struct {
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
	Buchfahrplan struct {
		Text          string `xml:",chardata"`
		Gattung       string `xml:"Gattung,attr"`
		Nummer        string `xml:"Nummer,attr"`
		Zuglauf       string `xml:"Zuglauf,attr"`
		BR            string `xml:"BR,attr"`
		Masse         string `xml:"Masse,attr"`
		SpMax         string `xml:"spMax,attr"`
		Bremsh        string `xml:"Bremsh,attr"`
		MBrh          string `xml:"MBrh,attr"`
		Laenge        string `xml:"Laenge,attr"`
		KmStart       string `xml:"kmStart,attr"`
		Bremsstellung string `xml:"Bremsstellung,attr"`
		DateiFpn      struct {
			Text      string `xml:",chardata"`
			Dateiname string `xml:"Dateiname,attr"`
			NurInfo   string `xml:"NurInfo,attr"`
		} `xml:"Datei_fpn"`
		DateiTrn struct {
			Text      string `xml:",chardata"`
			Dateiname string `xml:"Dateiname,attr"`
			NurInfo   string `xml:"NurInfo,attr"`
		} `xml:"Datei_trn"`
		FplZeile []struct {
			Text           string `xml:",chardata"`
			FplLaufweg     string `xml:"FplLaufweg,attr"`
			FplRglGgl      string `xml:"FplRglGgl,attr"`
			FahrstrStrecke string `xml:"FahrstrStrecke,attr"`
			FplIcon        struct {
				Text      string `xml:",chardata"`
				FplIconNr string `xml:"FplIconNr,attr"`
			} `xml:"FplIcon"`
			FplName struct {
				Text        string `xml:",chardata"`
				FplNameText string `xml:"FplNameText,attr"`
			} `xml:"FplName"`
			FplvMax struct {
				Text string `xml:",chardata"`
				VMax string `xml:"vMax,attr"`
			} `xml:"FplvMax"`
			Fplkm []struct {
				Text      string `xml:",chardata"`
				Km        string `xml:"km,attr"`
				FplSprung string `xml:"FplSprung,attr"`
				FplkmNeu  string `xml:"FplkmNeu,attr"`
			} `xml:"Fplkm"`
			FplSignaltyp struct {
				Text           string `xml:",chardata"`
				FplSignaltypNr string `xml:"FplSignaltypNr,attr"`
			} `xml:"FplSignaltyp"`
			FplAnk []struct {
				Text       string `xml:",chardata"`
				Ank        string `xml:"Ank,attr"`
				FplEintrag string `xml:"FplEintrag,attr"`
			} `xml:"FplAnk"`
			FplAbf []struct {
				Text string `xml:",chardata"`
				Abf  string `xml:"Abf,attr"`
			} `xml:"FplAbf"`
			FplSaegelinien struct {
				Text      string `xml:",chardata"`
				FplAnzahl string `xml:"FplAnzahl,attr"`
			} `xml:"FplSaegelinien"`
			FplNameRechts struct {
				Text        string `xml:",chardata"`
				FplNameText string `xml:"FplNameText,attr"`
			} `xml:"FplNameRechts"`
			FplRichtungswechsel string `xml:"FplRichtungswechsel"`
		} `xml:"FplZeile"`
	} `xml:"Buchfahrplan"`
}
