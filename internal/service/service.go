package service

import (
	"context"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/Shemistan/uzum_admin/internal/storage"
	desc "github.com/Shemistan/uzum_admin/pkg/auth_v1"
)

type IService interface {
	AddProduct(ctx context.Context, req *models.Product) (int64, error)
	GetAllProducts(ctx context.Context) ([]*models.Product, error)
}

func NewService(storage storage.IStorage, authClient desc.AuthV1Client) IService {
	return &service{store: storage,
		authClient: authClient,
	}
}

type service struct {
	store      storage.IStorage
	authClient desc.AuthV1Client
}
