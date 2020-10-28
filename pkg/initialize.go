package pkg

import (
	"errors"
	"github.com/totoval/configer/internal"
	"github.com/totoval/configer/internal/viper"
)

type Config struct {
	config internal.Configer
}

func (c *Config) Component() interface{} {
	return c.config
}

func (c *Config) Use(driver string) Componentor {
	switch driver {
	case DriverViper:
		c.config = &viper.Config{}
	default:
		c.config = &viper.Config{}
	}
	return c
}

func (c *Config) Config(configuration map[string]interface{}) error {
	config := c.config.New()

	// PRIORITY: env > file

	// use config file firstly
	path, ok := configuration["path"].(string)
	if !ok {
		return errors.New("unknown configuration config path")
	}
	filename, ok := configuration["filename"].(string)
	if !ok {
		return errors.New("unknown configuration config filename")
	}
	extension, ok := configuration["extension"].(string)
	if !ok {
		return errors.New("unknown configuration config extension")
	}
	if err := config.SetConfigFile(path, filename, extension); err != nil {
		return err
	}

	// use config env secondly
	envPrefix, ok := configuration["env_prefix"].(string)
	if !ok {
		return errors.New("unknown configuration config env_prefix")
	}
	if err := config.SetEnvPrefix(envPrefix); err != nil {
		return err
	}

	return nil
}
