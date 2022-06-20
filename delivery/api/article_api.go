package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mohamadelabror.me/goapiblog/delivery/appreq"
	"mohamadelabror.me/goapiblog/usecase"
)

type ArticleApi struct {
	createArticleUseCase usecase.CreateArticleUseCase
	getAllArticleUseCase usecase.GetAllArticleUseCase
	getArticleUseCase    usecase.GetArticleUseCase
	deleteArticleUseCase usecase.DeleteArticleUseCase
}

func (a *ArticleApi) CreateArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newArticle appreq.NewArticleRequest
		date := time.Now().Format("01-02-2006")
		err := c.ShouldBindJSON(&newArticle)
		articleId := uuid.New().String()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"ERROR": "CANNOT BIND JSON"})
			return
		}

		err = a.createArticleUseCase.CreateArticle(newArticle.Title, newArticle.Desc, newArticle.Author, date, articleId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"ERROR": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"STATUS": "SUCCESS"})
	}
}

func (a *ArticleApi) GetAllArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var article []model.ArticleHead

		articles, err := a.getAllArticleUseCase.GetAllArticle()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Success", "Data": articles})

	}
}

func (a *ArticleApi) GetArticle() gin.HandlerFunc {
	return func(c *gin.Context) {

		articleId := c.Param("id")
		selectedArticle, err := a.getArticleUseCase.GetArticle(articleId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			return
		}

		c.JSON(http.StatusOK, selectedArticle)

	}
}

func (a *ArticleApi) DeleteArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		articleId := c.Param("id")
		err := a.deleteArticleUseCase.DeleteArticle(articleId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error: ": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Success Deleted :": articleId})
	}
}

func NewArticleApi(articleRoute *gin.RouterGroup, createArticleUseCase usecase.CreateArticleUseCase, getAllArticleUseCase usecase.GetAllArticleUseCase, getArticleUseCase usecase.GetArticleUseCase, deleteArticleUseCase usecase.DeleteArticleUseCase) {
	api := ArticleApi{
		createArticleUseCase: createArticleUseCase,
		getAllArticleUseCase: getAllArticleUseCase,
		getArticleUseCase:    getArticleUseCase,
		deleteArticleUseCase: deleteArticleUseCase,
	}
	articleRoute.GET("/:id", api.GetArticle())
	articleRoute.DELETE("/:id", api.DeleteArticle())
	articleRoute.POST("/post", api.CreateArticle())
	articleRoute.GET("/all", api.GetAllArticle())
}
