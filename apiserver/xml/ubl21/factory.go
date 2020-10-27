package ubl21

import (
	"encoding/xml"
	"strconv"

	"github.com/slovak-egov/einvoice/apiserver/db"
)

func Create(value string) (*db.Invoice, error) {
	inv := &Invoice{}
	err := xml.Unmarshal([]byte(value), &inv)
	if err != nil {
		return nil, err
	}

	price, _ := strconv.ParseFloat(inv.LegalMonetaryTotal.PayableAmount.Value, 64)
	return &db.Invoice{
		Sender:   inv.Supplier.Party.Name.Name,
		Receiver: inv.Customer.Party.Name.Name,
		Format:   db.UblFormat,
		Price:    price,
	}, nil
}
