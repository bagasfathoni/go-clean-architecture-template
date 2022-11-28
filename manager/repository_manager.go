package manager

import "github.com/bagasfathoni/go-clean-architecture-template/repository"

type repositoryManager struct {
	infra InfraManager
}

type RepositoryManager interface {
	ARepository() repository.ARepository
}

func (r *repositoryManager) ARepository() repository.ARepository {
	return repository.InitARepository(r.infra.SqlDb())
}

func InitRepositoryManager(i InfraManager) RepositoryManager {
	return &repositoryManager{infra: i}
}
