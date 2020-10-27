package db

import (
	"time"

	"github.com/go-pg/pg/v10"
)

const (
	UblFormat  = "ubl2.1"
	D16bFormat = "d16b"
)

type Invoice struct {
	tableName   struct{}          `pg:"invoices"`
	Id          int               `json:"id"`
	Sender      string            `json:"sender"`
	Receiver    string            `json:"receiver"`
	Format      string            `json:"format"`
	Price       float64           `json:"price"`
	CreatedAt   time.Time         `json:"created_at"`
	Attachments []Attachment      `json:"attachments" pg:"-"`
}

func (connector *Connector) GetInvoices(formats []string) ([]Invoice, error) {
	invoices := []Invoice{}
	err := connector.Db.Model(&invoices).Where("format IN (?)", pg.In(formats)).Select()
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

func (connector *Connector) GetInvoice(id int) (*Invoice, error) {
	inv := &Invoice{}
	err := connector.Db.Model(inv).Where("id = ?", id).Select(inv)
	if err != nil {
		return nil, err
	}

	return inv, nil
}

func (connector *Connector) CreateInvoice(invoice *Invoice) error {
	_, err := connector.Db.Model(invoice).Insert(invoice)

	return err
}

func (connector *Connector) GetInvoiceAttachments(id int) ([]Attachment, error) {
	ats := []Attachment{}
	err := connector.Db.Model(&ats).Where("invoice_id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return ats, nil
}
