package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"
)

type Sirket struct {
	XMLName     xml.Name   `xml:"sirket"`
	Personeller []Personel `xml:personel`
}

type Personel struct {
	XMLName  xml.Name `xml:"personel"`
	TCNO     int      `xml:"tcno"`
	PerAd    string   `xml:"personelAd"`
	PerSoyad string   `xml:"personelSoyad"`
	PerMevki string   `xml:"mevki"`
}

func main() {

	dosya, err := os.Create("personeller.xml")
	if err != nil {
		log.Fatal(err)
	}

	str := "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"
	dosya.WriteString(str)
	s := &Sirket{}
	s.Personeller = append(s.Personeller, Personel{TCNO: 3480374, PerAd: "beyza", PerSoyad: "atay", PerMevki: "yazılımcı"})
	s.Personeller = append(s.Personeller, Personel{TCNO: 44534, PerAd: "helin", PerSoyad: "atay", PerMevki: "yönetici"})

	xmlWriter := io.Writer(dosya)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	enc.Encode(s)
}
