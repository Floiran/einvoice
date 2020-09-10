package db

import (
	"github.com/filipsladek/einvoice/apiserver/invoice"
	"github.com/filipsladek/einvoice/common"
	"strconv"
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

type DBConnector interface {
	Connect(config ConnectionConfig)
	Close()
	InitDB() error
	GetAllInvoice() []invoice.Meta
	GetInvoiceMeta(id string) *invoice.Meta
	CreateInvoice(invoice *invoice.Meta) *invoice.Meta
}
