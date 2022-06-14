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
	getArticleUseCase    usecase.GetArticleUseCase
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

func (a *ArticleApi) GetArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var articleId appreq.ArticleRequest
		err := c.ShouldBindJSON(&articleId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"ERROR": "CANNOT BIND JSON"})
			return
		}

		selectedArticle, err := a.getArticleUseCase.GetArticle(articleId.ArticleId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"ERROR": err})
			return
		}

		c.JSON(http.StatusOK, selectedArticle)

	}
}

func NewArticleApi(articleRoute *gin.RouterGroup, getArticleUseCase usecase.GetArticleUseCase, createArticleUseCase usecase.CreateArticleUseCase) {
	api := ArticleApi{
		getArticleUseCase:    getArticleUseCase,
		createArticleUseCase: createArticleUseCase,
	}
	articleRoute.GET("/", api.GetArticle())
	articleRoute.POST("/post", api.CreateArticle())
}
