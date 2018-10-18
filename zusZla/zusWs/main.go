package main

import (
	"fmt"
	"log"
	"net/http"
	zusWSDL "zus_channel_zla_binder"

	"github.com/fiorix/wsdl2go/soap"
)

type pobierzDokument struct {
}

type esDokument struct {
	esBody string
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9443", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["method"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'method' is missing")
		return
	}

	key := keys[0]

	log.Println("Url Param 'method' is: " + string(key))

	/*decoder := json.NewDecoder(r.Body)
	var pd esDokument
	err := decoder.Decode(&pd)
	if err != nil {
		panic(err)
	}

	log.Println(pd.esBody)
	*/
	switch keys[0] {
	case "pobierzOswiadczenie":
		zusPobierzOswiadczenie()
	case "podpiszDokument":
		zusPodpiszDokument()
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Println("Missing method!")
	}

}

func zusPobierzOswiadczenie() {

	addr := "https://pue.zus.pl:8001/ws/zus.channel.gabinetowe:zla"
	//client := &http.Client{}
	cli := soap.Client{URL: addr}
	

	//body := []byte("{\n  \"title\": \"Buy cheese and bread for breakfast.\"\n}")

	//req, _ := http.NewRequest("POST", addr, bytes.NewBuffer(body))
	//req, _ := http.NewRequest("POST", addr, nil)

	//req.Header.Add("Content-Type", "application/json")
	zusClient := zusWSDL.NewZla_PortType(&cli)
	//fmt.Println(dupa)
	var b zusWSDL.PobierzOswiadczenie
	resp, err := zusClient.PobierzOswiadczenie(&b)
	if err != nil {
		fmt.Printf("Błąd: %s", err)
	}
	fmt.Println(resp)
	/*resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(respBody))*/
}

func zusPodpiszDokument() {

}
