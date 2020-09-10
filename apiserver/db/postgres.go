package db

import (
	"github.com/filipsladek/einvoice/apiserver/invoice"
	"github.com/go-pg/pg/v10"
	"io/ioutil"
	"strconv"
)

type dbConnector struct {
	db *pg.DB
}

func NewDBConnector() DBConnector {
	return &dbConnector{}
}

func (connector *dbConnector) Connect(config ConnectionConfig) {
	connector.db = pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
	})
}

func (connector *dbConnector) Close() {
	connector.db.Close()
}

func (connector *dbConnector) InitDB() error {

	query, err := ioutil.ReadFile("sql/setup.sql")
	if err != nil {
		return err
	}

	_, err = connector.db.Exec(string(query))

	return err
}

func (connector *dbConnector) GetAllInvoice() []invoice.Meta {
	var invoices []invoice.Meta
	err := connector.db.Model(&invoices).Select()
	if err != nil {
		panic(err)
	}
	return invoices
}

func (connector *dbConnector) GetInvoiceMeta(id string) *invoice.Meta {
	invoice := &invoice.Meta{}
	err := connector.db.Model(invoice).Where("id = ?", id).Select(invoice)
	if err != nil {
		panic(err)
	}
	return invoice
}

func (connector *dbConnector) CreateInvoice(invoice *invoice.Meta) *invoice.Meta {
	_, err := connector.db.Model(invoice).Insert(invoice)
	if err != nil {
		panic(err)
	}
	return invoice
}
