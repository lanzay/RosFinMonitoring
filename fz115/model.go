package fz115

import "encoding/xml"

type (
	FinMonitoring struct {
		XMLName       xml.Name    `xml:"TERRORISTS_CATALOG"`
		XMLID         string      `xml:"ID,attr"`
		XMLDate       string      `xml:"DATE,attr"`
		XMLNum        string      `xml:"NUM,attr"`
		TerroristList []Terrorist `xml:"TERRORISTS"`
		FileName      string
	}
	Terrorist struct {
		Name                string `xml:"TERRORISTS_NAME"`
		Num                 int    `xml:"NUM"`
		PersonType          int    `xml:"person_type"`
		BirthDate           string `xml:"birth_date"`
		Description         string `xml:"DESCRIPTION"`
		Addres              string `xml:"ADDRESS"`
		TerrorisrResolution string `xml:"TERRORISTS_RESOLUTION"`
		BirthPlase          string `xml:"BIRTH_PLACE"`
		PassportStr         string `xml:"PASSPORT"`
		Passport            Passport
		ID                  string `xml:"ID"`
	}
)

type Passport struct {
	Type   string
	NumSer string
	NumNum string
	Vidan  string
}
