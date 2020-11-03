package auth

import (
	log "github.com/sirupsen/logrus"
	"github.com/slovak-egov/einvoice/authproxy/cache"
	"github.com/slovak-egov/einvoice/authproxy/db"
	"github.com/slovak-egov/einvoice/authproxy/user"
	"github.com/slovak-egov/einvoice/random"
)

type UserManager interface {
	Create(id, name string) (*user.User, error)
	GetUserIdByToken(token string) (string, error)
	GetUser(id string) (*user.User, error)
	UpdateUser(updates *user.User) error

	CreateToken(user *user.User)
	RemoveToken(token string) bool
}

type userManager struct {
	db    db.AuthDb
	cache cache.Cache
}

func NewUserManager(db db.AuthDb, cache cache.Cache) UserManager {
	return &userManager{db, cache}
}

func (userManager *userManager) Create(id, name string) (*user.User, error) {
	usr := &user.User{Id: id, Name: name}

	err := userManager.db.AddUser(usr)
	if err != nil {
		log.WithField("error", err.Error()).Error("manager.user.create.failed")
		return nil, err
	}
	return usr, nil
}

func (userManager *userManager) GetUserIdByToken(token string) (string, error) {
	id, err := userManager.cache.GetUserId(token)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (userManager *userManager) GetUser(id string) (*user.User, error) {
	usr, err := userManager.db.GetUser(id)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (userManager *userManager) UpdateUser(updates *user.User) error {
	return userManager.db.UpdateUser(updates)
}

func (userManager *userManager) CreateToken(user *user.User) {
	user.Token = random.String(50)
	userManager.cache.SaveToken(user.Token, user.Id)
}

func (userManager *userManager) RemoveToken(token string) bool {
	return userManager.cache.RemoveToken(token)
}
