package redis

type RedisStringValue struct {
	Container RedisContainer
	Key string
}

func (r RedisStringValue) Set(value string) error {
	return r.Container.redis.Set(r.Container.Decorate(r.Key), value, 0).Err()
}

func (r RedisStringValue) Get() string {
	return r.Container.redis.Get(r.Container.Decorate(r.Key)).Val()
}

