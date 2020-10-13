package storage

import (
	"errors"
	"fmt"
	. "github.com/slovak-egov/einvoice/apiserver/config"
)

type Storage interface {
	SaveObject(path, value string) error
	ReadObject(path string) (string, error)
}

func InitStorage() Storage {
	var storage Storage
	var storageType = Config.SlowStorageType

	switch storageType {
	case "local":
		storage = NewLocalStorage()
	case "gcs":
		storage = NewGSC()
	default:
		panic(errors.Unwrap(fmt.Errorf("unsupported storage type %q. Supported values are local, gcs, s3", storageType)))
	}

	return storage
}
