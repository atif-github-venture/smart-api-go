package config

import "smartapigo/utility"

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
	var c utility.Prop
	c.ReadProperty()

	return &Config{
		DB: &DBConfig{
			MongoUrl:   c.MongoDBUrl,
			Username:   "admin",
			Password:   "admin",
			Name:       "smart-mongo-core",
			Collection: "projects",
		},
	}
}
