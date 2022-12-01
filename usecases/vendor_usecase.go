package usecases

import (
	"context"
	"time"

	"github.com/bagasfathoni/go-clean-architecture-template/model"
	"github.com/bagasfathoni/go-clean-architecture-template/repository"
)

type vendorUsecase struct {
	repo    repository.VendorRepository
	timeout time.Duration
}
type VendorUsecase interface {
	GetVendorByName(c context.Context, name string) (*model.Vendor, error)
	GetAllVendor(c context.Context) (*[]model.Vendor, error)
}

func (v *vendorUsecase) GetVendorByName(c context.Context, name string) (*model.Vendor, error) {
	ctx, cancel := context.WithTimeout(c, v.timeout)
	defer cancel()

	res, err := v.repo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (v *vendorUsecase) GetAllVendor(c context.Context) (*[]model.Vendor, error) {
	ctx, cancel := context.WithTimeout(c, v.timeout)
	defer cancel()

	res, err := v.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewVendorUsecase(repo repository.VendorRepository, timeout time.Duration) VendorUsecase {
	newUsecase := new(vendorUsecase)
	newUsecase.repo = repo
	newUsecase.timeout = timeout
	return newUsecase
}
