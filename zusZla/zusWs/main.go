package main

import (
	"log"

	"../zusWSDL"
)

func main() {
	con := zusWSDL.PobierzPlatnikowUbezpieczonego(zusWSDL.PobierzPlatnikowUbezpieczonego{IdSesji: nil, Ubezpieczony: nil})

	log.Println(con)
}

/*Oryginal *DocTypeRef_Dokument `xml:"Oryginal,omitempty" json:"Oryginal,omitempty" yaml:"Oryginal,omitempty"`
Kopia    *DocTypeRef_Dokument `xml:"Kopia,omitempty" json:"Kopia,omitempty" yaml:"Kopia,omitempty"`
NrRef    Inne                 `xml:"NrRef,attr,omitempty" json:"NrRef,attr,omitempty" yaml:"NrRef,attr,omitempty"`
*/
