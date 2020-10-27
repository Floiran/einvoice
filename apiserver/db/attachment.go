package db

import (
	"time"
)

type Attachment struct {
	tableName struct{}  `pg:"attachments"`
	InvoiceId int       `json:"-"`
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (connector *Connector) GetAttachment(id int) (*Attachment, error) {
	attachment := &Attachment{}
	err := connector.Db.Model(attachment).Where("id = ?", id).Select(attachment)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

func (connector *Connector) CreateAttachment(invoiceId int, name string) (*Attachment, error) {
	at := &Attachment{
		InvoiceId: invoiceId,
		Name:      name,
	}
	_, err := connector.Db.Model(at).Insert(at)
	if err != nil {
		return nil, err
	}

	return at, nil
}
