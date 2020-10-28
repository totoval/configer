package viper

func (c *Config) Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return c.Get(envName, defaultValue[0])
	}
	return c.Get(envName)
}
