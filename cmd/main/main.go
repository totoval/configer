package main

import (
	"fmt"
	"github.com/totoval/configer/pkg"
	"github.com/totoval/configer/pkg/facade"
)

var config facade.Configer

func main() {
	// use driver then config
	c := &pkg.Config{}
	if err := c.Use(pkg.DriverViper).Config(map[string]interface{}{
		"path":       "configs",
		"filename":   ".env",
		"extension":  "json",
		"env_prefix": "totoval",
	}); err != nil {
		panic(err)
	}
	// get the facade
	config = c.Component().(facade.Configer)

	fmt.Println(config.Get("hello"))
}
