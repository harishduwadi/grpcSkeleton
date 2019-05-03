package config

import "encoding/xml"

type Configuration struct {
	XMLName    xml.Name `xml:"Configuration"`
	ListenPort string   `xml:"ListenPort"`
	LogFile    string   `xml:"LogFile"`
	PublicKey  string   `xml:"PublicKeyLocation"`
	PrivateKey string   `xml:"PrivateKeyLocation"`
}
