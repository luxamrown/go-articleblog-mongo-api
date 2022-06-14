package usecase

import (
	"mohamadelabror.me/goapiblog/model"
	"mohamadelabror.me/goapiblog/repository"
)

type GetArticleUseCase interface {
	GetArticle(id string) (model.Article, error)
}

type getArticleUseCase struct {
	repo repository.ArticleRepo
}

func (g *getArticleUseCase) GetArticle(id string) (model.Article, error) {
	return g.repo.GetArticle(id)
}

func NewGetArticleUseCase(repo repository.ArticleRepo) GetArticleUseCase {
	return &getArticleUseCase{
		repo: repo,
	}
}
