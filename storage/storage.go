package storage

import (
	"errors"
	"fmt"
	"github.com/filipsladek/einvoice/common"
)

type Storage interface {
	SaveObject(path, value string) error
	ReadObject(path string) (string, error)
}

func InitStorage() Storage {
	var storage Storage
	var storageType = common.GetRequiredEnvVariable("SLOW_STORAGE_TYPE")

	switch storageType {
	case "local":
		storage = NewLocalStorage()
	default:
		_ =errors.Unwrap(fmt.Errorf("Unsupported storage type %w. Supported values are local, gcs, s3", storageType))
	}

	return storage
}
