package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mohamadelabror.me/goapiblog/delivery/appreq"
	"mohamadelabror.me/goapiblog/usecase"
)

type ArticleApi struct {
	getArticleUseCase usecase.GetArticleUseCase
}

func (a *ArticleApi) GetArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var articleId appreq.ArticleRequest
		err := c.ShouldBindJSON(&articleId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"ERROR": "CANNOT BIND JSON"})
		}

		selectedArticle, err := a.getArticleUseCase.GetArticle(articleId.ArticleId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"ERROR": err})
		}

		c.JSON(http.StatusOK, selectedArticle)

	}
}

func NewArticleApi(articleRoute *gin.RouterGroup, getArticleUseCase usecase.GetArticleUseCase) {
	api := ArticleApi{
		getArticleUseCase: getArticleUseCase,
	}
	articleRoute.GET("/", api.GetArticle())
}
