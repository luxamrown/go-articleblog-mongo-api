package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mohamadelabror.me/goapiblog/config"
	"mohamadelabror.me/goapiblog/delivery/api"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine *gin.Engine
	cfg          config.Config
}

func (a *appServer) initHandlers() {
	a.v1()
}

func (a *appServer) v1() {
	articleApiGroup := a.routerEngine.Group("/article")
	api.NewArticleApi(articleApiGroup, a.cfg.UseCaseManager.CreateArticleUseCase(), a.cfg.UseCaseManager.GetAllArticlUseCase(), a.cfg.UseCaseManager.GetArticleUseCase(), a.cfg.UseCaseManager.DeleteArticleUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.routerEngine.Run(a.cfg.ApiConfig.Url)
	if err != nil {
		panic(err)
	}
}

func Server() AppServer {
	r := gin.Default()
	r.Use(cors.Default())
	c := config.NewConfig()
	return &appServer{
		routerEngine: r,
		cfg:          c,
	}
}
