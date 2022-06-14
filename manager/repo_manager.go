package manager

import "mohamadelabror.me/goapiblog/repository"

type RepoManager interface {
	ArticleRepo() repository.ArticleRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) ArticleRepo() repository.ArticleRepo {
	return repository.NewArticleRepo(r.infra.MongoDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
