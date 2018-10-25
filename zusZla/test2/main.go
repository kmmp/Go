package main

import (
	"fmt"
	"log"
	"net/http"

	soap "github.com/fiorix/wsdl2go/soap2"
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
	addr := "https://pue.zus.pl:8001/ws/zus.channel.gabinetowe:zla"
	client := soap.NewClient(addr, true, nil)
	//req := zusWSDL.PobierzOswiadczenie{}
	//res := zusWSDL.PobierzOswiadczenieResponse{}
	err := client.Call("pobierzOswiadczenie", nil, nil, nil)
	fmt.Println(err)
}

func zusPodpiszDokument() {

}
