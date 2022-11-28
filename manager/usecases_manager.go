package manager

import "github.com/bagasfathoni/go-clean-architecture-template/usecases"

type usecasesManager struct {
	repoManager RepositoryManager
}

type UsecasesManager interface {
	AUsecases() usecases.AUsecases
}

func (u *usecasesManager) AUsecases() usecases.AUsecases {
	return usecases.InitAUsecases(u.repoManager.ARepository())
}

func InitUsecasesManager(r RepositoryManager) UsecasesManager {
	return &usecasesManager{repoManager: r}
}
