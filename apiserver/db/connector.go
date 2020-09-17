package db

import (
	"github.com/slovak-egov/einvoice/apiserver/invoice"
	"github.com/slovak-egov/einvoice/common"
)

type DBConnector interface {
	Connect()
	Close()
	InitDB() error
	GetAllInvoice() ([]invoice.Meta, error)
	GetInvoiceMeta(id string) (*invoice.Meta, error)
	CreateInvoice(invoice *invoice.Meta) error
}
