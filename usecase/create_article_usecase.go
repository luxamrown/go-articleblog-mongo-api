package usecase

import (
	"mohamadelabror.me/goapiblog/repository"
)

type CreateArticleUseCase interface {
	CreateArticle(title, desc, author, date, article_id string) error
}

type createArticleUseCase struct {
	articleRepo repository.ArticleRepo
}

func (c *createArticleUseCase) CreateArticle(title, desc, author, date, article_id string) error {
	return c.articleRepo.CreateArticle(title, desc, author, date, article_id)
}

func NewCreateArticleUseCase(articleRepo repository.ArticleRepo) CreateArticleUseCase {
	return &createArticleUseCase{
		articleRepo: articleRepo,
	}
}
