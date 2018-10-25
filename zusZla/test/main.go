package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//wsdl "zus_channel_zla_binder"
)

func main() {
	client := &http.Client{}

	//body := []byte("{\n  \"title\": \"Buy cheese and bread for breakfast.\"\n}")

	req, _ := http.NewRequest("POST", "https://pue.zus.pl:8001/ws/zus.channel.gabinetowe:zla", nil)
	req.SetBasicAuth("ezla_ag", "ezla_ag")

	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("SOAPAction", "pobierzOswiadczenie")
	fmt.Println(req)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
}
