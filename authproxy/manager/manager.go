package manager

import (
	"errors"

	"github.com/slovak-egov/einvoice/authproxy/cache"
	"github.com/slovak-egov/einvoice/authproxy/config"
	"github.com/slovak-egov/einvoice/authproxy/db"
	"github.com/slovak-egov/einvoice/random"
)

type Manager struct {
	Cache cache.Cache
	Db    db.Connector
}

func Init(appConfig config.Configuration) Manager {
	return Manager{
		cache.NewRedis(appConfig.RedisUrl, appConfig.TokenExpiration),
		db.Connect(appConfig.Db),
	}
}

func (m *Manager) GetUserIdByToken(token string) (string, error) {
	id, err := m.Cache.GetUserId(token)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (m *Manager) GetUser(id string) (*db.User, error) {
	user, err := m.Db.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *Manager) UpdateUser(user *db.UserUpdate) (*db.User, error) {
	return m.Db.UpdateUser(user)
}

func (m *Manager) LogoutUser(token string) error {
	deleted := m.Cache.RemoveUserToken(token)
	if !deleted {
		return errors.New("Token not found")
	}

	return nil
}

func (m *Manager) CreateUser(id, name string) (*db.User, error) {
	user := &db.User{Id: id, Name: name}
	err := m.Db.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *Manager) CreateUserToken(userId string) string {
	token := random.String(50)
	m.Cache.SaveUserToken(token, userId)
	return token
}
