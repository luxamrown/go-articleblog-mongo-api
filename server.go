package main

import (
	"os"

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
	api.NewArticleApi(articleApiGroup, a.cfg.UseCaseManager.CreateArticleUseCase(), a.cfg.UseCaseManager.GetAllArticlUseCase(), a.cfg.UseCaseManager.GetArticleUseCase())
}

func (a *appServer) Run() {
	port := os.Getenv("PORT")
	a.initHandlers()
	err := a.routerEngine.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.NewConfig()
	return &appServer{
		routerEngine: r,
		cfg:          c,
	}
}
