package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	zus "zus_channel_zla_binder"
	//"github.com/fiorix/wsdl2go/soap"
)

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

	// FIXME: encoding
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
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		println("Error POSTing HTTP request:", err.Error())
		return
	}
	if resp.StatusCode != 200 {
		println("Error:", resp.Status)
	}
	// FIXME: check Content-Type
	// FIXME: encoding

	// responseEnvelope := SoapEnvelope{}
	bodyElement, err := DecodeResponseBody(resp.Body)
	if err != nil {
		println("Error decoding body:", err.Error())
		return
	}
	println("Decoded body!")

	if bodyElement.Oswiadczenie == nil {
		println("Oswiadczenie is nil")
		return
	}

	/*for _, project := range bodyElement.GetAllProjectsResult.ProjectReference1 {
		println("Project:", *project.Slug, *project.Uri, *project.DisplayText)
	}*/
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
				fmt.Printf("%s", responseBody)
				return &responseBody, nil
			}
			if startElement.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && startElement.Name.Local == "Body" {
				nextElementIsBody = true
			}
		}
	}

	return nil, errors.New("Did not find SOAP body element")
}