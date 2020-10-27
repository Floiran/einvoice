package manager

import (
	"github.com/slovak-egov/einvoice/apiserver/db"
)

type Attachment struct {
	Name    string
	Content []byte
}

func (m *Manager) GetAttachment(id int) (string, string, error) {
	metadata, err := m.Db.GetAttachment(id)
	if err != nil {
		return "", "", err
	}
	at, err := m.storage.GetAttachment(id)
	if err != nil {
		return "", "", err
	}
	return at, metadata.Name, nil
}

func (m *Manager) createAttachments(invoiceId int, attachments []*Attachment) ([]db.Attachment, error) {
	attachmentsMeta := []db.Attachment{}
	for _, at := range attachments {
		atMeta, err := m.Db.CreateAttachment(invoiceId, at.Name)
		if err != nil {
			return nil, err
		}
		if err = m.storage.SaveAttachment(atMeta.Id, string(at.Content)); err != nil {
			return nil, err
		}
		attachmentsMeta = append(attachmentsMeta, *atMeta)
	}
	return attachmentsMeta, nil
}
