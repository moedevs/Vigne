package database

import (
	"github.com/moedevs/Vigne/database/interfaces"
	"github.com/moedevs/Vigne/database/redis"
)

type Config struct {
	tokenValue interfaces.StringValue
	regexValue interfaces.StringValue
	modsSet interfaces.SetValue
}

func (d *Database) createConfig() *Config {
	config := Config{}
	cfgHandler := redis.NewConfigHandler(d.Container)
	config.tokenValue = cfgHandler.RequiredValue("token", "Bot 123456789.abcdEFGH")
	config.regexValue = cfgHandler.RequiredValue("commandRegex", `^(?:[-]{2,}>?|[sv]!|—|/|->)\s*([^\s]+)(?:\s(.*))?$`)
	config.modsSet = d.Set("mods")
	return &config
}

func (d *Database) Config() *Config {
	if d.config == nil {
		d.config = d.createConfig()
	}
	return d.config
}

func (config Config) Token() string {
	return config.tokenValue.Get()
}

func (config Config) CommandRegex() string {
	return config.regexValue.Get()
}

func (config Config) IsMod(id string) bool {
	return config.modsSet.Contains(id)
}