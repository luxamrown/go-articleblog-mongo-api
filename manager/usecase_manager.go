package manager

import "mohamadelabror.me/goapiblog/usecase"

type UseCaseManager interface {
	GetArticleUseCase() usecase.GetArticleUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) GetArticleUseCase() usecase.GetArticleUseCase {
	return usecase.NewGetArticleUseCase(u.repo.ArticleRepo())
}

func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
