package db

import (
	"fmt"

	"github.com/go-pg/pg/v10"

	"github.com/slovak-egov/einvoice/authproxy/config"
)

type Connector struct {
	db *pg.DB
}

func Connect(dbConfig config.DbConfiguration) Connector {
	return Connector{
		db: pg.Connect(&pg.Options{
			Addr:     fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
			User:     dbConfig.User,
			Password: dbConfig.Password,
			Database: dbConfig.Name,
		}),
	}
}

func (connector *Connector) Close() {
	connector.db.Close()
}
