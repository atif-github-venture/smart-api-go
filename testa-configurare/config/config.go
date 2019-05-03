package config

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
	return &Config{
		DB: &DBConfig{
			MongoUrl:   "mongo:27017",
			Username:   "admin",
			Password:   "admin",
			Collection: "test-configuration",
		},
	}
}
