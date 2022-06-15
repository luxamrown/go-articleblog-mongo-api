package config

import (
	"fmt"
	"os"

	"mohamadelabror.me/goapiblog/manager"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	Uri string
}

type Manager struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

type Config struct {
	Manager
	ApiConfig
	DbConfig
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func (c Config) SetConfig() Config {
	c.ApiConfig = ApiConfig{Url: GetPort()}
	c.DbConfig = DbConfig{
		Uri: "mongodb+srv://luxamrown:%40Bulungan2018@blog.yzqxr.mongodb.net/?retryWrites=true&w=majority",
	}
	fmt.Println(c.ApiConfig)
	return c
}

func NewConfig() Config {
	cfg := Config{}
	cfg = cfg.SetConfig()

	dataSourceName := cfg.DbConfig.Uri
	cfg.InfraManager = manager.NewInfra(dataSourceName)
	cfg.RepoManager = manager.NewRepoManager(cfg.InfraManager)
	cfg.UseCaseManager = manager.NewUseCaseManager(cfg.RepoManager)
	return cfg
}
