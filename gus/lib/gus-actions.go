package lib

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//GetSid method comment
func GetSid() string {

	v := "<soap:Envelope xmlns:soap=\"http://www.w3.org/2003/05/soap-envelope\" xmlns:ns=\"http://CIS/BIR/PUBL/2014/07\">" +
		"<soap:Header xmlns:wsa=\"http://www.w3.org/2005/08/addressing\">" +
		"<wsa:Action>http://CIS/BIR/PUBL/2014/07/IUslugaBIRzewnPubl/Zaloguj</wsa:Action>" +
		"<wsa:To>https://wyszukiwarkaregon.stat.gov.pl/wsBIR/UslugaBIRzewnPubl.svc</wsa:To>" +
		"</soap:Header>" +
		"<soap:Body>" +
		"<ns:Zaloguj>" +
		"<ns:pKluczUzytkownika>d83f52f498b44dfda2f2</ns:pKluczUzytkownika>" +
		"</ns:Zaloguj>" +
		"</soap:Body>" +
		"</soap:Envelope>"

	buffer := bytes.Buffer{}
	buffer.Write([]byte(v))
	client := http.Client{}
	req := getReq(&buffer)

	resp, err := client.Do(req)
	if err != nil {
		println("Error POSTing HTTP request:", err.Error())
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Error: %v", resp.Status)
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	r, _ := regexp.Compile("<ZalogujResult>(.*)</ZalogujResult>")
	return r.FindStringSubmatch(string(contents))[1]
}

func getReq(buffer *bytes.Buffer) *http.Request {
	req, err := http.NewRequest("POST", "https://wyszukiwarkaregon.stat.gov.pl/wsBIR/UslugaBIRzewnPubl.svc", buffer)
	if err != nil {
		println("Error creating HTTP request:", err.Error())
		return nil
	}

	req.Header.Add("Content-Type", "application/soap+xml")

	//fmt.Printf("%v", req)
	return req
}

//GetFirmBasicData method comment
func GetFirmBasicData(sid, typ, wartosc string) (regon, raportTyp string) {
	inputNip := "<soap:Envelope xmlns:soap=\"http://www.w3.org/2003/05/soap-envelope\" xmlns:ns=\"http://CIS/BIR/PUBL/2014/07\" xmlns:dat=\"http://CIS/BIR/PUBL/2014/07/DataContract\">" +
		"<soap:Header xmlns:wsa=\"http://www.w3.org/2005/08/addressing\">" +
		"<wsa:Action>http://CIS/BIR/PUBL/2014/07/IUslugaBIRzewnPubl/DaneSzukaj</wsa:Action>" +
		"<wsa:To>https://wyszukiwarkaregon.stat.gov.pl/wsBIR/UslugaBIRzewnPubl.svc</wsa:To>" +
		"</soap:Header>" +
		"<soap:Body>" +
		"<ns:DaneSzukaj>" +
		"<ns:pParametryWyszukiwania>" +
		"<dat:Nip>" + wartosc + "</dat:Nip>" +
		"</ns:pParametryWyszukiwania>" +
		"</ns:DaneSzukaj>" +
		"</soap:Body>" +
		"</soap:Envelope>"

	inputRegon := "<soap:Envelope xmlns:soap=\"http://www.w3.org/2003/05/soap-envelope\" xmlns:ns=\"http://CIS/BIR/PUBL/2014/07\" xmlns:dat=\"http://CIS/BIR/PUBL/2014/07/DataContract\">" +
		"<soap:Header xmlns:wsa=\"http://www.w3.org/2005/08/addressing\">" +
		"<wsa:Action>http://CIS/BIR/PUBL/2014/07/IUslugaBIRzewnPubl/DaneSzukaj</wsa:Action>" +
		"<wsa:To>https://wyszukiwarkaregon.stat.gov.pl/wsBIR/UslugaBIRzewnPubl.svc</wsa:To>" +
		"</soap:Header>" +
		"<soap:Body>" +
		"<ns:DaneSzukaj>" +
		"<ns:pParametryWyszukiwania>" +
		"<dat:Regon>" + wartosc + "</dat:Regon>" +
		"</ns:pParametryWyszukiwania>" +
		"</ns:DaneSzukaj>" +
		"</soap:Body>" +
		"</soap:Envelope>"

	buffer := bytes.Buffer{}
	switch typ {
	case "NIP":
		buffer.Write([]byte(inputNip))
	case "REGON":
		buffer.Write([]byte(inputRegon))
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://wyszukiwarkaregon.stat.gov.pl/wsBIR/UslugaBIRzewnPubl.svc", &buffer)
	if err != nil {
		println("Error creating HTTP request:", err.Error())
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/soap+xml")
	req.Header.Add("sid", sid)

	resp, err := client.Do(req)
	if err != nil {
		println("Error POSTing HTTP request:", err.Error())
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Error: %v", resp.Status)
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Println(string(contents))
	replaceResult := strings.Replace(string(contents), "&lt;", "<", -1)
	replaceResult = strings.Replace(replaceResult, "&gt;&#xD;", ">", -1)
	replaceResult = strings.Replace(replaceResult, "&gt;", ">", -1)
	fmt.Println(replaceResult)

	r, _ := regexp.Compile("(?s)(<root>.*</root>)")
	fmt.Println(r.FindStringSubmatch(replaceResult)[1])

	res := &root{}
	err = xml.Unmarshal([]byte(r.FindStringSubmatch(replaceResult)[1]), res)

	regon = *res.Dane.Regon
	raportTyp = *res.Dane.Typ

	return regon, raportTyp
}

//GetFirmFullData method comment
func GetFirmFullData(sid, regon, typ string) string {

	var typRaportu string
	switch typ {
	case "P":
		typRaportu = "PublDaneRaportPrawna"
		break
	case "F":
		typRaportu = "PublDaneRaportFizycznaOsoba"
		break
	}

	inputPelnyRaport := "<soap:Envelope xmlns:soap=\"http://www.w3.org/2003/05/soap-envelope\" xmlns:ns=\"http://CIS/BIR/PUBL/2014/07\">" +
		"<soap:Header xmlns:wsa=\"http://www.w3.org/2005/08/addressing\">" +
		"<wsa:Action>http://CIS/BIR/PUBL/2014/07/IUslugaBIRzewnPubl/DanePobierzPelnyRaport</wsa:Action>" +
		"<wsa:To>https://wyszukiwarkaregon.stat.gov.pl/wsBIR/UslugaBIRzewnPubl.svc</wsa:To>" +
		"</soap:Header>" +
		"<soap:Body>" +
		"<ns:DanePobierzPelnyRaport>" +
		"<ns:pRegon>" + regon + "</ns:pRegon>" +
		"<ns:pNazwaRaportu>" + typRaportu + "</ns:pNazwaRaportu>" +
		"</ns:DanePobierzPelnyRaport>" +
		"</soap:Body>" +
		"</soap:Envelope>"

	buffer := bytes.Buffer{}
	buffer.Write([]byte(inputPelnyRaport))
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://wyszukiwarkaregon.stat.gov.pl/wsBIR/UslugaBIRzewnPubl.svc", &buffer)
	if err != nil {
		println("Error creating HTTP request:", err.Error())
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/soap+xml")
	req.Header.Add("sid", sid)

	resp, err := client.Do(req)
	if err != nil {
		println("Error POSTing HTTP request:", err.Error())
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Error: %v", resp.Status)
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Println(string(contents))
	replaceResult := strings.Replace(string(contents), "&lt;", "<", -1)
	replaceResult = strings.Replace(replaceResult, "&gt;&#xD;", ">", -1)
	replaceResult = strings.Replace(replaceResult, "&gt;", ">", -1)
	fmt.Println(replaceResult)

	r, _ := regexp.Compile("(?s)(<root>.*</root>)")
	fmt.Println(r.FindStringSubmatch(replaceResult)[1])

	return r.FindStringSubmatch(replaceResult)[1]
}

func Logout(sid string) {
	input := "<soap:Envelope xmlns:soap=\"http://www.w3.org/2003/05/soap-envelope\" xmlns:ns=\"http://CIS/BIR/PUBL/2014/07\">" +
		"<soap:Header xmlns:wsa=\"http://www.w3.org/2005/08/addressing\">" +
		"<wsa:To>https://wyszukiwarkaregontest.stat.gov.pl/wsBIR/UslugaBIRzewnPubl.svc</wsa:To>" +
		"<wsa:Action>http://CIS/BIR/PUBL/2014/07/IUslugaBIRzewnPubl/Wyloguj</wsa:Action>" +
		"</soap:Header>" +
		"<soap:Body>" +
		"<ns:Wyloguj>" +
		"<ns:pIdentyfikatorSesji>" + sid + "</ns:pIdentyfikatorSesji>" +
		"</ns:Wyloguj>" +
		"</soap:Body>" +
		"</soap:Envelope>"

	buffer := bytes.Buffer{}
	buffer.Write([]byte(input))
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://wyszukiwarkaregon.stat.gov.pl/wsBIR/UslugaBIRzewnPubl.svc", &buffer)
	if err != nil {
		println("Error creating HTTP request:", err.Error())
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/soap+xml")
	req.Header.Add("sid", sid)

	resp, err := client.Do(req)
	if err != nil {
		println("Error POSTing HTTP request:", err.Error())
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Error: %v", resp.Status)
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Println(string(contents))
}
