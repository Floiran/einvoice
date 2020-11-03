package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/slovak-egov/einvoice/authproxy/config"
	"github.com/slovak-egov/einvoice/authproxy/user"
)

type postgresDb struct {
	Db *pg.DB
}

type UserModel struct {
	tableName         struct{} `pg:"users"`
	Id                string   `pg:"id"`
	Name              string   `pg:"name"`
	ServiceAccountKey string   `pg:"service_account_key"`
	Email             string   `pg:"email"`
}

func New() AuthDb {
	return &postgresDb{
		Db: pg.Connect(&pg.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Config.Db.Host, config.Config.Db.Port),
			User:     config.Config.Db.User,
			Password: config.Config.Db.Password,
			Database: config.Config.Db.Name,
		}),
	}
}

func (psql *postgresDb) Close() error {
	return psql.Db.Close()
}

func (psql *postgresDb) AddUser(usr *user.User) error {
	data := &UserModel{
		Id:                usr.Id,
		Name:              usr.Name,
		ServiceAccountKey: usr.ServiceAccountKey,
		Email:             usr.Email,
	}
	_, err := psql.Db.Model(data).Insert(data)
	return err
}

func (psql *postgresDb) GetUser(id string) (*user.User, error) {
	usr := &UserModel{}
	err := psql.Db.Model(usr).Where("id = ?", id).Select(usr)
	if err != nil {
		return nil, err
	}
	return &user.User{
		Id:                usr.Id,
		Name:              usr.Name,
		ServiceAccountKey: usr.ServiceAccountKey,
		Email:             usr.Email,
	}, nil
}

func (psql *postgresDb) UpdateUser(usr *user.User) error {
	data := &UserModel{
		Id:                usr.Id,
		Name:              usr.Name,
		ServiceAccountKey: usr.ServiceAccountKey,
		Email:             usr.Email,
	}
	q := psql.Db.Model(data).Where("id = ?", data.Id).Returning("*")
	if data.ServiceAccountKey != "" {
		q = q.Set("service_account_key = ?", data.ServiceAccountKey)
	}
	if data.Email != "" {
		q = q.Set("email = ?", data.Email)
	}
	_, err := q.Update(data)

	usr.Name = data.Name
	usr.ServiceAccountKey = data.ServiceAccountKey
	usr.Email = data.Email

	return err
}
