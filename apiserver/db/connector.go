package db

import (
	"github.com/slovak-egov/einvoice/apiserver/attachment"
	"github.com/slovak-egov/einvoice/apiserver/invoice"
)

type DBConnector interface {
	Connect()
	Close()
	GetAllInvoice(formats []string) ([]invoice.Meta, error)
	GetInvoiceMeta(id int) (*invoice.Meta, error)
	CreateInvoice(invoice *invoice.Meta) error
	CreateAttachment(invoiceId int, name string) (*attachment.Meta, error)
	GetAttachment(id int) (*attachment.Meta, error)
	GetAttachmentForInvoice(id int) ([]attachment.Meta, error)
}
