package lib

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var username, password string = "ezla_ag", "ezla_ag"

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

func GetOswiadczenie() *PobierzOswiadczenieResponse {

	buffer := &bytes.Buffer{}
	requestEnvelope := CreateSoapEnvelope()
	encoder := xml.NewEncoder(buffer)
	err := encoder.Encode(requestEnvelope)
	if err != nil {
		println("Error encoding document:", err.Error())
		return nil
	}

	client := http.Client{}
	req := getReq("zus_channel_zla_Binder_pobierzOswiadczenie", buffer)
	resp, err := client.Do(req)
	if err != nil {
		println("Error POSTing HTTP request:", err.Error())
		return nil
	}
	if resp.StatusCode != 200 {
		println("Error:", resp.Status)
	}

	bodyElement, err := decodeResponseBody(resp.Body)
	if err != nil {
		println("Error decoding body:", err.Error())
		return nil
	}
	println("Decoded body!")

	if bodyElement.Oswiadczenie == nil {
		println("Oswiadczenie is nil")
		return nil
	}

	return bodyElement
}

func getSession(po *PobierzOswiadczenieResponse) *ZalogujPodpisemResponse {

	data := &OswiadczenieResponse{}
	xml.Unmarshal([]byte(*po.Oswiadczenie), &data)

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
	var auth MetodaUwierzytelnienia = "certyfikat"
	podpisaneOswiadczenie := ZalogujPodpisem{
		PodpisaneOswiadczenie: &send,
		MetodaWeryfikacji:     &auth,
	}

	client := http.Client{}
	buffer := &bytes.Buffer{}
	requestEnvelope := CreateSoapEnvelope()
	requestEnvelope.Body.Body = podpisaneOswiadczenie
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(requestEnvelope)
	if err != nil {
		println("Error encoding document:", err.Error())
		return nil
	}

	req, err := http.NewRequest("POST", "https://pue.zus.pl:8001/ws/zus.channel.gabinetowe:zla", buffer)
	if err != nil {
		println("Error creating HTTP request:", err.Error())
		return nil
	}
	if username != "" && password != "" {
		println("Autheticating")
		req.SetBasicAuth(username, password)
	}
	req.Header.Add("SOAPAction", "zus_channel_zla_Binder_zalogujPodpisem")
	req.Header.Add("Content-Type", "text/xml")

	resp, err := client.Do(req)
	fmt.Println("REQ: ", req)
	if err != nil {
		println("Error POSTing HTTP request:", err.Error())
		return nil
	}
	if resp.StatusCode != 200 {
		println("Error:", resp.Status)
	}

	bodyElement, err := decodeResponseZaloguj(resp.Body)
	if err != nil {
		println("Error decoding body:", err.Error())
		return nil
	}
	println("Decoded body!")

	if bodyElement.IdSesji == nil {
		println("Id sesji is nil")
		return nil
	}
	return bodyElement
}

func decodeResponseBody(body io.Reader) (*PobierzOswiadczenieResponse, error) {
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
				responseBody := PobierzOswiadczenieResponse{}
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

func decodeResponseZaloguj(body io.Reader) (*ZalogujPodpisemResponse, error) {
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
				responseBody := ZalogujPodpisemResponse{}
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

func getReq(soapAction string, buffer *bytes.Buffer) *http.Request {
	req, err := http.NewRequest("POST", "https://pue.zus.pl:8001/ws/zus.channel.gabinetowe:zla", buffer)
	if err != nil {
		println("Error creating HTTP request:", err.Error())
		return nil
	}
	if username != "" && password != "" {
		println("Autheticating")
		req.SetBasicAuth(username, password)
	}
	req.Header.Add("SOAPAction", soapAction)
	req.Header.Add("Content-Type", "text/xml")

	return req
}
