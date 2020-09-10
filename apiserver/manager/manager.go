package manager

import (
	"encoding/json"
	"github.com/filipsladek/einvoice/apiserver/db"
	"github.com/filipsladek/einvoice/apiserver/invoice"
	"github.com/filipsladek/einvoice/apiserver/storage"
	"github.com/filipsladek/einvoice/apiserver/xml/d16b"
	"github.com/filipsladek/einvoice/apiserver/xml/ubl21"
)

type Manager interface {
	Create(invoice *invoice.Invoice) (error, *invoice.Meta)
	CreateUBL(value string) (error, *invoice.Meta)
	CreateD16B(value string) (error, *invoice.Meta)
	CreateJSON(value string) (error, *invoice.Meta)

	GetFull(id string, format string) (error, string)
	GetMeta(id string) (error, *invoice.Meta)
	GetAllInvoiceMeta() []invoice.Meta
}

type managerImpl struct {
	db      db.DBConnector
	storage storage.Storage
}

func NewManager(db db.DBConnector, storage storage.Storage) Manager {
	return &managerImpl{db, storage}
}

func (manager *managerImpl) Create(invoice *invoice.Invoice) (error, *invoice.Meta) {
	meta := invoice.GetMeta()
	manager.db.CreateInvoice(meta)

	jsonString, err := json.Marshal(invoice)
	if err != nil {
		return err, nil
	}
	err = manager.storage.SaveObject("invoice-"+meta.Id+".json", string(jsonString))
	if err != nil {
		return err, nil
	}
	return nil, meta
}

func (manager *managerImpl) CreateJSON(value string) (error, *invoice.Meta) {
	var inv = new(invoice.Invoice)

	if err := json.Unmarshal([]byte(value), &inv); err != nil {
		return err, nil
	}

	meta := inv.GetMeta()
	manager.db.CreateInvoice(meta)

	err := manager.storage.SaveObject("invoice-"+meta.Id+".json", value)
	if err != nil {
		return err, nil
	}
	return nil, meta
}

func (manager *managerImpl) CreateUBL(value string) (error, *invoice.Meta) {
	err, meta := ubl21.Create(value)
	if err != nil {
		return err, nil
	}

	manager.db.CreateInvoice(meta)

	err = manager.storage.SaveObject("invoice-"+meta.Id+".xml", value)
	if err != nil {
		return err, nil
	}
	return nil, meta
}

func (manager *managerImpl) CreateD16B(value string) (error, *invoice.Meta) {
	err, meta := d16b.Create(value)
	if err != nil {
		return err, nil
	}

	manager.db.CreateInvoice(meta)

	err = manager.storage.SaveObject("invoice-"+meta.Id+".xml", value)
	if err != nil {
		return err, nil
	}
	return nil, meta
}

func (manager *managerImpl) GetMeta(id string) (error, *invoice.Meta) {
	return nil, manager.db.GetInvoiceMeta(id)
}

func (manager *managerImpl) GetFull(id string, format string) (error, string) {
	extension := "json"
	if format != invoice.JsonFormat {
		extension = "xml"
	}
	invoiceStr, err := manager.storage.ReadObject("invoice-" + id + "." + extension)
	if err != nil {
		return err, ""
	}
	return nil, invoiceStr
}

func (manager *managerImpl) GetAllInvoiceMeta() []invoice.Meta {
	return manager.db.GetAllInvoice()
}
