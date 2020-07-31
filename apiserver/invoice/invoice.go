package invoice

type Invoice struct {
	Id       string `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

func NewInvoice(id, sender, receiver string) Invoice {
	return Invoice{id, sender, receiver}
}
