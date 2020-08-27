package d16b

import (
	"encoding/xml"
	"github.com/filipsladek/einvoice/apiserver/invoice"
)

func Create(value string) (error, *invoice.Invoice) {
	d16bInvoice := &CrossIndustryInvoice{}
	err := xml.Unmarshal([]byte(value), &d16bInvoice)
	if err != nil {
		return err, nil
	}

	return nil, &invoice.Invoice{
		Sender:   d16bInvoice.SupplyChainTradeTransaction.ApplicableHeaderTradeAgreement.SellerTradeParty.Name,
		Receiver: d16bInvoice.SupplyChainTradeTransaction.ApplicableHeaderTradeAgreement.BuyerTradeParty.Name,
	}
}
