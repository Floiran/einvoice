package storage

import (
	. "github.com/slovak-egov/einvoice/apiserver/config"
	"io/ioutil"
)

type LocalStorage struct {
	basePath string
}

func (storage *LocalStorage) SaveObject(path, value string) error {
	err := ioutil.WriteFile(storage.basePath+path, []byte(value), 0644)
	return err
}

func (storage *LocalStorage) ReadObject(path string) (string, error) {
	bytes, err := ioutil.ReadFile(storage.basePath + path)
	return string(bytes), err
}

func NewLocalStorage() *LocalStorage {
	var basePath = Config.LocalStorageBasePath
	if basePath[len(basePath)-1] != '/' {
		basePath = basePath + "/"
	}

	return &LocalStorage{basePath}
}
