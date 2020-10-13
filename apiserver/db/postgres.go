package db

import (
	"github.com/go-pg/pg/v10"
	. "github.com/slovak-egov/einvoice/apiserver/config"
	"github.com/slovak-egov/einvoice/apiserver/invoice"
)

type dbConnector struct {
	db *pg.DB
}

func NewDBConnector() DBConnector {
	return &dbConnector{}
}

func (connector *dbConnector) Connect() {
    if Config.Db.InstanceConnectionName != "" {
        connector.db = pg.Connect(&pg.Options{
            Addr:     "/cloudsql/" + Config.Db.InstanceConnectionName + "/.s.PGSQL.5432",
            Network:  "unix",
            User:     Config.Db.User,
            Password: Config.Db.Password,
            Database: Config.Db.Name,
        })
	} else {
        connector.db = pg.Connect(&pg.Options{
            Addr:     Config.Db.Host + ":" + Config.Db.Port,
            User:     Config.Db.User,
            Password: Config.Db.Password,
            Database: Config.Db.Name,
        })
	}
}

func (connector *dbConnector) Close() {
	connector.db.Close()
}

func (connector *dbConnector) GetAllInvoice() ([]invoice.Meta, error) {
	var invoices []invoice.Meta
	err := connector.db.Model(&invoices).Select()
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (connector *dbConnector) GetInvoiceMeta(id string) (*invoice.Meta, error) {
	inv := &invoice.Meta{}
	err := connector.db.Model(inv).Where("id = ?", id).Select(inv)
	if err != nil {
		return nil, err
	}
	return inv, nil
}

func (connector *dbConnector) CreateInvoice(invoice *invoice.Meta) error {
	_, err := connector.db.Model(invoice).Insert(invoice)

	return err
}
