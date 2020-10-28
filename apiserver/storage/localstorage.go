package storage

import (
	"fmt"
	"io/ioutil"

	"github.com/slovak-egov/einvoice/apiserver/config"
)

type LocalStorage struct {
	basePath string
}

func (storage *LocalStorage) invoiceFilename(id int) string {
	return fmt.Sprintf("%s/invoice-%d.xml",storage.basePath, id)
}

func (storage *LocalStorage) GetInvoice(id int) (string, error) {
	return storage.readObject(storage.invoiceFilename(id))
}

func (storage *LocalStorage) SaveInvoice(id int, value string) error {
	return storage.saveObject(storage.invoiceFilename(id), value)
}

func (storage *LocalStorage) attachmentFilename(id int) string {
	return fmt.Sprintf("%s/attachment-%d.xml",storage.basePath, id)
}

func (storage *LocalStorage) GetAttachment(id int) (string, error) {
	return storage.readObject(storage.attachmentFilename(id))
}

func (storage *LocalStorage) SaveAttachment(id int, value string) error {
	return storage.saveObject(storage.attachmentFilename(id), value)
}

func (storage *LocalStorage) saveObject(path, value string) error {
	err := ioutil.WriteFile(path, []byte(value), 0644)
	return err
}

func (storage *LocalStorage) readObject(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	return string(bytes), err
}

func NewLocalStorage(appConfig config.Configuration) *LocalStorage {
	return &LocalStorage{
		appConfig.LocalStorageBasePath,
	}
}
