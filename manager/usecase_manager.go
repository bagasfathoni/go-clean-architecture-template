package manager

import (
	"time"

	"github.com/bagasfathoni/go-clean-architecture-template/usecases"
)

type usecaseManager struct {
	repoManager RepositoryManager
}

type UsecaseManager interface {
	VendorUsecase() usecases.VendorUsecase
}

func (u *usecaseManager) VendorUsecase() usecases.VendorUsecase {
	return usecases.NewVendorUsecase(u.repoManager.VendorRepo(), 5*time.Second)
}

func NewUsecaseManager(repoManager RepositoryManager) UsecaseManager {
	newUsecaseManager := new(usecaseManager)
	newUsecaseManager.repoManager = repoManager
	return newUsecaseManager
}
