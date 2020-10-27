package d16b

import (
	"encoding/xml"
	"strconv"

	"github.com/slovak-egov/einvoice/apiserver/db"
)

func Create(value string) (*db.Invoice, error) {
	inv := &CrossIndustryInvoice{}
	err := xml.Unmarshal([]byte(value), &inv)
	if err != nil {
		return nil, err
	}

	price, _ := strconv.ParseFloat(inv.SupplyChainTradeTransaction.ApplicableHeaderTradeSettlement.SpecifiedTradeSettlementHeaderMonetarySummation.LineTotalAmount.Value, 64)
	return &db.Invoice{
		Sender:   inv.SupplyChainTradeTransaction.ApplicableHeaderTradeAgreement.SellerTradeParty.Name,
		Receiver: inv.SupplyChainTradeTransaction.ApplicableHeaderTradeAgreement.BuyerTradeParty.Name,
		Format:   db.D16bFormat,
		Price:    price,
	}, nil
}
