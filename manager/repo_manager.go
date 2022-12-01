package manager

import "github.com/bagasfathoni/go-clean-architecture-template/repository"

type RepositoryManager interface {
	VendorRepo() repository.VendorRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) VendorRepo() repository.VendorRepository {
	return repository.NewVendorRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	newRepoManager := new(repositoryManager)
	newRepoManager.infra = infra
	return newRepoManager
}
