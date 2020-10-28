package storage

import (
	"context"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/storage"

	"github.com/slovak-egov/einvoice/apiserver/config"
)

type Gcs struct {
	bkt *storage.BucketHandle
	ctx context.Context
}

func (storage *Gcs) invoiceFilename(id int) string {
	return fmt.Sprintf("invoice-%d", id)
}

func (storage *Gcs) GetInvoice(id int) (string, error) {
	return storage.readObject(storage.invoiceFilename(id))
}

func (storage *Gcs) SaveInvoice(id int, value string) error {
	return storage.saveObject(storage.invoiceFilename(id), value)
}

func (storage *Gcs) attachmentFilename(id int) string {
	return fmt.Sprintf("attachment-%d", id)
}

func (storage *Gcs) GetAttachment(id int) (string, error) {
	return storage.readObject(storage.attachmentFilename(id))
}

func (storage *Gcs) SaveAttachment(id int, value string) error {
	return storage.saveObject(storage.attachmentFilename(id), value)
}

func (storage *Gcs) saveObject(path, value string) error {
	obj := storage.bkt.Object(path)
	w := obj.NewWriter(storage.ctx)

	if _, err := fmt.Fprint(w, value); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

func (storage *Gcs) readObject(path string) (string, error) {
	obj := storage.bkt.Object(path)
	r, err := obj.NewReader(storage.ctx)
	if err != nil {
		return "", err
	}
	defer r.Close()

	res, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func NewGcs(appConfig config.Configuration) *Gcs {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		panic(err)
	}

	bktName := appConfig.GcsBucket
	bkt := client.Bucket(bktName)

	return &Gcs{bkt, ctx}
}
