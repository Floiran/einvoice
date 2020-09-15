package db

import (
	"github.com/slovak-egov/einvoice/apiserver/invoice"
	"github.com/slovak-egov/einvoice/common"
)

type ConnectionConfig struct {
	ConnectionName     string
	User     string
	Password string
	Database string
}

func NewConnectionConfig() ConnectionConfig {
	return ConnectionConfig{
		ConnectionName:     common.GetRequiredEnvVariable("DB_INSTANCE_CONNECTION_NAME"),
		User:     common.GetRequiredEnvVariable("DB_USER"),
		Password: common.GetRequiredEnvVariable("DB_PASSWORD"),
		Database: common.GetRequiredEnvVariable("DB_NAME"),
	}
}

type DBConnector interface {
	Connect(config ConnectionConfig)
	Close()
	InitDB() error
	GetAllInvoice() ([]invoice.Meta, error)
	GetInvoiceMeta(id string) (*invoice.Meta, error)
	CreateInvoice(invoice *invoice.Meta) error
}
