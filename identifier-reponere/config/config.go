package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	MongoUrl   string
	Username   string
	Password   string
	Name       string
	Collection string
}

func GetConfig() *Config {
	var c Prop
	c.ReadProperty()
	return &Config{
		DB: &DBConfig{
			MongoUrl:   c.MongoDBUrl,
			Username:   "admin",
			Password:   "admin",
			Collection: "object-repository",
		},
	}
}

type Prop struct {
	MongoDBUrl string `yaml:"mongourl"`
}

func (c *Prop) ReadProperty() *Prop {
	absPath, _ := filepath.Abs("../smartapigo/property.yml")
	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		print(err)
		absPath, _ = filepath.Abs("../property.yml")
		yamlFile, err = ioutil.ReadFile(absPath)
		if err != nil {
			panic(err)
		}
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}
	return c
}
