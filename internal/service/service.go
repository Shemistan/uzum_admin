package service

import (
	"context"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/Shemistan/uzum_admin/internal/storage"
)

type IService interface {
	AddProduct(ctx context.Context, req *models.Product) (int64, error)
	GetAllProducts(ctx context.Context) ([]*models.Product, error)
}

func NewService(storage storage.IStorage) IService {
	return &service{store: storage}
}

type service struct {
	store storage.IStorage
}
