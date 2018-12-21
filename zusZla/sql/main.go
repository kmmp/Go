package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "github.com/a-palchikov/sqlago"
)

type ezla struct {
	IDEzwolnienia            int            `json:"id_ezwolnienia"`
	PeselPacjenta            sql.NullInt64  `json:"pesel_pacjenta"`
	PaszportPacjenta         sql.NullString `json:"paszport_pacjenta"`
	DataUrodzeniaPacjenta    sql.NullString `json:"data_urodzenia_pacjenta"`
	ImiePacjenta             sql.NullString `json:"imie_pacjenta"`
	NazwiskoPacjenta         sql.NullString `json:"nazwisko_pacjenta"`
	Ubezpieczajacy           sql.NullInt64  `json:"ubezpieczajacy"`
	KodPocztowyPacjenta      sql.NullString `json:"kod_pocztowy_pacjenta"`
	MiejscowoscPacjenta      sql.NullString `json:"miejscowosc_pacjenta"`
	UlicaPacjenta            sql.NullString `json:"ulica_pacjenta"`
	NrDomuPacjenta           sql.NullString `json:"nr_domu_pacjenta"`
	NrMieszkaniaPacjenta     sql.NullString `json:"nr_mieszkania_pacjenta"`
	NiezdolnoscOd            sql.NullString `json:"niezdolnoscod"`
	NiezdolnoscDo            sql.NullString `json:"niezdolnoscdo"`
	SzpitalOd                sql.NullString `json:"szpitalod"`
	SzpitalDo                sql.NullString `json:"szpitaldo"`
	Wskazanie                sql.NullString `json:"wskazanie"`
	NrStatystyczny           sql.NullString `json:"nr_statystyczny"`
	Kod1                     sql.NullString `json:"kod1"`
	Kod2                     sql.NullString `json:"kod2"`
	Kod3                     sql.NullString `json:"kod3"`
	Kod4                     sql.NullString `json:"kod4"`
	Pokrewienstwo            sql.NullString `json:"pokrewienstwo"`
	DataUrodzeniaKrewnego    sql.NullString `json:"data_urodzenia_krewnego"`
	RodzajPlatnika           sql.NullString `json:"rodzaj_platnika"`
	IDPlatnika               sql.NullString `json:"id_platnika"`
	NazwaPlacowki            sql.NullString `json:"nazwa_placowki"`
	KodPocztowyPlacowki      sql.NullString `json:"kod_pocztowy_placowki"`
	MiejscowoscPlacowki      sql.NullString `json:"miejscowosc_placowki"`
	UlicaPlacowki            sql.NullString `json:"ulica_placowki"`
	NrDomuPlacowki           sql.NullString `json:"nr_domu_placowki"`
	NrMieszkaniaPlacowki     sql.NullString `json:"nr_mieszkania_placowki"`
	NrPrawaWykonywaniaZawodu sql.NullString `json:"nr_prawa_wykonywania_zawodu"`
	ImieLekarza              sql.NullString `json:"imie_lekarza"`
	NazwiskoLekarza          sql.NullString `json:"nazwisko_lekarza"`
	DataWystawienia          sql.NullString `json:"data_wystawienia"`
	UzasadnienieWstecznego   sql.NullString `json:"uzasadnienie_wstecznego"`
	StacjonarnyZOZ           sql.NullString `json:"stacjonarny_ZOZ"`
	NieInformowacPlatnika    sql.NullString `json:"nie_informowac_platnika"`
	NipPlacowki              sql.NullString `json:"nip_placowki"`
	IDPobytu                 sql.NullString `json:"id_pobytu"`
	Seria                    sql.NullString `json:"seria"`
	Numer                    sql.NullString `json:"numer"`
	IDDokumentu              sql.NullString `json:"id_dokumentu"`
	IDDokumentuKopia         sql.NullString `json:"id_dokumentu_kopia"`
	IDPacjenta               sql.NullString `json:"id_pacjenta"`
	IDLekarza                sql.NullString `json:"id_lekarza"`
	StatusZwolnienia         sql.NullString `json:"status_zwolnienia"`
	IDUzytkownikaWstaw       sql.NullString `json:"id_uzytkownika_wstaw"`
	DataWstawnienia          sql.NullString `json:"data_wstawienia"`
	IDUzytkownikaModyf       sql.NullString `json:"id_uzytkownika_modyf"`
	DataModyfikacji          sql.NullString `json:"data_modyfikacji"`
	KodAnulowania            sql.NullString `json:"kod_anulowania"`
	OpisAnulowania           sql.NullString `json:"opis_anulowania"`
}

type configuration struct {
	UID  string `json:"uid"`
	PWD  string `json:"pwd"`
	HOST string `json:"host"`
	ENG  string `json:"eng"`
	DBN  string `json:"dbn"`
}

var idIDEzwolnienia int

func main() {
	//Open configuration file
	file, err := os.Open("conf.json")
	if err != nil {
		log.Fatalf("Error open file: %s", err)
	}
	defer file.Close()

	//Parse json
	var config configuration
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Error parse json: %s", err)
	}

	//Connect to database
	connectionString := "uid=" + config.UID + ";pwd=" + config.PWD + ";host=" + config.HOST + ";eng=" + config.ENG + ";dbn=" + config.DBN
	db, err := sql.Open("sqlany", connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s", err)
	}

	//Get ezwolenie
	eszla := ezla{}
	idIDEzwolnienia = 100000045
	// Run query
	err = db.QueryRow("select * from ezwolnienie where id_ezwolnienia=?", idIDEzwolnienia).Scan(&eszla.IDEzwolnienia,
		&eszla.PeselPacjenta, &eszla.PaszportPacjenta, &eszla.DataUrodzeniaPacjenta, &eszla.ImiePacjenta, &eszla.NazwiskoPacjenta,
		&eszla.Ubezpieczajacy, &eszla.KodPocztowyPacjenta, &eszla.MiejscowoscPacjenta, &eszla.UlicaPacjenta, &eszla.NrDomuPacjenta,
		&eszla.NrMieszkaniaPacjenta, &eszla.NiezdolnoscOd, &eszla.NiezdolnoscDo, &eszla.SzpitalOd, &eszla.SzpitalDo, &eszla.Wskazanie,
		&eszla.NrStatystyczny, &eszla.Kod1, &eszla.Kod2, &eszla.Kod3, &eszla.Kod4, &eszla.Pokrewienstwo, &eszla.DataUrodzeniaKrewnego,
		&eszla.RodzajPlatnika, &eszla.IDPlatnika, &eszla.NazwaPlacowki, &eszla.KodPocztowyPlacowki, &eszla.MiejscowoscPlacowki,
		&eszla.UlicaPlacowki, &eszla.NrDomuPlacowki, &eszla.NrMieszkaniaPlacowki, &eszla.NrPrawaWykonywaniaZawodu, &eszla.ImieLekarza,
		&eszla.NazwiskoLekarza, &eszla.DataWystawienia, &eszla.UzasadnienieWstecznego, &eszla.StacjonarnyZOZ, &eszla.NieInformowacPlatnika,
		&eszla.NipPlacowki, &eszla.IDPobytu, &eszla.Seria, &eszla.Numer, &eszla.IDDokumentu, &eszla.IDDokumentuKopia, &eszla.IDPacjenta,
		&eszla.IDLekarza, &eszla.StatusZwolnienia, &eszla.IDUzytkownikaWstaw, &eszla.DataWstawnienia, &eszla.IDUzytkownikaModyf, &eszla.DataModyfikacji,
		&eszla.KodAnulowania, &eszla.OpisAnulowania)
	if err != nil {
		log.Fatalf("Select failed: %s", err)
	}
	log.Println(eszla)
}
