package config

import "os"

type mongoConfig struct {
	Host   string
	DbName string
}

type grpcClientConfig struct {
}

type Configuration struct {
	Mongo mongoConfig
}

var Env = Configuration{
	Mongo: mongoConfig{
		Host:   os.Getenv("HOST_MONGO"),
		DbName: os.Getenv("DB_MONGO_NAME"),
	},
}
