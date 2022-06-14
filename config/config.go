package config

import "mohamadelabror.me/goapiblog/manager"

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

func (c Config) SetConfig() Config {
	c.ApiConfig = ApiConfig{Url: "localhost:6666"}
	c.DbConfig = DbConfig{
		Uri: "mongodb+srv://<username>:<password>@<host>:<port>/?retryWrites=true&w=majority",
	}
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
