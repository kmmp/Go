package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../lib"
)

type requestParms struct {
	ID   string `json:"id"`
	Cert string `json:"cert"`
}

type respMessage struct {
	IsError bool   `json:"error"`
	Message string `json:"message"`
}

func main() {

	http.HandleFunc("/ws/sendZLA", sendHandler)
	http.HandleFunc("/ws/cancelZLA", cancelHandler)
	log.Fatal(http.ListenAndServe("localhost:9443", nil))
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	parms := getBody(r)
	i, err := strconv.Atoi(parms.ID)
	if err != nil {
		panic(err)
	}
	docParm := lib.GetDocParms(i)
	log.Println(docParm)
	_, err = createSendDok(docParm)
	if err != nil {
		panic(err)
	}

	//Response
	resp := respMessage{
		IsError: false,
		Message: "Possij :D",
	}
	respJSON, errJSON := json.Marshal(resp)
	if errJSON != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	parms := getBody(r)
	log.Println(parms)
}

func getBody(r *http.Request) requestParms {
	decoder := json.NewDecoder(r.Body)
	var parms requestParms
	err := decoder.Decode(&parms)
	if err != nil {
		panic(err)
	}
	return parms
}

func createSendDok(ezla lib.Ezla) (lib.DokumentSend, error) {
	pozycje := lib.DokumentSend{
		TrescDokumentu: lib.Pozycje{
			Ogolne: lib.OgolneParm{
				Identyfikator: lib.IdentParm{
					Seria: ezla.Seria.String,
					Numer: ezla.Numer.String,
				},
				OryginalKopia: "ORYGINAŁ",
			},
			DanePacjenta: lib.DanePacjentaParm{
				Pesel:         ezla.PeselPacjenta.Int64,
				Imie:          ezla.ImiePacjenta.String,
				Nazwisko:      ezla.NazwiskoPacjenta.String,
				Instytucja:    ezla.Ubezpieczajacy.Int64,
				Paszport:      ezla.PaszportPacjenta.String,
				DataUrodzenia: ezla.DataUrodzeniaPacjenta.String,
			},
			AdresPacjenta: lib.AdresPacjentaParm{
				KodPocztowy:  ezla.KodPocztowyPacjenta.String,
				Miejscowosc:  ezla.MiejscowoscPacjenta.String,
				Ulica:        ezla.UlicaPacjenta.String,
				NrDomu:       ezla.NrDomuPacjenta.String,
				NrMieszkania: ezla.NrMieszkaniaPacjenta.String,
			},
			Zwolnienie: lib.ZwolnienieParm{
				Niezdolny: lib.NiezdolnyParm{
					Od: ezla.NiezdolnoscOd.String,
					Do: ezla.NiezdolnoscDo.String,
				},
				Szpital: lib.SzpitalParm{
					SzpitalOd: ezla.SzpitalOd.String,
					SzpitalDo: ezla.SzpitalDo.String,
				},
				Wskazanie:   ezla.Wskazanie.String,
				Kod1:        ezla.Kod1.String,
				Kod2:        ezla.Kod2.String,
				Kod3:        ezla.Kod3.String,
				Kod4:        ezla.Kod4.String,
				Rozpoznanie: ezla.NrStatystyczny.String,
				Pokrewienstwo: lib.PokrewienstwoParm{
					KodPokrewienstwa:            ezla.Pokrewienstwo.String,
					DataUrodzeniaOsobyPodOpieka: ezla.DataUrodzeniaKrewnego.String,
				},
			},
			Platnik: lib.PlatnikParm{
				Rodzaj:        ezla.RodzajPlatnika.String,
				Identyfikator: ezla.IDPlatnika.String,
			},
			Przychodnia: lib.PrzychodniaParm{
				Nazwa:        ezla.NazwaPlacowki.String,
				KodPocztowy:  ezla.KodPocztowyPlacowki.String,
				Miejscowosc:  ezla.MiejscowoscPlacowki.String,
				Ulica:        ezla.UlicaPlacowki.String,
				NrDomu:       ezla.NrDomuPlacowki.String,
				NrMieszkania: ezla.NrMieszkaniaPlacowki.String,
			},
			Lekarz: lib.LekarzParm{
				NumerPrawa: ezla.NrPrawaWykonywaniaZawodu.String,
				Imie:       ezla.ImieLekarza.String,
				Nazwisko:   ezla.NazwiskoLekarza.String,
			},
			Podsumowanie: lib.PodsumowanieParm{
				DataWystawieniaDokumentu: ezla.DataWystawienia.String,
				DataElektronizacji:       ezla.DataWystawienia.String,
				WsteczneWystawienie:      ezla.UzasadnienieWstecznego.String,
				Anulowanie: lib.AnulowanieParm{
					Seria: ezla.Seria.String,
					Numer: ezla.Numer.String,
				},
				Powiazanie: lib.PowiazanieParm{
					Seria: ezla.Seria.String,
					Numer: ezla.Numer.String,
				},
				PobytStajonarny: ezla.StacjonarnyZOZ.Bool,
				InfoDlaPlatnika: ezla.NieInformowacPlatnika.Bool,
				NipPwdl:         ezla.NipPlacowki.Int64,
			},
		},
	}
	output, err := xml.MarshalIndent(pozycje, " ", " ")
	if err != nil {
		fmt.Printf("Error create xml doc: %v\n", err)
	}
	fmt.Printf("Output: %s", output)
	return pozycje, nil
}

func sendZLA(lib.DokumentSend) error {
	return nil
}

//func getTokenZLA() lib.
