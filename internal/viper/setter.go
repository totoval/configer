package viper

func (c *Config) Add(name string, configuration map[string]interface{}) {
	c.viper.Set(name, configuration)
}
