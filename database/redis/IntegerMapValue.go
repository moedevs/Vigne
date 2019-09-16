package redis

import (
	"github.com/moedevs/Vigne/database/interfaces"
)

type RedisIntegerMapValue struct {
	RedisIntegerMap
	Field string
}

func (r RedisIntegerMapValue) Set(value int) error {
	return r.redis.HSet(r.Decorate(r.Key), r.Field, int64(value)).Err()
}

func (r RedisIntegerMapValue) Get() (int, error) {
	return r.redis.HGet(r.Decorate(r.Key), r.Field).Int()
}

func (r RedisIntegerMapValue) Add(amount int) error {
	return r.redis.HIncrBy(r.Decorate(r.Key), r.Field, int64(amount)).Err()
}

type RedisIntegerMap struct {
	RedisContainer
	Key string
}


func (r RedisIntegerMap) Get(field string) interfaces.IntegerValue {
	return &RedisIntegerMapValue{
		Field: field,
		RedisIntegerMap: r,
	}
}

func (r RedisIntegerMap) Contains(field string) bool {
	return r.redis.HExists(r.Decorate(r.Key), field).Val()
}
