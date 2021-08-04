package office

import (
	"encoding/xml"
	"strings"
)

type officeCoreProperty struct {
	XMLName        xml.Name `xml:"coreProperties"`
	Creator        string   `xml:"creator"`
	LastModifiedBy string   `xml:"lastModifiedBy"`
}

type officeAppProperty struct {
	XMLName     xml.Name `xml:"Properties"`
	Application string   `xml:"Application"`
	Company     string   `xml:"Company"`
	Version     string   `xml:"AppVersion"`
}

var officeVersions = map[string]string{
	"16": "2016",
	"15": "2013",
	"14": "2010",
	"12": "2007",
	"11": "2003",
}

func (a *officeAppProperty) GetMajorVersion() string {
	tokens := strings.Split(a.Version, ".")

	if len(tokens) < 2 {
		return "Unknown"
	}

	v, ok := officeVersions[tokens[0]]
	if !ok {
		return "Unknown"
	}
	return v
}
