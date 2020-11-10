package db

import (
	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/json"
)

type User struct {
	tableName         struct{} `pg:"users"`
	Id                string   `json:"id"`
	Name              *string  `json:"name"`
	ServiceAccountKey *string  `json:"serviceAccountKey"`
	Email             *string  `json:"email"`
	TaxId             *int     `json:"taxId"`
	VatNumber         *int     `json:"vatNumber"`
}

type UserUpdate struct {
	ServiceAccountKey *string      `json:"serviceAccountKey"`
	Email             *string      `json:"email"`
	TaxId             json.JSONInt `json:"taxId"`
	VatNumber         json.JSONInt `json:"vatNumber"`
}

func (connector *Connector) GetUser(id string) (*User, error) {
	user := &User{}
	err := connector.Db.Model(user).Where("id = ?", id).Select(user)
	if err != nil {
		log.WithField("error", err.Error()).Warn("db.get_user")
		return nil, err
	}
	return user, nil
}

// TODO: refactor this
func (connector *Connector) UpdateUser(id string, userUpdate *UserUpdate) (*User, error) {
	user := &User{}
	query := connector.Db.Model(user).Where("id = ?", id).Returning("*")
	if userUpdate.ServiceAccountKey != nil {
		query = query.Set("service_account_key = ?", *userUpdate.ServiceAccountKey)
	}
	if userUpdate.Email != nil {
		query = query.Set("email = ?", *userUpdate.Email)
	}
	if userUpdate.TaxId.Set {
		if userUpdate.TaxId.Value == nil {
			query = query.Set("tax_id = NULL")
		} else {
			query = query.Set("tax_id = ?", *userUpdate.TaxId.Value)
		}
	}
	if userUpdate.VatNumber.Set {
		if userUpdate.VatNumber.Value == nil {
			query = query.Set("vat_number = NULL")
		} else {
			query = query.Set("vat_number = ?", *userUpdate.VatNumber.Value)
		}
	}
	_, err := query.Update()

	if err != nil {
		log.WithField("error", err.Error()).Warn("db.update_user")
	}

	return user, err
}

func (connector *Connector) CreateUser(user *User) (*User, error) {
	_, err := connector.Db.Model(user).Insert(user)

	if err != nil {
		log.WithField("error", err.Error()).Warn("db.create_user")
	}

	return user, err
}
