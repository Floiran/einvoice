package db

type User struct {
	tableName         struct{} `pg:"users"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	ServiceAccountKey string   `json:"serviceAccountKey"`
	Email             string   `json:"email"`
}

type UserUpdate struct {
	UserId            string
	ServiceAccountKey *string `json:"serviceAccountKey"`
	Email             *string `json:"email"`
}

func (user *UserUpdate) IsEmpty() bool {
	return user.ServiceAccountKey == nil &&
		user.Email == nil
}

func (connector *Connector) GetUser(id string) (*User, error) {
	user := &User{}
	err := connector.Db.Model(user).Where("id = ?", id).Select(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (connector *Connector) UpdateUser(updatedUserData *UserUpdate) (*User, error) {
	user := &User{}
	query := connector.Db.Model(user).Where("id = ?", updatedUserData.UserId).Returning("*")
	if updatedUserData.ServiceAccountKey != nil {
		query = query.Set("service_account_key = ?", *updatedUserData.ServiceAccountKey)
	}
	if updatedUserData.Email != nil {
		query = query.Set("email = ?", *updatedUserData.Email)
	}
	_, err := query.Update()

	return user, err
}

func (connector *Connector) CreateUser(user *User) error {
	_, err := connector.Db.Model(user).Insert(user)
	return err
}
