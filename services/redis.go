package services

type RedisService struct {
}

func (RedisService) GetUserKey(account string) string {
	return "user_" + account
}
