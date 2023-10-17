package service

import (
	"context"

	"github.com/Shemistan/uzum_admin/internal/models"
)

func (s *service) AddProduct(ctx context.Context, req *models.Product) (int64, error) {
	req.Price = req.Price * 100
	return s.store.AddProduct(ctx, req)
}
