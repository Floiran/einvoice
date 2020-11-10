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
	return m.Cache.GetUserId(token)
}

func (m *Manager) GetUser(id string) (*db.User, error) {
	return m.Db.GetUser(id)
}

func (m *Manager) UpdateUser(id string, user *db.UserUpdate) (*db.User, error) {
	return m.Db.UpdateUser(id, user)
}

func (m *Manager) LogoutUser(token string) error {
	deleted := m.Cache.RemoveUserToken(token)
	if !deleted {
		return errors.New("Token not found")
	}

	return nil
}

func (m *Manager) CreateUser(id, name string) (*db.User, error) {
	user := &db.User{Id: id, Name: &name}
	return m.Db.CreateUser(user)
}

func (m *Manager) CreateUserToken(userId string) string {
	token := random.String(50)
	m.Cache.SaveUserToken(token, userId)
	return token
}
