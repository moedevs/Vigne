package database

import (
	"github.com/moedevs/Vigne/errors"
	"strings"
)

type Roles struct {
	Database *Database
}

func (d *Database) Roles() (*Roles, error) {
	roles := Roles{}
	roles.Database = d
	//Does it have a role map?
	hasConfig := d.redis.Exists(d.Decorate("roles")).Val()
	if hasConfig != 1{
		return nil, errors.NoRoles
	}

	return &roles, nil
}

func (r Roles) GetRoleIDFromName(name string) (string, error) {
	return r.Database.redis.HGet(r.Database.Decorate("roles"), strings.ToLower(name)).Result()
}

func (r Roles) GetAllRoles() map[string]string {
	mapped := r.Database.redis.HGetAll(r.Database.Decorate("roles")).Val()
	return mapped
}