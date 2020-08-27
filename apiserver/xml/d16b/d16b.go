package d16b

import "encoding/xml"

type CrossIndustryInvoice struct {
	XMLName                     xml.Name                    `xml:"CrossIndustryInvoice"`
	SupplyChainTradeTransaction SupplyChainTradeTransaction `xml:"SupplyChainTradeTransaction"`
}

type SupplyChainTradeTransaction struct {
	ApplicableHeaderTradeAgreement HeaderTradeDeliveryType `xml:"ApplicableHeaderTradeAgreement"`
}

type HeaderTradeDeliveryType struct {
	SellerTradeParty TradePartyType `xml:"SellerTradeParty"`
	BuyerTradeParty  TradePartyType `xml:"BuyerTradeParty"`
}

type TradePartyType struct {
	Name string `xml:"Name"`
}
