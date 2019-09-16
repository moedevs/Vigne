package redis

type RedisIntegerValue struct {
	Container RedisContainer
	Key string
}

func (r RedisIntegerValue) Set(value int) error {
	return r.Container.redis.Set(r.Container.Decorate(r.Key), value, 0).Err()
}

func (r RedisIntegerValue) Get() (int, error) {
	return r.Container.redis.Get(r.Container.Decorate(r.Key)).Int()
}

func (r RedisIntegerValue) Add(amount int) error {
	return r.Container.redis.IncrBy(r.Container.Decorate(r.Key), int64(amount)).Err()
}
