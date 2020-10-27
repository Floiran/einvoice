package manager

import (
	"github.com/slovak-egov/einvoice/apiserver/db"
	"github.com/slovak-egov/einvoice/apiserver/xml/d16b"
	"github.com/slovak-egov/einvoice/apiserver/xml/ubl21"
)

func (m *Manager) GetInvoices(formats []string) ([]db.Invoice, error) {
	return m.Db.GetInvoices(formats)
}

func (m *Manager) GetInvoice(id int) (*db.Invoice, error) {
	inv, err := m.Db.GetInvoice(id)
	if err != nil {
		return nil, err
	}
	ats, err := m.Db.GetInvoiceAttachments(inv.Id)
	if err != nil {
		return nil, err
	}
	inv.Attachments = ats
	return inv, nil
}

func (m *Manager) GetInvoiceDetail(id int) (string, error) {
	return m.storage.GetInvoice(id)
}

func (m *Manager) CreateInvoice(format, data string, attachments []*Attachment) (*db.Invoice, error) {
	var metadata *db.Invoice
	var err error
	switch format {
	case db.UblFormat:
		metadata, err = ubl21.Create(data)
	case db.D16bFormat:
		metadata, err = d16b.Create(data)
	}

	if err != nil {
		return nil, err
	}
	// TODO: make DB+storage+attachments saving atomic
	if err := m.Db.CreateInvoice(metadata); err != nil {
		return nil, err
	}
	if err := m.storage.SaveInvoice(metadata.Id, data); err != nil {
		return nil, err
	}

	metadata.Attachments, err = m.createAttachments(metadata.Id, attachments)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}
