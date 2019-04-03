package database

import (
	"github.com/go-redis/redis"
	"github.com/moedevs/Vigne/database/interfaces"
	redis2 "github.com/moedevs/Vigne/database/redis"
)

type Database struct {
	//redis      *redis.Client
	Identifier string
	config     *Config
	interfaces.Container

	redis *redis.Client
}

func NewDatabase(identifier, address, password string) *Database {
	d := Database{}
	d.Identifier = identifier
	redis := redis.NewClient(&redis.Options{
		Addr: address,
		Password: password,
	})
	//Setup container
	d.Container = redis2.CreateContainer(identifier, redis)
	d.redis = redis
	return &d
}