package postgres

import (
	"github.com/filipsladek/einvoice/apiserver/invoice"
	"github.com/filipsladek/einvoice/common"
	"io/ioutil"
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

func InitDB(connector *DBConnector) {

	query, err := ioutil.ReadFile("sql/setup.sql")
	if err != nil {
		panic(err)
	}

	if _, err := connector.DB.Exec(string(query)); err != nil {
		panic(err)
	}
}

type DBConnector struct {
	DB *pg.DB
}

func (connector *DBConnector) GetAllInvoice() []invoice.Meta {
	var invoices []invoice.Meta
	err := connector.DB.Model(&invoices).Select()
	if err != nil {
		panic(err)
	}
	return invoices
}

func (connector *DBConnector) GetInvoiceMeta(id string) *invoice.Meta {
	invoice := &invoice.Meta{}
	err := connector.DB.Model(invoice).Where("id = ?", id).Select(invoice)
	if err != nil {
		panic(err)
	}
	return invoice
}

func (connector *DBConnector) CreateInvoice(invoice *invoice.Meta) *invoice.Meta {
	_, err := connector.DB.Model(invoice).Insert(invoice)
	if err != nil {
		panic(err)
	}
	return invoice
}
