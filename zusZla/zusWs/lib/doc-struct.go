package lib

//DokumentSend comment
type DokumentSend struct {
	TrescDokumentu Pozycje `xml:"TrescDokumentu"`
}

//Pozycje comment
type Pozycje struct {
	Ogolne        OgolneParm        `xml:"I"`
	DanePacjenta  DanePacjentaParm  `xml:"II"`
	AdresPacjenta AdresPacjentaParm `xml:"III"`
	Zwolnienie    ZwolnienieParm    `xml:"IV"`
	Platnik       PlatnikParm       `xml:"V"`
	Przychodnia   PrzychodniaParm   `xml:"VI"`
	Lekarz        LekarzParm        `xml:"VII"`
	Podsumowanie  PodsumowanieParm  `xml:"VIII"`
}

//OgolneParm comment
type OgolneParm struct {
	Identyfikator IdentParm `xml:"p1,omitempty"`
	OryginalKopia string    `xml:"p2"`
}

//IdentParm comment
type IdentParm struct {
	Seria string `xml:"p1"`
	Numer string `xml:"p2"`
}

//DanePacjentaParm comment
type DanePacjentaParm struct {
	Pesel         int64  `xml:"p1,omitempty"`
	Imie          string `xml:"p2,omitempty"`
	Nazwisko      string `xml:"p3,omitempty"`
	Instytucja    int64  `xml:"p4"`
	Paszport      string `xml:"p5,omitempty"`
	DataUrodzenia string `xml:"p6,omitempty"`
}

//AdresPacjentaParm comment
type AdresPacjentaParm struct {
	KodPocztowy  string `xml:"p1,omitempty"`
	Miejscowosc  string `xml:"p2,omitempty"`
	Ulica        string `xml:"p3,omitempty"`
	NrDomu       string `xml:"p4"`
	NrMieszkania string `xml:"p5,omitempty"`
}

//ZwolnienieParm comment
type ZwolnienieParm struct {
	Niezdolny     NiezdolnyParm     `xml:"p1,omitempty"`
	Szpital       SzpitalParm       `xml:"p2,omitempty"`
	Wskazanie     string            `xml:"p3"`
	Kod1          string            `xml:"p4,omitempty"`
	Kod2          string            `xml:"p5,omitempty"`
	Kod3          string            `xml:"p6,omitempty"`
	Kod4          string            `xml:"p7,omitempty"`
	Rozpoznanie   string            `xml:"p8,omitempty"`
	Pokrewienstwo PokrewienstwoParm `xml:"p9,omitempty"`
}

//NiezdolnyParm comment
type NiezdolnyParm struct {
	Od string `xml:"p1"`
	Do string `xml:"p2"`
}

//SzpitalParm comment
type SzpitalParm struct {
	SzpitalOd string `xml:"p1"`
	SzpitalDo string `xml:"p2"`
}

//PokrewienstwoParm comment
type PokrewienstwoParm struct {
	KodPokrewienstwa            string `xml:"p1"`
	DataUrodzeniaOsobyPodOpieka string `xml:"p2"`
}

//PlatnikParm comment
type PlatnikParm struct {
	Rodzaj        string `xml:"p1,omitempty"`
	Identyfikator string `xml:"p2,omitempty"`
}

//PrzychodniaParm comment
type PrzychodniaParm struct {
	Nazwa        string `xml:"p1,omitempty"`
	KodPocztowy  string `xml:"p2,omitempty"`
	Miejscowosc  string `xml:"p3,omitempty"`
	Ulica        string `xml:"p4,omitempty"`
	NrDomu       string `xml:"p5"`
	NrMieszkania string `xml:"p6,omitempty"`
}

//LekarzParm comment
type LekarzParm struct {
	NumerPrawa string `xml:"p1"`
	Imie       string `xml:"p2"`
	Nazwisko   string `xml:"p3"`
}

//PodsumowanieParm comment
type PodsumowanieParm struct {
	DataWystawieniaDokumentu string         `xml:"p1"`
	DataElektronizacji       string         `xml:"p2,omitempty"`
	WsteczneWystawienie      string         `xml:"p3,omitempty"`
	Anulowanie               AnulowanieParm `xml:"p4,omitempty"`
	Powiazanie               PowiazanieParm `xml:"p5,omitempty"`
	PobytStajonarny          bool           `xml:"p6,omitempty"`
	InfoDlaPlatnika          bool           `xml:"p7"`
	NipPwdl                  int64          `xml:"p8,omitempty"`
}

//AnulowanieParm comment
type AnulowanieParm struct {
	Seria string `xml:"p1"`
	Numer string `xml:"p2"`
}

//PowiazanieParm comment
type PowiazanieParm struct {
	Seria string `xml:"p1"`
	Numer string `xml:"p2"`
}
