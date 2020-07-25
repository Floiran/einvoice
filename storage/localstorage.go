package storage

import (
	"io/ioutil"
	"github.com/filipsladek/einvoice/common"
)

type LocalStorage struct {
	basePath string
}

func (storage LocalStorage) SaveObject(path, value string) error {
	err := ioutil.WriteFile(storage.basePath+path, []byte(value), 0644)
	return err
}

func (storage LocalStorage) ReadObject(path string) (string, error) {
	_, err := ioutil.ReadFile(storage.basePath + path)
	return "", err
}

func InitLocalStorage() LocalStorage {
	var basePath = common.GetRequiredEnvVariable("LOCAL_STORAGE_BASE_PATH")
	if basePath[len(basePath)-1] != '/' {
		basePath = basePath + "/"
	}

	return LocalStorage{basePath}
}