package cache

type Cache interface {
	SaveUserToken(token, userId string)
	GetUserId(token string) (string, error)
	RemoveUserToken(token string) bool
}
