package main

import (
	"fmt"
	"log"
	"net/http"
<<<<<<< HEAD
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
}

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
=======
	zus "zus_channel_zla_binder"

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
	}

}

func zusPobierzOswiadczenie() {
	client := soap.Client{
		URL: "https://pue.zus.pl:8001/ws/zus.channel.gabinetowe:zla"}

	fmt.Println(zus.NewZla_PortType(&client).PobierzOswiadczenie(&zus.PobierzOswiadczenie{}))
}

func zusPodpiszDokument() {

>>>>>>> 89d2fc4f8ac288bd4c673636700bb7365f84917d
}
