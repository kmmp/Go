package lib

//root comment
type root struct {
	Dane DaneParm `xml:"dane"`
}

//DaneParm comment
type DaneParm struct {
	Regon       *string `xml:"Regon"`
	RegonLink   *string `xml:"RegonLink"`
	Nazwa       *string `xml:"Nazwa"`
	Wojewodztwo *string `xml:"Wojewodztwo"`
	Powiat      *string `xml:"Powiat"`
	Gmina       *string `xml:"Gmina"`
	Miejscowosc *string `xml:"Miejscowosc"`
	KodPocztowy *string `xml:"KodPocztowy"`
	Ulica       *string `xml:"Ulica"`
	Typ         *string `xml:"Typ"`
	SilosID     *string `xml:"SilosID"`
}
