package db

import (
	"fmt"

	"github.com/go-pg/pg/v10"

	"github.com/slovak-egov/einvoice/apiserver/config"
)

type Connector struct {
	Db *pg.DB
}

func Connect(dbConfig config.DbConfiguration) Connector {
	return Connector{
		Db: pg.Connect(&pg.Options{
			Addr:     fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
			User:     dbConfig.User,
			Password: dbConfig.Password,
			Database: dbConfig.Name,
		}),
	}
}

func (connector *Connector) Close() {
	connector.Db.Close()
}

// Note: Used only in tests
func (connector *Connector) ClearData() {
	connector.Db.Model(&Invoice{}).Where("TRUE").Delete()
	connector.Db.Model(&Attachment{}).Where("TRUE").Delete()
}
