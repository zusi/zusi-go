package st3

import "encoding/xml"

type File struct {
	XMLName xml.Name `xml:"Zusi"`
	Text    string   `xml:",chardata"`
	Info    struct {
		Text         string `xml:",chardata"`
		DateiTyp     string `xml:"DateiTyp,attr"`
		Version      string `xml:"Version,attr"`
		MinVersion   string `xml:"MinVersion,attr"`
		Beschreibung string `xml:"Beschreibung,attr"`
		AutorEintrag struct {
			Text       string `xml:",chardata"`
			AutorID    string `xml:"AutorID,attr"`
			AutorName  string `xml:"AutorName,attr"`
			AutorEmail string `xml:"AutorEmail,attr"`
		} `xml:"AutorEintrag"`
		Datei []struct {
			Text      string `xml:",chardata"`
			Dateiname string `xml:"Dateiname,attr"`
		} `xml:"Datei"`
	} `xml:"Info"`
	Strecke struct {
		Text          string `xml:",chardata"`
		Himmelsmodell string `xml:"Himmelsmodell,attr"`
		Datei         struct {
			Text      string `xml:",chardata"`
			Dateiname string `xml:"Dateiname,attr"`
		} `xml:"Datei"`
		HintergrundDatei     string `xml:"HintergrundDatei"`
		BefehlsKonfiguration struct {
			Text      string `xml:",chardata"`
			Dateiname string `xml:"Dateiname,attr"`
		} `xml:"BefehlsKonfiguration"`
		Kachelpfad struct {
			Text      string `xml:",chardata"`
			Dateiname string `xml:"Dateiname,attr"`
			NurInfo   string `xml:"NurInfo,attr"`
		} `xml:"Kachelpfad"`
		UTM struct {
			Text     string `xml:",chardata"`
			UTMWE    string `xml:"UTM_WE,attr"`
			UTMNS    string `xml:"UTM_NS,attr"`
			UTMZone  string `xml:"UTM_Zone,attr"`
			UTMZone2 string `xml:"UTM_Zone2,attr"`
		} `xml:"UTM"`
		Huellkurve struct {
			Text     string `xml:",chardata"`
			PunktXYZ []struct {
				Text string `xml:",chardata"`
				X    string `xml:"X,attr"`
				Y    string `xml:"Y,attr"`
			} `xml:"PunktXYZ"`
		} `xml:"Huellkurve"`
		SkyDome struct {
			Text      string `xml:",chardata"`
			HimmelTex struct {
				Text      string `xml:",chardata"`
				Dateiname string `xml:"Dateiname,attr"`
			} `xml:"HimmelTex"`
			SonneTex struct {
				Text      string `xml:",chardata"`
				Dateiname string `xml:"Dateiname,attr"`
			} `xml:"SonneTex"`
			SonneHorizontTex struct {
				Text      string `xml:",chardata"`
				Dateiname string `xml:"Dateiname,attr"`
			} `xml:"SonneHorizontTex"`
			MondTex struct {
				Text      string `xml:",chardata"`
				Dateiname string `xml:"Dateiname,attr"`
			} `xml:"MondTex"`
			SternTex struct {
				Text      string `xml:",chardata"`
				Dateiname string `xml:"Dateiname,attr"`
			} `xml:"SternTex"`
		} `xml:"SkyDome"`
		StreckenStandort []struct {
			Text    string `xml:",chardata"`
			StrInfo string `xml:"StrInfo,attr"`
			P       struct {
				Text string `xml:",chardata"`
				X    string `xml:"X,attr"`
				Y    string `xml:"Y,attr"`
				Z    string `xml:"Z,attr"`
			} `xml:"p"`
			Lookat struct {
				Text string `xml:",chardata"`
				X    string `xml:"X,attr"`
				Y    string `xml:"Y,attr"`
				Z    string `xml:"Z,attr"`
			} `xml:"lookat"`
			Up struct {
				Text string `xml:",chardata"`
				X    string `xml:"X,attr"`
				Y    string `xml:"Y,attr"`
				Z    string `xml:"Z,attr"`
			} `xml:"up"`
		} `xml:"StreckenStandort"`
		ModulDateien []struct {
			Text  string `xml:",chardata"`
			Datei struct {
				Text      string `xml:",chardata"`
				Dateiname string `xml:"Dateiname,attr"`
				NurInfo   string `xml:"NurInfo,attr"`
			} `xml:"Datei"`
		} `xml:"ModulDateien"`
		ReferenzElemente []struct {
			Text       string `xml:",chardata"`
			ReferenzNr string `xml:"ReferenzNr,attr"`
			StrElement string `xml:"StrElement,attr"`
			RefTyp     string `xml:"RefTyp,attr"`
			Info       string `xml:"Info,attr"`
			StrNorm    string `xml:"StrNorm,attr"`
		} `xml:"ReferenzElemente"`
		StrElement []struct {
			Text             string `xml:",chardata"`
			Nr               string `xml:"Nr,attr"`
			SpTrass          string `xml:"spTrass,attr"`
			Oberbau          string `xml:"Oberbau,attr"`
			Volt             string `xml:"Volt,attr"`
			Drahthoehe       string `xml:"Drahthoehe,attr"`
			Kr               string `xml:"kr,attr"`
			Anschluss        string `xml:"Anschluss,attr"`
			Ueberh           string `xml:"Ueberh,attr"`
			Zwangshelligkeit string `xml:"Zwangshelligkeit,attr"`
			Fkt              string `xml:"Fkt,attr"`
			G                struct {
				Text string `xml:",chardata"`
				X    string `xml:"X,attr"`
				Y    string `xml:"Y,attr"`
				Z    string `xml:"Z,attr"`
			} `xml:"g"`
			B struct {
				Text string `xml:",chardata"`
				X    string `xml:"X,attr"`
				Y    string `xml:"Y,attr"`
				Z    string `xml:"Z,attr"`
			} `xml:"b"`
			InfoNormRichtung struct {
				Text             string `xml:",chardata"`
				VMax             string `xml:"vMax,attr"`
				Km               string `xml:"km,attr"`
				Pos              string `xml:"pos,attr"`
				Reg              string `xml:"Reg,attr"`
				KoppelWeicheNr   string `xml:"KoppelWeicheNr,attr"`
				KoppelWeicheNorm string `xml:"KoppelWeicheNorm,attr"`
				Ereignis         []struct {
					Text   string `xml:",chardata"`
					Er     string `xml:"Er,attr"`
					Beschr string `xml:"Beschr,attr"`
					Wert   string `xml:"Wert,attr"`
				} `xml:"Ereignis"`
				Signal *struct {
					Text               string `xml:",chardata"`
					SignalTyp          string `xml:"SignalTyp,attr"`
					BoundingR          string `xml:"BoundingR,attr"`
					NameBetriebsstelle string `xml:"NameBetriebsstelle,attr"`
					Stellwerk          string `xml:"Stellwerk,attr"`
					Signalname         string `xml:"Signalname,attr"`
					SignalFlags        string `xml:"SignalFlags,attr"`
					P                  struct {
						Text string `xml:",chardata"`
						X    string `xml:"X,attr"`
						Y    string `xml:"Y,attr"`
						Z    string `xml:"Z,attr"`
					} `xml:"p"`
					Phi struct {
						Text string `xml:",chardata"`
						Z    string `xml:"Z,attr"`
						X    string `xml:"X,attr"`
						Y    string `xml:"Y,attr"`
					} `xml:"phi"`
					SignalFrame []struct {
						Text                         string `xml:",chardata"`
						WeichenbaugruppeIndex        string `xml:"WeichenbaugruppeIndex,attr"`
						WeichenbaugruppeNr           string `xml:"WeichenbaugruppeNr,attr"`
						WeichenbaugruppePos0         string `xml:"WeichenbaugruppePos0,attr"`
						WeichenbaugruppePos1         string `xml:"WeichenbaugruppePos1,attr"`
						WeichenbaugruppeBeschreibung string `xml:"WeichenbaugruppeBeschreibung,attr"`
						P                            struct {
							Text string `xml:",chardata"`
							X    string `xml:"X,attr"`
							Y    string `xml:"Y,attr"`
							Z    string `xml:"Z,attr"`
						} `xml:"p"`
						Phi struct {
							Text string `xml:",chardata"`
							X    string `xml:"X,attr"`
							Y    string `xml:"Y,attr"`
							Z    string `xml:"Z,attr"`
						} `xml:"phi"`
						Datei struct {
							Text      string `xml:",chardata"`
							Dateiname string `xml:"Dateiname,attr"`
						} `xml:"Datei"`
					} `xml:"SignalFrame"`
					HsigBegriff []struct {
						Text       string `xml:",chardata"`
						HsigGeschw string `xml:"HsigGeschw,attr"`
						FahrstrTyp string `xml:"FahrstrTyp,attr"`
					} `xml:"HsigBegriff"`
					VsigBegriff []struct {
						Text       string `xml:",chardata"`
						VsigGeschw string `xml:"VsigGeschw,attr"`
					} `xml:"VsigBegriff"`
					MatrixEintrag []struct {
						Text         string `xml:",chardata"`
						Signalbild   string `xml:"Signalbild,attr"`
						MatrixGeschw string `xml:"MatrixGeschw,attr"`
						Ereignis     []struct {
							Text   string `xml:",chardata"`
							Er     string `xml:"Er,attr"`
							Wert   string `xml:"Wert,attr"`
							Beschr string `xml:"Beschr,attr"`
						} `xml:"Ereignis"`
					} `xml:"MatrixEintrag"`
					Ersatzsignal []struct {
						Text                 string `xml:",chardata"`
						ErsatzsigBezeichnung string `xml:"ErsatzsigBezeichnung,attr"`
						ErsatzsigID          string `xml:"ErsatzsigID,attr"`
						MatrixEintrag        struct {
							Text         string `xml:",chardata"`
							Signalbild   string `xml:"Signalbild,attr"`
							MatrixGeschw string `xml:"MatrixGeschw,attr"`
							Ereignis     []struct {
								Text   string `xml:",chardata"`
								Er     string `xml:"Er,attr"`
								Wert   string `xml:"Wert,attr"`
								Beschr string `xml:"Beschr,attr"`
							} `xml:"Ereignis"`
						} `xml:"MatrixEintrag"`
					} `xml:"Ersatzsignal"`
					KoppelSignal struct {
						Text       string `xml:",chardata"`
						ReferenzNr string `xml:"ReferenzNr,attr"`
						Datei      struct {
							Text      string `xml:",chardata"`
							Dateiname string `xml:"Dateiname,attr"`
							NurInfo   string `xml:"NurInfo,attr"`
						} `xml:"Datei"`
					} `xml:"KoppelSignal"`
				} `xml:"Signal"`
			} `xml:"InfoNormRichtung"`
			InfoGegenRichtung struct {
				Text     string `xml:",chardata"`
				VMax     string `xml:"vMax,attr"`
				Km       string `xml:"km,attr"`
				Pos      string `xml:"pos,attr"`
				Reg      string `xml:"Reg,attr"`
				Ereignis []struct {
					Text   string `xml:",chardata"`
					Er     string `xml:"Er,attr"`
					Beschr string `xml:"Beschr,attr"`
					Wert   string `xml:"Wert,attr"`
				} `xml:"Ereignis"`
				Signal *struct {
					Text               string `xml:",chardata"`
					NameBetriebsstelle string `xml:"NameBetriebsstelle,attr"`
					Stellwerk          string `xml:"Stellwerk,attr"`
					Signalname         string `xml:"Signalname,attr"`
					SignalTyp          string `xml:"SignalTyp,attr"`
					BoundingR          string `xml:"BoundingR,attr"`
					SignalFlags        string `xml:"SignalFlags,attr"`
					ZufallsWert        string `xml:"ZufallsWert,attr"`
					P                  struct {
						Text string `xml:",chardata"`
						X    string `xml:"X,attr"`
						Y    string `xml:"Y,attr"`
						Z    string `xml:"Z,attr"`
					} `xml:"p"`
					Phi struct {
						Text string `xml:",chardata"`
						X    string `xml:"X,attr"`
						Y    string `xml:"Y,attr"`
						Z    string `xml:"Z,attr"`
					} `xml:"phi"`
					SignalFrame []struct {
						Text                  string `xml:",chardata"`
						WeichenbaugruppeIndex string `xml:"WeichenbaugruppeIndex,attr"`
						WeichenbaugruppeNr    string `xml:"WeichenbaugruppeNr,attr"`
						WeichenbaugruppePos0  string `xml:"WeichenbaugruppePos0,attr"`
						WeichenbaugruppePos1  string `xml:"WeichenbaugruppePos1,attr"`
						P                     struct {
							Text string `xml:",chardata"`
							X    string `xml:"X,attr"`
							Y    string `xml:"Y,attr"`
							Z    string `xml:"Z,attr"`
						} `xml:"p"`
						Phi struct {
							Text string `xml:",chardata"`
							Z    string `xml:"Z,attr"`
							X    string `xml:"X,attr"`
						} `xml:"phi"`
						Datei struct {
							Text      string `xml:",chardata"`
							Dateiname string `xml:"Dateiname,attr"`
						} `xml:"Datei"`
					} `xml:"SignalFrame"`
					HsigBegriff []struct {
						Text       string `xml:",chardata"`
						HsigGeschw string `xml:"HsigGeschw,attr"`
						FahrstrTyp string `xml:"FahrstrTyp,attr"`
					} `xml:"HsigBegriff"`
					VsigBegriff []struct {
						Text       string `xml:",chardata"`
						VsigGeschw string `xml:"VsigGeschw,attr"`
					} `xml:"VsigBegriff"`
					MatrixEintrag []struct {
						Text         string `xml:",chardata"`
						Signalbild   string `xml:"Signalbild,attr"`
						MatrixGeschw string `xml:"MatrixGeschw,attr"`
						Ereignis     []struct {
							Text   string `xml:",chardata"`
							Er     string `xml:"Er,attr"`
							Wert   string `xml:"Wert,attr"`
							Beschr string `xml:"Beschr,attr"`
						} `xml:"Ereignis"`
					} `xml:"MatrixEintrag"`
					Ersatzsignal []struct {
						Text                 string `xml:",chardata"`
						ErsatzsigBezeichnung string `xml:"ErsatzsigBezeichnung,attr"`
						ErsatzsigID          string `xml:"ErsatzsigID,attr"`
						MatrixEintrag        struct {
							Text         string `xml:",chardata"`
							Signalbild   string `xml:"Signalbild,attr"`
							MatrixGeschw string `xml:"MatrixGeschw,attr"`
							Ereignis     []struct {
								Text string `xml:",chardata"`
								Er   string `xml:"Er,attr"`
								Wert string `xml:"Wert,attr"`
							} `xml:"Ereignis"`
						} `xml:"MatrixEintrag"`
					} `xml:"Ersatzsignal"`
					KoppelSignal struct {
						Text       string `xml:",chardata"`
						ReferenzNr string `xml:"ReferenzNr,attr"`
						Datei      struct {
							Text      string `xml:",chardata"`
							Dateiname string `xml:"Dateiname,attr"`
							NurInfo   string `xml:"NurInfo,attr"`
						} `xml:"Datei"`
					} `xml:"KoppelSignal"`
				} `xml:"Signal"`
			} `xml:"InfoGegenRichtung"`
			NachNorm []struct {
				Text string `xml:",chardata"`
				Nr   string `xml:"Nr,attr"`
			} `xml:"NachNorm"`
			NachGegenModul struct {
				Text  string `xml:",chardata"`
				Nr    string `xml:"Nr,attr"`
				Datei struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
					NurInfo   string `xml:"NurInfo,attr"`
				} `xml:"Datei"`
			} `xml:"NachGegenModul"`
			NachGegen struct {
				Text string `xml:",chardata"`
				Nr   string `xml:"Nr,attr"`
			} `xml:"NachGegen"`
			NachNormModul struct {
				Text  string `xml:",chardata"`
				Nr    string `xml:"Nr,attr"`
				Datei struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
					NurInfo   string `xml:"NurInfo,attr"`
				} `xml:"Datei"`
			} `xml:"NachNormModul"`
		} `xml:"StrElement"`
		Fahrstrasse []struct {
			Text           string `xml:",chardata"`
			FahrstrName    string `xml:"FahrstrName,attr"`
			FahrstrTyp     string `xml:"FahrstrTyp,attr"`
			Laenge         string `xml:"Laenge,attr"`
			FahrstrStrecke string `xml:"FahrstrStrecke,attr"`
			RglGgl         string `xml:"RglGgl,attr"`
			ZufallsWert    string `xml:"ZufallsWert,attr"`
			FahrstrStart   struct {
				Text  string `xml:",chardata"`
				Ref   string `xml:"Ref,attr"`
				Datei struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
					NurInfo   string `xml:"NurInfo,attr"`
				} `xml:"Datei"`
			} `xml:"FahrstrStart"`
			FahrstrZiel struct {
				Text  string `xml:",chardata"`
				Ref   string `xml:"Ref,attr"`
				Datei struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
					NurInfo   string `xml:"NurInfo,attr"`
				} `xml:"Datei"`
			} `xml:"FahrstrZiel"`
			FahrstrRegister []struct {
				Text  string `xml:",chardata"`
				Ref   string `xml:"Ref,attr"`
				Datei struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
					NurInfo   string `xml:"NurInfo,attr"`
				} `xml:"Datei"`
			} `xml:"FahrstrRegister"`
			FahrstrAufloesung []struct {
				Text  string `xml:",chardata"`
				Ref   string `xml:"Ref,attr"`
				Datei struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
					NurInfo   string `xml:"NurInfo,attr"`
				} `xml:"Datei"`
			} `xml:"FahrstrAufloesung"`
			FahrstrSigHaltfall []struct {
				Text  string `xml:",chardata"`
				Ref   string `xml:"Ref,attr"`
				Datei struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
					NurInfo   string `xml:"NurInfo,attr"`
				} `xml:"Datei"`
			} `xml:"FahrstrSigHaltfall"`
			FahrstrTeilaufloesung []struct {
				Text  string `xml:",chardata"`
				Ref   string `xml:"Ref,attr"`
				Datei struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
					NurInfo   string `xml:"NurInfo,attr"`
				} `xml:"Datei"`
			} `xml:"FahrstrTeilaufloesung"`
			FahrstrSignal []struct {
				Text                      string `xml:",chardata"`
				Ref                       string `xml:"Ref,attr"`
				FahrstrSignalZeile        string `xml:"FahrstrSignalZeile,attr"`
				FahrstrSignalErsatzsignal string `xml:"FahrstrSignalErsatzsignal,attr"`
				Datei                     struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
					NurInfo   string `xml:"NurInfo,attr"`
				} `xml:"Datei"`
			} `xml:"FahrstrSignal"`
			FahrstrVSignal []struct {
				Text                string `xml:",chardata"`
				Ref                 string `xml:"Ref,attr"`
				FahrstrSignalSpalte string `xml:"FahrstrSignalSpalte,attr"`
				Datei               struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
					NurInfo   string `xml:"NurInfo,attr"`
				} `xml:"Datei"`
			} `xml:"FahrstrVSignal"`
			FahrstrWeiche []struct {
				Text               string `xml:",chardata"`
				Ref                string `xml:"Ref,attr"`
				FahrstrWeichenlage string `xml:"FahrstrWeichenlage,attr"`
				Datei              struct {
					Text      string `xml:",chardata"`
					Dateiname string `xml:"Dateiname,attr"`
					NurInfo   string `xml:"NurInfo,attr"`
				} `xml:"Datei"`
			} `xml:"FahrstrWeiche"`
		} `xml:"Fahrstrasse"`
	} `xml:"Strecke"`
}
