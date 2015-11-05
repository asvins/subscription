package main

import (
	"fmt"

	"github.com/asvins/common_db/postgres"
	"gopkg.in/gcfg.v1"
)

// Config struct for this service
type Config struct {
	Server struct {
		Addr string
		Port string
	}
	Service struct {
		Env string
	}
	Database struct {
		User    string
		DbName  string
		SSLMode string
	}
}

func LoadConfig() Config {
	cfg := Config{}
	err := gcfg.ReadFileInto(&cfg, "subscription_config.gcfg")
	if err != nil {
		fmt.Println("Error while loading config: %s", err.Error())
		return Config{}
	}
	return cfg
}

func DBConfig() *postgres.Config {
	var pcfg postgres.Config
	pcfg = postgres.Config(LoadConfig().Database)
	return &pcfg
}
