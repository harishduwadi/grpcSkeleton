package config

import "encoding/xml"

type Configuration struct {
	XMLName           xml.Name `xml:"Configuration"`
	ListenPort        string   `xml:"ListenAddress"`
	LogFile           string   `xml:"LogFile"`
	PublicKeyLocation string   `xml:"PublicKeyLocation"`
}
