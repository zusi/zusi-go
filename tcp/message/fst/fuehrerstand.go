//go:generate stringer -type FuehrerstandId fuehrerstand.go

package fst

type FuehrerstandId uint16

const (
	_keineFunktion                                   FuehrerstandId = iota // Single
	Geschwindigkeit                                                        // Single: m/s
	DruckHauptluftleitung                                                  // Single: bar
	DruckBremszylinder                                                     // Single: bar
	DruckHauptluftbehaelter                                                // Single: bar
	LuftpresserLaeuft                                                      // Single: aus/an
	LuftstromFvb                                                           // Single: -1..0..+1
	LuftstromZbv                                                           // Single: -1..0..+1
	LuefterAn                                                              // Single: aus/an
	ZugkraftGesamt                                                         // Single: N
	ZugkraftProAchse                                                       // Single: N
	ZugkraftSollGesamt                                                     // Single: N
	ZugkraftSollProAchse                                                   // Single: N
	Oberstrom                                                              // Single: A
	Fahrleitungsspannung                                                   // Single: V
	Motordrehzahl                                                          // Single: 1/min
	UhrzeitStunde                                                          // Single: Stunde (Zeigerposition Analoguhren)
	UhrzeitMinute                                                          // Single: Minute (Zeigerposition Analoguhren)
	UhrzeitSekunde                                                         // Single: Sekunde (Zeigerposition Analoguhren)
	Hauptschalter                                                          // Single: aus/an
	Trennschuetz                                                           // Single: aus/an
	Fahrstufe                                                              // Single: 1
	_3DFenster                                                             // Single: Nicht sinnvoll zu verwenden
	AfbSollgeschwindigkeit                                                 // Single: m/s
	DruckHilfsluftbehaelter                                                // Single: bar
	ZurueckgelegterGesamtweg                                               // Single: m
	LmGetriebe                                                             // Single: aus/an
	LmSchleudern                                                           // Single: aus/an
	LmGleiten                                                              // Single: aus/an
	LmMgBremse                                                             // Single: aus/an
	LmHBremse                                                              // Single: aus/an
	LmRBremse                                                              // Single: aus/an
	LmHochabbremsung                                                       // Single: aus/an
	LmSchnellbremsung                                                      // Single: aus/an
	StatusNotbremsung                                                      // Siehe extra Abschnitt
	LmUhrzeitDigital                                                       // Single: 0...1 (0:00 bis 24:00 Uhr)
	LmDrehzahlverstellung                                                  // Single: aus/an
	LmFahrtrichtungVor                                                     // Single: aus/an
	LmFahrtrichtungZurueck                                                 // Single: aus/an
	LmFahrtrichtungM                                                       // Single: aus/an
	_Hintergrundbild                                                       // Single: Nicht sinnvoll zu verwenden
	Motordrehmoment                                                        // Single: Nm
	MotorlastNormiert                                                      // Single: 1 (0...1)
	Tunnel                                                                 // Single: aus/an
	SchienenstossWeiche                                                    // Single: aus/an
	Stahlbruecke                                                           // Single: aus/an
	Steinbruecke                                                           // Single: aus/an
	XKoordinate                                                            // Single: m (Bez. Strecken-UTM-Punkt)
	YKoordinate                                                            // Single: m (Bez. Strecken-UTM-Punkt)
	ZKoordinate                                                            // Single: m
	UTMReferenzpunktXkm                                                    // Single: km
	UTMReferenzpunktYkm                                                    // Single: km
	UTMZone                                                                // Single
	UTMZone2                                                               // Single
	AFBAn                                                                  // Single: aus/an
	Individuell01                                                          // Single
	Individuell02                                                          // Single
	Individuell03                                                          // Single
	Individuell04                                                          // Single
	Individuell05                                                          // Single
	Individuell06                                                          // Single
	Individuell07                                                          // Single
	Individuell08                                                          // Single
	Individuell09                                                          // Single
	Individuell10                                                          // Single
	Individuell11                                                          // Single
	Individuell12                                                          // Single
	Individuell13                                                          // Single
	Individuell14                                                          // Single
	Individuell15                                                          // Single
	Individuell16                                                          // Single
	Individuell17                                                          // Single
	Individuell18                                                          // Single
	Individuell19                                                          // Single
	Individuell20                                                          // Single
	Datum                                                                  // Single (Tage mit 0 = 30.12.1899)
	Gleiskruemmung                                                         // Single 1000/m
	Streckenhoechstgeschwindigkeit                                         // Single: m/s
	ZugkraftvorschlagAutopilot                                             // Single: N
	BeschleunigungX                                                        // Single: m/s^2
	BeschleunigungY                                                        // Single: m/s^2
	BeschleunigungZ                                                        // Single: m/s^2
	DrehbeschleunigungXAchse                                               // Single: rad/s^2
	DrehbeschleunigungYAchse                                               // Single: rad/s^2
	DrehbeschleunigungZAchse                                               // Single: rad/s^2
	Stromabnehmer                                                          // Single: 2x4 bit (4bit: 1 fuer SA=oben; 4 bit: 1 fuer SA hebt sich gerade)
	LmFederspeicherbremse                                                  // Single: aus/an
	ZustandFederspeicherbremse                                             // Single -1, 0, 1, 2 (nicht vorhanden, aus, an, blinkend)
	SteuerwagenLmGetriebe                                                  // Single: aus/an
	SteuerwagenLmSchleudern                                                // Single: aus/an
	SteuerwagenLmGleiten                                                   // Single: aus/an
	SteuerwagenLmHBremse                                                   // Single: aus/an
	SteuerwagenLmRBremse                                                   // Single: aus/an
	SteuerwagenLmDrehzahlverstellung                                       // Single: aus/an
	DruckZweitbehaelter                                                    // Single: bar
	GeschwindigkeitAbsolut                                                 // Single: m/s
	ZugIstEntgleist                                                        // Single: aus/an
	KilometrierungZugspitze                                                // Single: km
	Motorstrom                                                             // Single: A
	Motorspannung                                                          // Single: V
	StatusSifa                                                             // Siehe eigener Abschnitt
	StatusZugsicherung                                                     // Siehe eigener Abschnitt
	StatusTueren                                                           // Siehe eigener Abschnitt
	Individuell21                                                          // Single
	Individuell22                                                          // Single
	Individuell23                                                          // Single
	Individuell24                                                          // Single
	Individuell25                                                          // Single
	Individuell26                                                          // Single
	Individuell27                                                          // Single
	Individuell28                                                          // Single
	Individuell29                                                          // Single
	Individuell30                                                          // Single
	Individuell31                                                          // Single
	Individuell32                                                          // Single
	Individuell33                                                          // Single
	Individuell34                                                          // Single
	Individuell35                                                          // Single
	Individuell36                                                          // Single
	Individuell37                                                          // Single
	Individuell38                                                          // Single
	Individuell39                                                          // Single
	Individuell40                                                          // Single
	SteuerwagenLuefterAn                                                   // Single: aus/an
	SteuerwagenZugkraftGesamt                                              // Single: N
	SteuerwagenZugraftProAchse                                             // Single: N
	SteuerwagenZugkraftSollGesamt                                          // Single: N
	SteuerwagenZugraftSollProAchse                                         // Single: N
	SteuerwagenOberstrom                                                   // Single: N
	SteuerwagenFahrleitungsspannung                                        // Single V
	SteuerwagenMotordrehzahl                                               // Single 1/min
	SteuerwagenHauptschalter                                               // Single aus/an
	SteuerwagenTrennschuetz                                                // Single: aus/an
	SteuerwagenFahrstufe                                                   // Single: 1
	SteuerwagenMotordrehmoment                                             // Single: Nm
	SteuerwagenMotorlastNormiert                                           // Single: 1 (0...1)
	SteuerwagenStromabnehmer                                               // Single
	SteuerwagenMotorstrom                                                  // Single: A
	SteuerwagenMotorspannung                                               // Single: V
	GeschwindigkeitAbsolutInklSchleudern                                   // Single: m/s
	BatteriehauptschalterAus                                               // Single: aus/an
	StatusFahrzeug                                                         // s. eigener Abschnitt
	StatusZugverband                                                       // s. eigener Abschnitt
	Bremsprobenfunktion                                                    // Single: 0 (aus, >0 aktiv)
	ZugUndBremsGesamtkraftSollNormiert                                     // Single: 1 ((0...1) normiert auf aktuelle Fmax
	SteuerwagenZugUndBremsGesamtkraftSollNormiert                          // Single: 1 ((0...1) normiert auf aktuelle Fmax
	StatusWeichen                                                          // s. eigener Abschnitt
	ZugUndBremsGesamtkraftAbsolutNormiert                                  // Single: 1 ((0...1) normiert auf aktuelle Fmax
	SteuerwagenZugUndBremsGesamtkraftAbsolutNormiert                       // Single: 1 ((0...1) normiert auf aktuelle Fmax
)
