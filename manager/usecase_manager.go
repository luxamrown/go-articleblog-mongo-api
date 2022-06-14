package manager

import "mohamadelabror.me/goapiblog/usecase"

type UseCaseManager interface {
	CreateArticleUseCase() usecase.CreateArticleUseCase
	GetAllArticlUseCase() usecase.GetAllArticleUseCase
	GetArticleUseCase() usecase.GetArticleUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) CreateArticleUseCase() usecase.CreateArticleUseCase {
	return usecase.NewCreateArticleUseCase(u.repo.ArticleRepo())
}

func (u *useCaseManager) GetAllArticlUseCase() usecase.GetAllArticleUseCase {
	return usecase.NewGetAllArticleUseCase(u.repo.ArticleRepo())

}
func (u *useCaseManager) GetArticleUseCase() usecase.GetArticleUseCase {
	return usecase.NewGetArticleUseCase(u.repo.ArticleRepo())
}

func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
