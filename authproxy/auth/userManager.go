package auth

import (
	"github.com/slovak-egov/einvoice/authproxy/db"
	"github.com/slovak-egov/einvoice/authproxy/user"
	"github.com/slovak-egov/einvoice/random"
)

type UserManager interface {
	Create(id, name string) *user.User
	GetUserByToken(token string) *user.User
	GetUser(id string) *user.User
	UpdateUser(user, updates *user.User)

	CreateToken(user *user.User) error
	RemoveToken(user *user.User) error
}

type userManager struct {
	db db.AuthDB
}

func NewUserManager(db db.AuthDB) UserManager {
	return userManager{db}
}

func (userManager userManager) Create(id, name string) *user.User {
	usr := &user.User{Id: id, Name: name}

	userManager.db.SaveUser(usr)

	return usr
}

func (userManager userManager) GetUserByToken(token string) *user.User {
	return userManager.db.GetUserByToken(token)
}

func (userManager userManager) GetUser(id string) *user.User {
	return userManager.db.GetUser(id)
}

func (userManager userManager) UpdateUser(user, updates *user.User) {
	if updates.Email != "" {
		user.Email = updates.Email
	}
	if updates.ServiceAccountKey != "" {
		user.ServiceAccountKey = updates.ServiceAccountKey
	}
	userManager.db.SaveUser(user)
}

func (userManager userManager) CreateToken(user *user.User) error {
	user.Token = random.String(50)
	return userManager.db.AddToken(user.Id, user.Token)
}

func (userManager userManager) RemoveToken(user *user.User) error {
	return userManager.db.RemoveToken(user.Id, user.Token)
}
