package invoice

import (
	"encoding/xml"
)

type Invoice struct {
	XMLName  xml.Name `xml:"invoice"`
	Id       string   `xml:"id" json:"id"`
	Sender   string   `xml:"sender" json:"sender"`
	Receiver string   `xml:"receiver" json:"receiver"`
}

type Meta struct {
	Id       string `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

func (invoice *Invoice) GetMeta() *Meta {
	return &Meta{Id: invoice.Id, Sender: invoice.Sender, Receiver: invoice.Receiver}
}
