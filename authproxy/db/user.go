package db

type User struct {
	tableName         struct{} `pg:"users"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	ServiceAccountKey string   `json:"serviceAccountKey,omitempty" pg:"service_account_key"`
	Email             string   `json:"email,omitempty"`
}

func (connector *Connector) GetUser(id string) (*User, error) {
	user := &User{}
	err := connector.db.Model(user).Where("id = ?", id).Select(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (connector *Connector) UpdateUser(user *User) (*User, error) {
	query := connector.db.Model(user).Where("id = ?", user.Id).Returning("*")
	if user.ServiceAccountKey != "" {
		query= query.Set("service_account_key = ?", user.ServiceAccountKey)
	}
	if user.Email != "" {
		query = query.Set("email = ?", user.Email)
	}
	_, err := query.Update()

	return user, err
}

func (connector *Connector) CreateUser(user *User) error {
	_, err := connector.db.Model(user).Insert(user)
	return err
}

