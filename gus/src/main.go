package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../lib"
)

var accessKeyValue = "ESodX123#"

//PORT comment
const PORT = 8080

//Incoming comment
type Incoming struct {
	AccsessKey string `json:"key,omitempty"`
	Type       string `json:"type,omitempt"`
	Value      string `json:"value,omitempt"`
}

func main() {
	log.Println("Starting server on port: " + strconv.Itoa(PORT))
	http.HandleFunc("/", getFirm)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}

//GetFirm comment 5262281008
func getFirm(w http.ResponseWriter, r *http.Request) {
	byteData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Błędny opbiekt JSON"))
		return
	}
	var incoming Incoming
	err = json.Unmarshal(byteData, &incoming)
	key := incoming.AccsessKey
	typeP := incoming.Type
	value := incoming.Value

	if key == "ESodX123#" {
		sid := lib.GetSid()
		regon, typ := lib.GetFirmBasicData(sid, typeP, value)
		firmData := lib.GetFirmFullData(sid, regon, typ)
		lib.Logout(sid)
		w.Write([]byte(firmData))
	} else {
		w.Write([]byte("Błędny klucz dostępu"))
	}
}
