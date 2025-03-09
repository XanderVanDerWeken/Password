package config

import (
	"log"

	"github.com/greencoda/confiq"
	confiqjson "github.com/greencoda/confiq/loaders/json"
)

type ConfigStruct struct {
	ServerPort       int    `cfg:"serverPort"`
	ConnectionString string `cfg:"connectionString"`
}

func LoadConfig() ConfigStruct {
	configSet := confiq.New()

	if err := configSet.Load(
		confiqjson.Load().FromFile("./config.json"),
	); err != nil {
		log.Fatal(err)
	}

	var cfg ConfigStruct

	if err := configSet.Decode(&cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
