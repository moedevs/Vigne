package redis

import (
	"github.com/go-redis/redis"
	"github.com/moedevs/Vigne/database/interfaces"
)

type RedisContainer struct{
	prefix string
	redis *redis.Client
}



func CreateContainer(prefix string, redis *redis.Client) *RedisContainer {
	container := RedisContainer{}

	container.prefix = prefix
	container.redis = redis

	return &container
}

func (r RedisContainer) GetContainer(key string) interfaces.Container {
	return CreateContainer(r.Decorate(key), r.redis)
}

func (r RedisContainer) Decorate(key string) string {
	return r.prefix + ":" + key
}
func (r RedisContainer) Value(key string) interfaces.StringValue {
	return &RedisStringValue{
		Key: key,
		Container:r,
	}
}

func (r RedisContainer) Integer(key string) interfaces.IntegerValue {
	return &RedisIntegerValue{
		Key: key,
		Container:r,
	}
}

func (r RedisContainer) Map(key string) interfaces.MapValue {
	return &RedisMapValue{
		Key:key,
		RedisContainer: r,
	}
}

func (r RedisContainer) Set(key string) interfaces.SetValue {
	return &RedisSetValue{
		Key:key,
		RedisContainer:r,
	}
}
