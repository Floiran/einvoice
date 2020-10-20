package manager

import (
	"errors"
	"github.com/slovak-egov/einvoice/apiserver/attachment"
	"github.com/slovak-egov/einvoice/apiserver/db"
	"github.com/slovak-egov/einvoice/apiserver/invoice"
	"github.com/slovak-egov/einvoice/apiserver/storage"
	"github.com/slovak-egov/einvoice/apiserver/xml"
	"github.com/slovak-egov/einvoice/apiserver/xml/d16b"
	"github.com/slovak-egov/einvoice/apiserver/xml/ubl21"
	"strconv"
)

type Manager interface {
	CreateInvoice(format, data string, attachments []*attachment.PostAttachment) (*invoice.Meta, error)

	GetFull(id int) (string, error)
	GetMeta(id int) (*invoice.Meta, error)
	GetAllInvoiceMeta(formats []string) ([]invoice.Meta, error)

	GetAttachmentFile(id int) (string, error)
	GetAttachmentMeta(id int) (*attachment.Meta, error)
}

type managerImpl struct {
	db        db.DBConnector
	storage   storage.Storage
	validator xml.Validator
}

func NewManager(db db.DBConnector, storage storage.Storage, validator xml.Validator) Manager {
	return &managerImpl{db, storage, validator}
}

func invoiceFileName(id int) string {
	return strconv.Itoa(id) + "-invoice"
}

func attachmentFileName(id int) string {
	return strconv.Itoa(id) + "-attachment"
}

// todo: must be atomic
func (m *managerImpl) CreateInvoice(format, data string, attachments []*attachment.PostAttachment) (*invoice.Meta, error) {
	var meta *invoice.Meta
	var err error

	switch format {
	case invoice.UblFormat:
		if err = m.validator.ValidateUBL21([]byte(data)); err != nil {
			return nil, err
		}
		err, meta = ubl21.Create(data)
	case invoice.D16bFormat:
		if err = m.validator.ValidateD16B([]byte(data)); err != nil {
			return nil, err
		}
		err, meta = d16b.Create(data)
	default:
		return nil, errors.New("Unknown format")
	}
	if err != nil {
		return nil, err
	}

	if err = m.db.CreateInvoice(meta); err != nil {
		return nil, err
	}
	if err = m.storage.SaveObject(invoiceFileName(meta.Id), data); err != nil {
		return nil, err
	}

	meta.Attachments = []attachment.Meta{}
	for _, at := range attachments {
		atMeta, err := m.db.CreateAttachment(meta.Id, at.Name)
		if err != nil {
			return nil, err
		}
		if err = m.storage.SaveObject(attachmentFileName(atMeta.Id), string(at.Content)); err != nil {
			return nil, err
		}
		meta.Attachments = append(meta.Attachments, *atMeta)
	}

	return meta, nil
}

func (m *managerImpl) GetMeta(id int) (*invoice.Meta, error) {
	meta, err := m.db.GetInvoiceMeta(id)
	if err != nil {
		return nil, err
	}
	ats, err := m.db.GetAttachmentForInvoice(meta.Id)
	if err != nil {
		return nil, err
	}
	meta.Attachments = ats
	return meta, nil
}

func (m *managerImpl) GetAttachmentFile(id int) (string, error) {
	at, err := m.storage.ReadObject(attachmentFileName(id))
	if err != nil {
		return "", err
	}

	return at, nil
}

func (m *managerImpl) GetAttachmentMeta(id int) (*attachment.Meta, error) {
	return m.db.GetAttachment(id)
}

func (m *managerImpl) GetFull(id int) (string, error) {
	meta, err := m.db.GetInvoiceMeta(id)
	if err != nil {
		return "", err
	}
	invoiceStr, err := m.storage.ReadObject(invoiceFileName(meta.Id))
	if err != nil {
		return "", err
	}
	return invoiceStr, nil
}

func (m *managerImpl) GetAllInvoiceMeta(formats []string) ([]invoice.Meta, error) {
	return m.db.GetAllInvoice(formats)
}
