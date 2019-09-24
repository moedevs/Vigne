package redis

import (
	"github.com/go-redis/redis"
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
	val, err :=  r.redis.HGet(r.Decorate(r.Key), r.Field).Int()
	if err == redis.Nil {
		return 0, nil
	}
	return val, err
}

func (r RedisIntegerMapValue) Add(amount int) error {
	return r.redis.HIncrBy(r.Decorate(r.Key), r.Field, int64(amount)).Err()
}

type RedisIntegerMap struct {
	RedisContainer
	Key string
}

func (r RedisIntegerMap) GetAll() (map[string]interfaces.IntegerValue, error) {
	result, err := r.redis.HGetAll(r.Decorate(r.Key)).Result()
	if err != nil {
		return nil, err
	}
	o := map[string]interfaces.IntegerValue{}
	for field, _ := range result {
		o[field] = r.Get(field)
	}
	return o, nil
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
