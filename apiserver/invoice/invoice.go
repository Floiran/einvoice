package invoice

import (
	"github.com/slovak-egov/einvoice/apiserver/attachment"
	"time"
)

const (
	UblFormat  = "ubl2.1"
	D16bFormat = "d16b"
)

type Meta struct {
	tableName   struct{}          `pg:"invoices"`
	Id          int               `json:"id"`
	Sender      string            `json:"sender"`
	Receiver    string            `json:"receiver"`
	Format      string            `json:"format"`
	Price       float64           `json:"price"`
	CreatedAt   time.Time         `json:"created_at"`
	Attachments []attachment.Meta `json:"attachments" pg:"-"`
}
