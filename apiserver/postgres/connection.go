package postgres

import (
	. "github.com/filipsladek/einvoice/apiserver/invoice"
	"github.com/filipsladek/einvoice/common"
	"strconv"

	"github.com/go-pg/pg/v10"
)

type ConnectionConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewConnectionConfig() ConnectionConfig {
	port, _ := strconv.Atoi(common.GetRequiredEnvVariable("DB_PORT"))
	return ConnectionConfig{
		Host:     common.GetRequiredEnvVariable("DB_HOST"),
		Port:     port,
		User:     common.GetRequiredEnvVariable("DB_USER"),
		Password: common.GetRequiredEnvVariable("DB_PASSWORD"),
		Database: common.GetRequiredEnvVariable("DB_NAME"),
	}
}

func Connect(config ConnectionConfig) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
	})

	return db
}

type DBConnector struct {
	DB *pg.DB
}

func (connector DBConnector) GetAllInvoice() []Invoice {
	var invoices []Invoice
	err := connector.DB.Model(&invoices).Select()
	if err != nil {
		panic(err)
	}
	return invoices
}

func (connector DBConnector) GetInvoice(id string) *Invoice {
	invoice := &Invoice{}
	err := connector.DB.Model(invoice).Where("id = ?", id).Select(invoice)
	if err != nil {
		panic(err)
	}
	return invoice
}

func (connector DBConnector) CreateInvoice(invoice *Invoice) *Invoice {
	_, err := connector.DB.Model(invoice).Insert(invoice)
	if err != nil {
		panic(err)
	}
	return invoice
}
