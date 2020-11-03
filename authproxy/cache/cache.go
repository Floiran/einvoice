package cache

type Cache interface {
	SaveToken(token, userId string)
	GetUserId(token string) (string, error)
	RemoveToken(token string) bool
}
