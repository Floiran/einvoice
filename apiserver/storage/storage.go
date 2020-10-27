package storage

import (
	"github.com/slovak-egov/einvoice/apiserver/config"
)

type Storage interface {
	SaveInvoice(id int, value string) error
	GetInvoice(id int) (string, error)
	SaveAttachment(id int, value string) error
	GetAttachment(id int) (string, error)
}

func Init(appConfig config.Configuration) Storage {
	var storage Storage

	switch appConfig.SlowStorageType {
	case "local":
		storage = NewLocalStorage(appConfig)
	case "gcs":
		storage = NewGcs(appConfig)
	}

	return storage
}
