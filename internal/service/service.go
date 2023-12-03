package service

import (
	"context"
	"github.com/Shemistan/uzum_admin/generated/protos/login_v1"
	"github.com/Shemistan/uzum_admin/internal/db"
	"github.com/google/uuid"
)

type (
	IStorage interface {
		CreateProduct(ctx context.Context, arg db.CreateProductParams) error
		UpdateProduct(ctx context.Context, arg db.UpdateProductParams) error
		DeleteProduct(ctx context.Context, id uuid.UUID) error
		GetProduct(ctx context.Context, id uuid.UUID) (db.Product, error)
		GetProducts(ctx context.Context) ([]db.Product, error)
		GetBaskets(ctx context.Context) ([]db.Basket, error)
	}
	AdminService struct {
		storage     IStorage
		loginClient login_v1.LoginV1Client
	}
)

func NewAdminService(storage IStorage, loginClient login_v1.LoginV1Client) *AdminService {
	return &AdminService{
		storage:     storage,
		loginClient: loginClient,
	}
}
