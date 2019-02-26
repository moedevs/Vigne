package database


type Config struct {
	tokenValue StringValue
	regexValue StringValue

	//TODO: Remove. Legacy for isMod
	Database *Database
}

func (d *Database) createConfig() *Config {
	config := Config{}
	config.Database = d
	cfgHandler := d.NewConfigHandler()
	config.tokenValue = cfgHandler.RequiredValue("token", "Bot 123456789.abcdEFGH")
	config.regexValue = cfgHandler.RequiredValue("commandRegex", `^(?:[-]{2,}>?|[sv]!|—|/|->)\s*([^\s]+)(?:\s(.*))?$`)
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
	return config.Database.Redis.SIsMember(config.Database.Decorate("mods"), id).Val()
}