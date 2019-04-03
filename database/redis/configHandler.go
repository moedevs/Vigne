package redis

import "github.com/moedevs/Vigne/database/interfaces"

func NewConfigHandler(container interfaces.Container) interfaces.Config {
	cfg := ConfigHandler{}

	cfg.Map = container.Map("config")

	return &cfg
}

type ConfigHandler struct {
	Map interfaces.MapValue
}

func (cfg ConfigHandler) OptionalValue(name string) (value interfaces.StringValue, exists bool) {
	return cfg.Map.Get(name), cfg.Map.Contains(name)
}

func (cfg ConfigHandler) RequiredValue(name string, defaultValue string) interfaces.StringValue {
	if !cfg.Map.Contains(name) {
		cfg.Map.Get(name).Set(defaultValue)
	}
	value, _ := cfg.OptionalValue(name)
	return value
}