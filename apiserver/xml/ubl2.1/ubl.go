package ubl2_1

import "encoding/xml"

type Invoice struct {
	XMLName  xml.Name                `xml:"Invoice"`
	ID       string                  `xml:"ID"`
	Supplier AccountingSupplierParty `xml:"AccountingSupplierParty"`
	Customer AccountingCustomerParty `xml:"AccountingCustomerParty"`
}

type AccountingSupplierParty struct {
	Party Party `xml:"Party"`
}

type AccountingCustomerParty struct {
	Party Party `xml:"Party"`
}

type Party struct {
	Name PartyName `xml:"PartyName"`
}

type PartyName struct {
	Name string `xml:"Name"`
}
