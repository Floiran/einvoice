package manager

import (
	"encoding/json"
	"github.com/filipsladek/einvoice/apiserver/invoice"
	"github.com/filipsladek/einvoice/apiserver/postgres"
	"github.com/filipsladek/einvoice/apiserver/storage"
	"github.com/filipsladek/einvoice/apiserver/xml/d16b"
	"github.com/filipsladek/einvoice/apiserver/xml/ubl21"
)

type Manager interface {
	Create(invoice *invoice.Invoice) (error, *invoice.Invoice, *invoice.Meta)
	CreateUBL(value string) (error, *invoice.Invoice, *invoice.Meta)
	CreateD16B(value string) (error, *invoice.Invoice, *invoice.Meta)
	CreateJSON(value string) (error, *invoice.Invoice, *invoice.Meta)

	Get(id string) (error, *invoice.Invoice)
	GetMeta(id string) (error, *invoice.Meta)
	GetAllInvoiceMeta() []invoice.Meta
}

type managerImpl struct {
	db      postgres.DBConnector
	storage storage.Storage
}

func NewManager(db postgres.DBConnector, storage storage.Storage) Manager {
	return &managerImpl{db, storage}
}

func (manager managerImpl) Create(invoice *invoice.Invoice) (error, *invoice.Invoice, *invoice.Meta) {
	meta := invoice.GetMeta()
	manager.db.CreateInvoice(meta)
	invoice.Id = meta.Id
	jsonString, err := json.Marshal(invoice)
	if err != nil {
		return err, nil, nil
	}
	err = manager.storage.SaveObject("invoice-"+invoice.Id, string(jsonString))
	if err != nil {
		return err, nil, nil
	}
	return nil, invoice, meta
}

func (manager managerImpl) CreateJSON(value string) (error, *invoice.Invoice, *invoice.Meta) {
	var invoice = new(invoice.Invoice)

	if err := json.Unmarshal([]byte(value), &invoice); err != nil {
		return err, nil, nil
	}

	meta := invoice.GetMeta()
	manager.db.CreateInvoice(meta)
	invoice.Id = meta.Id

	jsonString, err := json.Marshal(invoice)
	if err != nil {
		return err, nil, nil
	}
	err = manager.storage.SaveObject("invoice-"+invoice.Id, string(jsonString))
	if err != nil {
		return err, nil, nil
	}
	return nil, invoice, meta
}

func (manager managerImpl) CreateUBL(value string) (error, *invoice.Invoice, *invoice.Meta) {
	err, inv := ubl21.Create(value)
	if err != nil {
		return err, nil, nil
	}

	meta := inv.GetMeta()
	manager.db.CreateInvoice(meta)
	inv.Id = meta.Id

	jsonString, err := json.Marshal(inv)
	if err != nil {
		return err, nil, nil
	}
	err = manager.storage.SaveObject("invoice-"+inv.Id, string(jsonString))
	if err != nil {
		return err, nil, nil
	}
	return nil, inv, meta
}

func (manager managerImpl) CreateD16B(value string) (error, *invoice.Invoice, *invoice.Meta) {
	err, inv := d16b.Create(value)
	if err != nil {
		return err, nil, nil
	}

	meta := inv.GetMeta()
	manager.db.CreateInvoice(meta)
	inv.Id = meta.Id

	jsonString, err := json.Marshal(inv)
	if err != nil {
		return err, nil, nil
	}
	err = manager.storage.SaveObject("invoice-"+inv.Id, string(jsonString))
	if err != nil {
		return err, nil, nil
	}
	return nil, inv, meta
}

func (manager managerImpl) GetMeta(id string) (error, *invoice.Meta) {
	return nil, manager.db.GetInvoiceMeta(id)
}

func (manager managerImpl) Get(id string) (error, *invoice.Invoice) {
	invoiceStr, err := manager.storage.ReadObject("invoice-" + id)
	if err != nil {
		return err, nil
	}
	invoice := &invoice.Invoice{}
	if err := json.Unmarshal([]byte(invoiceStr), &invoice); err != nil {
		return err, nil
	}
	return nil, invoice
}

func (manager managerImpl) GetAllInvoiceMeta() []invoice.Meta {
	return manager.db.GetAllInvoice()
}
