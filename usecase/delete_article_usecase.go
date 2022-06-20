package usecase

import "mohamadelabror.me/goapiblog/repository"

type DeleteArticleUseCase interface {
	DeleteArticle(id string) error
}

type deleteArticleUseCase struct {
	articleDb repository.ArticleRepo
}

func (d *deleteArticleUseCase) DeleteArticle(id string) error {
	return d.articleDb.DeleteArticle(id)
}

func NewDeleteArticleUseCase(articleDb repository.ArticleRepo) DeleteArticleUseCase {
	return &deleteArticleUseCase{
		articleDb: articleDb,
	}
}
