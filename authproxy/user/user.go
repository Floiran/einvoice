package user

type User struct {
	Token             string `json:"token,omitempty"`
	Id                string `json:"id"`
	Name              string `json:"name"`
	ServiceAccountKey string `json:"serviceAccountKey,omitempty"`
	Email             string `json:"email,omitempty"`
}
