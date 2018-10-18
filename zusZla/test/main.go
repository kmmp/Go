package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	body := []byte("{\n  \"title\": \"Buy cheese and bread for breakfast.\"\n}")

	req, _ := http.NewRequest("POST", "https://pue.zus.pl:8001/ws/zus.channel.gabinetowe:zla", bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("SOAPAction", "pobierzDokument")
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
