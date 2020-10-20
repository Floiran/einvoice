package attachment

import "time"

type PostAttachment struct {
	Name    string
	Content []byte
}

type Meta struct {
	tableName struct{}  `pg:"attachments"`
	InvoiceId int       `json:"-"`
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
