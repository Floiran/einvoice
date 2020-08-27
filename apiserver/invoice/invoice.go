package invoice

type Invoice struct {
	Id       string `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

type Meta struct {
	Id       string `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

func (invoice *Invoice) GetMeta() *Meta {
	return &Meta{Id: invoice.Id, Sender: invoice.Sender, Receiver: invoice.Receiver}
}
