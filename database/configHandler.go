package database

func (d *Database) NewConfigHandler() *ConfigHandler {
	cfg := ConfigHandler{}

	cfg.d = d
	cfg.key = "config"

	return &cfg
}

type ConfigHandler struct {
	d   *Database
	key string
}

func (cfg ConfigHandler) OptionalValue(name string) (value StringValue, exists bool) {
	return ConfigStringValue{key: cfg.key, field:name, d: cfg.d}, cfg.d.Redis.HExists(cfg.d.Decorate(cfg.key), name).Val()
}

func (cfg ConfigHandler) RequiredValue(name string, defaultValue string) StringValue {
	if !cfg.d.Redis.HExists(cfg.d.Decorate(cfg.key), name).Val() {
		cfg.d.Redis.HSet(cfg.d.Decorate(cfg.key), name, defaultValue)
	}
	value, _ := cfg.OptionalValue(name)
	return value
}

type ConfigStringValue struct {
	key   string
	field string
	d     *Database
}

func (v ConfigStringValue) Set(value string) error {
	return v.d.Redis.HSet(v.d.Decorate(v.key), v.field, value).Err()
}

func (v ConfigStringValue) Get() string {
	value, err := v.d.Redis.HGet(v.d.Decorate(v.key), v.field).Result()
	if err != nil {
		return ""
	}
	return value
}