package redis

import (
	"github.com/moedevs/Vigne/database/interfaces"
)

type RedisMapValueString struct {
	RedisMapValue
	Field string
}

func (r RedisMapValueString) Set(value string) error {
	return r.redis.HSet(r.Decorate(r.RedisMapValue.Key), r.Field, value).Err()
}

func (r RedisMapValueString) Get() string {
	return r.redis.HGet(r.Decorate(r.RedisMapValue.Key), r.Field).Val()
}

type RedisMapValue struct {
	RedisContainer
	Key string
}

func (r RedisMapValue) Get(field string) interfaces.StringValue {
	return &RedisMapValueString{
		Field: field,
		RedisMapValue: r,
	}
}

func (r RedisMapValue) Contains(field string) bool {
	return r.redis.HExists(r.Decorate(r.Key), field).Val()
}
