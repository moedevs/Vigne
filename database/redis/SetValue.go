package redis

type RedisSetValue struct {
	RedisContainer
	Key string
}

func (r RedisSetValue) Add(member string) error {
	return r.redis.SAdd(r.Decorate(r.Key), member).Err()
}

func (r RedisSetValue) Contains(member string) bool {
	return r.redis.SIsMember(r.Decorate(r.Key), member).Val()
}

func (r RedisSetValue) Remove(member string) error {
	return r.redis.SRem(r.Decorate(r.Key), member).Err()
}
