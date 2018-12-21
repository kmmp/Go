package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"

	zus "../zus_channel_zla_binder"
	//"github.com/fiorix/wsdl2go/soap"
)

var metoda123 zus.MetodaUwierzytelnienia

type ZalogujPodpisem struct {
	PodpisaneOswiadczenie string `xml:"PodpisaneOswiadczenie"`
	MetodaWeryfikacji     string `xml:"MetodaWeryfikacji,omitempty"`
}

type PodpisaneOswiadczenie struct {
	Oswiadczenie OswiadczenieResponse `xml:"Oswiadczenie"`
}
type OswiadczenieResponse struct {
	Tresc  *string `xml:"Tresc"`
	Data   *string `xml:"Data"`
	Czas   *string `xml:"Czas"`
	Token  *string `xml:"Token"`
	Podpis string  `xml:"ds:Signature,omitempty" xmlns:"ds=http://www.w3.org/2000/09/xml,omitempty"`
}
type SoapEnvelope struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	EncodingStyle string   `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
	Body          SoapBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type SoapBody struct {
	Body interface{}
}

func CreateSoapEnvelope() *SoapEnvelope {
	retval := &SoapEnvelope{}
	retval.EncodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"
	return retval
}

func main() {
	var username = flag.String("username", "ezla_ag", "HTTP basic auth username")
	var password = flag.String("password", "ezla_ag", "HTTP basic auth password")
	flag.Parse()

	buffer := &bytes.Buffer{}
	requestEnvelope := CreateSoapEnvelope()
	///requestEnvelope.Body.Body = zus.PobierzOswiadczenie{}
	encoder := xml.NewEncoder(buffer)
	err := encoder.Encode(requestEnvelope)
	if err != nil {
		println("Error encoding document:", err.Error())
		return
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", "https://pue.zus.pl:8001/ws/zus.channel.gabinetowe:zla", buffer)
	if err != nil {
		println("Error creating HTTP request:", err.Error())
		return
	}
	if username != nil && password != nil && *username != "" && *password != "" {
		println("Autheticating")
		req.SetBasicAuth(*username, *password)
	}
	req.Header.Add("SOAPAction", "zus_channel_zla_Binder_pobierzOswiadczenie")
	req.Header.Add("Content-Type", "text/xml")
	resp, err := client.Do(req)
	if err != nil {
		println("Error POSTing HTTP request:", err.Error())
		return
	}
	if resp.StatusCode != 200 {
		println("Error:", resp.Status)
	}

	// responseEnvelope := SoapEnvelope{}
	bodyElement, err := DecodeResponseBody(resp.Body)
	if err != nil {
		println("Error decoding body:", err.Error())
		return
	}
	println("Decoded body!")
	data := &OswiadczenieResponse{}
	xml.Unmarshal([]byte(*bodyElement.Oswiadczenie), &data)
	//fmt.Println("Token: ", *data.Token)

	if bodyElement.Oswiadczenie == nil {
		println("Oswiadczenie is nil")
		return
	}
	//////////////////////////////////////
	clientPodpis := http.Client{}
	buffer = &bytes.Buffer{}
	requestEnvelope = CreateSoapEnvelope()

	v := PodpisaneOswiadczenie{
		Oswiadczenie: OswiadczenieResponse{
			Tresc:  data.Czas,
			Data:   data.Data,
			Czas:   data.Czas,
			Token:  data.Token,
			Podpis: "test",
		},
	}

	output, err := xml.MarshalIndent(v, " ", " ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	send := string(output[:])
	metoda123 = "certyfikat"
	podpisaneOswiadczenie := zus.ZalogujPodpisem{
		PodpisaneOswiadczenie: &send,
		MetodaWeryfikacji:     &metoda123,
	}
	//fmt.Printf("Output: %s", output)
	requestEnvelope.Body.Body = podpisaneOswiadczenie
	encoder = xml.NewEncoder(buffer)
	err = encoder.Encode(requestEnvelope)
	if err != nil {
		println("Error encoding document:", err.Error())
		return
	}

	req, err = http.NewRequest("POST", "https://pue.zus.pl:8001/ws/zus.channel.gabinetowe:zla", buffer)
	if err != nil {
		println("Error creating HTTP request:", err.Error())
		return
	}
	if username != nil && password != nil && *username != "" && *password != "" {
		println("Autheticating")
		req.SetBasicAuth(*username, *password)
	}
	req.Header.Add("SOAPAction", "zus_channel_zla_Binder_zalogujPodpisem")
	req.Header.Add("Content-Type", "text/xml")

	resp, err = clientPodpis.Do(req)
	fmt.Println("REQ: ", req)
	if err != nil {
		println("Error POSTing HTTP request:", err.Error())
		return
	}
	if resp.StatusCode != 200 {
		println("Error:", resp.Status)
	}

	// responseEnvelope := SoapEnvelope{}
	bodyElement2, err := DecodeResponseZaloguj(resp.Body)
	if err != nil {
		println("Error decoding body:", err.Error())
		return
	}
	println("Decoded body!")
	//data2 := zus.ZalogujPodpisemResponse{}
	//xml.Unmarshal([]byte(&bodyElement2), &data2)
	//fmt.Println("Sesja: ", *data2.IdSesji)

	if bodyElement2.IdSesji == nil {
		println("Oswiadczenie is nil")
		return
	}
}

func DecodeResponseBody(body io.Reader) (*zus.PobierzOswiadczenieResponse, error) {
	decoder := xml.NewDecoder(body)
	nextElementIsBody := false
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			if nextElementIsBody {
				fmt.Println("Jestem w body")
				responseBody := zus.PobierzOswiadczenieResponse{}
				err = decoder.DecodeElement(&responseBody, &startElement)
				if err != nil {
					return nil, err
				}
				fmt.Printf("%v", *responseBody.Oswiadczenie)
				return &responseBody, nil
			}
			if startElement.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && startElement.Name.Local == "Body" {
				nextElementIsBody = true
			}
		}
	}
	return nil, errors.New("Did not find SOAP body element")
}

func DecodeResponseZaloguj(body io.Reader) (*zus.ZalogujPodpisemResponse, error) {
	decoder := xml.NewDecoder(body)
	nextElementIsBody := false
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			if nextElementIsBody {
				fmt.Println("Jestem w body logowania")
				responseBody := zus.ZalogujPodpisemResponse{}
				err = decoder.DecodeElement(&responseBody, &startElement)
				if err != nil {
					return nil, err
				}
				fmt.Printf("%v%v", *responseBody.Rezultat.OpisBledu, *responseBody.Rezultat.KodBledu)
				return &responseBody, nil
			}
			if startElement.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && startElement.Name.Local == "Body" {
				nextElementIsBody = true
			}
		}
	}

	return nil, errors.New("Did not find SOAP body element")
}
