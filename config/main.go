package config

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var Config config

type config struct {
	Debug    bool     `yaml:"debug"`
	Database Database `yaml:"database"`
}

type Environment struct {
	Development config `yaml:"development"`
	Production  config `yaml:"production"`
}

type Database struct {
	Path string `yaml:"path"`
}

func (c *config) IsDev() bool {
	return c.Debug == true
}

func (c *config) IsProduction() bool {
	return c.Debug == false
}

func init() {
	env := "development"
	flag.Parse()
	if args := flag.Args(); 0 < len(args) && args[0] == "production" {
		env = "production"
	}

	buf, err := ioutil.ReadFile("application.yml")
	if err != nil {
		panic(err)
	}

	var environment Environment
	err = yaml.Unmarshal(buf, &environment)
	if err != nil {
		panic(err)
	}

	if env == "production" {
		Config = environment.Production
	} else {
		Config = environment.Development
	}
}
