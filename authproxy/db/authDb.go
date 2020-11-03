package db

import "github.com/slovak-egov/einvoice/authproxy/user"

type AuthDb interface {
	AddUser(user *user.User) error
	GetUser(id string) (*user.User, error)
	UpdateUser(user *user.User) error

	Close() error
}
