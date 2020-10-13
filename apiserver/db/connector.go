package db

import (
	"github.com/slovak-egov/einvoice/apiserver/invoice"
)

type DBConnector interface {
	Connect()
	Close()
	GetAllInvoice() ([]invoice.Meta, error)
	GetInvoiceMeta(id string) (*invoice.Meta, error)
	CreateInvoice(invoice *invoice.Meta) error
}
