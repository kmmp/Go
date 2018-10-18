package zus_channel_zla_binder

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://zus.pl/b2b/zus/channel/gabinetowe"

// NewZla_PortType creates an initializes a Zla_PortType.
func NewZla_PortType(cli *soap.Client) Zla_PortType {
	return &zla_PortType{cli}
}

// Zla_PortType was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Zla_PortType interface {
	// NadajSeriaNumerZla was auto-generated from WSDL.
	NadajSeriaNumerZla(parameters *NadajSeriaNumerZla) (*NadajSeriaNumerZlaResponse, error)

	// NotyfikujAnulowanieZla was auto-generated from WSDL.
	NotyfikujAnulowanieZla(parameters *NotyfikujAnulowanieZla) (*NotyfikujAnulowanieZlaResponse, error)

	// PobierzAdresyPlacowek was auto-generated from WSDL.
	PobierzAdresyPlacowek(parameters *PobierzAdresyPlacowek) (*PobierzAdresyPlacowekResponse, error)

	// PobierzCzlonkowRodzinyUbezpieczonego was auto-generated from
	// WSDL.
	PobierzCzlonkowRodzinyUbezpieczonego(parameters *PobierzCzlonkowRodzinyUbezpieczonego) (*PobierzCzlonkowRodzinyUbezpieczonegoResponse, error)

	// PobierzDaneLekarza was auto-generated from WSDL.
	PobierzDaneLekarza(parameters *PobierzDaneLekarza) (*PobierzDaneLekarzaResponse, error)

	// PobierzDanePlatnika was auto-generated from WSDL.
	PobierzDanePlatnika(parameters *PobierzDanePlatnika) (*PobierzDanePlatnikaResponse, error)

	// PobierzDaneUbezpieczonego was auto-generated from WSDL.
	PobierzDaneUbezpieczonego(parameters *PobierzDaneUbezpieczonego) (*PobierzDaneUbezpieczonegoResponse, error)

	// PobierzDokument was auto-generated from WSDL.
	PobierzDokument(parameters *PobierzDokument) (*PobierzDokumentResponse, error)

	// PobierzIdentyfikatorDokumentu was auto-generated from WSDL.
	PobierzIdentyfikatorDokumentu(parameters *PobierzIdentyfikatorDokumentu) (*PobierzIdentyfikatorDokumentuResponse, error)

	// PobierzKodChoroby was auto-generated from WSDL.
	PobierzKodChoroby(parameters *PobierzKodChoroby) (*PobierzKodChorobyResponse, error)

	// PobierzListeBiezacychZlaLekarza was auto-generated from WSDL.
	PobierzListeBiezacychZlaLekarza(parameters *PobierzListeBiezacychZlaLekarza) (*PobierzListeBiezacychZlaLekarzaResponse, error)

	// PobierzListePowiazanychZla was auto-generated from WSDL.
	PobierzListePowiazanychZla(parameters *PobierzListePowiazanychZla) (*PobierzListePowiazanychZlaResponse, error)

	// PobierzListePr4Ubezpieczonego was auto-generated from WSDL.
	PobierzListePr4Ubezpieczonego(parameters *PobierzListePr4Ubezpieczonego) (*PobierzListePr4UbezpieczonegoResponse, error)

	// PobierzListePrzetworzonychZlaLekarza was auto-generated from
	// WSDL.
	PobierzListePrzetworzonychZlaLekarza(parameters *PobierzListePrzetworzonychZlaLekarza) (*PobierzListePrzetworzonychZlaLekarzaResponse, error)

	// PobierzListeZlaUbezpieczonego was auto-generated from WSDL.
	PobierzListeZlaUbezpieczonego(parameters *PobierzListeZlaUbezpieczonego) (*PobierzListeZlaUbezpieczonegoResponse, error)

	// PobierzLiterowyKodChoroby was auto-generated from WSDL.
	PobierzLiterowyKodChoroby(parameters *PobierzLiterowyKodChoroby) (*PobierzLiterowyKodChorobyResponse, error)

	// PobierzMiejsceWykonywaniaZawodu was auto-generated from WSDL.
	PobierzMiejsceWykonywaniaZawodu(parameters *PobierzMiejsceWykonywaniaZawodu) (*PobierzMiejsceWykonywaniaZawoduResponse, error)

	// PobierzOpisChoroby was auto-generated from WSDL.
	PobierzOpisChoroby(parameters *PobierzOpisChoroby) (*PobierzOpisChorobyResponse, error)

	// PobierzOswiadczenie was auto-generated from WSDL.
	PobierzOswiadczenie(parameters *PobierzOswiadczenie) (*PobierzOswiadczenieResponse, error)

	// PobierzPlatnikowUbezpieczonego was auto-generated from WSDL.
	PobierzPlatnikowUbezpieczonego(parameters *PobierzPlatnikowUbezpieczonego) (*PobierzPlatnikowUbezpieczonegoResponse, error)

	// PobierzSlownikKodowPokrewienstwa was auto-generated from WSDL.
	PobierzSlownikKodowPokrewienstwa(parameters *PobierzSlownikKodowPokrewienstwa) (*PobierzSlownikKodowPokrewienstwaResponse, error)

	// PobierzSlownikPrzyczynAnulowania was auto-generated from WSDL.
	PobierzSlownikPrzyczynAnulowania(parameters *PobierzSlownikPrzyczynAnulowania) (*PobierzSlownikPrzyczynAnulowaniaResponse, error)

	// PobierzSlownikPrzyczynUniewaznienia was auto-generated from
	// WSDL.
	PobierzSlownikPrzyczynUniewaznienia(parameters *PobierzSlownikPrzyczynUniewaznienia) (*PobierzSlownikPrzyczynUniewaznieniaResponse, error)

	// PobierzStatusyDrukowZla was auto-generated from WSDL.
	PobierzStatusyDrukowZla(parameters *PobierzStatusyDrukowZla) (*PobierzStatusyDrukowZlaResponse, error)

	// PobierzSzczegolyZlaBiezace was auto-generated from WSDL.
	PobierzSzczegolyZlaBiezace(parameters *PobierzSzczegolyZlaBiezace) (*PobierzSzczegolyZlaBiezaceResponse, error)

	// PobierzSzczegolyZlaPrzetworzone was auto-generated from WSDL.
	PobierzSzczegolyZlaPrzetworzone(parameters *PobierzSzczegolyZlaPrzetworzone) (*PobierzSzczegolyZlaPrzetworzoneResponse, error)

	// PobierzUppDlaDokumentu was auto-generated from WSDL.
	PobierzUppDlaDokumentu(parameters *PobierzUppDlaDokumentu) (*PobierzUppDlaDokumentuResponse, error)

	// PobierzUprawnieniaNaDzien was auto-generated from WSDL.
	PobierzUprawnieniaNaDzien(parameters *PobierzUprawnieniaNaDzien) (*PobierzUprawnieniaNaDzienResponse, error)

	// RezerwujSeriaNumerZla was auto-generated from WSDL.
	RezerwujSeriaNumerZla(parameters *RezerwujSeriaNumerZla) (*RezerwujSeriaNumerZlaResponse, error)

	// SprawdzMozliwoscAnulowania was auto-generated from WSDL.
	SprawdzMozliwoscAnulowania(parameters *SprawdzMozliwoscAnulowania) (*SprawdzMozliwoscAnulowaniaResponse, error)

	// SprawdzMozliwoscElektronizacji was auto-generated from WSDL.
	SprawdzMozliwoscElektronizacji(parameters *SprawdzMozliwoscElektronizacji) (*SprawdzMozliwoscElektronizacjiResponse, error)

	// SprawdzMozliwoscUniewaznienia was auto-generated from WSDL.
	SprawdzMozliwoscUniewaznienia(parameters *SprawdzMozliwoscUniewaznienia) (*SprawdzMozliwoscUniewaznieniaResponse, error)

	// SprawdzProfilRehabilitacji was auto-generated from WSDL.
	SprawdzProfilRehabilitacji(parameters *SprawdzProfilRehabilitacji) (*SprawdzProfilRehabilitacjiResponse, error)

	// UsunSesje was auto-generated from WSDL.
	UsunSesje(parameters *UsunSesje) (*UsunSesjeResponse, error)

	// WalidujDokumenty was auto-generated from WSDL.
	WalidujDokumenty(parameters *WalidujDokumenty) (*WalidujDokumentyResponse, error)

	// WalidujWniosek was auto-generated from WSDL.
	WalidujWniosek(parameters *WalidujWniosek) (*WalidujWniosekResponse, error)

	// WyslijDokumenty was auto-generated from WSDL.
	WyslijDokumenty(parameters *WyslijDokumenty) (*WyslijDokumentyResponse, error)

	// WyslijWniosek was auto-generated from WSDL.
	WyslijWniosek(parameters *WyslijWniosek) (*WyslijWniosekResponse, error)

	// ZalogujPodpisem was auto-generated from WSDL.
	ZalogujPodpisem(parameters *ZalogujPodpisem) (*ZalogujPodpisemResponse, error)
}

// Date in WSDL format.
type Date string

// Data was auto-generated from WSDL.
type Data string

// DrukZlaKolumnyEnumeracja was auto-generated from WSDL.
type DrukZlaKolumnyEnumeracja string

// Validate validates DrukZlaKolumnyEnumeracja.
func (v DrukZlaKolumnyEnumeracja) Validate() bool {
	for _, vv := range []string{
		"Status",
		"SeriaNumer",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ID was auto-generated from WSDL.
type ID string

// Inne was auto-generated from WSDL.
type Inne string

// KEDU was auto-generated from WSDL.
type KEDU string

// KodChorobyEnumeracja was auto-generated from WSDL.
type KodChorobyEnumeracja string

// Validate validates KodChorobyEnumeracja.
func (v KodChorobyEnumeracja) Validate() bool {
	for _, vv := range []string{
		"StatystycznyKodChoroby",
		"NazwaChoroby",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// KodyPokrewienstwaEnumeracja was auto-generated from WSDL.
type KodyPokrewienstwaEnumeracja string

// Validate validates KodyPokrewienstwaEnumeracja.
func (v KodyPokrewienstwaEnumeracja) Validate() bool {
	for _, vv := range []string{
		"1",
		"2",
		"3",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// KontekstDostepu was auto-generated from WSDL.
type KontekstDostepu string

// Validate validates KontekstDostepu.
func (v KontekstDostepu) Validate() bool {
	for _, vv := range []string{
		"WystawienieZla",
		"WyszukanieZla",
		"SzczegolyAnulowanie",
		"SzczegolyNowe",
		"SzczegolyPlatnik",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// MetodaUwierzytelnienia was auto-generated from WSDL.
type MetodaUwierzytelnienia string

// Validate validates MetodaUwierzytelnienia.
func (v MetodaUwierzytelnienia) Validate() bool {
	for _, vv := range []string{
		"certyfikat",
		"ePuap",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// NIP was auto-generated from WSDL.
type NIP string

// NumerZla was auto-generated from WSDL.
type NumerZla string

// OperatorAndOrEnum was auto-generated from WSDL.
type OperatorAndOrEnum string

// Validate validates OperatorAndOrEnum.
func (v OperatorAndOrEnum) Validate() bool {
	for _, vv := range []string{
		"AND",
		"OR",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// OperatorEnum was auto-generated from WSDL.
type OperatorEnum string

// Validate validates OperatorEnum.
func (v OperatorEnum) Validate() bool {
	for _, vv := range []string{
		"equalTo",
		"isNotEmpty",
		"lessThan",
		"lessThanOrEqualTo",
		"largerThan",
		"largerThanOrEqualTo",
		"contains",
		"startsWith",
		"isEmpty",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// Pesel was auto-generated from WSDL.
type Pesel string

// Pesel2 was auto-generated from WSDL.
type Pesel2 string

// RezultatWalidacjiEnumeracja was auto-generated from WSDL.
type RezultatWalidacjiEnumeracja string

// Validate validates RezultatWalidacjiEnumeracja.
func (v RezultatWalidacjiEnumeracja) Validate() bool {
	for _, vv := range []string{
		"POZYTYWNY",
		"NEGATYWNY",
		"OSTRZEZENIE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// RodzajBleduWalidacjiEnumeracja was auto-generated from WSDL.
type RodzajBleduWalidacjiEnumeracja string

// Validate validates RodzajBleduWalidacjiEnumeracja.
func (v RodzajBleduWalidacjiEnumeracja) Validate() bool {
	for _, vv := range []string{
		"OSTRZEZENIE",
		"BLAD",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// RodzajDokumentuEnumeracja was auto-generated from WSDL.
type RodzajDokumentuEnumeracja string

// Validate validates RodzajDokumentuEnumeracja.
func (v RodzajDokumentuEnumeracja) Validate() bool {
	for _, vv := range []string{
		"ZLA",
		"KOPIA_ZLA",
		"AZLA",
		"UZLA",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// RodzajWnioskuEnumeracja was auto-generated from WSDL.
type RodzajWnioskuEnumeracja string

// Validate validates RodzajWnioskuEnumeracja.
func (v RodzajWnioskuEnumeracja) Validate() bool {
	for _, vv := range []string{
		"FZLA",
		"PR_4",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// SeriaZla was auto-generated from WSDL.
type SeriaZla string

// SortowanieEnumeracja was auto-generated from WSDL.
type SortowanieEnumeracja string

// Validate validates SortowanieEnumeracja.
func (v SortowanieEnumeracja) Validate() bool {
	for _, vv := range []string{
		"ASC",
		"DESC",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StatusDrukuEnumeracja was auto-generated from WSDL.
type StatusDrukuEnumeracja string

// Validate validates StatusDrukuEnumeracja.
func (v StatusDrukuEnumeracja) Validate() bool {
	for _, vv := range []string{
		"UNIEWAZNIONY",
		"ELEKTRONIZOWANY",
		"WYDRUKOWANY",
		"ZWERYFIKOWANY",
		"WYKORZYSTANY",
		"ROBOCZY",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StatystycznyKodChoroby was auto-generated from WSDL.
type StatystycznyKodChoroby string

// TrybWystawianiaEnumeracja was auto-generated from WSDL.
type TrybWystawianiaEnumeracja string

// Validate validates TrybWystawianiaEnumeracja.
func (v TrybWystawianiaEnumeracja) Validate() bool {
	for _, vv := range []string{
		"Biezacy",
		"Alternatywny",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// TypIdUbezpieczonegoEnumeracja was auto-generated from WSDL.
type TypIdUbezpieczonegoEnumeracja string

// Validate validates TypIdUbezpieczonegoEnumeracja.
func (v TypIdUbezpieczonegoEnumeracja) Validate() bool {
	for _, vv := range []string{
		"PESEL",
		"SeriaNumerPaszportu",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// UPP was auto-generated from WSDL.
type UPP string

// UzasadnienieZaswiadczeniaWstecznego was auto-generated from
// WSDL.
type UzasadnienieZaswiadczeniaWstecznego string

// ZaswiadczenieLekarskieKsiKolumnyEnumeracja was auto-generated
// from WSDL.
type ZaswiadczenieLekarskieKsiKolumnyEnumeracja string

// Validate validates ZaswiadczenieLekarskieKsiKolumnyEnumeracja.
func (v ZaswiadczenieLekarskieKsiKolumnyEnumeracja) Validate() bool {
	for _, vv := range []string{
		"SeriaZla",
		"NumerZla",
		"PeselUbezpieczonego",
		"NumerDokumentuUbezpieczonego",
		"ImieUbezpieczonego",
		"NazwiskoUbezpieczonego",
		"NiezdolnoscDoPracyOd",
		"NiezdolnoscDoPracyDo",
		"LiczbaDniPobytuWSzpitalu",
		"KodPokrewienstwa",
		"DataUrodzeniaOsobySpokrewnionej",
		"StatystycznyNumerChoroby",
		"ImieLekarza",
		"NazwiskoLekarza",
		"IdentyfikatorLekarza",
		"DataWystawieniaZla",
		"Status",
		"IdentyfikatorPlatnika",
		"TypIdentyfikatoraPlatnika",
		"PobytWSzpitaluOd",
		"PobytWSzpitaluDo",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ZaswiadczenieLekarskiePueKolumnyEnumeracja was auto-generated
// from WSDL.
type ZaswiadczenieLekarskiePueKolumnyEnumeracja string

// Validate validates ZaswiadczenieLekarskiePueKolumnyEnumeracja.
func (v ZaswiadczenieLekarskiePueKolumnyEnumeracja) Validate() bool {
	for _, vv := range []string{
		"IdentyfikatorUbezpieczonego",
		"TypIdentyfikatoraUbezpieczonego",
		"ImieUbezpieczonego",
		"NazwiskoUbezpieczonego",
		"NiezdolnoscDoPracyOd",
		"NiezdolnoscDoPracyDo",
		"LiczbaDniPobytuWSzpitalu",
		"KodPokrewienstwa",
		"DataUrodzeniaOsobySpokrewnionej",
		"IdentyfikatorPlatnika",
		"TypIdentyfikatoraPlatnika",
		"StatystycznyNumerChoroby",
		"NpwzLekarza",
		"ImieLekarza",
		"NazwiskoLekarza",
		"Status",
		"DataWystawieniaZla",
		"SeriaNumerZla",
		"PobytWSzpitaluOd",
		"PobytWSzpitaluDo",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// DocTypeRef_AdresyPlacowek was auto-generated from WSDL.
type DocTypeRef_AdresyPlacowek struct {
	KodPocztowy *string `xml:"KodPocztowy,omitempty" json:"KodPocztowy,omitempty" yaml:"KodPocztowy,omitempty"`
	Miejscowosc *string `xml:"Miejscowosc,omitempty" json:"Miejscowosc,omitempty" yaml:"Miejscowosc,omitempty"`
	Ulica       *string `xml:"Ulica,omitempty" json:"Ulica,omitempty" yaml:"Ulica,omitempty"`
	NumerDomu   *string `xml:"NumerDomu,omitempty" json:"NumerDomu,omitempty" yaml:"NumerDomu,omitempty"`
	NumerLokalu *string `xml:"NumerLokalu,omitempty" json:"NumerLokalu,omitempty" yaml:"NumerLokalu,omitempty"`
}

// DocTypeRef_BladWalidacji was auto-generated from WSDL.
type DocTypeRef_BladWalidacji struct {
	NrRef       []*string                       `xml:"NrRef,omitempty" json:"NrRef,omitempty" yaml:"NrRef,omitempty"`
	Rodzaj      *RodzajBleduWalidacjiEnumeracja `xml:"Rodzaj,omitempty" json:"Rodzaj,omitempty" yaml:"Rodzaj,omitempty"`
	KodBledu    *string                         `xml:"KodBledu,omitempty" json:"KodBledu,omitempty" yaml:"KodBledu,omitempty"`
	OpisBledu   *string                         `xml:"OpisBledu,omitempty" json:"OpisBledu,omitempty" yaml:"OpisBledu,omitempty"`
	Lokalizacja *string                         `xml:"Lokalizacja,omitempty" json:"Lokalizacja,omitempty" yaml:"Lokalizacja,omitempty"`
}

// DocTypeRef_CzlonekRodzinyUbezpieczonego was auto-generated from
// WSDL.
type DocTypeRef_CzlonekRodzinyUbezpieczonego struct {
	Imie          *string `xml:"Imie,omitempty" json:"Imie,omitempty" yaml:"Imie,omitempty"`
	Nazwisko      *string `xml:"Nazwisko,omitempty" json:"Nazwisko,omitempty" yaml:"Nazwisko,omitempty"`
	DataUrodzenia *Date   `xml:"DataUrodzenia,omitempty" json:"DataUrodzenia,omitempty" yaml:"DataUrodzenia,omitempty"`
}

// DocTypeRef_Dokument was auto-generated from WSDL.
type DocTypeRef_Dokument struct {
	KEDU  *DocTypeRef_KEDU `xml:"KEDU,omitempty" json:"KEDU,omitempty" yaml:"KEDU,omitempty"`
	NrRef Inne             `xml:"NrRef,attr,omitempty" json:"NrRef,attr,omitempty" yaml:"NrRef,attr,omitempty"`
}

// DocTypeRef_DrukZlaFiltrowanie was auto-generated from WSDL.
type DocTypeRef_DrukZlaFiltrowanie struct {
	WarunekFiltrowania        []*DocTypeRef_DrukZlaFiltrowanieWarunek `xml:"WarunekFiltrowania,omitempty" json:"WarunekFiltrowania,omitempty" yaml:"WarunekFiltrowania,omitempty"`
	LogiczneZlaczenieWarunkow *OperatorAndOrEnum                      `xml:"LogiczneZlaczenieWarunkow,omitempty" json:"LogiczneZlaczenieWarunkow,omitempty" yaml:"LogiczneZlaczenieWarunkow,omitempty"`
}

// DocTypeRef_DrukZlaFiltrowanieWarunek was auto-generated from
// WSDL.
type DocTypeRef_DrukZlaFiltrowanieWarunek struct {
	Kolumna  *DrukZlaKolumnyEnumeracja `xml:"Kolumna,omitempty" json:"Kolumna,omitempty" yaml:"Kolumna,omitempty"`
	Operator *OperatorEnum             `xml:"Operator,omitempty" json:"Operator,omitempty" yaml:"Operator,omitempty"`
	Wartosc  *string                   `xml:"Wartosc,omitempty" json:"Wartosc,omitempty" yaml:"Wartosc,omitempty"`
}

// DocTypeRef_DrukZlaSortowanie was auto-generated from WSDL.
type DocTypeRef_DrukZlaSortowanie struct {
	WarunekSortowania []*DocTypeRef_DrukZlaSortowanieWarunek `xml:"WarunekSortowania,omitempty" json:"WarunekSortowania,omitempty" yaml:"WarunekSortowania,omitempty"`
}

// DocTypeRef_DrukZlaSortowanieWarunek was auto-generated from
// WSDL.
type DocTypeRef_DrukZlaSortowanieWarunek struct {
	Kolumna  *DrukZlaKolumnyEnumeracja `xml:"Kolumna,omitempty" json:"Kolumna,omitempty" yaml:"Kolumna,omitempty"`
	Kierunek *SortowanieEnumeracja     `xml:"Kierunek,omitempty" json:"Kierunek,omitempty" yaml:"Kierunek,omitempty"`
}

// DocTypeRef_IdentyfikacjaPlatnika was auto-generated from WSDL.
type DocTypeRef_IdentyfikacjaPlatnika struct {
	Pesel               *Pesel `xml:"Pesel,omitempty" json:"Pesel,omitempty" yaml:"Pesel,omitempty"`
	SeriaNumerPaszportu *ID    `xml:"SeriaNumerPaszportu,omitempty" json:"SeriaNumerPaszportu,omitempty" yaml:"SeriaNumerPaszportu,omitempty"`
	Nip                 *NIP   `xml:"Nip,omitempty" json:"Nip,omitempty" yaml:"Nip,omitempty"`
}

// DocTypeRef_IdentyfikacjaUbezpieczonego was auto-generated from
// WSDL.
type DocTypeRef_IdentyfikacjaUbezpieczonego struct {
	Pesel               *Pesel `xml:"Pesel,omitempty" json:"Pesel,omitempty" yaml:"Pesel,omitempty"`
	SeriaNumerPaszportu *ID    `xml:"SeriaNumerPaszportu,omitempty" json:"SeriaNumerPaszportu,omitempty" yaml:"SeriaNumerPaszportu,omitempty"`
}

// DocTypeRef_KEDU was auto-generated from WSDL.
type DocTypeRef_KEDU struct {
	KEDU *KEDU `xml:"KEDU,omitempty" json:"KEDU,omitempty" yaml:"KEDU,omitempty"`
}

// DocTypeRef_KodChorobyFiltrowanie was auto-generated from WSDL.
type DocTypeRef_KodChorobyFiltrowanie struct {
	WarunekFiltrowania        []*DocTypeRef_KodChorobyFiltrowanieWarunek `xml:"WarunekFiltrowania,omitempty" json:"WarunekFiltrowania,omitempty" yaml:"WarunekFiltrowania,omitempty"`
	LogiczneZlaczenieWarunkow *OperatorAndOrEnum                         `xml:"LogiczneZlaczenieWarunkow,omitempty" json:"LogiczneZlaczenieWarunkow,omitempty" yaml:"LogiczneZlaczenieWarunkow,omitempty"`
}

// DocTypeRef_KodChorobyFiltrowanieWarunek was auto-generated from
// WSDL.
type DocTypeRef_KodChorobyFiltrowanieWarunek struct {
	Kolumna  *KodChorobyEnumeracja `xml:"Kolumna,omitempty" json:"Kolumna,omitempty" yaml:"Kolumna,omitempty"`
	Operator *OperatorEnum         `xml:"Operator,omitempty" json:"Operator,omitempty" yaml:"Operator,omitempty"`
	Wartosc  *string               `xml:"Wartosc,omitempty" json:"Wartosc,omitempty" yaml:"Wartosc,omitempty"`
}

// DocTypeRef_KodChorobySortowanie was auto-generated from WSDL.
type DocTypeRef_KodChorobySortowanie struct {
	WarunekSortowania []*DocTypeRef_KodChorobySortowanieWarunek `xml:"WarunekSortowania,omitempty" json:"WarunekSortowania,omitempty" yaml:"WarunekSortowania,omitempty"`
}

// DocTypeRef_KodChorobySortowanieWarunek was auto-generated from
// WSDL.
type DocTypeRef_KodChorobySortowanieWarunek struct {
	Kolumna  *KodChorobyEnumeracja `xml:"Kolumna,omitempty" json:"Kolumna,omitempty" yaml:"Kolumna,omitempty"`
	Kierunek *SortowanieEnumeracja `xml:"Kierunek,omitempty" json:"Kierunek,omitempty" yaml:"Kierunek,omitempty"`
}

// DocTypeRef_KodPokrewienstwa was auto-generated from WSDL.
type DocTypeRef_KodPokrewienstwa struct {
	Kod  *string `xml:"Kod,omitempty" json:"Kod,omitempty" yaml:"Kod,omitempty"`
	Opis *string `xml:"Opis,omitempty" json:"Opis,omitempty" yaml:"Opis,omitempty"`
}

// DocTypeRef_ListaDokumentow was auto-generated from WSDL.
type DocTypeRef_ListaDokumentow struct {
	Dokument []*DocTypeRef_Dokument `xml:"Dokument,omitempty" json:"Dokument,omitempty" yaml:"Dokument,omitempty"`
}

// DocTypeRef_NotyfikujAnulowanieZlaIn was auto-generated from
// WSDL.
type DocTypeRef_NotyfikujAnulowanieZlaIn struct {
	SeriaNumerZla            []*DocTypeRef_SeriaNumerZla `xml:"SeriaNumerZla,omitempty" json:"SeriaNumerZla,omitempty" yaml:"SeriaNumerZla,omitempty"`
	MiejsceWykonywaniaZawodu *string                     `xml:"MiejsceWykonywaniaZawodu,omitempty" json:"MiejsceWykonywaniaZawodu,omitempty" yaml:"MiejsceWykonywaniaZawodu,omitempty"`
	DodatkoweInformacje      *string                     `xml:"DodatkoweInformacje,omitempty" json:"DodatkoweInformacje,omitempty" yaml:"DodatkoweInformacje,omitempty"`
}

// DocTypeRef_NumerStatystycznyChoroby was auto-generated from
// WSDL.
type DocTypeRef_NumerStatystycznyChoroby struct {
	StatystycznyKodChoroby *string `xml:"StatystycznyKodChoroby,omitempty" json:"StatystycznyKodChoroby,omitempty" yaml:"StatystycznyKodChoroby,omitempty"`
	Nazwa                  *string `xml:"Nazwa,omitempty" json:"Nazwa,omitempty" yaml:"Nazwa,omitempty"`
	Opis                   *string `xml:"Opis,omitempty" json:"Opis,omitempty" yaml:"Opis,omitempty"`
}

// DocTypeRef_ParaDokumentowZla was auto-generated from WSDL.
type DocTypeRef_ParaDokumentowZla struct {
	Oryginal *DocTypeRef_Dokument `xml:"Oryginal,omitempty" json:"Oryginal,omitempty" yaml:"Oryginal,omitempty"`
	Kopia    *DocTypeRef_Dokument `xml:"Kopia,omitempty" json:"Kopia,omitempty" yaml:"Kopia,omitempty"`
	NrRef    Inne                 `xml:"NrRef,attr,omitempty" json:"NrRef,attr,omitempty" yaml:"NrRef,attr,omitempty"`
}

// DocTypeRef_Platnik was auto-generated from WSDL.
type DocTypeRef_Platnik struct {
	PlatnikIstnieje *bool   `xml:"PlatnikIstnieje,omitempty" json:"PlatnikIstnieje,omitempty" yaml:"PlatnikIstnieje,omitempty"`
	Nazwa           *string `xml:"Nazwa,omitempty" json:"Nazwa,omitempty" yaml:"Nazwa,omitempty"`
	Imie            *string `xml:"Imie,omitempty" json:"Imie,omitempty" yaml:"Imie,omitempty"`
	Nazwisko        *string `xml:"Nazwisko,omitempty" json:"Nazwisko,omitempty" yaml:"Nazwisko,omitempty"`
}

// DocTypeRef_Platnik_Ubezpieczonego was auto-generated from WSDL.
type DocTypeRef_Platnik_Ubezpieczonego struct {
	PlatnikIstnieje     *bool   `xml:"PlatnikIstnieje,omitempty" json:"PlatnikIstnieje,omitempty" yaml:"PlatnikIstnieje,omitempty"`
	Nazwa               *string `xml:"Nazwa,omitempty" json:"Nazwa,omitempty" yaml:"Nazwa,omitempty"`
	Imie                *string `xml:"Imie,omitempty" json:"Imie,omitempty" yaml:"Imie,omitempty"`
	Nazwisko            *string `xml:"Nazwisko,omitempty" json:"Nazwisko,omitempty" yaml:"Nazwisko,omitempty"`
	NIP                 *string `xml:"NIP,omitempty" json:"NIP,omitempty" yaml:"NIP,omitempty"`
	Pesel               *Pesel2 `xml:"Pesel,omitempty" json:"Pesel,omitempty" yaml:"Pesel,omitempty"`
	SeriaNumerPaszportu *string `xml:"SeriaNumerPaszportu,omitempty" json:"SeriaNumerPaszportu,omitempty" yaml:"SeriaNumerPaszportu,omitempty"`
	ProfilPUE           *bool   `xml:"ProfilPUE,omitempty" json:"ProfilPUE,omitempty" yaml:"ProfilPUE,omitempty"`
}

// DocTypeRef_PowiazaneZla was auto-generated from WSDL.
type DocTypeRef_PowiazaneZla struct {
	Seria                     *string `xml:"Seria,omitempty" json:"Seria,omitempty" yaml:"Seria,omitempty"`
	Numer                     *string `xml:"Numer,omitempty" json:"Numer,omitempty" yaml:"Numer,omitempty"`
	StatusPrzetwarzania       *string `xml:"StatusPrzetwarzania,omitempty" json:"StatusPrzetwarzania,omitempty" yaml:"StatusPrzetwarzania,omitempty"`
	DataWystawienia           *string `xml:"DataWystawienia,omitempty" json:"DataWystawienia,omitempty" yaml:"DataWystawienia,omitempty"`
	IdentyfikatorPlatnika     *string `xml:"IdentyfikatorPlatnika,omitempty" json:"IdentyfikatorPlatnika,omitempty" yaml:"IdentyfikatorPlatnika,omitempty"`
	TypIdentyfikatoraPlatnika *string `xml:"TypIdentyfikatoraPlatnika,omitempty" json:"TypIdentyfikatoraPlatnika,omitempty" yaml:"TypIdentyfikatoraPlatnika,omitempty"`
}

// DocTypeRef_Przyczyny was auto-generated from WSDL.
type DocTypeRef_Przyczyny struct {
	Kod   *string `xml:"Kod,omitempty" json:"Kod,omitempty" yaml:"Kod,omitempty"`
	Nazwa *string `xml:"Nazwa,omitempty" json:"Nazwa,omitempty" yaml:"Nazwa,omitempty"`
	Opis  *string `xml:"Opis,omitempty" json:"Opis,omitempty" yaml:"Opis,omitempty"`
}

// DocTypeRef_Rezultat was auto-generated from WSDL.
type DocTypeRef_Rezultat struct {
	KodBledu  *string `xml:"KodBledu,omitempty" json:"KodBledu,omitempty" yaml:"KodBledu,omitempty"`
	OpisBledu *string `xml:"OpisBledu,omitempty" json:"OpisBledu,omitempty" yaml:"OpisBledu,omitempty"`
}

// DocTypeRef_RezultatWalidacji was auto-generated from WSDL.
type DocTypeRef_RezultatWalidacji struct {
	Rezultat      *RezultatWalidacjiEnumeracja `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	BladWalidacji []*DocTypeRef_BladWalidacji  `xml:"BladWalidacji,omitempty" json:"BladWalidacji,omitempty" yaml:"BladWalidacji,omitempty"`
}

// DocTypeRef_SeriaNumerZla was auto-generated from WSDL.
type DocTypeRef_SeriaNumerZla struct {
	Seria *SeriaZla `xml:"Seria,omitempty" json:"Seria,omitempty" yaml:"Seria,omitempty"`
	Numer *NumerZla `xml:"Numer,omitempty" json:"Numer,omitempty" yaml:"Numer,omitempty"`
}

// DocTypeRef_StatusDruku was auto-generated from WSDL.
type DocTypeRef_StatusDruku struct {
	SeriaNumer *string                `xml:"SeriaNumer,omitempty" json:"SeriaNumer,omitempty" yaml:"SeriaNumer,omitempty"`
	Status     *StatusDrukuEnumeracja `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// DocTypeRef_Stronicowanie was auto-generated from WSDL.
type DocTypeRef_Stronicowanie struct {
	RekordyOd      *int64 `xml:"RekordyOd,omitempty" json:"RekordyOd,omitempty" yaml:"RekordyOd,omitempty"`
	LiczbaRekordow *int64 `xml:"LiczbaRekordow,omitempty" json:"LiczbaRekordow,omitempty" yaml:"LiczbaRekordow,omitempty"`
}

// DocTypeRef_UPP was auto-generated from WSDL.
type DocTypeRef_UPP struct {
	UPP *UPP `xml:"UPP,omitempty" json:"UPP,omitempty" yaml:"UPP,omitempty"`
}

// DocTypeRef_WeryfikacjaAkcjiZla was auto-generated from WSDL.
type DocTypeRef_WeryfikacjaAkcjiZla struct {
	SeriaNumerZla     *DocTypeRef_SeriaNumerZla `xml:"SeriaNumerZla,omitempty" json:"SeriaNumerZla,omitempty" yaml:"SeriaNumerZla,omitempty"`
	DopuszczalnaAkcja *bool                     `xml:"DopuszczalnaAkcja,omitempty" json:"DopuszczalnaAkcja,omitempty" yaml:"DopuszczalnaAkcja,omitempty"`
}

// DocTypeRef_WniosekPr4 was auto-generated from WSDL.
type DocTypeRef_WniosekPr4 struct {
	DataWystawienia          *string `xml:"DataWystawienia,omitempty" json:"DataWystawienia,omitempty" yaml:"DataWystawienia,omitempty"`
	NumerStatystycznyChoroby *string `xml:"NumerStatystycznyChoroby,omitempty" json:"NumerStatystycznyChoroby,omitempty" yaml:"NumerStatystycznyChoroby,omitempty"`
	NpwzLekarza              *string `xml:"NpwzLekarza,omitempty" json:"NpwzLekarza,omitempty" yaml:"NpwzLekarza,omitempty"`
	ImieLekarza              *string `xml:"ImieLekarza,omitempty" json:"ImieLekarza,omitempty" yaml:"ImieLekarza,omitempty"`
	NazwiskoLekarza          *string `xml:"NazwiskoLekarza,omitempty" json:"NazwiskoLekarza,omitempty" yaml:"NazwiskoLekarza,omitempty"`
}

// DocTypeRef_WynikNotyfikacjiPlatnik was auto-generated from WSDL.
type DocTypeRef_WynikNotyfikacjiPlatnik struct {
	PlatnikStatusWysylki *string `xml:"PlatnikStatusWysylki,omitempty" json:"PlatnikStatusWysylki,omitempty" yaml:"PlatnikStatusWysylki,omitempty"`
	Nip                  *string `xml:"Nip,omitempty" json:"Nip,omitempty" yaml:"Nip,omitempty"`
	Pesel                *string `xml:"Pesel,omitempty" json:"Pesel,omitempty" yaml:"Pesel,omitempty"`
	Paszport             *string `xml:"Paszport,omitempty" json:"Paszport,omitempty" yaml:"Paszport,omitempty"`
}

// DocTypeRef_ZaswiadczenieLekarskie was auto-generated from WSDL.
type DocTypeRef_ZaswiadczenieLekarskie struct {
	PeselUbezpieczonego             *string                      `xml:"PeselUbezpieczonego,omitempty" json:"PeselUbezpieczonego,omitempty" yaml:"PeselUbezpieczonego,omitempty"`
	PaszportUbezpieczonego          *string                      `xml:"PaszportUbezpieczonego,omitempty" json:"PaszportUbezpieczonego,omitempty" yaml:"PaszportUbezpieczonego,omitempty"`
	ImieUbezpieczonego              *string                      `xml:"ImieUbezpieczonego,omitempty" json:"ImieUbezpieczonego,omitempty" yaml:"ImieUbezpieczonego,omitempty"`
	NazwiskoUbezpieczonego          *string                      `xml:"NazwiskoUbezpieczonego,omitempty" json:"NazwiskoUbezpieczonego,omitempty" yaml:"NazwiskoUbezpieczonego,omitempty"`
	NiezdolnoscDoPracyOd            *string                      `xml:"NiezdolnoscDoPracyOd,omitempty" json:"NiezdolnoscDoPracyOd,omitempty" yaml:"NiezdolnoscDoPracyOd,omitempty"`
	NiezdolnoscDoPracyDo            *string                      `xml:"NiezdolnoscDoPracyDo,omitempty" json:"NiezdolnoscDoPracyDo,omitempty" yaml:"NiezdolnoscDoPracyDo,omitempty"`
	LiczbaDniPobytuWSzpitalu        *int64                       `xml:"LiczbaDniPobytuWSzpitalu,omitempty" json:"LiczbaDniPobytuWSzpitalu,omitempty" yaml:"LiczbaDniPobytuWSzpitalu,omitempty"`
	PobytWSzpitaluOd                *string                      `xml:"PobytWSzpitaluOd,omitempty" json:"PobytWSzpitaluOd,omitempty" yaml:"PobytWSzpitaluOd,omitempty"`
	PobytWSzpitaluDo                *string                      `xml:"PobytWSzpitaluDo,omitempty" json:"PobytWSzpitaluDo,omitempty" yaml:"PobytWSzpitaluDo,omitempty"`
	KodPokrewienstwa                *KodyPokrewienstwaEnumeracja `xml:"KodPokrewienstwa,omitempty" json:"KodPokrewienstwa,omitempty" yaml:"KodPokrewienstwa,omitempty"`
	DataUrodzeniaOsobySpokrewnionej *string                      `xml:"DataUrodzeniaOsobySpokrewnionej,omitempty" json:"DataUrodzeniaOsobySpokrewnionej,omitempty" yaml:"DataUrodzeniaOsobySpokrewnionej,omitempty"`
	IdentyfikatorPlatnika           *string                      `xml:"IdentyfikatorPlatnika,omitempty" json:"IdentyfikatorPlatnika,omitempty" yaml:"IdentyfikatorPlatnika,omitempty"`
	TypIdentyfikatoraPlatnika       *string                      `xml:"TypIdentyfikatoraPlatnika,omitempty" json:"TypIdentyfikatoraPlatnika,omitempty" yaml:"TypIdentyfikatoraPlatnika,omitempty"`
	StatystycznyNumerChoroby        *string                      `xml:"StatystycznyNumerChoroby,omitempty" json:"StatystycznyNumerChoroby,omitempty" yaml:"StatystycznyNumerChoroby,omitempty"`
	NpwzLekarza                     *string                      `xml:"NpwzLekarza,omitempty" json:"NpwzLekarza,omitempty" yaml:"NpwzLekarza,omitempty"`
	ImieLekarza                     *string                      `xml:"ImieLekarza,omitempty" json:"ImieLekarza,omitempty" yaml:"ImieLekarza,omitempty"`
	NazwiskoLekarza                 *string                      `xml:"NazwiskoLekarza,omitempty" json:"NazwiskoLekarza,omitempty" yaml:"NazwiskoLekarza,omitempty"`
	StatusZla                       *string                      `xml:"StatusZla,omitempty" json:"StatusZla,omitempty" yaml:"StatusZla,omitempty"`
	DataWystawieniaZla              *string                      `xml:"DataWystawieniaZla,omitempty" json:"DataWystawieniaZla,omitempty" yaml:"DataWystawieniaZla,omitempty"`
	SeriaZla                        *string                      `xml:"SeriaZla,omitempty" json:"SeriaZla,omitempty" yaml:"SeriaZla,omitempty"`
	NumerZla                        *string                      `xml:"NumerZla,omitempty" json:"NumerZla,omitempty" yaml:"NumerZla,omitempty"`
}

// DocTypeRef_ZaswiadczenieLekarskieDanePelne was auto-generated
// from WSDL.
type DocTypeRef_ZaswiadczenieLekarskieDanePelne struct {
	PeselUbezpieczonego                 *string                              `xml:"PeselUbezpieczonego,omitempty" json:"PeselUbezpieczonego,omitempty" yaml:"PeselUbezpieczonego,omitempty"`
	PaszportUbezpieczonego              *string                              `xml:"PaszportUbezpieczonego,omitempty" json:"PaszportUbezpieczonego,omitempty" yaml:"PaszportUbezpieczonego,omitempty"`
	MiejsceUbezpieczenia                *string                              `xml:"MiejsceUbezpieczenia,omitempty" json:"MiejsceUbezpieczenia,omitempty" yaml:"MiejsceUbezpieczenia,omitempty"`
	ImieUbezpieczonego                  *string                              `xml:"ImieUbezpieczonego,omitempty" json:"ImieUbezpieczonego,omitempty" yaml:"ImieUbezpieczonego,omitempty"`
	NazwiskoUbezpieczonego              *string                              `xml:"NazwiskoUbezpieczonego,omitempty" json:"NazwiskoUbezpieczonego,omitempty" yaml:"NazwiskoUbezpieczonego,omitempty"`
	Ulica                               *string                              `xml:"Ulica,omitempty" json:"Ulica,omitempty" yaml:"Ulica,omitempty"`
	NrDomu                              *string                              `xml:"NrDomu,omitempty" json:"NrDomu,omitempty" yaml:"NrDomu,omitempty"`
	NrLokalu                            *string                              `xml:"NrLokalu,omitempty" json:"NrLokalu,omitempty" yaml:"NrLokalu,omitempty"`
	KodPocztowy                         *string                              `xml:"KodPocztowy,omitempty" json:"KodPocztowy,omitempty" yaml:"KodPocztowy,omitempty"`
	Miejscowosc                         *string                              `xml:"Miejscowosc,omitempty" json:"Miejscowosc,omitempty" yaml:"Miejscowosc,omitempty"`
	DataUrodzeniaUbezpieczonego         *Date                                `xml:"DataUrodzeniaUbezpieczonego,omitempty" json:"DataUrodzeniaUbezpieczonego,omitempty" yaml:"DataUrodzeniaUbezpieczonego,omitempty"`
	NiezdolnoscDoPracyOd                *Date                                `xml:"NiezdolnoscDoPracyOd,omitempty" json:"NiezdolnoscDoPracyOd,omitempty" yaml:"NiezdolnoscDoPracyOd,omitempty"`
	NiezdolnoscDoPracyDo                *Date                                `xml:"NiezdolnoscDoPracyDo,omitempty" json:"NiezdolnoscDoPracyDo,omitempty" yaml:"NiezdolnoscDoPracyDo,omitempty"`
	PobytWSzpitaluOd                    *Date                                `xml:"PobytWSzpitaluOd,omitempty" json:"PobytWSzpitaluOd,omitempty" yaml:"PobytWSzpitaluOd,omitempty"`
	PobytWSzpitaluDo                    *Date                                `xml:"PobytWSzpitaluDo,omitempty" json:"PobytWSzpitaluDo,omitempty" yaml:"PobytWSzpitaluDo,omitempty"`
	KodPokrewienstwa                    *KodyPokrewienstwaEnumeracja         `xml:"KodPokrewienstwa,omitempty" json:"KodPokrewienstwa,omitempty" yaml:"KodPokrewienstwa,omitempty"`
	DataUrodzeniaOsobySpokrewnionej     *Date                                `xml:"DataUrodzeniaOsobySpokrewnionej,omitempty" json:"DataUrodzeniaOsobySpokrewnionej,omitempty" yaml:"DataUrodzeniaOsobySpokrewnionej,omitempty"`
	IdentyfikatorPlatnika               *string                              `xml:"IdentyfikatorPlatnika,omitempty" json:"IdentyfikatorPlatnika,omitempty" yaml:"IdentyfikatorPlatnika,omitempty"`
	TypIdentyfikatoraPlatnika           *string                              `xml:"TypIdentyfikatoraPlatnika,omitempty" json:"TypIdentyfikatoraPlatnika,omitempty" yaml:"TypIdentyfikatoraPlatnika,omitempty"`
	StatystycznyNumerChoroby            *string                              `xml:"StatystycznyNumerChoroby,omitempty" json:"StatystycznyNumerChoroby,omitempty" yaml:"StatystycznyNumerChoroby,omitempty"`
	KodChorobyPierwszy                  *string                              `xml:"KodChorobyPierwszy,omitempty" json:"KodChorobyPierwszy,omitempty" yaml:"KodChorobyPierwszy,omitempty"`
	KodChorobyDrugi                     *string                              `xml:"KodChorobyDrugi,omitempty" json:"KodChorobyDrugi,omitempty" yaml:"KodChorobyDrugi,omitempty"`
	KodChorobyTrzeci                    *string                              `xml:"KodChorobyTrzeci,omitempty" json:"KodChorobyTrzeci,omitempty" yaml:"KodChorobyTrzeci,omitempty"`
	KodChorobyCzwarty                   *string                              `xml:"KodChorobyCzwarty,omitempty" json:"KodChorobyCzwarty,omitempty" yaml:"KodChorobyCzwarty,omitempty"`
	KodChorobyPiaty                     *string                              `xml:"KodChorobyPiaty,omitempty" json:"KodChorobyPiaty,omitempty" yaml:"KodChorobyPiaty,omitempty"`
	WskazanieLekarskie                  *string                              `xml:"WskazanieLekarskie,omitempty" json:"WskazanieLekarskie,omitempty" yaml:"WskazanieLekarskie,omitempty"`
	PobytStacjonarnyZoz                 *bool                                `xml:"PobytStacjonarnyZoz,omitempty" json:"PobytStacjonarnyZoz,omitempty" yaml:"PobytStacjonarnyZoz,omitempty"`
	NpwzLekarza                         *string                              `xml:"NpwzLekarza,omitempty" json:"NpwzLekarza,omitempty" yaml:"NpwzLekarza,omitempty"`
	NazwiskoLekarza                     *string                              `xml:"NazwiskoLekarza,omitempty" json:"NazwiskoLekarza,omitempty" yaml:"NazwiskoLekarza,omitempty"`
	ImieLekarza                         *string                              `xml:"ImieLekarza,omitempty" json:"ImieLekarza,omitempty" yaml:"ImieLekarza,omitempty"`
	PwdlNip                             *string                              `xml:"PwdlNip,omitempty" json:"PwdlNip,omitempty" yaml:"PwdlNip,omitempty"`
	PwdlNazwaSkrocona                   *string                              `xml:"PwdlNazwaSkrocona,omitempty" json:"PwdlNazwaSkrocona,omitempty" yaml:"PwdlNazwaSkrocona,omitempty"`
	PwdlKodPocztowy                     *string                              `xml:"PwdlKodPocztowy,omitempty" json:"PwdlKodPocztowy,omitempty" yaml:"PwdlKodPocztowy,omitempty"`
	PwdlMiejscowosc                     *string                              `xml:"PwdlMiejscowosc,omitempty" json:"PwdlMiejscowosc,omitempty" yaml:"PwdlMiejscowosc,omitempty"`
	PwdlUlica                           *string                              `xml:"PwdlUlica,omitempty" json:"PwdlUlica,omitempty" yaml:"PwdlUlica,omitempty"`
	PwdlNrDomu                          *string                              `xml:"PwdlNrDomu,omitempty" json:"PwdlNrDomu,omitempty" yaml:"PwdlNrDomu,omitempty"`
	PwdlNrLokalu                        *string                              `xml:"PwdlNrLokalu,omitempty" json:"PwdlNrLokalu,omitempty" yaml:"PwdlNrLokalu,omitempty"`
	StatusZla                           *string                              `xml:"StatusZla,omitempty" json:"StatusZla,omitempty" yaml:"StatusZla,omitempty"`
	DataWystawieniaZla                  *Date                                `xml:"DataWystawieniaZla,omitempty" json:"DataWystawieniaZla,omitempty" yaml:"DataWystawieniaZla,omitempty"`
	SeriaZla                            *string                              `xml:"SeriaZla,omitempty" json:"SeriaZla,omitempty" yaml:"SeriaZla,omitempty"`
	NumerZla                            *string                              `xml:"NumerZla,omitempty" json:"NumerZla,omitempty" yaml:"NumerZla,omitempty"`
	UzasadnienieZaswiadczeniaWstecznego *UzasadnienieZaswiadczeniaWstecznego `xml:"UzasadnienieZaswiadczeniaWstecznego,omitempty" json:"UzasadnienieZaswiadczeniaWstecznego,omitempty" yaml:"UzasadnienieZaswiadczeniaWstecznego,omitempty"`
	UkryteDlaPlatnika                   *bool                                `xml:"UkryteDlaPlatnika,omitempty" json:"UkryteDlaPlatnika,omitempty" yaml:"UkryteDlaPlatnika,omitempty"`
	KodPrzyczynyAnulowania              *string                              `xml:"KodPrzyczynyAnulowania,omitempty" json:"KodPrzyczynyAnulowania,omitempty" yaml:"KodPrzyczynyAnulowania,omitempty"`
}

// DocTypeRef_ZaswiadczenieLekarskieKsiFiltrowanie was auto-generated
// from WSDL.
type DocTypeRef_ZaswiadczenieLekarskieKsiFiltrowanie struct {
	WarunekFiltrowania        []*DocTypeRef_ZaswiadczenieLekarskieKsiFiltrowanieWarunek `xml:"WarunekFiltrowania,omitempty" json:"WarunekFiltrowania,omitempty" yaml:"WarunekFiltrowania,omitempty"`
	LogiczneZlaczenieWarunkow *OperatorAndOrEnum                                        `xml:"LogiczneZlaczenieWarunkow,omitempty" json:"LogiczneZlaczenieWarunkow,omitempty" yaml:"LogiczneZlaczenieWarunkow,omitempty"`
}

// DocTypeRef_ZaswiadczenieLekarskieKsiFiltrowanieWarunek was auto-generated
// from WSDL.
type DocTypeRef_ZaswiadczenieLekarskieKsiFiltrowanieWarunek struct {
	Kolumna  *ZaswiadczenieLekarskieKsiKolumnyEnumeracja `xml:"Kolumna,omitempty" json:"Kolumna,omitempty" yaml:"Kolumna,omitempty"`
	Operator *OperatorEnum                               `xml:"Operator,omitempty" json:"Operator,omitempty" yaml:"Operator,omitempty"`
	Wartosc  *string                                     `xml:"Wartosc,omitempty" json:"Wartosc,omitempty" yaml:"Wartosc,omitempty"`
}

// DocTypeRef_ZaswiadczenieLekarskieKsiKorygujace was auto-generated
// from WSDL.
type DocTypeRef_ZaswiadczenieLekarskieKsiKorygujace struct {
	IdentyfikatorUbezpieczonego     *string                        `xml:"IdentyfikatorUbezpieczonego,omitempty" json:"IdentyfikatorUbezpieczonego,omitempty" yaml:"IdentyfikatorUbezpieczonego,omitempty"`
	TypIdentyfikatoraUbezpieczonego *TypIdUbezpieczonegoEnumeracja `xml:"TypIdentyfikatoraUbezpieczonego,omitempty" json:"TypIdentyfikatoraUbezpieczonego,omitempty" yaml:"TypIdentyfikatoraUbezpieczonego,omitempty"`
	ImieUbezpieczonego              *string                        `xml:"ImieUbezpieczonego,omitempty" json:"ImieUbezpieczonego,omitempty" yaml:"ImieUbezpieczonego,omitempty"`
	NazwiskoUbezpieczonego          *string                        `xml:"NazwiskoUbezpieczonego,omitempty" json:"NazwiskoUbezpieczonego,omitempty" yaml:"NazwiskoUbezpieczonego,omitempty"`
	Ulica                           *string                        `xml:"Ulica,omitempty" json:"Ulica,omitempty" yaml:"Ulica,omitempty"`
	NrDomu                          *string                        `xml:"NrDomu,omitempty" json:"NrDomu,omitempty" yaml:"NrDomu,omitempty"`
	NrLokalu                        *string                        `xml:"NrLokalu,omitempty" json:"NrLokalu,omitempty" yaml:"NrLokalu,omitempty"`
	KodPocztowy                     *string                        `xml:"KodPocztowy,omitempty" json:"KodPocztowy,omitempty" yaml:"KodPocztowy,omitempty"`
	Miejscowosc                     *string                        `xml:"Miejscowosc,omitempty" json:"Miejscowosc,omitempty" yaml:"Miejscowosc,omitempty"`
	DataUrodzeniaUbezpieczonego     *Date                          `xml:"DataUrodzeniaUbezpieczonego,omitempty" json:"DataUrodzeniaUbezpieczonego,omitempty" yaml:"DataUrodzeniaUbezpieczonego,omitempty"`
	DataUstaniaNiezdolnosciDoPracy  *Date                          `xml:"DataUstaniaNiezdolnosciDoPracy,omitempty" json:"DataUstaniaNiezdolnosciDoPracy,omitempty" yaml:"DataUstaniaNiezdolnosciDoPracy,omitempty"`
	NazwaTjo                        *string                        `xml:"NazwaTjo,omitempty" json:"NazwaTjo,omitempty" yaml:"NazwaTjo,omitempty"`
	IdentyfikatorPlatnika           *string                        `xml:"IdentyfikatorPlatnika,omitempty" json:"IdentyfikatorPlatnika,omitempty" yaml:"IdentyfikatorPlatnika,omitempty"`
	TypIdentyfikatoraPlatnika       *string                        `xml:"TypIdentyfikatoraPlatnika,omitempty" json:"TypIdentyfikatoraPlatnika,omitempty" yaml:"TypIdentyfikatoraPlatnika,omitempty"`
	SeriaNumerZla                   *string                        `xml:"SeriaNumerZla,omitempty" json:"SeriaNumerZla,omitempty" yaml:"SeriaNumerZla,omitempty"`
	DataWystawieniaZlaKorygujacego  *Date                          `xml:"DataWystawieniaZlaKorygujacego,omitempty" json:"DataWystawieniaZlaKorygujacego,omitempty" yaml:"DataWystawieniaZlaKorygujacego,omitempty"`
}

// DocTypeRef_ZaswiadczenieLekarskieKsiSortowanie was auto-generated
// from WSDL.
type DocTypeRef_ZaswiadczenieLekarskieKsiSortowanie struct {
	WarunekSortowania []*DocTypeRef_ZaswiadczenieLekarskieKsiSortowanieWarunek `xml:"WarunekSortowania,omitempty" json:"WarunekSortowania,omitempty" yaml:"WarunekSortowania,omitempty"`
}

// DocTypeRef_ZaswiadczenieLekarskieKsiSortowanieWarunek was auto-generated
// from WSDL.
type DocTypeRef_ZaswiadczenieLekarskieKsiSortowanieWarunek struct {
	Kolumna  *ZaswiadczenieLekarskieKsiKolumnyEnumeracja `xml:"Kolumna,omitempty" json:"Kolumna,omitempty" yaml:"Kolumna,omitempty"`
	Kierunek *SortowanieEnumeracja                       `xml:"Kierunek,omitempty" json:"Kierunek,omitempty" yaml:"Kierunek,omitempty"`
}

// DocTypeRef_ZaswiadczenieLekarskiePueFiltrowanie was auto-generated
// from WSDL.
type DocTypeRef_ZaswiadczenieLekarskiePueFiltrowanie struct {
	WarunekFiltrowania        []*DocTypeRef_ZaswiadczenieLekarskiePueFiltrowanieWarunek `xml:"WarunekFiltrowania,omitempty" json:"WarunekFiltrowania,omitempty" yaml:"WarunekFiltrowania,omitempty"`
	LogiczneZlaczenieWarunkow *OperatorAndOrEnum                                        `xml:"LogiczneZlaczenieWarunkow,omitempty" json:"LogiczneZlaczenieWarunkow,omitempty" yaml:"LogiczneZlaczenieWarunkow,omitempty"`
}

// DocTypeRef_ZaswiadczenieLekarskiePueFiltrowanieWarunek was auto-generated
// from WSDL.
type DocTypeRef_ZaswiadczenieLekarskiePueFiltrowanieWarunek struct {
	Kolumna  *ZaswiadczenieLekarskiePueKolumnyEnumeracja `xml:"Kolumna,omitempty" json:"Kolumna,omitempty" yaml:"Kolumna,omitempty"`
	Operator *OperatorEnum                               `xml:"Operator,omitempty" json:"Operator,omitempty" yaml:"Operator,omitempty"`
	Wartosc  *string                                     `xml:"Wartosc,omitempty" json:"Wartosc,omitempty" yaml:"Wartosc,omitempty"`
}

// DocTypeRef_ZaswiadczenieLekarskiePueSortowanie was auto-generated
// from WSDL.
type DocTypeRef_ZaswiadczenieLekarskiePueSortowanie struct {
	WarunekSortowania []*DocTypeRef_ZaswiadczenieLekarskiePueSortowanieWarunek `xml:"WarunekSortowania,omitempty" json:"WarunekSortowania,omitempty" yaml:"WarunekSortowania,omitempty"`
}

// DocTypeRef_ZaswiadczenieLekarskiePueSortowanieWarunek was auto-generated
// from WSDL.
type DocTypeRef_ZaswiadczenieLekarskiePueSortowanieWarunek struct {
	Kolumna  *ZaswiadczenieLekarskiePueKolumnyEnumeracja `xml:"Kolumna,omitempty" json:"Kolumna,omitempty" yaml:"Kolumna,omitempty"`
	Kierunek *SortowanieEnumeracja                       `xml:"Kierunek,omitempty" json:"Kierunek,omitempty" yaml:"Kierunek,omitempty"`
}

// DocTypeRef_kmdk_Adres was auto-generated from WSDL.
type DocTypeRef_kmdk_Adres struct {
	KodPocztowy *string `xml:"KodPocztowy,omitempty" json:"KodPocztowy,omitempty" yaml:"KodPocztowy,omitempty"`
	Miejscowosc *string `xml:"Miejscowosc,omitempty" json:"Miejscowosc,omitempty" yaml:"Miejscowosc,omitempty"`
	Ulica       *string `xml:"Ulica,omitempty" json:"Ulica,omitempty" yaml:"Ulica,omitempty"`
	NrDomu      *string `xml:"NrDomu,omitempty" json:"NrDomu,omitempty" yaml:"NrDomu,omitempty"`
	NrLokalu    *string `xml:"NrLokalu,omitempty" json:"NrLokalu,omitempty" yaml:"NrLokalu,omitempty"`
}

// DocTypeRef_kmdk_AdresLekarza was auto-generated from WSDL.
type DocTypeRef_kmdk_AdresLekarza struct {
	Ulica       *string `xml:"Ulica,omitempty" json:"Ulica,omitempty" yaml:"Ulica,omitempty"`
	NrDomu      *string `xml:"NrDomu,omitempty" json:"NrDomu,omitempty" yaml:"NrDomu,omitempty"`
	NrLokalu    *string `xml:"NrLokalu,omitempty" json:"NrLokalu,omitempty" yaml:"NrLokalu,omitempty"`
	KodPocztowy *string `xml:"KodPocztowy,omitempty" json:"KodPocztowy,omitempty" yaml:"KodPocztowy,omitempty"`
	Miejscowosc *string `xml:"Miejscowosc,omitempty" json:"Miejscowosc,omitempty" yaml:"Miejscowosc,omitempty"`
}

// DocTypeRef_kmdk_Lekarz_DanePelne was auto-generated from WSDL.
type DocTypeRef_kmdk_Lekarz_DanePelne struct {
	Imie                         *string                                     `xml:"Imie,omitempty" json:"Imie,omitempty" yaml:"Imie,omitempty"`
	Nazwisko                     *string                                     `xml:"Nazwisko,omitempty" json:"Nazwisko,omitempty" yaml:"Nazwisko,omitempty"`
	Pesel                        *string                                     `xml:"Pesel,omitempty" json:"Pesel,omitempty" yaml:"Pesel,omitempty"`
	NumerPrawaWykonywaniaZawodu  *string                                     `xml:"NumerPrawaWykonywaniaZawodu,omitempty" json:"NumerPrawaWykonywaniaZawodu,omitempty" yaml:"NumerPrawaWykonywaniaZawodu,omitempty"`
	NazwaOkregowejIzbyLekarskiej *string                                     `xml:"NazwaOkregowejIzbyLekarskiej,omitempty" json:"NazwaOkregowejIzbyLekarskiej,omitempty" yaml:"NazwaOkregowejIzbyLekarskiej,omitempty"`
	StatusLekarza                *DocTypeRef_kmdk_StatusLekarza              `xml:"StatusLekarza,omitempty" json:"StatusLekarza,omitempty" yaml:"StatusLekarza,omitempty"`
	PosiadanaSpecjalizacja       []*DocTypeRef_kmdk_PosiadanaSpecjalizacja   `xml:"PosiadanaSpecjalizacja,omitempty" json:"PosiadanaSpecjalizacja,omitempty" yaml:"PosiadanaSpecjalizacja,omitempty"`
	Adres                        []*DocTypeRef_kmdk_AdresLekarza             `xml:"Adres,omitempty" json:"Adres,omitempty" yaml:"Adres,omitempty"`
	MiejsceWykonywaniaZawodu     []*DocTypeRef_kmdk_MiejsceWykonywaniaZawodu `xml:"MiejsceWykonywaniaZawodu,omitempty" json:"MiejsceWykonywaniaZawodu,omitempty" yaml:"MiejsceWykonywaniaZawodu,omitempty"`
}

// DocTypeRef_kmdk_MiejsceWykonywaniaZawodu was auto-generated
// from WSDL.
type DocTypeRef_kmdk_MiejsceWykonywaniaZawodu struct {
	Nazwa         *string `xml:"Nazwa,omitempty" json:"Nazwa,omitempty" yaml:"Nazwa,omitempty"`
	NIP           *string `xml:"NIP,omitempty" json:"NIP,omitempty" yaml:"NIP,omitempty"`
	AdresEMail    *string `xml:"AdresEMail,omitempty" json:"AdresEMail,omitempty" yaml:"AdresEMail,omitempty"`
	NumerFaksu    *string `xml:"NumerFaksu,omitempty" json:"NumerFaksu,omitempty" yaml:"NumerFaksu,omitempty"`
	NumerTelefonu *string `xml:"NumerTelefonu,omitempty" json:"NumerTelefonu,omitempty" yaml:"NumerTelefonu,omitempty"`
	Ulica         *string `xml:"Ulica,omitempty" json:"Ulica,omitempty" yaml:"Ulica,omitempty"`
	NrDomu        *string `xml:"NrDomu,omitempty" json:"NrDomu,omitempty" yaml:"NrDomu,omitempty"`
	NrLokalu      *string `xml:"NrLokalu,omitempty" json:"NrLokalu,omitempty" yaml:"NrLokalu,omitempty"`
	KodPocztowy   *string `xml:"KodPocztowy,omitempty" json:"KodPocztowy,omitempty" yaml:"KodPocztowy,omitempty"`
	Miejscowosc   *string `xml:"Miejscowosc,omitempty" json:"Miejscowosc,omitempty" yaml:"Miejscowosc,omitempty"`
}

// DocTypeRef_kmdk_PosiadanaSpecjalizacja was auto-generated from
// WSDL.
type DocTypeRef_kmdk_PosiadanaSpecjalizacja struct {
	KodSpecjalizacji     *string `xml:"KodSpecjalizacji,omitempty" json:"KodSpecjalizacji,omitempty" yaml:"KodSpecjalizacji,omitempty"`
	NazwaSpecjalizacji   *string `xml:"NazwaSpecjalizacji,omitempty" json:"NazwaSpecjalizacji,omitempty" yaml:"NazwaSpecjalizacji,omitempty"`
	StopienSpecjalizacji *string `xml:"StopienSpecjalizacji,omitempty" json:"StopienSpecjalizacji,omitempty" yaml:"StopienSpecjalizacji,omitempty"`
}

// DocTypeRef_kmdk_RezultatWysylkiPakietu was auto-generated from
// WSDL.
type DocTypeRef_kmdk_RezultatWysylkiPakietu struct {
	NrRef        []*string                    `xml:"NrRef,omitempty" json:"NrRef,omitempty" yaml:"NrRef,omitempty"`
	Wynik        *RezultatWalidacjiEnumeracja `xml:"Wynik,omitempty" json:"Wynik,omitempty" yaml:"Wynik,omitempty"`
	UtworzoneUPP []*DocTypeRef_UPP            `xml:"UtworzoneUPP,omitempty" json:"UtworzoneUPP,omitempty" yaml:"UtworzoneUPP,omitempty"`
	Blad         *DocTypeRef_Rezultat         `xml:"Blad,omitempty" json:"Blad,omitempty" yaml:"Blad,omitempty"`
}

// DocTypeRef_kmdk_StatusLekarza was auto-generated from WSDL.
type DocTypeRef_kmdk_StatusLekarza struct {
	DataWydaniaDecyzji *string `xml:"DataWydaniaDecyzji,omitempty" json:"DataWydaniaDecyzji,omitempty" yaml:"DataWydaniaDecyzji,omitempty"`
	Status             *string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// DocTypeRef_kmdk_Ubezpieczony was auto-generated from WSDL.
type DocTypeRef_kmdk_Ubezpieczony struct {
	Imie          *string                  `xml:"Imie,omitempty" json:"Imie,omitempty" yaml:"Imie,omitempty"`
	Nazwisko      *string                  `xml:"Nazwisko,omitempty" json:"Nazwisko,omitempty" yaml:"Nazwisko,omitempty"`
	DataUrodzenia *string                  `xml:"DataUrodzenia,omitempty" json:"DataUrodzenia,omitempty" yaml:"DataUrodzenia,omitempty"`
	DataZgonu     *string                  `xml:"DataZgonu,omitempty" json:"DataZgonu,omitempty" yaml:"DataZgonu,omitempty"`
	Adres         []*DocTypeRef_kmdk_Adres `xml:"Adres,omitempty" json:"Adres,omitempty" yaml:"Adres,omitempty"`
}

// NadajSeriaNumerZla was auto-generated from WSDL.
type NadajSeriaNumerZla struct {
	IdSesji        *ID                             `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	ParaDokumentow []*DocTypeRef_ParaDokumentowZla `xml:"ParaDokumentow,omitempty" json:"ParaDokumentow,omitempty" yaml:"ParaDokumentow,omitempty"`
}

// NadajSeriaNumerZlaResponse was auto-generated from WSDL.
type NadajSeriaNumerZlaResponse struct {
	ParaDokumentow    []*DocTypeRef_ParaDokumentowZla `xml:"ParaDokumentow,omitempty" json:"ParaDokumentow,omitempty" yaml:"ParaDokumentow,omitempty"`
	RezultatWalidacji *DocTypeRef_RezultatWalidacji   `xml:"RezultatWalidacji,omitempty" json:"RezultatWalidacji,omitempty" yaml:"RezultatWalidacji,omitempty"`
	Rezultat          *DocTypeRef_Rezultat            `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// NotyfikujAnulowanieZla was auto-generated from WSDL.
type NotyfikujAnulowanieZla struct {
	IdSesji         *ID                                  `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	DaneNotyfikacji *DocTypeRef_NotyfikujAnulowanieZlaIn `xml:"DaneNotyfikacji,omitempty" json:"DaneNotyfikacji,omitempty" yaml:"DaneNotyfikacji,omitempty"`
}

// NotyfikujAnulowanieZlaResponse was auto-generated from WSDL.
type NotyfikujAnulowanieZlaResponse struct {
	Rezultat                  *DocTypeRef_Rezultat                  `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	UbezpieczonyStatusWysylki *string                               `xml:"UbezpieczonyStatusWysylki,omitempty" json:"UbezpieczonyStatusWysylki,omitempty" yaml:"UbezpieczonyStatusWysylki,omitempty"`
	PlatnikStatusWysylki      []*DocTypeRef_WynikNotyfikacjiPlatnik `xml:"PlatnikStatusWysylki,omitempty" json:"PlatnikStatusWysylki,omitempty" yaml:"PlatnikStatusWysylki,omitempty"`
}

// PobierzAdresyPlacowek was auto-generated from WSDL.
type PobierzAdresyPlacowek struct {
	IdSesji *ID  `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	NIP     *NIP `xml:"NIP,omitempty" json:"NIP,omitempty" yaml:"NIP,omitempty"`
}

// PobierzAdresyPlacowekResponse was auto-generated from WSDL.
type PobierzAdresyPlacowekResponse struct {
	NazwaPlatnika  *string                      `xml:"NazwaPlatnika,omitempty" json:"NazwaPlatnika,omitempty" yaml:"NazwaPlatnika,omitempty"`
	AdresyPlacowek []*DocTypeRef_AdresyPlacowek `xml:"AdresyPlacowek,omitempty" json:"AdresyPlacowek,omitempty" yaml:"AdresyPlacowek,omitempty"`
	Rezultat       *DocTypeRef_Rezultat         `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// PobierzCzlonkowRodzinyUbezpieczonego was auto-generated from
// WSDL.
type PobierzCzlonkowRodzinyUbezpieczonego struct {
	IdSesji         *ID                                     `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Ubezpieczony    *DocTypeRef_IdentyfikacjaUbezpieczonego `xml:"Ubezpieczony,omitempty" json:"Ubezpieczony,omitempty" yaml:"Ubezpieczony,omitempty"`
	KontekstDostepu *KontekstDostepu                        `xml:"KontekstDostepu,omitempty" json:"KontekstDostepu,omitempty" yaml:"KontekstDostepu,omitempty"`
}

// PobierzCzlonkowRodzinyUbezpieczonegoResponse was auto-generated
// from WSDL.
type PobierzCzlonkowRodzinyUbezpieczonegoResponse struct {
	Rezultat                     *DocTypeRef_Rezultat                       `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	CzlonekRodzinyUbezpieczonego []*DocTypeRef_CzlonekRodzinyUbezpieczonego `xml:"CzlonekRodzinyUbezpieczonego,omitempty" json:"CzlonekRodzinyUbezpieczonego,omitempty" yaml:"CzlonekRodzinyUbezpieczonego,omitempty"`
}

// PobierzDaneLekarza was auto-generated from WSDL.
type PobierzDaneLekarza struct {
	IdSesji *ID `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
}

// PobierzDaneLekarzaResponse was auto-generated from WSDL.
type PobierzDaneLekarzaResponse struct {
	DaneLekarza *DocTypeRef_kmdk_Lekarz_DanePelne `xml:"DaneLekarza,omitempty" json:"DaneLekarza,omitempty" yaml:"DaneLekarza,omitempty"`
	Rezultat    *DocTypeRef_Rezultat              `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// PobierzDanePlatnika was auto-generated from WSDL.
type PobierzDanePlatnika struct {
	IdSesji *ID                               `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Platnik *DocTypeRef_IdentyfikacjaPlatnika `xml:"Platnik,omitempty" json:"Platnik,omitempty" yaml:"Platnik,omitempty"`
}

// PobierzDanePlatnikaResponse was auto-generated from WSDL.
type PobierzDanePlatnikaResponse struct {
	DanePlatnika *DocTypeRef_Platnik  `xml:"DanePlatnika,omitempty" json:"DanePlatnika,omitempty" yaml:"DanePlatnika,omitempty"`
	Rezultat     *DocTypeRef_Rezultat `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// PobierzDaneUbezpieczonego was auto-generated from WSDL.
type PobierzDaneUbezpieczonego struct {
	IdSesji         *ID                                     `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Ubezpieczony    *DocTypeRef_IdentyfikacjaUbezpieczonego `xml:"Ubezpieczony,omitempty" json:"Ubezpieczony,omitempty" yaml:"Ubezpieczony,omitempty"`
	KontekstDostepu *KontekstDostepu                        `xml:"KontekstDostepu,omitempty" json:"KontekstDostepu,omitempty" yaml:"KontekstDostepu,omitempty"`
}

// PobierzDaneUbezpieczonegoResponse was auto-generated from WSDL.
type PobierzDaneUbezpieczonegoResponse struct {
	DaneUbezpieczonego *DocTypeRef_kmdk_Ubezpieczony `xml:"DaneUbezpieczonego,omitempty" json:"DaneUbezpieczonego,omitempty" yaml:"DaneUbezpieczonego,omitempty"`
	Rezultat           *DocTypeRef_Rezultat          `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// PobierzDokument was auto-generated from WSDL.
type PobierzDokument struct {
	IdSesji         *ID                        `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	RodzajDokumentu *RodzajDokumentuEnumeracja `xml:"RodzajDokumentu,omitempty" json:"RodzajDokumentu,omitempty" yaml:"RodzajDokumentu,omitempty"`
	SeriaNumerZla   *DocTypeRef_SeriaNumerZla  `xml:"SeriaNumerZla,omitempty" json:"SeriaNumerZla,omitempty" yaml:"SeriaNumerZla,omitempty"`
	KontekstDostepu *KontekstDostepu           `xml:"KontekstDostepu,omitempty" json:"KontekstDostepu,omitempty" yaml:"KontekstDostepu,omitempty"`
}

// PobierzDokumentResponse was auto-generated from WSDL.
type PobierzDokumentResponse struct {
	WersjaKedu *string              `xml:"WersjaKedu,omitempty" json:"WersjaKedu,omitempty" yaml:"WersjaKedu,omitempty"`
	Dokument   *DocTypeRef_KEDU     `xml:"Dokument,omitempty" json:"Dokument,omitempty" yaml:"Dokument,omitempty"`
	Rezultat   *DocTypeRef_Rezultat `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// PobierzIdentyfikatorDokumentu was auto-generated from WSDL.
type PobierzIdentyfikatorDokumentu struct {
	IdSesji              *ID   `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	IloscIdentyfikatorow *uint `xml:"IloscIdentyfikatorow,omitempty" json:"IloscIdentyfikatorow,omitempty" yaml:"IloscIdentyfikatorow,omitempty"`
}

// PobierzIdentyfikatorDokumentuResponse was auto-generated from
// WSDL.
type PobierzIdentyfikatorDokumentuResponse struct {
	IdDokumentu []*string            `xml:"IdDokumentu,omitempty" json:"IdDokumentu,omitempty" yaml:"IdDokumentu,omitempty"`
	Rezultat    *DocTypeRef_Rezultat `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// PobierzKodChoroby was auto-generated from WSDL.
type PobierzKodChoroby struct {
	IdSesji       *ID                               `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Stronicowanie *DocTypeRef_Stronicowanie         `xml:"Stronicowanie,omitempty" json:"Stronicowanie,omitempty" yaml:"Stronicowanie,omitempty"`
	Sortowanie    *DocTypeRef_KodChorobySortowanie  `xml:"Sortowanie,omitempty" json:"Sortowanie,omitempty" yaml:"Sortowanie,omitempty"`
	Filtrowanie   *DocTypeRef_KodChorobyFiltrowanie `xml:"Filtrowanie,omitempty" json:"Filtrowanie,omitempty" yaml:"Filtrowanie,omitempty"`
}

// PobierzKodChorobyResponse was auto-generated from WSDL.
type PobierzKodChorobyResponse struct {
	Rezultat                 *DocTypeRef_Rezultat                   `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	LiczbaWszystkichRekordow *int64                                 `xml:"LiczbaWszystkichRekordow,omitempty" json:"LiczbaWszystkichRekordow,omitempty" yaml:"LiczbaWszystkichRekordow,omitempty"`
	NumerStatystycznyChoroby []*DocTypeRef_NumerStatystycznyChoroby `xml:"NumerStatystycznyChoroby,omitempty" json:"NumerStatystycznyChoroby,omitempty" yaml:"NumerStatystycznyChoroby,omitempty"`
}

// PobierzListeBiezacychZlaLekarza was auto-generated from WSDL.
type PobierzListeBiezacychZlaLekarza struct {
	IdSesji         *ID                                              `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Stronicowanie   *DocTypeRef_Stronicowanie                        `xml:"Stronicowanie,omitempty" json:"Stronicowanie,omitempty" yaml:"Stronicowanie,omitempty"`
	Sortowanie      *DocTypeRef_ZaswiadczenieLekarskiePueSortowanie  `xml:"Sortowanie,omitempty" json:"Sortowanie,omitempty" yaml:"Sortowanie,omitempty"`
	Filtrowanie     *DocTypeRef_ZaswiadczenieLekarskiePueFiltrowanie `xml:"Filtrowanie,omitempty" json:"Filtrowanie,omitempty" yaml:"Filtrowanie,omitempty"`
	KontekstDostepu *KontekstDostepu                                 `xml:"KontekstDostepu,omitempty" json:"KontekstDostepu,omitempty" yaml:"KontekstDostepu,omitempty"`
	NipPlacowki     *NIP                                             `xml:"NipPlacowki,omitempty" json:"NipPlacowki,omitempty" yaml:"NipPlacowki,omitempty"`
}

// PobierzListeBiezacychZlaLekarzaResponse was auto-generated from
// WSDL.
type PobierzListeBiezacychZlaLekarzaResponse struct {
	Rezultat                 *DocTypeRef_Rezultat                 `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	LiczbaWszystkichRekordow *int64                               `xml:"LiczbaWszystkichRekordow,omitempty" json:"LiczbaWszystkichRekordow,omitempty" yaml:"LiczbaWszystkichRekordow,omitempty"`
	ZaswiadczenieLekarskie   []*DocTypeRef_ZaswiadczenieLekarskie `xml:"ZaswiadczenieLekarskie,omitempty" json:"ZaswiadczenieLekarskie,omitempty" yaml:"ZaswiadczenieLekarskie,omitempty"`
}

// PobierzListePowiazanychZla was auto-generated from WSDL.
type PobierzListePowiazanychZla struct {
	IdSesji *ID                       `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Zla     *DocTypeRef_SeriaNumerZla `xml:"Zla,omitempty" json:"Zla,omitempty" yaml:"Zla,omitempty"`
}

// PobierzListePowiazanychZlaResponse was auto-generated from WSDL.
type PobierzListePowiazanychZlaResponse struct {
	Rezultat             *DocTypeRef_Rezultat       `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	PowiazaneZla         []*DocTypeRef_PowiazaneZla `xml:"PowiazaneZla,omitempty" json:"PowiazaneZla,omitempty" yaml:"PowiazaneZla,omitempty"`
	LiczbaPowiazanychZla *int64                     `xml:"LiczbaPowiazanychZla,omitempty" json:"LiczbaPowiazanychZla,omitempty" yaml:"LiczbaPowiazanychZla,omitempty"`
}

// PobierzListePr4Ubezpieczonego was auto-generated from WSDL.
type PobierzListePr4Ubezpieczonego struct {
	IdSesji         *ID                                     `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Stronicowanie   *DocTypeRef_Stronicowanie               `xml:"Stronicowanie,omitempty" json:"Stronicowanie,omitempty" yaml:"Stronicowanie,omitempty"`
	Ubezpieczony    *DocTypeRef_IdentyfikacjaUbezpieczonego `xml:"Ubezpieczony,omitempty" json:"Ubezpieczony,omitempty" yaml:"Ubezpieczony,omitempty"`
	DataOd          *Data                                   `xml:"DataOd,omitempty" json:"DataOd,omitempty" yaml:"DataOd,omitempty"`
	DataDo          *Data                                   `xml:"DataDo,omitempty" json:"DataDo,omitempty" yaml:"DataDo,omitempty"`
	KontekstDostepu *KontekstDostepu                        `xml:"KontekstDostepu,omitempty" json:"KontekstDostepu,omitempty" yaml:"KontekstDostepu,omitempty"`
}

// PobierzListePr4UbezpieczonegoResponse was auto-generated from
// WSDL.
type PobierzListePr4UbezpieczonegoResponse struct {
	Rezultat                 *DocTypeRef_Rezultat     `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	WniosekPr4               []*DocTypeRef_WniosekPr4 `xml:"WniosekPr4,omitempty" json:"WniosekPr4,omitempty" yaml:"WniosekPr4,omitempty"`
	LiczbaWszystkichRekordow *int64                   `xml:"LiczbaWszystkichRekordow,omitempty" json:"LiczbaWszystkichRekordow,omitempty" yaml:"LiczbaWszystkichRekordow,omitempty"`
}

// PobierzListePrzetworzonychZlaLekarza was auto-generated from
// WSDL.
type PobierzListePrzetworzonychZlaLekarza struct {
	IdSesji         *ID                                              `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Stronicowanie   *DocTypeRef_Stronicowanie                        `xml:"Stronicowanie,omitempty" json:"Stronicowanie,omitempty" yaml:"Stronicowanie,omitempty"`
	Sortowanie      *DocTypeRef_ZaswiadczenieLekarskieKsiSortowanie  `xml:"Sortowanie,omitempty" json:"Sortowanie,omitempty" yaml:"Sortowanie,omitempty"`
	Filtrowanie     *DocTypeRef_ZaswiadczenieLekarskieKsiFiltrowanie `xml:"Filtrowanie,omitempty" json:"Filtrowanie,omitempty" yaml:"Filtrowanie,omitempty"`
	KontekstDostepu *KontekstDostepu                                 `xml:"KontekstDostepu,omitempty" json:"KontekstDostepu,omitempty" yaml:"KontekstDostepu,omitempty"`
	NipPlacowki     *NIP                                             `xml:"NipPlacowki,omitempty" json:"NipPlacowki,omitempty" yaml:"NipPlacowki,omitempty"`
}

// PobierzListePrzetworzonychZlaLekarzaResponse was auto-generated
// from WSDL.
type PobierzListePrzetworzonychZlaLekarzaResponse struct {
	Rezultat                 *DocTypeRef_Rezultat                 `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	LiczbaWszystkichRekordow *int64                               `xml:"LiczbaWszystkichRekordow,omitempty" json:"LiczbaWszystkichRekordow,omitempty" yaml:"LiczbaWszystkichRekordow,omitempty"`
	ZaswiadczenieLekarskie   []*DocTypeRef_ZaswiadczenieLekarskie `xml:"ZaswiadczenieLekarskie,omitempty" json:"ZaswiadczenieLekarskie,omitempty" yaml:"ZaswiadczenieLekarskie,omitempty"`
}

// PobierzListeZlaUbezpieczonego was auto-generated from WSDL.
type PobierzListeZlaUbezpieczonego struct {
	IdSesji         *ID                                              `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Ubezpieczony    *DocTypeRef_IdentyfikacjaUbezpieczonego          `xml:"Ubezpieczony,omitempty" json:"Ubezpieczony,omitempty" yaml:"Ubezpieczony,omitempty"`
	Stronicowanie   *DocTypeRef_Stronicowanie                        `xml:"Stronicowanie,omitempty" json:"Stronicowanie,omitempty" yaml:"Stronicowanie,omitempty"`
	Sortowanie      *DocTypeRef_ZaswiadczenieLekarskieKsiSortowanie  `xml:"Sortowanie,omitempty" json:"Sortowanie,omitempty" yaml:"Sortowanie,omitempty"`
	Filtrowanie     *DocTypeRef_ZaswiadczenieLekarskieKsiFiltrowanie `xml:"Filtrowanie,omitempty" json:"Filtrowanie,omitempty" yaml:"Filtrowanie,omitempty"`
	KontekstDostepu *KontekstDostepu                                 `xml:"KontekstDostepu,omitempty" json:"KontekstDostepu,omitempty" yaml:"KontekstDostepu,omitempty"`
}

// PobierzListeZlaUbezpieczonegoResponse was auto-generated from
// WSDL.
type PobierzListeZlaUbezpieczonegoResponse struct {
	Rezultat                 *DocTypeRef_Rezultat                 `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	LiczbaWszystkichRekordow *int64                               `xml:"LiczbaWszystkichRekordow,omitempty" json:"LiczbaWszystkichRekordow,omitempty" yaml:"LiczbaWszystkichRekordow,omitempty"`
	ZaswiadczenieLekarskie   []*DocTypeRef_ZaswiadczenieLekarskie `xml:"ZaswiadczenieLekarskie,omitempty" json:"ZaswiadczenieLekarskie,omitempty" yaml:"ZaswiadczenieLekarskie,omitempty"`
}

// PobierzLiterowyKodChoroby was auto-generated from WSDL.
type PobierzLiterowyKodChoroby struct {
	IdSesji                  *ID                                     `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	DataPoczatkuNiezdolnosci *Data                                   `xml:"DataPoczatkuNiezdolnosci,omitempty" json:"DataPoczatkuNiezdolnosci,omitempty" yaml:"DataPoczatkuNiezdolnosci,omitempty"`
	StatystycznyKodChoroby   *StatystycznyKodChoroby                 `xml:"StatystycznyKodChoroby,omitempty" json:"StatystycznyKodChoroby,omitempty" yaml:"StatystycznyKodChoroby,omitempty"`
	Ubezpieczony             *DocTypeRef_IdentyfikacjaUbezpieczonego `xml:"Ubezpieczony,omitempty" json:"Ubezpieczony,omitempty" yaml:"Ubezpieczony,omitempty"`
}

// PobierzLiterowyKodChorobyResponse was auto-generated from WSDL.
type PobierzLiterowyKodChorobyResponse struct {
	Rezultat *DocTypeRef_Rezultat `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	KodA     *bool                `xml:"KodA,omitempty" json:"KodA,omitempty" yaml:"KodA,omitempty"`
	KodD     *bool                `xml:"KodD,omitempty" json:"KodD,omitempty" yaml:"KodD,omitempty"`
}

// PobierzMiejsceWykonywaniaZawodu was auto-generated from WSDL.
type PobierzMiejsceWykonywaniaZawodu struct {
	IdSesji *ID `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
}

// PobierzMiejsceWykonywaniaZawoduResponse was auto-generated from
// WSDL.
type PobierzMiejsceWykonywaniaZawoduResponse struct {
	Rezultat                 *DocTypeRef_Rezultat                        `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	MiejsceWykonywaniaZawodu []*DocTypeRef_kmdk_MiejsceWykonywaniaZawodu `xml:"MiejsceWykonywaniaZawodu,omitempty" json:"MiejsceWykonywaniaZawodu,omitempty" yaml:"MiejsceWykonywaniaZawodu,omitempty"`
}

// PobierzOpisChoroby was auto-generated from WSDL.
type PobierzOpisChoroby struct {
	IdSesji                *ID                     `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	StatystycznyKodChoroby *StatystycznyKodChoroby `xml:"StatystycznyKodChoroby,omitempty" json:"StatystycznyKodChoroby,omitempty" yaml:"StatystycznyKodChoroby,omitempty"`
}

// PobierzOpisChorobyResponse was auto-generated from WSDL.
type PobierzOpisChorobyResponse struct {
	Rezultat     *DocTypeRef_Rezultat `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	NazwaChoroby *string              `xml:"NazwaChoroby,omitempty" json:"NazwaChoroby,omitempty" yaml:"NazwaChoroby,omitempty"`
	OpisChoroby  *string              `xml:"OpisChoroby,omitempty" json:"OpisChoroby,omitempty" yaml:"OpisChoroby,omitempty"`
}

// PobierzOswiadczenie was auto-generated from WSDL.
type PobierzOswiadczenie struct {
}

// PobierzOswiadczenieResponse was auto-generated from WSDL.
type PobierzOswiadczenieResponse struct {
	Oswiadczenie *string `xml:"Oswiadczenie,omitempty" json:"Oswiadczenie,omitempty" yaml:"Oswiadczenie,omitempty"`
}

// PobierzPlatnikowUbezpieczonego was auto-generated from WSDL.
type PobierzPlatnikowUbezpieczonego struct {
	IdSesji      *ID                                     `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Ubezpieczony *DocTypeRef_IdentyfikacjaUbezpieczonego `xml:"Ubezpieczony,omitempty" json:"Ubezpieczony,omitempty" yaml:"Ubezpieczony,omitempty"`
}

// PobierzPlatnikowUbezpieczonegoResponse was auto-generated from
// WSDL.
type PobierzPlatnikowUbezpieczonegoResponse struct {
	Rezultat *DocTypeRef_Rezultat                 `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	Platnik  []*DocTypeRef_Platnik_Ubezpieczonego `xml:"Platnik,omitempty" json:"Platnik,omitempty" yaml:"Platnik,omitempty"`
}

// PobierzSlownikKodowPokrewienstwa was auto-generated from WSDL.
type PobierzSlownikKodowPokrewienstwa struct {
	IdSesji *ID `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
}

// PobierzSlownikKodowPokrewienstwaResponse was auto-generated
// from WSDL.
type PobierzSlownikKodowPokrewienstwaResponse struct {
	KodPokrewienstwa []*DocTypeRef_KodPokrewienstwa `xml:"KodPokrewienstwa,omitempty" json:"KodPokrewienstwa,omitempty" yaml:"KodPokrewienstwa,omitempty"`
	Rezultat         *DocTypeRef_Rezultat           `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// PobierzSlownikPrzyczynAnulowania was auto-generated from WSDL.
type PobierzSlownikPrzyczynAnulowania struct {
	IdSesji *ID `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
}

// PobierzSlownikPrzyczynAnulowaniaResponse was auto-generated
// from WSDL.
type PobierzSlownikPrzyczynAnulowaniaResponse struct {
	Przyczyna []*DocTypeRef_Przyczyny `xml:"Przyczyna,omitempty" json:"Przyczyna,omitempty" yaml:"Przyczyna,omitempty"`
	Rezultat  *DocTypeRef_Rezultat    `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// PobierzSlownikPrzyczynUniewaznienia was auto-generated from
// WSDL.
type PobierzSlownikPrzyczynUniewaznienia struct {
	IdSesji *ID `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
}

// PobierzSlownikPrzyczynUniewaznieniaResponse was auto-generated
// from WSDL.
type PobierzSlownikPrzyczynUniewaznieniaResponse struct {
	Przyczyna []*DocTypeRef_Przyczyny `xml:"Przyczyna,omitempty" json:"Przyczyna,omitempty" yaml:"Przyczyna,omitempty"`
	Rezultat  *DocTypeRef_Rezultat    `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// PobierzStatusyDrukowZla was auto-generated from WSDL.
type PobierzStatusyDrukowZla struct {
	IdSesji       *ID                            `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Stronicowanie *DocTypeRef_Stronicowanie      `xml:"Stronicowanie,omitempty" json:"Stronicowanie,omitempty" yaml:"Stronicowanie,omitempty"`
	Sortowanie    *DocTypeRef_DrukZlaSortowanie  `xml:"Sortowanie,omitempty" json:"Sortowanie,omitempty" yaml:"Sortowanie,omitempty"`
	Filtrowanie   *DocTypeRef_DrukZlaFiltrowanie `xml:"Filtrowanie,omitempty" json:"Filtrowanie,omitempty" yaml:"Filtrowanie,omitempty"`
	NipPlacowki   *NIP                           `xml:"NipPlacowki,omitempty" json:"NipPlacowki,omitempty" yaml:"NipPlacowki,omitempty"`
}

// PobierzStatusyDrukowZlaResponse was auto-generated from WSDL.
type PobierzStatusyDrukowZlaResponse struct {
	Rezultat                 *DocTypeRef_Rezultat      `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	StatusDruku              []*DocTypeRef_StatusDruku `xml:"StatusDruku,omitempty" json:"StatusDruku,omitempty" yaml:"StatusDruku,omitempty"`
	LiczbaWszystkichRekordow *string                   `xml:"LiczbaWszystkichRekordow,omitempty" json:"LiczbaWszystkichRekordow,omitempty" yaml:"LiczbaWszystkichRekordow,omitempty"`
}

// PobierzSzczegolyZlaBiezace was auto-generated from WSDL.
type PobierzSzczegolyZlaBiezace struct {
	IdSesji         *ID                       `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	SeriaNumerZla   *DocTypeRef_SeriaNumerZla `xml:"SeriaNumerZla,omitempty" json:"SeriaNumerZla,omitempty" yaml:"SeriaNumerZla,omitempty"`
	KontekstDostepu *KontekstDostepu          `xml:"KontekstDostepu,omitempty" json:"KontekstDostepu,omitempty" yaml:"KontekstDostepu,omitempty"`
}

// PobierzSzczegolyZlaBiezaceResponse was auto-generated from WSDL.
type PobierzSzczegolyZlaBiezaceResponse struct {
	Rezultat                      *DocTypeRef_Rezultat                        `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	ZaswiadczenieBiezaceSzczegoly *DocTypeRef_ZaswiadczenieLekarskieDanePelne `xml:"ZaswiadczenieBiezaceSzczegoly,omitempty" json:"ZaswiadczenieBiezaceSzczegoly,omitempty" yaml:"ZaswiadczenieBiezaceSzczegoly,omitempty"`
}

// PobierzSzczegolyZlaPrzetworzone was auto-generated from WSDL.
type PobierzSzczegolyZlaPrzetworzone struct {
	IdSesji         *ID                       `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	SeriaNumerZla   *DocTypeRef_SeriaNumerZla `xml:"SeriaNumerZla,omitempty" json:"SeriaNumerZla,omitempty" yaml:"SeriaNumerZla,omitempty"`
	KontekstDostepu *KontekstDostepu          `xml:"KontekstDostepu,omitempty" json:"KontekstDostepu,omitempty" yaml:"KontekstDostepu,omitempty"`
}

// PobierzSzczegolyZlaPrzetworzoneResponse was auto-generated from
// WSDL.
type PobierzSzczegolyZlaPrzetworzoneResponse struct {
	Rezultat                           *DocTypeRef_Rezultat                            `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	ZaswiadczeniePrzetworzoneSzczegoly *DocTypeRef_ZaswiadczenieLekarskieDanePelne     `xml:"ZaswiadczeniePrzetworzoneSzczegoly,omitempty" json:"ZaswiadczeniePrzetworzoneSzczegoly,omitempty" yaml:"ZaswiadczeniePrzetworzoneSzczegoly,omitempty"`
	ZaswiadczenieLekarskieKorygujace   *DocTypeRef_ZaswiadczenieLekarskieKsiKorygujace `xml:"ZaswiadczenieLekarskieKorygujace,omitempty" json:"ZaswiadczenieLekarskieKorygujace,omitempty" yaml:"ZaswiadczenieLekarskieKorygujace,omitempty"`
}

// PobierzUppDlaDokumentu was auto-generated from WSDL.
type PobierzUppDlaDokumentu struct {
	IdSesji         *ID                        `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	RodzajDokumentu *RodzajDokumentuEnumeracja `xml:"RodzajDokumentu,omitempty" json:"RodzajDokumentu,omitempty" yaml:"RodzajDokumentu,omitempty"`
	SeriaNumerZla   *DocTypeRef_SeriaNumerZla  `xml:"SeriaNumerZla,omitempty" json:"SeriaNumerZla,omitempty" yaml:"SeriaNumerZla,omitempty"`
}

// PobierzUppDlaDokumentuResponse was auto-generated from WSDL.
type PobierzUppDlaDokumentuResponse struct {
	Rezultat *DocTypeRef_Rezultat `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	UPP      *DocTypeRef_UPP      `xml:"UPP,omitempty" json:"UPP,omitempty" yaml:"UPP,omitempty"`
}

// PobierzUprawnieniaNaDzien was auto-generated from WSDL.
type PobierzUprawnieniaNaDzien struct {
	IdSesji *ID   `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Data    *Date `xml:"Data,omitempty" json:"Data,omitempty" yaml:"Data,omitempty"`
}

// PobierzUprawnieniaNaDzienResponse was auto-generated from WSDL.
type PobierzUprawnieniaNaDzienResponse struct {
	Rezultat                          *DocTypeRef_Rezultat `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	MozliwoscWystawianiaZaswiadczenia *bool                `xml:"MozliwoscWystawianiaZaswiadczenia,omitempty" json:"MozliwoscWystawianiaZaswiadczenia,omitempty" yaml:"MozliwoscWystawianiaZaswiadczenia,omitempty"`
}

// RezerwujSeriaNumerZla was auto-generated from WSDL.
type RezerwujSeriaNumerZla struct {
	IdSesji *ID   `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Liczba  *uint `xml:"Liczba,omitempty" json:"Liczba,omitempty" yaml:"Liczba,omitempty"`
}

// RezerwujSeriaNumerZlaResponse was auto-generated from WSDL.
type RezerwujSeriaNumerZlaResponse struct {
	Rezultat      *DocTypeRef_Rezultat        `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	SeriaNumerZla []*DocTypeRef_SeriaNumerZla `xml:"SeriaNumerZla,omitempty" json:"SeriaNumerZla,omitempty" yaml:"SeriaNumerZla,omitempty"`
}

// SprawdzMozliwoscAnulowania was auto-generated from WSDL.
type SprawdzMozliwoscAnulowania struct {
	IdSesji       *ID                         `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	SeriaNumerZla []*DocTypeRef_SeriaNumerZla `xml:"SeriaNumerZla,omitempty" json:"SeriaNumerZla,omitempty" yaml:"SeriaNumerZla,omitempty"`
}

// SprawdzMozliwoscAnulowaniaResponse was auto-generated from WSDL.
type SprawdzMozliwoscAnulowaniaResponse struct {
	Rezultat            *DocTypeRef_Rezultat              `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	MozliwoscAnulowania []*DocTypeRef_WeryfikacjaAkcjiZla `xml:"MozliwoscAnulowania,omitempty" json:"MozliwoscAnulowania,omitempty" yaml:"MozliwoscAnulowania,omitempty"`
}

// SprawdzMozliwoscElektronizacji was auto-generated from WSDL.
type SprawdzMozliwoscElektronizacji struct {
	IdSesji       *ID                         `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	SeriaNumerZla []*DocTypeRef_SeriaNumerZla `xml:"SeriaNumerZla,omitempty" json:"SeriaNumerZla,omitempty" yaml:"SeriaNumerZla,omitempty"`
}

// SprawdzMozliwoscElektronizacjiResponse was auto-generated from
// WSDL.
type SprawdzMozliwoscElektronizacjiResponse struct {
	Rezultat                *DocTypeRef_Rezultat              `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	MozliwoscElektronizacji []*DocTypeRef_WeryfikacjaAkcjiZla `xml:"MozliwoscElektronizacji,omitempty" json:"MozliwoscElektronizacji,omitempty" yaml:"MozliwoscElektronizacji,omitempty"`
}

// SprawdzMozliwoscUniewaznienia was auto-generated from WSDL.
type SprawdzMozliwoscUniewaznienia struct {
	IdSesji       *ID                         `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	SeriaNumerZla []*DocTypeRef_SeriaNumerZla `xml:"SeriaNumerZla,omitempty" json:"SeriaNumerZla,omitempty" yaml:"SeriaNumerZla,omitempty"`
}

// SprawdzMozliwoscUniewaznieniaResponse was auto-generated from
// WSDL.
type SprawdzMozliwoscUniewaznieniaResponse struct {
	Rezultat               *DocTypeRef_Rezultat              `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	MozliwoscUniewaznienia []*DocTypeRef_WeryfikacjaAkcjiZla `xml:"MozliwoscUniewaznienia,omitempty" json:"MozliwoscUniewaznienia,omitempty" yaml:"MozliwoscUniewaznienia,omitempty"`
}

// SprawdzProfilRehabilitacji was auto-generated from WSDL.
type SprawdzProfilRehabilitacji struct {
	IdSesji    *ID                     `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	KodChoroby *StatystycznyKodChoroby `xml:"KodChoroby,omitempty" json:"KodChoroby,omitempty" yaml:"KodChoroby,omitempty"`
}

// SprawdzProfilRehabilitacjiResponse was auto-generated from WSDL.
type SprawdzProfilRehabilitacjiResponse struct {
	Rezultat            *DocTypeRef_Rezultat `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	ProfilRehabilitacji *string              `xml:"ProfilRehabilitacji,omitempty" json:"ProfilRehabilitacji,omitempty" yaml:"ProfilRehabilitacji,omitempty"`
}

// UsunSesje was auto-generated from WSDL.
type UsunSesje struct {
	IdSesji *ID `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
}

// UsunSesjeResponse was auto-generated from WSDL.
type UsunSesjeResponse struct {
	Rezultat *DocTypeRef_Rezultat `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// WalidujDokumenty was auto-generated from WSDL.
type WalidujDokumenty struct {
	IdSesji         *ID                         `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Tryb            *TrybWystawianiaEnumeracja  `xml:"Tryb,omitempty" json:"Tryb,omitempty" yaml:"Tryb,omitempty"`
	ListaDokumentow *DocTypeRef_ListaDokumentow `xml:"ListaDokumentow,omitempty" json:"ListaDokumentow,omitempty" yaml:"ListaDokumentow,omitempty"`
	Dokument        *DocTypeRef_Dokument        `xml:"Dokument,omitempty" json:"Dokument,omitempty" yaml:"Dokument,omitempty"`
}

// WalidujDokumentyResponse was auto-generated from WSDL.
type WalidujDokumentyResponse struct {
	Rezultat          *DocTypeRef_Rezultat          `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	RezultatWalidacji *DocTypeRef_RezultatWalidacji `xml:"RezultatWalidacji,omitempty" json:"RezultatWalidacji,omitempty" yaml:"RezultatWalidacji,omitempty"`
}

// WalidujWniosek was auto-generated from WSDL.
type WalidujWniosek struct {
	IdSesji      *ID                      `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	DokumentXml  *string                  `xml:"DokumentXml,omitempty" json:"DokumentXml,omitempty" yaml:"DokumentXml,omitempty"`
	TypDokumentu *RodzajWnioskuEnumeracja `xml:"TypDokumentu,omitempty" json:"TypDokumentu,omitempty" yaml:"TypDokumentu,omitempty"`
}

// WalidujWniosekResponse was auto-generated from WSDL.
type WalidujWniosekResponse struct {
	Rezultat          *DocTypeRef_Rezultat          `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	RezultatWalidacji *DocTypeRef_RezultatWalidacji `xml:"RezultatWalidacji,omitempty" json:"RezultatWalidacji,omitempty" yaml:"RezultatWalidacji,omitempty"`
}

// WyslijDokumenty was auto-generated from WSDL.
type WyslijDokumenty struct {
	IdSesji  *ID                        `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Tryb     *TrybWystawianiaEnumeracja `xml:"Tryb,omitempty" json:"Tryb,omitempty" yaml:"Tryb,omitempty"`
	Dokument []*DocTypeRef_Dokument     `xml:"Dokument,omitempty" json:"Dokument,omitempty" yaml:"Dokument,omitempty"`
}

// WyslijDokumentyResponse was auto-generated from WSDL.
type WyslijDokumentyResponse struct {
	Rezultat          *DocTypeRef_Rezultat                      `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	RezultatWalidacji *DocTypeRef_RezultatWalidacji             `xml:"RezultatWalidacji,omitempty" json:"RezultatWalidacji,omitempty" yaml:"RezultatWalidacji,omitempty"`
	RezultatWysylki   []*DocTypeRef_kmdk_RezultatWysylkiPakietu `xml:"RezultatWysylki,omitempty" json:"RezultatWysylki,omitempty" yaml:"RezultatWysylki,omitempty"`
}

// WyslijWniosek was auto-generated from WSDL.
type WyslijWniosek struct {
	IdSesji      *ID                      `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	DokumentXml  *string                  `xml:"DokumentXml,omitempty" json:"DokumentXml,omitempty" yaml:"DokumentXml,omitempty"`
	TypDokumentu *RodzajWnioskuEnumeracja `xml:"TypDokumentu,omitempty" json:"TypDokumentu,omitempty" yaml:"TypDokumentu,omitempty"`
}

// WyslijWniosekResponse was auto-generated from WSDL.
type WyslijWniosekResponse struct {
	Rezultat          *DocTypeRef_Rezultat          `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
	RezultatWalidacji *DocTypeRef_RezultatWalidacji `xml:"RezultatWalidacji,omitempty" json:"RezultatWalidacji,omitempty" yaml:"RezultatWalidacji,omitempty"`
	WynikWyslania     *string                       `xml:"WynikWyslania,omitempty" json:"WynikWyslania,omitempty" yaml:"WynikWyslania,omitempty"`
	UtworzoneUPP      *DocTypeRef_UPP               `xml:"UtworzoneUPP,omitempty" json:"UtworzoneUPP,omitempty" yaml:"UtworzoneUPP,omitempty"`
}

// ZalogujPodpisem was auto-generated from WSDL.
type ZalogujPodpisem struct {
	PodpisaneOswiadczenie *string                 `xml:"PodpisaneOswiadczenie,omitempty" json:"PodpisaneOswiadczenie,omitempty" yaml:"PodpisaneOswiadczenie,omitempty"`
	MetodaWeryfikacji     *MetodaUwierzytelnienia `xml:"MetodaWeryfikacji,omitempty" json:"MetodaWeryfikacji,omitempty" yaml:"MetodaWeryfikacji,omitempty"`
}

// ZalogujPodpisemResponse was auto-generated from WSDL.
type ZalogujPodpisemResponse struct {
	IdSesji  *string              `xml:"IdSesji,omitempty" json:"IdSesji,omitempty" yaml:"IdSesji,omitempty"`
	Rezultat *DocTypeRef_Rezultat `xml:"Rezultat,omitempty" json:"Rezultat,omitempty" yaml:"Rezultat,omitempty"`
}

// Operation wrapper for NadajSeriaNumerZla.
// OperationZla_PortType_nadajSeriaNumerZla was auto-generated
// from WSDL.
type OperationZla_PortType_nadajSeriaNumerZla struct {
	Parameters *NadajSeriaNumerZla `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for NadajSeriaNumerZla.
// OperationZla_PortType_nadajSeriaNumerZlaResponse was auto-generated
// from WSDL.
type OperationZla_PortType_nadajSeriaNumerZlaResponse struct {
	Parameters *NadajSeriaNumerZlaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for NotyfikujAnulowanieZla.
// OperationZla_PortType_notyfikujAnulowanieZla was auto-generated
// from WSDL.
type OperationZla_PortType_notyfikujAnulowanieZla struct {
	Parameters *NotyfikujAnulowanieZla `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for NotyfikujAnulowanieZla.
// OperationZla_PortType_notyfikujAnulowanieZlaResponse was auto-generated
// from WSDL.
type OperationZla_PortType_notyfikujAnulowanieZlaResponse struct {
	Parameters *NotyfikujAnulowanieZlaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzAdresyPlacowek.
// OperationZla_PortType_pobierzAdresyPlacowek was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzAdresyPlacowek struct {
	Parameters *PobierzAdresyPlacowek `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzAdresyPlacowek.
// OperationZla_PortType_pobierzAdresyPlacowekResponse was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzAdresyPlacowekResponse struct {
	Parameters *PobierzAdresyPlacowekResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzCzlonkowRodzinyUbezpieczonego.
// OperationZla_PortType_pobierzCzlonkowRodzinyUbezpieczonego was
// auto-generated from WSDL.
type OperationZla_PortType_pobierzCzlonkowRodzinyUbezpieczonego struct {
	Parameters *PobierzCzlonkowRodzinyUbezpieczonego `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzCzlonkowRodzinyUbezpieczonego.
// OperationZla_PortType_pobierzCzlonkowRodzinyUbezpieczonegoResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzCzlonkowRodzinyUbezpieczonegoResponse struct {
	Parameters *PobierzCzlonkowRodzinyUbezpieczonegoResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzDaneLekarza.
// OperationZla_PortType_pobierzDaneLekarza was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzDaneLekarza struct {
	Parameters *PobierzDaneLekarza `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzDaneLekarza.
// OperationZla_PortType_pobierzDaneLekarzaResponse was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzDaneLekarzaResponse struct {
	Parameters *PobierzDaneLekarzaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzDanePlatnika.
// OperationZla_PortType_pobierzDanePlatnika was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzDanePlatnika struct {
	Parameters *PobierzDanePlatnika `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzDanePlatnika.
// OperationZla_PortType_pobierzDanePlatnikaResponse was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzDanePlatnikaResponse struct {
	Parameters *PobierzDanePlatnikaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzDaneUbezpieczonego.
// OperationZla_PortType_pobierzDaneUbezpieczonego was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzDaneUbezpieczonego struct {
	Parameters *PobierzDaneUbezpieczonego `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzDaneUbezpieczonego.
// OperationZla_PortType_pobierzDaneUbezpieczonegoResponse was
// auto-generated from WSDL.
type OperationZla_PortType_pobierzDaneUbezpieczonegoResponse struct {
	Parameters *PobierzDaneUbezpieczonegoResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzDokument.
// OperationZla_PortType_pobierzDokument was auto-generated from
// WSDL.
type OperationZla_PortType_pobierzDokument struct {
	Parameters *PobierzDokument `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzDokument.
// OperationZla_PortType_pobierzDokumentResponse was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzDokumentResponse struct {
	Parameters *PobierzDokumentResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzIdentyfikatorDokumentu.
// OperationZla_PortType_pobierzIdentyfikatorDokumentu was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzIdentyfikatorDokumentu struct {
	Parameters *PobierzIdentyfikatorDokumentu `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzIdentyfikatorDokumentu.
// OperationZla_PortType_pobierzIdentyfikatorDokumentuResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzIdentyfikatorDokumentuResponse struct {
	Parameters *PobierzIdentyfikatorDokumentuResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzKodChoroby.
// OperationZla_PortType_pobierzKodChoroby was auto-generated from
// WSDL.
type OperationZla_PortType_pobierzKodChoroby struct {
	Parameters *PobierzKodChoroby `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzKodChoroby.
// OperationZla_PortType_pobierzKodChorobyResponse was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzKodChorobyResponse struct {
	Parameters *PobierzKodChorobyResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzListeBiezacychZlaLekarza.
// OperationZla_PortType_pobierzListeBiezacychZlaLekarza was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzListeBiezacychZlaLekarza struct {
	Parameters *PobierzListeBiezacychZlaLekarza `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzListeBiezacychZlaLekarza.
// OperationZla_PortType_pobierzListeBiezacychZlaLekarzaResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzListeBiezacychZlaLekarzaResponse struct {
	Parameters *PobierzListeBiezacychZlaLekarzaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzListePowiazanychZla.
// OperationZla_PortType_pobierzListePowiazanychZla was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzListePowiazanychZla struct {
	Parameters *PobierzListePowiazanychZla `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzListePowiazanychZla.
// OperationZla_PortType_pobierzListePowiazanychZlaResponse was
// auto-generated from WSDL.
type OperationZla_PortType_pobierzListePowiazanychZlaResponse struct {
	Parameters *PobierzListePowiazanychZlaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzListePr4Ubezpieczonego.
// OperationZla_PortType_pobierzListePr4Ubezpieczonego was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzListePr4Ubezpieczonego struct {
	Parameters *PobierzListePr4Ubezpieczonego `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzListePr4Ubezpieczonego.
// OperationZla_PortType_pobierzListePr4UbezpieczonegoResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzListePr4UbezpieczonegoResponse struct {
	Parameters *PobierzListePr4UbezpieczonegoResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzListePrzetworzonychZlaLekarza.
// OperationZla_PortType_pobierzListePrzetworzonychZlaLekarza was
// auto-generated from WSDL.
type OperationZla_PortType_pobierzListePrzetworzonychZlaLekarza struct {
	Parameters *PobierzListePrzetworzonychZlaLekarza `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzListePrzetworzonychZlaLekarza.
// OperationZla_PortType_pobierzListePrzetworzonychZlaLekarzaResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzListePrzetworzonychZlaLekarzaResponse struct {
	Parameters *PobierzListePrzetworzonychZlaLekarzaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzListeZlaUbezpieczonego.
// OperationZla_PortType_pobierzListeZlaUbezpieczonego was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzListeZlaUbezpieczonego struct {
	Parameters *PobierzListeZlaUbezpieczonego `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzListeZlaUbezpieczonego.
// OperationZla_PortType_pobierzListeZlaUbezpieczonegoResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzListeZlaUbezpieczonegoResponse struct {
	Parameters *PobierzListeZlaUbezpieczonegoResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzLiterowyKodChoroby.
// OperationZla_PortType_pobierzLiterowyKodChoroby was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzLiterowyKodChoroby struct {
	Parameters *PobierzLiterowyKodChoroby `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzLiterowyKodChoroby.
// OperationZla_PortType_pobierzLiterowyKodChorobyResponse was
// auto-generated from WSDL.
type OperationZla_PortType_pobierzLiterowyKodChorobyResponse struct {
	Parameters *PobierzLiterowyKodChorobyResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzMiejsceWykonywaniaZawodu.
// OperationZla_PortType_pobierzMiejsceWykonywaniaZawodu was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzMiejsceWykonywaniaZawodu struct {
	Parameters *PobierzMiejsceWykonywaniaZawodu `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzMiejsceWykonywaniaZawodu.
// OperationZla_PortType_pobierzMiejsceWykonywaniaZawoduResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzMiejsceWykonywaniaZawoduResponse struct {
	Parameters *PobierzMiejsceWykonywaniaZawoduResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzOpisChoroby.
// OperationZla_PortType_pobierzOpisChoroby was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzOpisChoroby struct {
	Parameters *PobierzOpisChoroby `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzOpisChoroby.
// OperationZla_PortType_pobierzOpisChorobyResponse was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzOpisChorobyResponse struct {
	Parameters *PobierzOpisChorobyResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzOswiadczenie.
// OperationZla_PortType_pobierzOswiadczenie was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzOswiadczenie struct {
	Parameters *PobierzOswiadczenie `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzOswiadczenie.
// OperationZla_PortType_pobierzOswiadczenieResponse was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzOswiadczenieResponse struct {
	Parameters *PobierzOswiadczenieResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzPlatnikowUbezpieczonego.
// OperationZla_PortType_pobierzPlatnikowUbezpieczonego was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzPlatnikowUbezpieczonego struct {
	Parameters *PobierzPlatnikowUbezpieczonego `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzPlatnikowUbezpieczonego.
// OperationZla_PortType_pobierzPlatnikowUbezpieczonegoResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzPlatnikowUbezpieczonegoResponse struct {
	Parameters *PobierzPlatnikowUbezpieczonegoResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzSlownikKodowPokrewienstwa.
// OperationZla_PortType_pobierzSlownikKodowPokrewienstwa was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzSlownikKodowPokrewienstwa struct {
	Parameters *PobierzSlownikKodowPokrewienstwa `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzSlownikKodowPokrewienstwa.
// OperationZla_PortType_pobierzSlownikKodowPokrewienstwaResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzSlownikKodowPokrewienstwaResponse struct {
	Parameters *PobierzSlownikKodowPokrewienstwaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzSlownikPrzyczynAnulowania.
// OperationZla_PortType_pobierzSlownikPrzyczynAnulowania was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzSlownikPrzyczynAnulowania struct {
	Parameters *PobierzSlownikPrzyczynAnulowania `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzSlownikPrzyczynAnulowania.
// OperationZla_PortType_pobierzSlownikPrzyczynAnulowaniaResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzSlownikPrzyczynAnulowaniaResponse struct {
	Parameters *PobierzSlownikPrzyczynAnulowaniaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzSlownikPrzyczynUniewaznienia.
// OperationZla_PortType_pobierzSlownikPrzyczynUniewaznienia was
// auto-generated from WSDL.
type OperationZla_PortType_pobierzSlownikPrzyczynUniewaznienia struct {
	Parameters *PobierzSlownikPrzyczynUniewaznienia `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzSlownikPrzyczynUniewaznienia.
// OperationZla_PortType_pobierzSlownikPrzyczynUniewaznieniaResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzSlownikPrzyczynUniewaznieniaResponse struct {
	Parameters *PobierzSlownikPrzyczynUniewaznieniaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzStatusyDrukowZla.
// OperationZla_PortType_pobierzStatusyDrukowZla was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzStatusyDrukowZla struct {
	Parameters *PobierzStatusyDrukowZla `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzStatusyDrukowZla.
// OperationZla_PortType_pobierzStatusyDrukowZlaResponse was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzStatusyDrukowZlaResponse struct {
	Parameters *PobierzStatusyDrukowZlaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzSzczegolyZlaBiezace.
// OperationZla_PortType_pobierzSzczegolyZlaBiezace was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzSzczegolyZlaBiezace struct {
	Parameters *PobierzSzczegolyZlaBiezace `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzSzczegolyZlaBiezace.
// OperationZla_PortType_pobierzSzczegolyZlaBiezaceResponse was
// auto-generated from WSDL.
type OperationZla_PortType_pobierzSzczegolyZlaBiezaceResponse struct {
	Parameters *PobierzSzczegolyZlaBiezaceResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzSzczegolyZlaPrzetworzone.
// OperationZla_PortType_pobierzSzczegolyZlaPrzetworzone was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzSzczegolyZlaPrzetworzone struct {
	Parameters *PobierzSzczegolyZlaPrzetworzone `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzSzczegolyZlaPrzetworzone.
// OperationZla_PortType_pobierzSzczegolyZlaPrzetworzoneResponse
// was auto-generated from WSDL.
type OperationZla_PortType_pobierzSzczegolyZlaPrzetworzoneResponse struct {
	Parameters *PobierzSzczegolyZlaPrzetworzoneResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzUppDlaDokumentu.
// OperationZla_PortType_pobierzUppDlaDokumentu was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzUppDlaDokumentu struct {
	Parameters *PobierzUppDlaDokumentu `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzUppDlaDokumentu.
// OperationZla_PortType_pobierzUppDlaDokumentuResponse was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzUppDlaDokumentuResponse struct {
	Parameters *PobierzUppDlaDokumentuResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzUprawnieniaNaDzien.
// OperationZla_PortType_pobierzUprawnieniaNaDzien was auto-generated
// from WSDL.
type OperationZla_PortType_pobierzUprawnieniaNaDzien struct {
	Parameters *PobierzUprawnieniaNaDzien `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for PobierzUprawnieniaNaDzien.
// OperationZla_PortType_pobierzUprawnieniaNaDzienResponse was
// auto-generated from WSDL.
type OperationZla_PortType_pobierzUprawnieniaNaDzienResponse struct {
	Parameters *PobierzUprawnieniaNaDzienResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for RezerwujSeriaNumerZla.
// OperationZla_PortType_rezerwujSeriaNumerZla was auto-generated
// from WSDL.
type OperationZla_PortType_rezerwujSeriaNumerZla struct {
	Parameters *RezerwujSeriaNumerZla `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for RezerwujSeriaNumerZla.
// OperationZla_PortType_rezerwujSeriaNumerZlaResponse was auto-generated
// from WSDL.
type OperationZla_PortType_rezerwujSeriaNumerZlaResponse struct {
	Parameters *RezerwujSeriaNumerZlaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for SprawdzMozliwoscAnulowania.
// OperationZla_PortType_sprawdzMozliwoscAnulowania was auto-generated
// from WSDL.
type OperationZla_PortType_sprawdzMozliwoscAnulowania struct {
	Parameters *SprawdzMozliwoscAnulowania `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for SprawdzMozliwoscAnulowania.
// OperationZla_PortType_sprawdzMozliwoscAnulowaniaResponse was
// auto-generated from WSDL.
type OperationZla_PortType_sprawdzMozliwoscAnulowaniaResponse struct {
	Parameters *SprawdzMozliwoscAnulowaniaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for SprawdzMozliwoscElektronizacji.
// OperationZla_PortType_sprawdzMozliwoscElektronizacji was auto-generated
// from WSDL.
type OperationZla_PortType_sprawdzMozliwoscElektronizacji struct {
	Parameters *SprawdzMozliwoscElektronizacji `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for SprawdzMozliwoscElektronizacji.
// OperationZla_PortType_sprawdzMozliwoscElektronizacjiResponse
// was auto-generated from WSDL.
type OperationZla_PortType_sprawdzMozliwoscElektronizacjiResponse struct {
	Parameters *SprawdzMozliwoscElektronizacjiResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for SprawdzMozliwoscUniewaznienia.
// OperationZla_PortType_sprawdzMozliwoscUniewaznienia was auto-generated
// from WSDL.
type OperationZla_PortType_sprawdzMozliwoscUniewaznienia struct {
	Parameters *SprawdzMozliwoscUniewaznienia `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for SprawdzMozliwoscUniewaznienia.
// OperationZla_PortType_sprawdzMozliwoscUniewaznieniaResponse
// was auto-generated from WSDL.
type OperationZla_PortType_sprawdzMozliwoscUniewaznieniaResponse struct {
	Parameters *SprawdzMozliwoscUniewaznieniaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for SprawdzProfilRehabilitacji.
// OperationZla_PortType_sprawdzProfilRehabilitacji was auto-generated
// from WSDL.
type OperationZla_PortType_sprawdzProfilRehabilitacji struct {
	Parameters *SprawdzProfilRehabilitacji `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for SprawdzProfilRehabilitacji.
// OperationZla_PortType_sprawdzProfilRehabilitacjiResponse was
// auto-generated from WSDL.
type OperationZla_PortType_sprawdzProfilRehabilitacjiResponse struct {
	Parameters *SprawdzProfilRehabilitacjiResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for UsunSesje.
// OperationZla_PortType_usunSesje was auto-generated from WSDL.
type OperationZla_PortType_usunSesje struct {
	Parameters *UsunSesje `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for UsunSesje.
// OperationZla_PortType_usunSesjeResponse was auto-generated from
// WSDL.
type OperationZla_PortType_usunSesjeResponse struct {
	Parameters *UsunSesjeResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for WalidujDokumenty.
// OperationZla_PortType_walidujDokumenty was auto-generated from
// WSDL.
type OperationZla_PortType_walidujDokumenty struct {
	Parameters *WalidujDokumenty `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for WalidujDokumenty.
// OperationZla_PortType_walidujDokumentyResponse was auto-generated
// from WSDL.
type OperationZla_PortType_walidujDokumentyResponse struct {
	Parameters *WalidujDokumentyResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for WalidujWniosek.
// OperationZla_PortType_walidujWniosek was auto-generated from
// WSDL.
type OperationZla_PortType_walidujWniosek struct {
	Parameters *WalidujWniosek `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for WalidujWniosek.
// OperationZla_PortType_walidujWniosekResponse was auto-generated
// from WSDL.
type OperationZla_PortType_walidujWniosekResponse struct {
	Parameters *WalidujWniosekResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for WyslijDokumenty.
// OperationZla_PortType_wyslijDokumenty was auto-generated from
// WSDL.
type OperationZla_PortType_wyslijDokumenty struct {
	Parameters *WyslijDokumenty `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for WyslijDokumenty.
// OperationZla_PortType_wyslijDokumentyResponse was auto-generated
// from WSDL.
type OperationZla_PortType_wyslijDokumentyResponse struct {
	Parameters *WyslijDokumentyResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for WyslijWniosek.
// OperationZla_PortType_wyslijWniosek was auto-generated from
// WSDL.
type OperationZla_PortType_wyslijWniosek struct {
	Parameters *WyslijWniosek `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for WyslijWniosek.
// OperationZla_PortType_wyslijWniosekResponse was auto-generated
// from WSDL.
type OperationZla_PortType_wyslijWniosekResponse struct {
	Parameters *WyslijWniosekResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for ZalogujPodpisem.
// OperationZla_PortType_zalogujPodpisem was auto-generated from
// WSDL.
type OperationZla_PortType_zalogujPodpisem struct {
	Parameters *ZalogujPodpisem `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for ZalogujPodpisem.
// OperationZla_PortType_zalogujPodpisemResponse was auto-generated
// from WSDL.
type OperationZla_PortType_zalogujPodpisemResponse struct {
	Parameters *ZalogujPodpisemResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// zla_PortType implements the Zla_PortType interface.
type zla_PortType struct {
	cli *soap.Client
}

// NadajSeriaNumerZla was auto-generated from WSDL.
func (p *zla_PortType) NadajSeriaNumerZla(parameters *NadajSeriaNumerZla) (*NadajSeriaNumerZlaResponse, error) {
	α := struct {
		OperationZla_PortType_nadajSeriaNumerZla `xml:"tns:nadajSeriaNumerZla"`
	}{
		OperationZla_PortType_nadajSeriaNumerZla{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_nadajSeriaNumerZlaResponse `xml:"nadajSeriaNumerZlaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_nadajSeriaNumerZla", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// NotyfikujAnulowanieZla was auto-generated from WSDL.
func (p *zla_PortType) NotyfikujAnulowanieZla(parameters *NotyfikujAnulowanieZla) (*NotyfikujAnulowanieZlaResponse, error) {
	α := struct {
		OperationZla_PortType_notyfikujAnulowanieZla `xml:"tns:notyfikujAnulowanieZla"`
	}{
		OperationZla_PortType_notyfikujAnulowanieZla{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_notyfikujAnulowanieZlaResponse `xml:"notyfikujAnulowanieZlaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_notyfikujAnulowanieZla", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzAdresyPlacowek was auto-generated from WSDL.
func (p *zla_PortType) PobierzAdresyPlacowek(parameters *PobierzAdresyPlacowek) (*PobierzAdresyPlacowekResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzAdresyPlacowek `xml:"tns:pobierzAdresyPlacowek"`
	}{
		OperationZla_PortType_pobierzAdresyPlacowek{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzAdresyPlacowekResponse `xml:"pobierzAdresyPlacowekResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzAdresyPlacowek", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzCzlonkowRodzinyUbezpieczonego was auto-generated from
// WSDL.
func (p *zla_PortType) PobierzCzlonkowRodzinyUbezpieczonego(parameters *PobierzCzlonkowRodzinyUbezpieczonego) (*PobierzCzlonkowRodzinyUbezpieczonegoResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzCzlonkowRodzinyUbezpieczonego `xml:"tns:pobierzCzlonkowRodzinyUbezpieczonego"`
	}{
		OperationZla_PortType_pobierzCzlonkowRodzinyUbezpieczonego{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzCzlonkowRodzinyUbezpieczonegoResponse `xml:"pobierzCzlonkowRodzinyUbezpieczonegoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzCzlonkowRodzinyUbezpieczonego", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzDaneLekarza was auto-generated from WSDL.
func (p *zla_PortType) PobierzDaneLekarza(parameters *PobierzDaneLekarza) (*PobierzDaneLekarzaResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzDaneLekarza `xml:"tns:pobierzDaneLekarza"`
	}{
		OperationZla_PortType_pobierzDaneLekarza{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzDaneLekarzaResponse `xml:"pobierzDaneLekarzaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzDaneLekarza", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzDanePlatnika was auto-generated from WSDL.
func (p *zla_PortType) PobierzDanePlatnika(parameters *PobierzDanePlatnika) (*PobierzDanePlatnikaResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzDanePlatnika `xml:"tns:pobierzDanePlatnika"`
	}{
		OperationZla_PortType_pobierzDanePlatnika{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzDanePlatnikaResponse `xml:"pobierzDanePlatnikaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzDanePlatnika", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzDaneUbezpieczonego was auto-generated from WSDL.
func (p *zla_PortType) PobierzDaneUbezpieczonego(parameters *PobierzDaneUbezpieczonego) (*PobierzDaneUbezpieczonegoResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzDaneUbezpieczonego `xml:"tns:pobierzDaneUbezpieczonego"`
	}{
		OperationZla_PortType_pobierzDaneUbezpieczonego{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzDaneUbezpieczonegoResponse `xml:"pobierzDaneUbezpieczonegoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzDaneUbezpieczonego", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzDokument was auto-generated from WSDL.
func (p *zla_PortType) PobierzDokument(parameters *PobierzDokument) (*PobierzDokumentResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzDokument `xml:"tns:pobierzDokument"`
	}{
		OperationZla_PortType_pobierzDokument{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzDokumentResponse `xml:"pobierzDokumentResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzDokument", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzIdentyfikatorDokumentu was auto-generated from WSDL.
func (p *zla_PortType) PobierzIdentyfikatorDokumentu(parameters *PobierzIdentyfikatorDokumentu) (*PobierzIdentyfikatorDokumentuResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzIdentyfikatorDokumentu `xml:"tns:pobierzIdentyfikatorDokumentu"`
	}{
		OperationZla_PortType_pobierzIdentyfikatorDokumentu{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzIdentyfikatorDokumentuResponse `xml:"pobierzIdentyfikatorDokumentuResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzIdentyfikatorDokumentu", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzKodChoroby was auto-generated from WSDL.
func (p *zla_PortType) PobierzKodChoroby(parameters *PobierzKodChoroby) (*PobierzKodChorobyResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzKodChoroby `xml:"tns:pobierzKodChoroby"`
	}{
		OperationZla_PortType_pobierzKodChoroby{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzKodChorobyResponse `xml:"pobierzKodChorobyResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzKodChoroby", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzListeBiezacychZlaLekarza was auto-generated from WSDL.
func (p *zla_PortType) PobierzListeBiezacychZlaLekarza(parameters *PobierzListeBiezacychZlaLekarza) (*PobierzListeBiezacychZlaLekarzaResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzListeBiezacychZlaLekarza `xml:"tns:pobierzListeBiezacychZlaLekarza"`
	}{
		OperationZla_PortType_pobierzListeBiezacychZlaLekarza{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzListeBiezacychZlaLekarzaResponse `xml:"pobierzListeBiezacychZlaLekarzaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzListeBiezacychZlaLekarza", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzListePowiazanychZla was auto-generated from WSDL.
func (p *zla_PortType) PobierzListePowiazanychZla(parameters *PobierzListePowiazanychZla) (*PobierzListePowiazanychZlaResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzListePowiazanychZla `xml:"tns:pobierzListePowiazanychZla"`
	}{
		OperationZla_PortType_pobierzListePowiazanychZla{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzListePowiazanychZlaResponse `xml:"pobierzListePowiazanychZlaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzListePowiazanychZla", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzListePr4Ubezpieczonego was auto-generated from WSDL.
func (p *zla_PortType) PobierzListePr4Ubezpieczonego(parameters *PobierzListePr4Ubezpieczonego) (*PobierzListePr4UbezpieczonegoResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzListePr4Ubezpieczonego `xml:"tns:pobierzListePr4Ubezpieczonego"`
	}{
		OperationZla_PortType_pobierzListePr4Ubezpieczonego{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzListePr4UbezpieczonegoResponse `xml:"pobierzListePr4UbezpieczonegoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzListePr4Ubezpieczonego", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzListePrzetworzonychZlaLekarza was auto-generated from
// WSDL.
func (p *zla_PortType) PobierzListePrzetworzonychZlaLekarza(parameters *PobierzListePrzetworzonychZlaLekarza) (*PobierzListePrzetworzonychZlaLekarzaResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzListePrzetworzonychZlaLekarza `xml:"tns:pobierzListePrzetworzonychZlaLekarza"`
	}{
		OperationZla_PortType_pobierzListePrzetworzonychZlaLekarza{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzListePrzetworzonychZlaLekarzaResponse `xml:"pobierzListePrzetworzonychZlaLekarzaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzListePrzetworzonychZlaLekarza", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzListeZlaUbezpieczonego was auto-generated from WSDL.
func (p *zla_PortType) PobierzListeZlaUbezpieczonego(parameters *PobierzListeZlaUbezpieczonego) (*PobierzListeZlaUbezpieczonegoResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzListeZlaUbezpieczonego `xml:"tns:pobierzListeZlaUbezpieczonego"`
	}{
		OperationZla_PortType_pobierzListeZlaUbezpieczonego{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzListeZlaUbezpieczonegoResponse `xml:"pobierzListeZlaUbezpieczonegoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzListeZlaUbezpieczonego", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzLiterowyKodChoroby was auto-generated from WSDL.
func (p *zla_PortType) PobierzLiterowyKodChoroby(parameters *PobierzLiterowyKodChoroby) (*PobierzLiterowyKodChorobyResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzLiterowyKodChoroby `xml:"tns:pobierzLiterowyKodChoroby"`
	}{
		OperationZla_PortType_pobierzLiterowyKodChoroby{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzLiterowyKodChorobyResponse `xml:"pobierzLiterowyKodChorobyResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzLiterowyKodChoroby", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzMiejsceWykonywaniaZawodu was auto-generated from WSDL.
func (p *zla_PortType) PobierzMiejsceWykonywaniaZawodu(parameters *PobierzMiejsceWykonywaniaZawodu) (*PobierzMiejsceWykonywaniaZawoduResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzMiejsceWykonywaniaZawodu `xml:"tns:pobierzMiejsceWykonywaniaZawodu"`
	}{
		OperationZla_PortType_pobierzMiejsceWykonywaniaZawodu{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzMiejsceWykonywaniaZawoduResponse `xml:"pobierzMiejsceWykonywaniaZawoduResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzMiejsceWykonywaniaZawodu", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzOpisChoroby was auto-generated from WSDL.
func (p *zla_PortType) PobierzOpisChoroby(parameters *PobierzOpisChoroby) (*PobierzOpisChorobyResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzOpisChoroby `xml:"tns:pobierzOpisChoroby"`
	}{
		OperationZla_PortType_pobierzOpisChoroby{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzOpisChorobyResponse `xml:"pobierzOpisChorobyResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzOpisChoroby", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzOswiadczenie was auto-generated from WSDL.
func (p *zla_PortType) PobierzOswiadczenie(parameters *PobierzOswiadczenie) (*PobierzOswiadczenieResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzOswiadczenie `xml:"tns:pobierzOswiadczenie"`
	}{
		OperationZla_PortType_pobierzOswiadczenie{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzOswiadczenieResponse `xml:"pobierzOswiadczenieResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzOswiadczenie", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzPlatnikowUbezpieczonego was auto-generated from WSDL.
func (p *zla_PortType) PobierzPlatnikowUbezpieczonego(parameters *PobierzPlatnikowUbezpieczonego) (*PobierzPlatnikowUbezpieczonegoResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzPlatnikowUbezpieczonego `xml:"tns:pobierzPlatnikowUbezpieczonego"`
	}{
		OperationZla_PortType_pobierzPlatnikowUbezpieczonego{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzPlatnikowUbezpieczonegoResponse `xml:"pobierzPlatnikowUbezpieczonegoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzPlatnikowUbezpieczonego", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzSlownikKodowPokrewienstwa was auto-generated from WSDL.
func (p *zla_PortType) PobierzSlownikKodowPokrewienstwa(parameters *PobierzSlownikKodowPokrewienstwa) (*PobierzSlownikKodowPokrewienstwaResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzSlownikKodowPokrewienstwa `xml:"tns:pobierzSlownikKodowPokrewienstwa"`
	}{
		OperationZla_PortType_pobierzSlownikKodowPokrewienstwa{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzSlownikKodowPokrewienstwaResponse `xml:"pobierzSlownikKodowPokrewienstwaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzSlownikKodowPokrewienstwa", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzSlownikPrzyczynAnulowania was auto-generated from WSDL.
func (p *zla_PortType) PobierzSlownikPrzyczynAnulowania(parameters *PobierzSlownikPrzyczynAnulowania) (*PobierzSlownikPrzyczynAnulowaniaResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzSlownikPrzyczynAnulowania `xml:"tns:pobierzSlownikPrzyczynAnulowania"`
	}{
		OperationZla_PortType_pobierzSlownikPrzyczynAnulowania{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzSlownikPrzyczynAnulowaniaResponse `xml:"pobierzSlownikPrzyczynAnulowaniaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzSlownikPrzyczynAnulowania", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzSlownikPrzyczynUniewaznienia was auto-generated from
// WSDL.
func (p *zla_PortType) PobierzSlownikPrzyczynUniewaznienia(parameters *PobierzSlownikPrzyczynUniewaznienia) (*PobierzSlownikPrzyczynUniewaznieniaResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzSlownikPrzyczynUniewaznienia `xml:"tns:pobierzSlownikPrzyczynUniewaznienia"`
	}{
		OperationZla_PortType_pobierzSlownikPrzyczynUniewaznienia{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzSlownikPrzyczynUniewaznieniaResponse `xml:"pobierzSlownikPrzyczynUniewaznieniaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzSlownikPrzyczynUniewaznienia", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzStatusyDrukowZla was auto-generated from WSDL.
func (p *zla_PortType) PobierzStatusyDrukowZla(parameters *PobierzStatusyDrukowZla) (*PobierzStatusyDrukowZlaResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzStatusyDrukowZla `xml:"tns:pobierzStatusyDrukowZla"`
	}{
		OperationZla_PortType_pobierzStatusyDrukowZla{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzStatusyDrukowZlaResponse `xml:"pobierzStatusyDrukowZlaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzStatusyDrukowZla", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzSzczegolyZlaBiezace was auto-generated from WSDL.
func (p *zla_PortType) PobierzSzczegolyZlaBiezace(parameters *PobierzSzczegolyZlaBiezace) (*PobierzSzczegolyZlaBiezaceResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzSzczegolyZlaBiezace `xml:"tns:pobierzSzczegolyZlaBiezace"`
	}{
		OperationZla_PortType_pobierzSzczegolyZlaBiezace{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzSzczegolyZlaBiezaceResponse `xml:"pobierzSzczegolyZlaBiezaceResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzSzczegolyZlaBiezace", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzSzczegolyZlaPrzetworzone was auto-generated from WSDL.
func (p *zla_PortType) PobierzSzczegolyZlaPrzetworzone(parameters *PobierzSzczegolyZlaPrzetworzone) (*PobierzSzczegolyZlaPrzetworzoneResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzSzczegolyZlaPrzetworzone `xml:"tns:pobierzSzczegolyZlaPrzetworzone"`
	}{
		OperationZla_PortType_pobierzSzczegolyZlaPrzetworzone{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzSzczegolyZlaPrzetworzoneResponse `xml:"pobierzSzczegolyZlaPrzetworzoneResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzSzczegolyZlaPrzetworzone", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzUppDlaDokumentu was auto-generated from WSDL.
func (p *zla_PortType) PobierzUppDlaDokumentu(parameters *PobierzUppDlaDokumentu) (*PobierzUppDlaDokumentuResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzUppDlaDokumentu `xml:"tns:pobierzUppDlaDokumentu"`
	}{
		OperationZla_PortType_pobierzUppDlaDokumentu{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzUppDlaDokumentuResponse `xml:"pobierzUppDlaDokumentuResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzUppDlaDokumentu", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// PobierzUprawnieniaNaDzien was auto-generated from WSDL.
func (p *zla_PortType) PobierzUprawnieniaNaDzien(parameters *PobierzUprawnieniaNaDzien) (*PobierzUprawnieniaNaDzienResponse, error) {
	α := struct {
		OperationZla_PortType_pobierzUprawnieniaNaDzien `xml:"tns:pobierzUprawnieniaNaDzien"`
	}{
		OperationZla_PortType_pobierzUprawnieniaNaDzien{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_pobierzUprawnieniaNaDzienResponse `xml:"pobierzUprawnieniaNaDzienResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_pobierzUprawnieniaNaDzien", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// RezerwujSeriaNumerZla was auto-generated from WSDL.
func (p *zla_PortType) RezerwujSeriaNumerZla(parameters *RezerwujSeriaNumerZla) (*RezerwujSeriaNumerZlaResponse, error) {
	α := struct {
		OperationZla_PortType_rezerwujSeriaNumerZla `xml:"tns:rezerwujSeriaNumerZla"`
	}{
		OperationZla_PortType_rezerwujSeriaNumerZla{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_rezerwujSeriaNumerZlaResponse `xml:"rezerwujSeriaNumerZlaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_rezerwujSeriaNumerZla", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// SprawdzMozliwoscAnulowania was auto-generated from WSDL.
func (p *zla_PortType) SprawdzMozliwoscAnulowania(parameters *SprawdzMozliwoscAnulowania) (*SprawdzMozliwoscAnulowaniaResponse, error) {
	α := struct {
		OperationZla_PortType_sprawdzMozliwoscAnulowania `xml:"tns:sprawdzMozliwoscAnulowania"`
	}{
		OperationZla_PortType_sprawdzMozliwoscAnulowania{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_sprawdzMozliwoscAnulowaniaResponse `xml:"sprawdzMozliwoscAnulowaniaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_sprawdzMozliwoscAnulowania", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// SprawdzMozliwoscElektronizacji was auto-generated from WSDL.
func (p *zla_PortType) SprawdzMozliwoscElektronizacji(parameters *SprawdzMozliwoscElektronizacji) (*SprawdzMozliwoscElektronizacjiResponse, error) {
	α := struct {
		OperationZla_PortType_sprawdzMozliwoscElektronizacji `xml:"tns:sprawdzMozliwoscElektronizacji"`
	}{
		OperationZla_PortType_sprawdzMozliwoscElektronizacji{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_sprawdzMozliwoscElektronizacjiResponse `xml:"sprawdzMozliwoscElektronizacjiResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_sprawdzMozliwoscElektronizacji", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// SprawdzMozliwoscUniewaznienia was auto-generated from WSDL.
func (p *zla_PortType) SprawdzMozliwoscUniewaznienia(parameters *SprawdzMozliwoscUniewaznienia) (*SprawdzMozliwoscUniewaznieniaResponse, error) {
	α := struct {
		OperationZla_PortType_sprawdzMozliwoscUniewaznienia `xml:"tns:sprawdzMozliwoscUniewaznienia"`
	}{
		OperationZla_PortType_sprawdzMozliwoscUniewaznienia{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_sprawdzMozliwoscUniewaznieniaResponse `xml:"sprawdzMozliwoscUniewaznieniaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_sprawdzMozliwoscUniewaznienia", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// SprawdzProfilRehabilitacji was auto-generated from WSDL.
func (p *zla_PortType) SprawdzProfilRehabilitacji(parameters *SprawdzProfilRehabilitacji) (*SprawdzProfilRehabilitacjiResponse, error) {
	α := struct {
		OperationZla_PortType_sprawdzProfilRehabilitacji `xml:"tns:sprawdzProfilRehabilitacji"`
	}{
		OperationZla_PortType_sprawdzProfilRehabilitacji{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_sprawdzProfilRehabilitacjiResponse `xml:"sprawdzProfilRehabilitacjiResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_sprawdzProfilRehabilitacji", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// UsunSesje was auto-generated from WSDL.
func (p *zla_PortType) UsunSesje(parameters *UsunSesje) (*UsunSesjeResponse, error) {
	α := struct {
		OperationZla_PortType_usunSesje `xml:"tns:usunSesje"`
	}{
		OperationZla_PortType_usunSesje{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_usunSesjeResponse `xml:"usunSesjeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_usunSesje", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// WalidujDokumenty was auto-generated from WSDL.
func (p *zla_PortType) WalidujDokumenty(parameters *WalidujDokumenty) (*WalidujDokumentyResponse, error) {
	α := struct {
		OperationZla_PortType_walidujDokumenty `xml:"tns:walidujDokumenty"`
	}{
		OperationZla_PortType_walidujDokumenty{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_walidujDokumentyResponse `xml:"walidujDokumentyResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_walidujDokumenty", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// WalidujWniosek was auto-generated from WSDL.
func (p *zla_PortType) WalidujWniosek(parameters *WalidujWniosek) (*WalidujWniosekResponse, error) {
	α := struct {
		OperationZla_PortType_walidujWniosek `xml:"tns:walidujWniosek"`
	}{
		OperationZla_PortType_walidujWniosek{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_walidujWniosekResponse `xml:"walidujWniosekResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_walidujWniosek", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// WyslijDokumenty was auto-generated from WSDL.
func (p *zla_PortType) WyslijDokumenty(parameters *WyslijDokumenty) (*WyslijDokumentyResponse, error) {
	α := struct {
		OperationZla_PortType_wyslijDokumenty `xml:"tns:wyslijDokumenty"`
	}{
		OperationZla_PortType_wyslijDokumenty{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_wyslijDokumentyResponse `xml:"wyslijDokumentyResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_wyslijDokumenty", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// WyslijWniosek was auto-generated from WSDL.
func (p *zla_PortType) WyslijWniosek(parameters *WyslijWniosek) (*WyslijWniosekResponse, error) {
	α := struct {
		OperationZla_PortType_wyslijWniosek `xml:"tns:wyslijWniosek"`
	}{
		OperationZla_PortType_wyslijWniosek{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_wyslijWniosekResponse `xml:"wyslijWniosekResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_wyslijWniosek", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// ZalogujPodpisem was auto-generated from WSDL.
func (p *zla_PortType) ZalogujPodpisem(parameters *ZalogujPodpisem) (*ZalogujPodpisemResponse, error) {
	α := struct {
		OperationZla_PortType_zalogujPodpisem `xml:"tns:zalogujPodpisem"`
	}{
		OperationZla_PortType_zalogujPodpisem{
			parameters,
		},
	}

	γ := struct {
		OperationZla_PortType_zalogujPodpisemResponse `xml:"zalogujPodpisemResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("zus_channel_zla_Binder_zalogujPodpisem", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}
