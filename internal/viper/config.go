package viper

import (
	"github.com/spf13/viper"
	"github.com/totoval/configer/internal"
)

type Config struct {
	viper *viper.Viper
}

func (c *Config) SetConfigFile(path string, filename string, extension string) error {
	c.viper.SetConfigName(filename)  // .env
	c.viper.SetConfigType(extension) // json
	c.viper.AddConfigPath(path)      // .

	err := c.viper.ReadInConfig() // Find and read the config file
	return err
}

func (c *Config) SetEnvPrefix(prefix string) error {
	c.viper.SetEnvPrefix(prefix) // totoval
	c.viper.AutomaticEnv()
	return nil
}

func (c *Config) New() internal.Configer {
	c.viper = viper.New()
	return c
}
