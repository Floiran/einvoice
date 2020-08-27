package ubl21

import (
	"encoding/xml"
	"github.com/filipsladek/einvoice/apiserver/invoice"
)

func Create(value string) (error, *invoice.Meta) {
	ublInvoice := &Invoice{}
	err := xml.Unmarshal([]byte(value), &ublInvoice)
	if err != nil {
		return err, nil
	}

	return nil, &invoice.Meta{
		Sender:   ublInvoice.Supplier.Party.Name.Name,
		Receiver: ublInvoice.Customer.Party.Name.Name,
		Format:   invoice.UblFormat,
	}
}