package viper

import (
	"github.com/spf13/cast"
)

func (c *Config) Get(path string, defaultValue ...interface{}) interface{} {
	if !c.viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return c.viper.Get(path)
}
func (c *Config) GetInterface(path string, defaultValue ...interface{}) (value interface{}) {

	if len(defaultValue) > 0 {
		value = c.Get(path, defaultValue[0])
	} else {
		value = c.Get(path)
	}

	return value
}
func (c *Config) GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(c.GetInterface(path, defaultValue...))
}
func (c *Config) GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(c.GetInterface(path, defaultValue...))
}
func (c *Config) GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(c.GetInterface(path, defaultValue...))
}
func (c *Config) GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(c.GetInterface(path, defaultValue...))
}
func (c *Config) GetBool(path string, defaultValue ...interface{}) bool {

	var value interface{}
	if len(defaultValue) > 0 {
		value = c.Get(path, defaultValue[0])
	} else {
		value = c.Get(path)
	}

	return cast.ToBool(value)
}
