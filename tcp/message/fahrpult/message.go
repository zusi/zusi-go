package fahrpult

// Anwendung 02 ("Fahrpult")
type FahrpultMessage struct {
	NeededData    *NeededDataMessage `zusi:"0003" json:"needed_data,omitempty"`
	AckNeededData *AckNeededData     `zusi:"0004" json:"ack_needed_data,omitempty"`
	DataFtd       *DataFtd           `zusi:"000A" json:"data_ftd,omitempty"`
	DataProg      *DataProg          `zusi:"000C" json:"data_prog,omitempty"`
	Control       *Control           `zusi:"010B" json:"control,omitempty"`
}

func (FahrpultMessage) IsToplevel() bool {
	return true
}

// 5.3.3.1 Befehl 00 03 - NEEDED_DATA (Client → Zusi)
type NeededDataMessage struct {
	Fuehrerstandsanzeigen  *Fuehrerstandsanzeigen  `zusi:"000A" json:"fuehrerstandsanzeigen,omitempty"`
	Fuehrerstandsbedienung *Fuehrerstandsbedienung `zusi:"000B" json:"fuehrerstandsbedienung,omitempty"`
	Programmdaten          *Programmdaten          `zusi:"000C" json:"programmdaten,omitempty"`
}

type Fuehrerstandsanzeigen struct {
	Anzeigen []uint16 `zusi:"0001" json:"anzeigen,omitempty"`
}

type Fuehrerstandsbedienung struct{}

type Programmdaten struct {
	Anzeigen []uint16 `zusi:"0001" json:"anzeigen,omitempty"`
}

// 5.3.3.2 Befehl 00 04 - ACK_NEEDED_DATA (Zusi → Client)
type AckNeededData struct {
	ErrorCode *byte `zusi:"0001" json:"error_code,omitempty"` // Der Befehl wurde akzeptiert, wenn das Byte auf 00 steht. Wird der Befehl nicht akzeptiert, wird stattdes- sen ein anderes Byte gesendet.
}

// 5.3.3.3 Befehl 00 0A - DATA_FTD (Zusi → Client)
type DataFtd struct {
	Geschwindigkeit         *float32 `zusi:"0001" json:"geschwindigkeit,omitempty"`          // Single: m/s
	DruckHauptluftleitung   *float32 `zusi:"0002" json:"druck_hauptluftleitung,omitempty"`   // Single: bar
	DruckBremszylinder      *float32 `zusi:"0003" json:"druck_bremszylinder,omitempty"`      // Single: bar
	DruckHauptluftbehaelter *float32 `zusi:"0004" json:"druck_hauptluftbehaelter,omitempty"` // Single: bar
	LuftpresserLaeuft       *float32 `zusi:"0005" json:"luftpresser_laeuft,omitempty"`       // Single: aus/an
	LuftstromFvb            *float32 `zusi:"0006" json:"luftstrom_fvb,omitempty"`            // Single: -1..0..+1
	LuftstromZbv            *float32 `zusi:"0007" json:"luftstrom_zbv,omitempty"`            // Single: -1..0..+1
	LuefterAn               *float32 `zusi:"0008" json:"luefter_an,omitempty"`               // Single: aus/an
	ZugkraftGesamt          *float32 `zusi:"0009" json:"zugkraft_gesamt,omitempty"`          // Single: N
	ZugkraftProAchse        *float32 `zusi:"000A" json:"zugkraft_pro_achse,omitempty"`       // Single: N
	ZugkraftSollGesamt      *float32 `zusi:"000B" json:"zugkraft_soll_gesamt,omitempty"`     // Single: N
	ZugkraftSollProAchse    *float32 `zusi:"000C" json:"zugkraft_soll_pro_achse,omitempty"`  // Single: N
	Oberstrom               *float32 `zusi:"000D" json:"oberstrom,omitempty"`                // Single: A
	Fahrleitungsspannung    *float32 `zusi:"000E" json:"fahrleitungsspannung,omitempty"`     // Single: V
	Motordrehzahl           *float32 `zusi:"000F" json:"motordrehzahl,omitempty"`            // Single: 1/min
	UhrzeitStunde           *float32 `zusi:"0010" json:"uhrzeit_stunde,omitempty"`           // Single: Stunde (Zeigerposition Analoguhren)
	UhrzeitMinute           *float32 `zusi:"0011" json:"uhrzeit_minute,omitempty"`           // Single: Minute (Zeigerposition Analoguhren)
	UhrzeitSekunde          *float32 `zusi:"0012" json:"uhrzeit_sekunde,omitempty"`          // Single: Sekunde (Zeigerposition Analoguhren)
	Hauptschalter           *float32 `zusi:"0013" json:"hauptschalter,omitempty"`            // Single: aus/an
	Trennschuetz            *float32 `zusi:"0014" json:"trennschuetz,omitempty"`             // Single: aus/an
	Fahrstufe               *float32 `zusi:"0015" json:"fahrstufe,omitempty"`                // Single: 1
	//d3DFenster                                       *float32              `zusi:"0016" json:"_3_d_fenster,omitempty"`              // Single: Nicht sinnvoll zu verwenden
	AfbSollgeschwindigkeit   *float32              `zusi:"0017" json:"afb_sollgeschwindigkeit,omitempty"`   // Single: m/s
	DruckHilfsluftbehaelter  *float32              `zusi:"0018" json:"druck_hilfsluftbehaelter,omitempty"`  // Single: bar
	ZurueckgelegterGesamtweg *float32              `zusi:"0019" json:"zurueckgelegter_gesamtweg,omitempty"` // Single: m
	LmGetriebe               *float32              `zusi:"001A" json:"lm_getriebe,omitempty"`               // Single: aus/an
	LmSchleudern             *float32              `zusi:"001B" json:"lm_schleudern,omitempty"`             // Single: aus/an
	LmGleiten                *float32              `zusi:"001C" json:"lm_gleiten,omitempty"`                // Single: aus/an
	LmMgBremse               *float32              `zusi:"001D" json:"lm_mg_bremse,omitempty"`              // Single: aus/an
	LmHBremse                *float32              `zusi:"001E" json:"lm_h_bremse,omitempty"`               // Single: aus/an
	LmRBremse                *float32              `zusi:"001F" json:"lm_r_bremse,omitempty"`               // Single: aus/an
	LmHochabbremsung         *float32              `zusi:"0020" json:"lm_hochabbremsung,omitempty"`         // Single: aus/an
	LmSchnellbremsung        *float32              `zusi:"0021" json:"lm_schnellbremsung,omitempty"`        // Single: aus/an
	StatusNotbremssystem     *StatusNotbremssystem `zusi:"0022" json:"status_notbremssystem,omitempty"`
	LmUhrzeitDigital         *float32              `zusi:"0023" json:"lm_uhrzeit_digital,omitempty"`       // Single: 0...1 (0:00 bis 24:00 Uhr)
	LmDrehzahlverstellung    *float32              `zusi:"0024" json:"lm_drehzahlverstellung,omitempty"`   // Single: aus/an
	LmFahrtrichtungVor       *float32              `zusi:"0025" json:"lm_fahrtrichtung_vor,omitempty"`     // Single: aus/an
	LmFahrtrichtungZurueck   *float32              `zusi:"0026" json:"lm_fahrtrichtung_zurueck,omitempty"` // Single: aus/an
	LmFahrtrichtungM         *float32              `zusi:"0027" json:"lm_fahrtrichtung_m,omitempty"`       // Single: aus/an
	//hintergrundbild                                 *float32              `zusi:"0028" json:"_hintergrundbild,omitempty"`                   // Single: Nicht sinnvoll zu verwenden
	Motordrehmoment                                  *float32                `zusi:"0029" json:"motordrehmoment,omitempty"`                    // Single: Nm
	MotorlastNormiert                                *float32                `zusi:"002A" json:"motorlast_normiert,omitempty"`                 // Single: 1 (0...1)
	Tunnel                                           *float32                `zusi:"002B" json:"tunnel,omitempty"`                             // Single: aus/an
	SchienenstossWeiche                              *float32                `zusi:"002C" json:"schienenstoss_weiche,omitempty"`               // Single: aus/an
	Stahlbruecke                                     *float32                `zusi:"002D" json:"stahlbruecke,omitempty"`                       // Single: aus/an
	Steinbruecke                                     *float32                `zusi:"002E" json:"steinbruecke,omitempty"`                       // Single: aus/an
	XKoordinate                                      *float32                `zusi:"002F" json:"x_koordinate,omitempty"`                       // Single: m (Bez. Strecken-UTM-Punkt)
	YKoordinate                                      *float32                `zusi:"0030" json:"y_koordinate,omitempty"`                       // Single: m (Bez. Strecken-UTM-Punkt)
	ZKoordinate                                      *float32                `zusi:"0031" json:"z_koordinate,omitempty"`                       // Single: m
	UTMReferenzpunktXkm                              *float32                `zusi:"0032" json:"utm_referenzpunkt_xkm,omitempty"`              // Single: km
	UTMReferenzpunktYkm                              *float32                `zusi:"0033" json:"utm_referenzpunkt_ykm,omitempty"`              // Single: km
	UTMZone                                          *float32                `zusi:"0034" json:"utm_zone,omitempty"`                           // Single
	UTMZone2                                         *float32                `zusi:"0035" json:"utm_zone_2,omitempty"`                         // Single
	AFBAn                                            *float32                `zusi:"0036" json:"afb_an,omitempty"`                             // Single: aus/an
	Individuell01                                    *float32                `zusi:"0037" json:"individuell_01,omitempty"`                     // Single
	Individuell02                                    *float32                `zusi:"0038" json:"individuell_02,omitempty"`                     // Single
	Individuell03                                    *float32                `zusi:"0039" json:"individuell_03,omitempty"`                     // Single
	Individuell04                                    *float32                `zusi:"003A" json:"individuell_04,omitempty"`                     // Single
	Individuell05                                    *float32                `zusi:"003B" json:"individuell_05,omitempty"`                     // Single
	Individuell06                                    *float32                `zusi:"003C" json:"individuell_06,omitempty"`                     // Single
	Individuell07                                    *float32                `zusi:"003D" json:"individuell_07,omitempty"`                     // Single
	Individuell08                                    *float32                `zusi:"003E" json:"individuell_08,omitempty"`                     // Single
	Individuell09                                    *float32                `zusi:"003F" json:"individuell_09,omitempty"`                     // Single
	Individuell10                                    *float32                `zusi:"0040" json:"individuell_10,omitempty"`                     // Single
	Individuell11                                    *float32                `zusi:"0041" json:"individuell_11,omitempty"`                     // Single
	Individuell12                                    *float32                `zusi:"0042" json:"individuell_12,omitempty"`                     // Single
	Individuell13                                    *float32                `zusi:"0043" json:"individuell_13,omitempty"`                     // Single
	Individuell14                                    *float32                `zusi:"0044" json:"individuell_14,omitempty"`                     // Single
	Individuell15                                    *float32                `zusi:"0045" json:"individuell_15,omitempty"`                     // Single
	Individuell16                                    *float32                `zusi:"0046" json:"individuell_16,omitempty"`                     // Single
	Individuell17                                    *float32                `zusi:"0047" json:"individuell_17,omitempty"`                     // Single
	Individuell18                                    *float32                `zusi:"0048" json:"individuell_18,omitempty"`                     // Single
	Individuell19                                    *float32                `zusi:"0049" json:"individuell_19,omitempty"`                     // Single
	Individuell20                                    *float32                `zusi:"004A" json:"individuell_20,omitempty"`                     // Single
	Datum                                            *float32                `zusi:"004B" json:"datum,omitempty"`                              // Single (Tage mit 0 = 30.12.1899)
	Gleiskruemmung                                   *float32                `zusi:"004C" json:"gleiskruemmung,omitempty"`                     // Single 1000/m
	Streckenhoechstgeschwindigkeit                   *float32                `zusi:"004D" json:"streckenhoechstgeschwindigkeit,omitempty"`     // Single: m/s
	ZugkraftvorschlagAutopilot                       *float32                `zusi:"004E" json:"zugkraftvorschlag_autopilot,omitempty"`        // Single: N
	BeschleunigungX                                  *float32                `zusi:"004F" json:"beschleunigung_x,omitempty"`                   // Single: m/s^2
	BeschleunigungY                                  *float32                `zusi:"0050" json:"beschleunigung_y,omitempty"`                   // Single: m/s^2
	BeschleunigungZ                                  *float32                `zusi:"0051" json:"beschleunigung_z,omitempty"`                   // Single: m/s^2
	DrehbeschleunigungXAchse                         *float32                `zusi:"0052" json:"drehbeschleunigung_x_achse,omitempty"`         // Single: rad/s^2
	DrehbeschleunigungYAchse                         *float32                `zusi:"0053" json:"drehbeschleunigung_y_achse,omitempty"`         // Single: rad/s^2
	DrehbeschleunigungZAchse                         *float32                `zusi:"0054" json:"drehbeschleunigung_z_achse,omitempty"`         // Single: rad/s^2
	Stromabnehmer                                    *float32                `zusi:"0055" json:"stromabnehmer,omitempty"`                      // Single: 2x4 bit (4bit: 1 fuer SA=oben; 4 bit: 1 fuer SA hebt sich gerade)
	LmFederspeicherbremse                            *float32                `zusi:"0056" json:"lm_federspeicherbremse,omitempty"`             // Single: aus/an
	ZustandFederspeicherbremse                       *float32                `zusi:"0057" json:"zustand_federspeicherbremse,omitempty"`        // Single -1, 0, 1, 2 (nicht vorhanden, aus, an, blinkend)
	SteuerwagenLmGetriebe                            *float32                `zusi:"0058" json:"steuerwagen_lm_getriebe,omitempty"`            // Single: aus/an
	SteuerwagenLmSchleudern                          *float32                `zusi:"0059" json:"steuerwagen_lm_schleudern,omitempty"`          // Single: aus/an
	SteuerwagenLmGleiten                             *float32                `zusi:"005A" json:"steuerwagen_lm_gleiten,omitempty"`             // Single: aus/an
	SteuerwagenLmHBremse                             *float32                `zusi:"005B" json:"steuerwagen_lm_h_bremse,omitempty"`            // Single: aus/an
	SteuerwagenLmRBremse                             *float32                `zusi:"005C" json:"steuerwagen_lm_r_bremse,omitempty"`            // Single: aus/an
	SteuerwagenLmDrehzahlverstellung                 *float32                `zusi:"005D" json:"steuerwagen_lm_drehzahlverstellung,omitempty"` // Single: aus/an
	DruckZweitbehaelter                              *float32                `zusi:"005E" json:"druck_zweitbehaelter,omitempty"`               // Single: bar
	GeschwindigkeitAbsolut                           *float32                `zusi:"005F" json:"geschwindigkeit_absolut,omitempty"`            // Single: m/s
	ZugIstEntgleist                                  *float32                `zusi:"0060" json:"zug_ist_entgleist,omitempty"`                  // Single: aus/an
	KilometrierungZugspitze                          *float32                `zusi:"0061" json:"kilometrierung_zugspitze,omitempty"`           // Single: km
	Motorstrom                                       *float32                `zusi:"0062" json:"motorstrom,omitempty"`                         // Single: A
	Motorspannung                                    *float32                `zusi:"0063" json:"motorspannung,omitempty"`                      // Single: V
	StatusSifa                                       *StatusSifa             `zusi:"0064" json:"status_sifa,omitempty"`
	StatusZugsicherung                               *StatusZugbeeinflussung `zusi:"0065" json:"status_zugsicherung,omitempty"`
	StatusTuersystem                                 *StatusTuersystem       `zusi:"0066" json:"status_tuersystem,omitempty"`
	Individuell21                                    *float32                `zusi:"0067" json:"individuell_21,omitempty"`                          // Single
	Individuell22                                    *float32                `zusi:"0068" json:"individuell_22,omitempty"`                          // Single
	Individuell23                                    *float32                `zusi:"0069" json:"individuell_23,omitempty"`                          // Single
	Individuell24                                    *float32                `zusi:"006A" json:"individuell_24,omitempty"`                          // Single
	Individuell25                                    *float32                `zusi:"006B" json:"individuell_25,omitempty"`                          // Single
	Individuell26                                    *float32                `zusi:"006C" json:"individuell_26,omitempty"`                          // Single
	Individuell27                                    *float32                `zusi:"006D" json:"individuell_27,omitempty"`                          // Single
	Individuell28                                    *float32                `zusi:"006E" json:"individuell_28,omitempty"`                          // Single
	Individuell29                                    *float32                `zusi:"006F" json:"individuell_29,omitempty"`                          // Single
	Individuell30                                    *float32                `zusi:"0070" json:"individuell_30,omitempty"`                          // Single
	Individuell31                                    *float32                `zusi:"0071" json:"individuell_31,omitempty"`                          // Single
	Individuell32                                    *float32                `zusi:"0072" json:"individuell_32,omitempty"`                          // Single
	Individuell33                                    *float32                `zusi:"0073" json:"individuell_33,omitempty"`                          // Single
	Individuell34                                    *float32                `zusi:"0074" json:"individuell_34,omitempty"`                          // Single
	Individuell35                                    *float32                `zusi:"0075" json:"individuell_35,omitempty"`                          // Single
	Individuell36                                    *float32                `zusi:"0076" json:"individuell_36,omitempty"`                          // Single
	Individuell37                                    *float32                `zusi:"0077" json:"individuell_37,omitempty"`                          // Single
	Individuell38                                    *float32                `zusi:"0078" json:"individuell_38,omitempty"`                          // Single
	Individuell39                                    *float32                `zusi:"0079" json:"individuell_39,omitempty"`                          // Single
	Individuell40                                    *float32                `zusi:"007A" json:"individuell_40,omitempty"`                          // Single
	SteuerwagenLuefterAn                             *float32                `zusi:"007B" json:"steuerwagen_luefter_an,omitempty"`                  // Single: aus/an
	SteuerwagenZugkraftGesamt                        *float32                `zusi:"007C" json:"steuerwagen_zugkraft_gesamt,omitempty"`             // Single: N
	SteuerwagenZugraftProAchse                       *float32                `zusi:"007D" json:"steuerwagen_zugraft_pro_achse,omitempty"`           // Single: N
	SteuerwagenZugkraftSollGesamt                    *float32                `zusi:"007E" json:"steuerwagen_zugkraft_soll_gesamt,omitempty"`        // Single: N
	SteuerwagenZugraftSollProAchse                   *float32                `zusi:"007F" json:"steuerwagen_zugraft_soll_pro_achse,omitempty"`      // Single: N
	SteuerwagenOberstrom                             *float32                `zusi:"0080" json:"steuerwagen_oberstrom,omitempty"`                   // Single: N
	SteuerwagenFahrleitungsspannung                  *float32                `zusi:"0081" json:"steuerwagen_fahrleitungsspannung,omitempty"`        // Single V
	SteuerwagenMotordrehzahl                         *float32                `zusi:"0082" json:"steuerwagen_motordrehzahl,omitempty"`               // Single 1/min
	SteuerwagenHauptschalter                         *float32                `zusi:"0083" json:"steuerwagen_hauptschalter,omitempty"`               // Single aus/an
	SteuerwagenTrennschuetz                          *float32                `zusi:"0084" json:"steuerwagen_trennschuetz,omitempty"`                // Single: aus/an
	SteuerwagenFahrstufe                             *float32                `zusi:"0085" json:"steuerwagen_fahrstufe,omitempty"`                   // Single: 1
	SteuerwagenMotordrehmoment                       *float32                `zusi:"0086" json:"steuerwagen_motordrehmoment,omitempty"`             // Single: Nm
	SteuerwagenMotorlastNormiert                     *float32                `zusi:"0087" json:"steuerwagen_motorlast_normiert,omitempty"`          // Single: 1 (0...1)
	SteuerwagenStromabnehmer                         *float32                `zusi:"0088" json:"steuerwagen_stromabnehmer,omitempty"`               // Single
	SteuerwagenMotorstrom                            *float32                `zusi:"0089" json:"steuerwagen_motorstrom,omitempty"`                  // Single: A
	SteuerwagenMotorspannung                         *float32                `zusi:"008A" json:"steuerwagen_motorspannung,omitempty"`               // Single: V
	GeschwindigkeitAbsolutInklSchleudern             *float32                `zusi:"008B" json:"geschwindigkeit_absolut_inkl_schleudern,omitempty"` // Single: m/s
	BatteriehauptschalterAus                         *float32                `zusi:"008C" json:"batteriehauptschalter_aus,omitempty"`               // Single: aus/an
	StatusFahrzeug                                   *StatusFahrzeug         `zusi:"008D" json:"status_fahrzeug,omitempty"`
	StatusZugverband                                 *StatusZugverband       `zusi:"008E" json:"status_zugverband,omitempty"`
	Bremsprobenfunktion                              *float32                `zusi:"008F" json:"bremsprobenfunktion,omitempty"`                                 // Single: 0 (aus, >0 aktiv)
	ZugUndBremsGesamtkraftSollNormiert               *float32                `zusi:"0090" json:"zug_und_brems_gesamtkraft_soll_normiert,omitempty"`             // Single: 1 ((0...1) normiert auf aktuelle Fmax
	SteuerwagenZugUndBremsGesamtkraftSollNormiert    *float32                `zusi:"0091" json:"steuerwagen_zug_und_brems_gesamtkraft_soll_normiert,omitempty"` // Single: 1 ((0...1) normiert auf aktuelle Fmax
	StatusWeichen                                    *StatusWeichen          `zusi:"0092" json:"status_weichen,omitempty"`
	ZugUndBremsGesamtkraftAbsolutNormiert            *float32                `zusi:"0093" json:"zug_und_brems_gesamtkraft_absolut_normiert,omitempty"`             // Single: 1 ((0...1) normiert auf aktuelle Fmax
	SteuerwagenZugUndBremsGesamtkraftAbsolutNormiert *float32                `zusi:"0094" json:"steuerwagen_zug_und_brems_gesamtkraft_absolut_normiert,omitempty"` // Single: 1 ((0...1) normiert auf aktuelle Fmax
}

// 5.3.3.3.2
type StatusNotbremssystem struct {
	Bauart               *string `zusi:"0001" json:"bauart,omitempty"`
	StatusNotbremssystem *byte   `zusi:"0002" json:"status_notbremssystem,omitempty"` // 0: NBÜ aus 1: NBÜ bereit 2: Notbremse gezogen 3: Notbremse wirkt (NBÜ bereit) 4: NBÜ durch Lokführer aktiviert 5 Notbremse wirkt (NBÜ aus)
	LmSystemBereit       *byte   `zusi:"0003" json:"lm_system_bereit,omitempty"`      // 1: an
	LmNotbremsung        *byte   `zusi:"0004" json:"lm_notbremsung,omitempty"`
	TestmodusAktiv       *byte   `zusi:"0005" json:"testmodus_aktiv,omitempty"`
}

// 5.3.3.3.3
type StatusSifa struct {
	Bauart              *string `zusi:"0001" json:"bauart,omitempty"`
	LmSifa              *byte   `zusi:"0002" json:"lm_sifa,omitempty"`
	SifaHupe            *byte   `zusi:"0003" json:"sifa_hupe,omitempty"`
	SifaHauptschalter   *byte   `zusi:"0004" json:"sifa_hauptschalter,omitempty"`
	SifaStoerschalter   *byte   `zusi:"0005" json:"sifa_stoerschalter,omitempty"`
	SifaLuftabsperrhahn *byte   `zusi:"0006" json:"sifa_luftabsperrhahn,omitempty"`
}

// 5.3.3.3.4 Status Zugbeeinflussung
type StatusZugbeeinflussung struct {
	Bauart              *string              `zusi:"0001" json:"bauart,omitempty"`
	IndusiEinstellungen *IndusiEinstellungen `zusi:"0002" json:"indusi_einstellungen,omitempty"`
	IndusiZustand       *IndusiZustand       `zusi:"0003" json:"indusi_zustand,omitempty"`
	EtcsEinstellungen   *EtcsEinstellungen   `zusi:"0004" json:"etcs_einstellungen,omitempty"`
	EtcsBetriebsdaten   *EtcsBetriebsdaten   `zusi:"0005" json:"etcs_betriebsdaten,omitempty"`
	ZubEinstellungen    *ZubEinstellungen    `zusi:"0006" json:"zub_einstellungen,omitempty"`
	ZubBetriebsdaten    *ZubBetriebsdaten    `zusi:"0007" json:"zub_betriebsdaten,omitempty"`
}

// 5.3.3.3.3.4.2 Indusi Analogsysteme und Basisdaten
type IndusiEinstellungen struct {
	Zugart           *byte     `zusi:"0001" json:"zugart,omitempty"` // Zugart:1: Zugart muss noch bestimmt werden 2: U 3: M 4: O	5: S-Bahn-Modus
	TfNummer         *string   `zusi:"0002" json:"tf_nummer,omitempty"`
	Zugnummer        *string   `zusi:"0003" json:"zugnummer,omitempty"`
	Grunddaten       *Zugdaten `zusi:"0004" json:"grunddaten,omitempty"` // LZB
	Ersatzzugdaten   *Zugdaten `zusi:"0005" json:"ersatzzugdaten,omitempty"`
	AktiveZugdaten   *Zugdaten `zusi:"0006" json:"aktive_zugdaten,omitempty"`
	Hauptschalter    *byte     `zusi:"0007" json:"hauptschalter,omitempty"`     // Hauptschalter   1: Indusi ausgeschaltet 2: Indusi eingeschaltet
	PzbStoerschalter *byte     `zusi:"0008" json:"pzb_stoerschalter,omitempty"` // Störschalter    1: Indusi abgeschaltet  2: Indusi eingeschaltet
	LzbStoerschalter *byte     `zusi:"0009" json:"lzb_stoerschalter,omitempty"` // LZB
	Luftabsperrhahn  *byte     `zusi:"000A" json:"luftabsperrhahn,omitempty"`   // Luftabsperrhahn 1: abgesperrt           2: offen
}

type Zugdaten struct {
	Bremshundertstel  *uint16 `zusi:"0001" json:"bremshundertstel,omitempty"`
	Bremsart          *uint16 `zusi:"0002" json:"bremsart,omitempty"`
	Zuglaenge         *uint16 `zusi:"0003" json:"zuglaenge,omitempty"` // LZB
	Vmax              *uint16 `zusi:"0004" json:"vmax,omitempty"`      // LZB in km/h
	Zugart            *byte   `zusi:"0005" json:"zugart,omitempty"`
	Modus             *byte   `zusi:"0006" json:"modus,omitempty"`             // Nur relevant für AktiveZugdaten. 5: Ersatzzugdaten 6: Normalbetrieb
	Klartextmeldungen *byte   `zusi:"000B" json:"klartextmeldungen,omitempty"` // 0: Keine Klartextmeldungen möglich 1: Keine Klartextmeldungen möglich aber nicht aktiv 2: Klartextmeldungen aktiv 3: nur Klartextmeldungen möglich
}

type IndusiZustand struct {
	Zugsicherung               *uint16                  `zusi:"0002" json:"zugsicherung,omitempty"`         // 1: Ausgeschaltet 2: abgeschaltet/gestört (1000Hz blinkt) 3: Hauptluftleitung unter 2,2 bar (1000Hz blinkt) 4: Aufforderung zur Zugdateneingabe 5: Normalbetrieb 6: Funktionsprüfung
	Zwangsbremsungsgrund       *uint16                  `zusi:"0003" json:"zwangsbremsungsgrund,omitempty"` // 0: keine Zwangsbremsung, sonst Zwansbremsung aktiv wegen: 1. Wachsam 2. 1000Hz 3. 500Hz 4. 2000Hz 5. Kein Halt nach Befreiung aus Zwangsbremsung 6. Fahrzeug-v-Max überschritten 7. Funktionsprüfung 8. 500Hz nach Befreiung 9. LZB-Halt überfahren 10. LZB-Rechnerausfall 11. LZB-Nothalt überfahren 12. Übertragungsausfall in verdeckter Aufnahme
	ZwangsbremsungsgrundString *string                  `zusi:"0004" json:"zwangsbremsungsgrund_*string,omitempty"`
	Lm1000Hz                   *byte                    `zusi:"0005" json:"lm_1000_hz,omitempty"`
	LmU                        *byte                    `zusi:"0006" json:"lm_u,omitempty"`
	LmM                        *byte                    `zusi:"0007" json:"lm_m,omitempty"`
	LmO                        *byte                    `zusi:"0008" json:"lm_o,omitempty"`
	IndusiHupe                 *byte                    `zusi:"0009" json:"indusi_hupe,omitempty"` // 1: Hupe 2: Zwangsbremsung
	Lm500Hz                    *byte                    `zusi:"000A" json:"lm_500_hz,omitempty"`
	LmBefehl                   *byte                    `zusi:"000B" json:"lm_befehl,omitempty"`
	ZusatzinfoMelderbild       *byte                    `zusi:"000C" json:"zusatzinfo_melderbild,omitempty"` // PZB90 0: Normal 1: 1000Hz nach 700m 2: Restriktiv 3: Restriktiv+1000Hz 4: Restriktiv+500Hz 5: Prüfablauf nach LZB-Übertragungsausfall
	Lzb                        *uint16                  `zusi:"000D" json:"lzb,omitempty"`                   // LZB 0: Keine LZB-Führung 1: Normale Fahrt 2: Nothalt 3: LZB-Halt überfahren 4: Rechnerausfall 5: Nachfahrauftrag 6: Funktionsprüfung
	LzbEndeVerfahren           *LzbAuftrag              `zusi:"000E" json:"lzb_ende_verfahren,omitempty"`
	LzbErsatzauftrag           *LzbAuftrag              `zusi:"000F" json:"lzb_ersatzauftrag,omitempty"`
	LzbFalschfahrauftrag       *LzbAuftrag              `zusi:"0010" json:"lzb_falschfahrauftrag,omitempty"`
	LzbVorsichtauftrag         *LzbAuftrag              `zusi:"0011" json:"lzb_vorsichtauftrag,omitempty"`
	LzbFahrtUeberHaltBefehl    *LzbAuftrag              `zusi:"0012" json:"lzb_fahrt_ueber_halt_befehl,omitempty"`
	LzbUebertragungsausfall    *LzbUebertragungsausfall `zusi:"0013" json:"lzb_uebertragungsausfall,omitempty"`
	LzbNothalt                 *LzbNothalt              `zusi:"0014" json:"lzb_nothalt,omitempty"`
	LzbRechnerausfall          *LzbRechnerausfall       `zusi:"0015" json:"lzb_rechnerausfall,omitempty"`
	LzbElAuftrag               *LzbElAuftrag            `zusi:"0016" json:"lzb_el_auftrag,omitempty"`
	LmH                        *byte                    `zusi:"0017" json:"lm_h,omitempty"`
	LmE40                      *byte                    `zusi:"0018" json:"lm_e40,omitempty"`
	LmEnde                     *byte                    `zusi:"0019" json:"lm_ende,omitempty"`
	LmB                        *byte                    `zusi:"001A" json:"lm_b,omitempty"`
	LmUe                       *byte                    `zusi:"001B" json:"lm_ue,omitempty"`
	LmG                        *byte                    `zusi:"001C" json:"lm_g,omitempty"`
	LmEL                       *byte                    `zusi:"001D" json:"lm_el,omitempty"`
	LmV40                      *byte                    `zusi:"001E" json:"lm_v40,omitempty"`
	LmS                        *byte                    `zusi:"001F" json:"lm_s,omitempty"`
	LmPruefStoer               *byte                    `zusi:"0020" json:"lm_pruefstoer,omitempty"`
	Sollgeschwindigkeit        *float32                 `zusi:"0021" json:"sollgeschwindigkeit,omitempty"` // m/s
	Zielgeschwindigkeit        *float32                 `zusi:"0022" json:"zielgeschwindigkeit,omitempty"` // m/s Wert < 0 = Dunkel
	Zielweg                    *float32                 `zusi:"0023" json:"zielweg,omitempty"`             // m Wert < 0 = Dunkel
	LmGBl                      *byte                    `zusi:"0024" json:"lm_g_bl,omitempty"`             // 0: aus 1: an 2: blinkt
	LmPruefStoerBl             *byte                    `zusi:"0025" json:"lm_pruefstoer_bl,omitempty"`    // 0: aus 1: an 2: blinkt
	CirElkeModus               *byte                    `zusi:"0026" json:"cir_elke_modus,omitempty"`      // 0: Normal 1: CIR-ELKE aktiv
	Anzeigemodus               *byte                    `zusi:"0027" json:"anzeigemodus,omitempty"`        // 0: Normal 1: Zugdatenanzeige im MFA
	LzbFunktionspruefung       *LzbFunktionspruefung    `zusi:"0028" json:"lzb_funktionspruefung,omitempty"`
	LmZugartLinks              *byte                    `zusi:"0029" json:"lm_zugart_links,omitempty"`  // PZB90 S-Bahn
	LmZugart65                 *byte                    `zusi:"002A" json:"lm_zugart_65,omitempty"`     // PZB90 S-Bahn
	LmZugartRechts             *byte                    `zusi:"002B" json:"lm_zugart_rechts,omitempty"` // PZB90 S-Bahn
}

type LzbAuftrag struct {
	Status *byte `zusi:"0001" json:"status,omitempty"` // 1: eingeleitet 2: quittiert bei Vorsichtauftrag: 3: Fahrt auf Sicht (V40 Melder Dauerlicht)
}

type LzbUebertragungsausfall struct {
	Zielgeschwindigkeit *float32 `zusi:"0001" json:"zielgeschwindigkeit,omitempty"` // m/s
	Status              *uint16  `zusi:"0002" json:"status,omitempty"`              // 1: eingeleitet 2: Ü Blinkt 3: erste Quittierung erfolt 4: Bedienung für 2. Quittierung gegeben 5: zweite Quittierung erfolgt 6: Ausfall nach verdeckter LZB Aufnahme (CE) 7: dito, Befehl blinkt
	Zielweg             *float32 `zusi:"0003" json:"zielweg,omitempty"`             // nur CIR-ELKE
}

type LzbNothalt struct {
	Status       *byte `zusi:"0001" json:"status,omitempty"`        // 1: empfangen 2: überfahren 3: aufgehoben
	WirdGesendet *byte `zusi:"0002" json:"wird_gesendet,omitempty"` // 1: wird gesendet
}

type LzbRechnerausfall struct {
	Status *byte `zusi:"0001" json:"status,omitempty"` // 1: alles dunkel 2: Befehlsmelder blinkt nach Neustart 3: Befehlsmelder Dauerlicht nach Quittierung
}

type LzbElAuftrag struct {
	Status *byte `zusi:"0001" json:"status,omitempty"` // 1: Hauptschalter aus (EL Dauerlicht) 2: Stromabnehmer senken (EL Blinkt)
}

type LzbFunktionspruefung struct {
	AlleMelderBlinken           *EmptyNode `zusi:"0001" json:"alle_melder_blinken,omitempty"`
	AnzeigeDerFuehrungsgroessen *EmptyNode `zusi:"0002" json:"anzeige_der_fuehrungsgroessen,omitempty"`
	BanUeaus                    *EmptyNode `zusi:"0003" json:"ban_ueaus,omitempty"`
	ZwangsbremsungAktiv         *EmptyNode `zusi:"0004" json:"zwangsbremsung_aktiv,omitempty"`
}

type EtcsEinstellungen struct {
	Zustand                  *byte         `zusi:"0001" json:"zustand,omitempty"` // Zusi -> Client
	Stm                      []EtcsStm     `zusi:"0002" json:"stm,omitempty"`     // Erstes STM = Aktives STM
	Zugdaten                 *EtcsZugdaten `zusi:"0003" json:"zugdaten,omitempty"`
	Spec                     *EtcsSpec     `zusi:"0004" json:"spec,omitempty"`
	EtcsStoerschalter        *byte         `zusi:"0005" json:"etcs_stoerschalter,omitempty"`         // 1: ETCS Abgeschaltet 2: ETCS Eingeschaltet
	EtcsHauptschalter        *byte         `zusi:"0006" json:"etcs_hauptschalter,omitempty"`         // 1: ETCS Abgeschaltet 2: ETCS Eingeschaltet
	Luftabsperrhahn          *byte         `zusi:"0007" json:"luftabsperrhahn,omitempty"`            // 1: Abgesperrt 2: Offen
	EtcsQuittierschalter     *byte         `zusi:"0008" json:"etcs_quittierschalter,omitempty"`      // 1: verlegt 2: grundstellung
	OverrideAnforderung      *byte         `zusi:"0009" json:"override_anforderung,omitempty"`       // 1: Override angefordert 2: Grundstellung
	Start                    *byte         `zusi:"000A" json:"start,omitempty"`                      // Nur Client -> Zusi : 1: Startkommando 2: Grundstellung
	LevelEinstellenAnfordern *byte         `zusi:"000B" json:"level_einstellen_anfordern,omitempty"` // Client -> Zusi: ETCS-Level (STM, 0, 1, usw.)
	StmSelectedIndex         *uint16       `zusi:"000C" json:"stm_selected_index,omitempty"`         // Client -> Zusi
	ModusEinstellenAnfordern *uint16       `zusi:"000D" json:"modus_einstellen_anfordern,omitempty"` // Client -> Zusi
	TafModus                 *byte         `zusi:"000E" json:"taf_modus,omitempty"`                  // Client -> Zusi 1: TAF Quittiert 2: Grundstellung 3: TAF Abgelehnt
	Zugneustart              *byte         `zusi:"000F" json:"zugneustart,omitempty"`                // Zusi -> Client 1: Zug wurde neu gestartet bzw. neu uebernommen.
	InfoTonAbspielen         *byte         `zusi:"0010" json:"info_ton,omitempty"`                   // Client -> Zusi: 1: Zusi soll den Info-Ton 1x abspielen
}

type EtcsStm struct {
	StmIndex *uint16 `zusi:"0001" json:"stm_index,omitempty"`
	StmName  *string `zusi:"0002" json:"stm_name,omitempty"`
}

type EtcsZugdaten struct {
	Bremshundertstel       *uint16 `zusi:"0001" json:"bremshundertstel,omitempty"` // in %
	Zugkategorie           *uint16 `zusi:"0002" json:"zugkategorie,omitempty"`
	Zuglaenge              *uint16 `zusi:"0003" json:"zuglaenge,omitempty"`              // in m
	Hoechstgeschwindigkeit *uint16 `zusi:"0004" json:"hoechstgeschwindigkeit,omitempty"` // in km/h
	Achslast               *uint16 `zusi:"0005" json:"achslast,omitempty"`               // in kg
	Zugnummer              *uint16 `zusi:"0006" json:"zugnummer,omitempty"`
	TfNummer               *uint16 `zusi:"0007" json:"tf_nummer,omitempty"`
}

type EtcsSpec struct {
	Reibwert *byte `zusi:"0001" json:"reibwert,omitempty"` // 1: vermindert 2: nicht vermindert
}

type EtcsBetriebsdaten struct {
	AktivesLevel *uint16 `zusi:"0001" json:"aktives_level,omitempty"` // 0: Undefiniert 1: STM 2: 0 3: 1 4: 2 5: 3
	/*
		00: Undefiniert
		01: FS
		02: OS
		03: SR
		04: SH
		05: UN
		06: SL
		07: SB
		08: TR
		09: PT
		10: SF
		11: IS
		12: NP
		13: NL
		14: SE
		15: SN
		16: RV
		17: LS
		18: PS
	*/
	AktiverModus               *uint16                `zusi:"0002" json:"aktiver_modus,omitempty"`
	BremsungsGrund             *uint16                `zusi:"0003" json:"bremsungs_grund,omitempty"`
	BremsungsGrundString       *string                `zusi:"0004" json:"bremsungs_grund_string,omitempty"`
	StmInfo                    *EtcsStmInfo           `zusi:"0005" json:"stm_info,omitempty"`
	LevelAnkuendigung          *EtcsLevelAnkuendigung `zusi:"0006" json:"level_ankuendigung,omitempty"`
	ModusAnkuendigung          *EtcsModusAnkuendigung `zusi:"0007" json:"modus_ankuendigung,omitempty"`
	Funkstatus                 *EtcsFunkstatus        `zusi:"0008" json:"funkstatus,omitempty"`
	Zielgeschwindigkeit        *float32               `zusi:"0009" json:"zielgeschwindigkeit,omitempty"`        // in m/s
	Zielweg                    *float32               `zusi:"000A" json:"zielweg,omitempty"`                    // in m, <0 = Dunkel
	AbstandBremseinsatzpunkt   *float32               `zusi:"000B" json:"abstand_bremseinsatzpunkt,omitempty"`  // in m
	Entlassungsgeschwindigkeit *float32               `zusi:"000C" json:"entlassungsgeschwindigkeit,omitempty"` // in m/s
	Sollgeschwindigkeit        *float32               `zusi:"000D" json:"sollgeschwindigkeit,omitempty"`        // in m/s
	Warngeschwindigkeit        *float32               `zusi:"000E" json:"warngeschwindigkeit,omitempty"`        // in m/s
	Bremsgeschwindigkeit       *float32               `zusi:"000F" json:"bremsgeschwindigkeit,omitempty"`       // in m/s
	Zwangsbremsgeschwindigkeit *float32               `zusi:"0010" json:"zwangsbremsgeschwindigkeit,omitempty"` // in m/s
	BremskurveLaeuft           *byte                  `zusi:"0011" json:"bremskurve_laeuft,omitempty"`          // 0: nein 1: ja
	Vorschaupunkte             []EtcsVorschaupunkt    `zusi:"0012" json:"vorschaupunkte,omitempty"`
	OverrideAktiv              *byte                  `zusi:"0013" json:"override_aktiv,omitempty"` // Zusi -> Client
	NotrufStatus               *byte                  `zusi:"0014" json:"notruf_status,omitempty"`
	Betriebszwangsbremsung     *byte                  `zusi:"0015" json:"betriebszwangsbremsung,omitempty"`
}

type EtcsStmInfo struct {
	StmIndex *uint16 `zusi:"0001" json:"stm_index,omitempty"` // Index des aktiven STM-System, von 1 beginnend gemäß Reihenfolge in der ftd-Datei
}

type EtcsLevelAnkuendigung struct {
	NeuesLevel  *uint16 `zusi:"0001" json:"neues_level,omitempty"`
	Quittierung *byte   `zusi:"0002" json:"quittierung,omitempty"`
}

type EtcsModusAnkuendigung struct {
	NeuerModus  *uint16 `zusi:"0001" json:"neuer_modus,omitempty"`
	Quittierung *byte   `zusi:"0002" json:"quittierung,omitempty"`
}

type EtcsFunkstatus struct {
	Zustand *byte `zusi:"0001" json:"zustand,omitempty"`
}

type EtcsVorschaupunkt struct {
	Herkunft        *uint16  `zusi:"0001" json:"herkunft,omitempty"`        // 1: Strecke 3: Hauptsignal 9: Rangiersignal 14: ETCS
	Geschwindigkeit *float32 `zusi:"0002" json:"geschwindigkeit,omitempty"` // in m/s, -1 = ETCS Ende
	Abstand         *float32 `zusi:"0003" json:"abstand,omitempty"`         // in m
	Hoehenwert      *float32 `zusi:"0004" json:"hoehenwert,omitempty"`      // in m
}

type ZubEinstellungen struct {
	BrhWert   *uint16 `zusi:"0001" json:"brh_wert,omitempty"`
	Zuglaenge *uint16 `zusi:"0003" json:"zuglaenge,omitempty"`
	Vmax      *uint16 `zusi:"0004" json:"vmax,omitempty"` // in km/h
}

type ZubBetriebsdaten struct {
	LmGnt                 *byte   `zusi:"0001" json:"lm_gnt,omitempty"`
	LmGntUe               *byte   `zusi:"0002" json:"lm_gnt_ue,omitempty"`
	LmGntG                *byte   `zusi:"0003" json:"lm_gnt_g,omitempty"`
	LmGntS                *byte   `zusi:"0004" json:"lm_gnt_s,omitempty"`
	LmGntGst              *byte   `zusi:"0005" json:"lm_gnt_gst,omitempty"`
	LmGntStoer            *byte   `zusi:"0006" json:"lm_gnt_stoer,omitempty"`
	StatusLmGntUe         *byte   `zusi:"0007" json:"status_lm_gnt_ue,omitempty"`
	StatusLmGntG          *byte   `zusi:"0008" json:"status_lm_gnt_g,omitempty"`
	StatusLmGntS          *byte   `zusi:"0009" json:"status_lm_gnt_s,omitempty"`
	Zwangsbremsung        *uint16 `zusi:"000A" json:"zwangsbremsung,omitempty"`
	BetriebsbremsungAktiv *byte   `zusi:"000B" json:"betriebsbremsung_aktiv,omitempty"`
}

// 5.3.3.3.5
type StatusTuersystem struct {
	Bauart                       *string `zusi:"0001" json:"bauart,omitempty"`
	StatusLinks                  *byte   `zusi:"0002" json:"status_links,omitempty"`          // 0: zu 1: öffnend 2: offen 3: Fahrgastwechsel abgeschlossen 4: schließend 5: gestört 6: blockiert
	StatusRechts                 *byte   `zusi:"0003" json:"status_rechts,omitempty"`         // 0: zu 1: öffnend 2: offen 3: Fahrgastwechsel abgeschlossen 4: schließend 5: gestört 6: blockiert
	TraktionssperreAktiv         *byte   `zusi:"0004" json:"traktionssperre_aktiv,omitempty"` // 1: an
	Seitenwahlschalter           *byte   `zusi:"0005" json:"seitenwahlschalter,omitempty"`    // 0: zu 1: links 2: rechts 3: beide
	LmLinks                      *byte   `zusi:"0006" json:"lm_links,omitempty"`
	LmRechts                     *byte   `zusi:"0007" json:"lm_rechts,omitempty"`
	StatusLmLinks                *byte   `zusi:"0008" json:"status_lm_links,omitempty"`
	StatusLmRechts               *byte   `zusi:"0009" json:"status_lm_rechts,omitempty"`
	LmZwangsschliessen           *byte   `zusi:"000A" json:"lm_zwangsschliessen,omitempty"`
	StatusLmZwangsschliessen     *byte   `zusi:"000B" json:"status_lm_zwangsschliessen,omitempty"`
	LmLinksUndRechts             *byte   `zusi:"000C" json:"lm_links_und_rechts,omitempty"`
	StatusLmLinksUndRechts       *byte   `zusi:"000D" json:"status_lm_links_und_rechts,omitempty"`
	ZentralesOeffnenLinks        *byte   `zusi:"000E" json:"zentrales_oeffnen_links,omitempty"`
	ZentralesOeffnenRechts       *byte   `zusi:"000F" json:"zentrales_oeffnen_rechts,omitempty"`
	StatusZentralesOeffnenLinks  *byte   `zusi:"0010" json:"status_zentrales_oeffnen_links,omitempty"`  // 0: Aus 1: Dauerlicht 2: Blinkend
	StatusZentralesOeffnenRechts *byte   `zusi:"0011" json:"status_zentrales_oeffnen_rechts,omitempty"` // 0: Aus 1: Dauerlicht 2: Blinkend
	LmGruenschleife              *byte   `zusi:"0012" json:"lm_gruenschleife,omitempty"`                // 1: an
}

// 5.3.3.3.6
type StatusFahrzeug struct {
	GrundNullstellungszwang   *uint16 `zusi:"0001" json:"grund_nullstellungszwang,omitempty"` // 0: nichts 1: niedriger HLL druck 2: dynamische Bremse 3: traktionssperre
	GrundTraktionssperre      *uint16 `zusi:"0002" json:"grund_traktionssperre,omitempty"`    // 0: nichts 1: Federspeichenbremse 2: Türsystem 3: Bremsprobe läuft
	StatusFahrschalter        *byte   `zusi:"0003" json:"status_fahrschalter,omitempty"`      // 1: deaktivert 2: normalzustand 3: gestört
	StatusDynamischeBremse    *byte   `zusi:"0004" json:"status_dynamische_bremse,omitempty"` // 1: deaktivert 2: normalzustand
	StatusSander              *byte   `zusi:"0006" json:"status_sander,omitempty"`
	StatusBremsprobe          *byte   `zusi:"0007" json:"status_bremsprobe,omitempty"`
	StellungRichtungsschalter *byte   `zusi:"0008" json:"stellung_richtungsschalter,omitempty"`
}

// 5.3.3.3.7
type StatusZugverband struct {
	StatusFahrzeug []Einzelfahrzeug `zusi:"0001" json:"status_fahrzeug,omitempty"`
}

type Einzelfahrzeug struct {
	Dateiname               *string                      `zusi:"0001" json:"dateiname,omitempty"`
	Beschreibung            *string                      `zusi:"0002" json:"beschreibung,omitempty"`
	Bremsstellung           *uint16                      `zusi:"0003" json:"bremsstellung,omitempty"` // 0: keine/undefiniert 1: G 2: P 3 : P+Mg 4: R 5: R+Mg
	Zugbeeinflussungssystem []EfzZugbeeinflussungssystem `zusi:"0004" json:"zugbeeinflussungssystem,omitempty"`
	Vmax                    *float32                     `zusi:"0005" json:"vmax,omitempty"` // m/s
}

type EfzZugbeeinflussungssystem struct {
	Bezeichnung *string `zusi:"0001" json:"bezeichnung,omitempty"`
}

// 5.3.3.3.8 - Status Weichen
type StatusWeichen struct {
	Weichen []Weiche `zusi:"0001" json:"weichen,omitempty"`
}

type Weiche struct {
	Bezeichnung                *string `zusi:"0001" json:"bezeichnung,omitempty"`
	Bauart                     *int32  `zusi:"0002" json:"bauart,omitempty"`
	Grundstellung              *int32  `zusi:"0003" json:"grundstellung,omitempty"`
	Lage                       *byte   `zusi:"0004" json:"lage,omitempty"`
	Fahrtrichtung              *byte   `zusi:"0005" json:"fahrtrichtung,omitempty"`
	UmlaufmodusStumpfbefahrung *byte   `zusi:"0006" json:"umlaufmodus_stumpfbefahrung,omitempty"`
}

// 5.3.3.4 Befehl 00 0B - DATA_OPERATION (Zusi → Client)

// 5.3.3.5 Befehl 00 0C - DATA_PROG (Zusi → Client)
type DataProg struct {
	Zugdatei          *string `zusi:"0001" json:"zugdatei,omitempty"`
	Zugnummer         *string `zusi:"0002" json:"zugnummer,omitempty"`
	Ladepause         *byte   `zusi:"0003" json:"ladepause,omitempty"`         // 0: Ende Ladepause
	Buchfahrplandatei *string `zusi:"0004" json:"buchfahrplandatei,omitempty"` // serialisertes xml
}

// 5.3.3.6 Befehl 01 0A - INPUT (Client → Zusi)

// 5.3.3.7 Befehl 01 0B - CONTROL (Client → Zusi)
type Control struct {
	Pause         *ShortNode       `zusi:"0001" json:"pause,omitempty"`
	Restart       *ControlFilename `zusi:"0002" json:"restart,omitempty"`
	Start         *ControlFilename `zusi:"0003" json:"start,omitempty"`
	Ende          *EmptyNode       `zusi:"0004" json:"ende,omitempty"`
	FplNeustart   *EmptyNode       `zusi:"0005" json:"fpl_neustart,omitempty"`
	ZugAuswaehlen *Train           `zusi:"0006" json:"zug_auswaehlen,omitempty"` // nach Fahrplan neu starten
	Zeitsprung    *ShortNode       `zusi:"0007" json:"zeitsprung,omitempty"`
	Zeitraffer    *ShortNode       `zusi:"0008" json:"zeitraffer,omitempty"`
	Nebel         *SingleNode      `zusi:"0009" json:"nebel,omitempty"`
	Helligkeit    *SingleNode      `zusi:"000A" json:"helligkeit,omitempty"`
	Reibwert      *SingleNode      `zusi:"000B" json:"reibwert,omitempty"`
	Autopilot     *ShortNode       `zusi:"000C" json:"autopilot,omitempty"`
}

// 5.3.3.8 Befehl 01 0C - GRAPHIC (Client → Zusi)

type ShortNode struct {
	Control *int16 `zusi:"0001" json:"control,omitempty"` // -1: umschalten, 0: aus, 1: ein
}

type SingleNode struct {
	Control *float32 `zusi:"0001" json:"control,omitempty"`
}

type ControlFilename struct {
	Dateiname *string `zusi:"0001" json:"dateiname,omitempty"` // Dateiname des Zuges relativ zum Zusi-Verzeichnis. Wird ein Leerstring übermittelt, startet der zuletzt gefahrene Zug
}

type EmptyNode struct{}

type Train struct {
	Zugnummer *string `zusi:"0001" json:"zugnummer,omitempty"`
}
