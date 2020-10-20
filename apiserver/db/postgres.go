package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/slovak-egov/einvoice/apiserver/attachment"
	. "github.com/slovak-egov/einvoice/apiserver/config"
	"github.com/slovak-egov/einvoice/apiserver/invoice"
	"io/ioutil"
)

type dbConnector struct {
	db *pg.DB
}

func NewDBConnector() DBConnector {
	return &dbConnector{}
}

func (connector *dbConnector) Connect() {
	connector.db = pg.Connect(&pg.Options{
		Addr:     Config.Db.Host + ":" + Config.Db.Port,
		User:     Config.Db.User,
		Password: Config.Db.Password,
		Database: Config.Db.Name,
	})
}

// todo: proper setup of DB migration https://github.com/filipsladek/einvoice/issues/60
func (connector *dbConnector) InitDB() error {

	query, err := ioutil.ReadFile("sql/setup.sql")
	if err != nil {
		return err
	}

	_, err = connector.db.Exec(string(query))
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("init")

	return nil
}

func (connector *dbConnector) Close() {
	connector.db.Close()
}

// todo: optimize
func (connector *dbConnector) GetAllInvoice() ([]invoice.Meta, error) {
	var invoices []invoice.Meta
	err := connector.db.Model(&invoices).Select()
	if err != nil {
		return nil, err
	}

	for i := range invoices {
		ats, err := connector.GetAttachmentForInvoice(invoices[i].Id)
		if err != nil {
			return nil, err
		}
		invoices[i].Attachments = ats
	}
	return invoices, nil
}

func (connector *dbConnector) GetInvoiceMeta(id int) (*invoice.Meta, error) {
	inv := &invoice.Meta{}
	err := connector.db.Model(inv).Where("id = ?", id).Select(inv)
	if err != nil {
		return nil, err
	}

	ats, err := connector.GetAttachmentForInvoice(inv.Id)
	if err != nil {
		return nil, err
	}
	inv.Attachments = ats

	return inv, nil
}

func (connector *dbConnector) CreateInvoice(invoice *invoice.Meta) error {
	_, err := connector.db.Model(invoice).Insert(invoice)

	return err
}

func (connector *dbConnector) CreateAttachment(invoiceId int, name string) (*attachment.Meta, error) {
	at := &attachment.Meta{
		InvoiceId: invoiceId,
		Name:      name,
	}
	_, err := connector.db.Model(at).Insert(at)
	if err != nil {
		return nil, err
	}

	return &attachment.Meta{
		Id:   at.Id,
		Name: name,
	}, nil
}

func (connector *dbConnector) GetAttachment(id int) (*attachment.Meta, error) {
	meta := &attachment.Meta{}
	err := connector.db.Model(meta).Where("id = ?", id).Select(meta)
	if err != nil {
		return nil, err
	}
	return &attachment.Meta{
		Id:   meta.Id,
		Name: meta.Name,
	}, nil
}

func (connector *dbConnector) GetAttachmentForInvoice(id int) ([]attachment.Meta, error) {
	var ats []attachment.Meta
	err := connector.db.Model(&ats).Where("invoice_id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	var res = []attachment.Meta{}
	for _, at := range ats {
		res = append(res, attachment.Meta{
			Id:   at.Id,
			Name: at.Name,
		})
	}
	return res, nil
}
