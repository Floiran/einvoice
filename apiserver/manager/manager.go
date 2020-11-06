package manager

import (
	"github.com/slovak-egov/einvoice/apiserver/config"
	"github.com/slovak-egov/einvoice/apiserver/db"
	"github.com/slovak-egov/einvoice/apiserver/storage"
)

type Manager struct {
	Db      db.Connector
	storage storage.Storage
}

func Init(appConfig config.Configuration) Manager {
	return Manager{
		Db:      db.Connect(appConfig.Db),
		storage: storage.Init(appConfig),
	}
}
