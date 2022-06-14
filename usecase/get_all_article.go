package usecase

import (
	"mohamadelabror.me/goapiblog/model"
	"mohamadelabror.me/goapiblog/repository"
)

type GetAllArticleUseCase interface {
	GetAllArticle() ([]model.ArticleHead, error)
}

type getAllArticleUseCase struct {
	articleDb repository.ArticleRepo
}

func (g *getAllArticleUseCase) GetAllArticle() ([]model.ArticleHead, error) {
	return g.articleDb.GetAllArticle()
}

func NewGetAllArticleUseCase(articleDb repository.ArticleRepo) GetAllArticleUseCase {
	return &getAllArticleUseCase{
		articleDb: articleDb,
	}
}
