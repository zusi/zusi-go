// Code generated by "stringer -type FuehrerstandId fuehrerstand.go"; DO NOT EDIT.

package fst

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[_keineFunktion-0]
	_ = x[Geschwindigkeit-1]
	_ = x[DruckHauptluftleitung-2]
	_ = x[DruckBremszylinder-3]
	_ = x[DruckHauptluftbehaelter-4]
	_ = x[LuftpresserLaeuft-5]
	_ = x[LuftstromFvb-6]
	_ = x[LuftstromZbv-7]
	_ = x[LuefterAn-8]
	_ = x[ZugkraftGesamt-9]
	_ = x[ZugkraftProAchse-10]
	_ = x[ZugkraftSollGesamt-11]
	_ = x[ZugkraftSollProAchse-12]
	_ = x[Oberstrom-13]
	_ = x[Fahrleitungsspannung-14]
	_ = x[Motordrehzahl-15]
	_ = x[UhrzeitStunde-16]
	_ = x[UhrzeitMinute-17]
	_ = x[UhrzeitSekunde-18]
	_ = x[Hauptschalter-19]
	_ = x[Trennschuetz-20]
	_ = x[Fahrstufe-21]
	_ = x[_3DFenster-22]
	_ = x[AfbSollgeschwindigkeit-23]
	_ = x[DruckHilfsluftbehaelter-24]
	_ = x[ZurueckgelegterGesamtweg-25]
	_ = x[LmGetriebe-26]
	_ = x[LmSchleudern-27]
	_ = x[LmGleiten-28]
	_ = x[LmMgBremse-29]
	_ = x[LmHBremse-30]
	_ = x[LmRBremse-31]
	_ = x[LmHochabbremsung-32]
	_ = x[LmSchnellbremsung-33]
	_ = x[StatusNotbremsung-34]
	_ = x[LmUhrzeitDigital-35]
	_ = x[LmDrehzahlverstellung-36]
	_ = x[LmFahrtrichtungVor-37]
	_ = x[LmFahrtrichtungZurueck-38]
	_ = x[LmFahrtrichtungM-39]
	_ = x[_Hintergrundbild-40]
	_ = x[Motordrehmoment-41]
	_ = x[MotorlastNormiert-42]
	_ = x[Tunnel-43]
	_ = x[SchienenstossWeiche-44]
	_ = x[Stahlbruecke-45]
	_ = x[Steinbruecke-46]
	_ = x[XKoordinate-47]
	_ = x[YKoordinate-48]
	_ = x[ZKoordinate-49]
	_ = x[UTMReferenzpunktXkm-50]
	_ = x[UTMReferenzpunktYkm-51]
	_ = x[UTMZone-52]
	_ = x[UTMZone2-53]
	_ = x[AFBAn-54]
	_ = x[Individuell01-55]
	_ = x[Individuell02-56]
	_ = x[Individuell03-57]
	_ = x[Individuell04-58]
	_ = x[Individuell05-59]
	_ = x[Individuell06-60]
	_ = x[Individuell07-61]
	_ = x[Individuell08-62]
	_ = x[Individuell09-63]
	_ = x[Individuell10-64]
	_ = x[Individuell11-65]
	_ = x[Individuell12-66]
	_ = x[Individuell13-67]
	_ = x[Individuell14-68]
	_ = x[Individuell15-69]
	_ = x[Individuell16-70]
	_ = x[Individuell17-71]
	_ = x[Individuell18-72]
	_ = x[Individuell19-73]
	_ = x[Individuell20-74]
	_ = x[Datum-75]
	_ = x[Gleiskruemmung-76]
	_ = x[Streckenhoechstgeschwindigkeit-77]
	_ = x[ZugkraftvorschlagAutopilot-78]
	_ = x[BeschleunigungX-79]
	_ = x[BeschleunigungY-80]
	_ = x[BeschleunigungZ-81]
	_ = x[DrehbeschleunigungXAchse-82]
	_ = x[DrehbeschleunigungYAchse-83]
	_ = x[DrehbeschleunigungZAchse-84]
	_ = x[Stromabnehmer-85]
	_ = x[LmFederspeicherbremse-86]
	_ = x[ZustandFederspeicherbremse-87]
	_ = x[SteuerwagenLmGetriebe-88]
	_ = x[SteuerwagenLmSchleudern-89]
	_ = x[SteuerwagenLmGleiten-90]
	_ = x[SteuerwagenLmHBremse-91]
	_ = x[SteuerwagenLmRBremse-92]
	_ = x[SteuerwagenLmDrehzahlverstellung-93]
	_ = x[DruckZweitbehaelter-94]
	_ = x[GeschwindigkeitAbsolut-95]
	_ = x[ZugIstEntgleist-96]
	_ = x[KilometrierungZugspitze-97]
	_ = x[Motorstrom-98]
	_ = x[Motorspannung-99]
	_ = x[StatusSifa-100]
	_ = x[StatusZugsicherung-101]
	_ = x[StatusTueren-102]
	_ = x[Individuell21-103]
	_ = x[Individuell22-104]
	_ = x[Individuell23-105]
	_ = x[Individuell24-106]
	_ = x[Individuell25-107]
	_ = x[Individuell26-108]
	_ = x[Individuell27-109]
	_ = x[Individuell28-110]
	_ = x[Individuell29-111]
	_ = x[Individuell30-112]
	_ = x[Individuell31-113]
	_ = x[Individuell32-114]
	_ = x[Individuell33-115]
	_ = x[Individuell34-116]
	_ = x[Individuell35-117]
	_ = x[Individuell36-118]
	_ = x[Individuell37-119]
	_ = x[Individuell38-120]
	_ = x[Individuell39-121]
	_ = x[Individuell40-122]
	_ = x[SteuerwagenLuefterAn-123]
	_ = x[SteuerwagenZugkraftGesamt-124]
	_ = x[SteuerwagenZugraftProAchse-125]
	_ = x[SteuerwagenZugkraftSollGesamt-126]
	_ = x[SteuerwagenZugraftSollProAchse-127]
	_ = x[SteuerwagenOberstrom-128]
	_ = x[SteuerwagenFahrleitungsspannung-129]
	_ = x[SteuerwagenMotordrehzahl-130]
	_ = x[SteuerwagenHauptschalter-131]
	_ = x[SteuerwagenTrennschuetz-132]
	_ = x[SteuerwagenFahrstufe-133]
	_ = x[SteuerwagenMotordrehmoment-134]
	_ = x[SteuerwagenMotorlastNormiert-135]
	_ = x[SteuerwagenStromabnehmer-136]
	_ = x[SteuerwagenMotorstrom-137]
	_ = x[SteuerwagenMotorspannung-138]
	_ = x[GeschwindigkeitAbsolutInklSchleudern-139]
	_ = x[BatteriehauptschalterAus-140]
	_ = x[StatusFahrzeug-141]
	_ = x[StatusZugverband-142]
	_ = x[Bremsprobenfunktion-143]
	_ = x[ZugUndBremsGesamtkraftSollNormiert-144]
	_ = x[SteuerwagenZugUndBremsGesamtkraftSollNormiert-145]
	_ = x[StatusWeichen-146]
	_ = x[ZugUndBremsGesamtkraftAbsolutNormiert-147]
	_ = x[SteuerwagenZugUndBremsGesamtkraftAbsolutNormiert-148]
}

const _FuehrerstandId_name = "_keineFunktionGeschwindigkeitDruckHauptluftleitungDruckBremszylinderDruckHauptluftbehaelterLuftpresserLaeuftLuftstromFvbLuftstromZbvLuefterAnZugkraftGesamtZugkraftProAchseZugkraftSollGesamtZugkraftSollProAchseOberstromFahrleitungsspannungMotordrehzahlUhrzeitStundeUhrzeitMinuteUhrzeitSekundeHauptschalterTrennschuetzFahrstufe_3DFensterAfbSollgeschwindigkeitDruckHilfsluftbehaelterZurueckgelegterGesamtwegLmGetriebeLmSchleudernLmGleitenLmMgBremseLmHBremseLmRBremseLmHochabbremsungLmSchnellbremsungStatusNotbremsungLmUhrzeitDigitalLmDrehzahlverstellungLmFahrtrichtungVorLmFahrtrichtungZurueckLmFahrtrichtungM_HintergrundbildMotordrehmomentMotorlastNormiertTunnelSchienenstossWeicheStahlbrueckeSteinbrueckeXKoordinateYKoordinateZKoordinateUTMReferenzpunktXkmUTMReferenzpunktYkmUTMZoneUTMZone2AFBAnIndividuell01Individuell02Individuell03Individuell04Individuell05Individuell06Individuell07Individuell08Individuell09Individuell10Individuell11Individuell12Individuell13Individuell14Individuell15Individuell16Individuell17Individuell18Individuell19Individuell20DatumGleiskruemmungStreckenhoechstgeschwindigkeitZugkraftvorschlagAutopilotBeschleunigungXBeschleunigungYBeschleunigungZDrehbeschleunigungXAchseDrehbeschleunigungYAchseDrehbeschleunigungZAchseStromabnehmerLmFederspeicherbremseZustandFederspeicherbremseSteuerwagenLmGetriebeSteuerwagenLmSchleudernSteuerwagenLmGleitenSteuerwagenLmHBremseSteuerwagenLmRBremseSteuerwagenLmDrehzahlverstellungDruckZweitbehaelterGeschwindigkeitAbsolutZugIstEntgleistKilometrierungZugspitzeMotorstromMotorspannungStatusSifaStatusZugsicherungStatusTuerenIndividuell21Individuell22Individuell23Individuell24Individuell25Individuell26Individuell27Individuell28Individuell29Individuell30Individuell31Individuell32Individuell33Individuell34Individuell35Individuell36Individuell37Individuell38Individuell39Individuell40SteuerwagenLuefterAnSteuerwagenZugkraftGesamtSteuerwagenZugraftProAchseSteuerwagenZugkraftSollGesamtSteuerwagenZugraftSollProAchseSteuerwagenOberstromSteuerwagenFahrleitungsspannungSteuerwagenMotordrehzahlSteuerwagenHauptschalterSteuerwagenTrennschuetzSteuerwagenFahrstufeSteuerwagenMotordrehmomentSteuerwagenMotorlastNormiertSteuerwagenStromabnehmerSteuerwagenMotorstromSteuerwagenMotorspannungGeschwindigkeitAbsolutInklSchleudernBatteriehauptschalterAusStatusFahrzeugStatusZugverbandBremsprobenfunktionZugUndBremsGesamtkraftSollNormiertSteuerwagenZugUndBremsGesamtkraftSollNormiertStatusWeichenZugUndBremsGesamtkraftAbsolutNormiertSteuerwagenZugUndBremsGesamtkraftAbsolutNormiert"

var _FuehrerstandId_index = [...]uint16{0, 14, 29, 50, 68, 91, 108, 120, 132, 141, 155, 171, 189, 209, 218, 238, 251, 264, 277, 291, 304, 316, 325, 335, 357, 380, 404, 414, 426, 435, 445, 454, 463, 479, 496, 513, 529, 550, 568, 590, 606, 622, 637, 654, 660, 679, 691, 703, 714, 725, 736, 755, 774, 781, 789, 794, 807, 820, 833, 846, 859, 872, 885, 898, 911, 924, 937, 950, 963, 976, 989, 1002, 1015, 1028, 1041, 1054, 1059, 1073, 1103, 1129, 1144, 1159, 1174, 1198, 1222, 1246, 1259, 1280, 1306, 1327, 1350, 1370, 1390, 1410, 1442, 1461, 1483, 1498, 1521, 1531, 1544, 1554, 1572, 1584, 1597, 1610, 1623, 1636, 1649, 1662, 1675, 1688, 1701, 1714, 1727, 1740, 1753, 1766, 1779, 1792, 1805, 1818, 1831, 1844, 1864, 1889, 1915, 1944, 1974, 1994, 2025, 2049, 2073, 2096, 2116, 2142, 2170, 2194, 2215, 2239, 2275, 2299, 2313, 2329, 2348, 2382, 2427, 2440, 2477, 2525}

func (i FuehrerstandId) String() string {
	if i >= FuehrerstandId(len(_FuehrerstandId_index)-1) {
		return "FuehrerstandId(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FuehrerstandId_name[_FuehrerstandId_index[i]:_FuehrerstandId_index[i+1]]
}
