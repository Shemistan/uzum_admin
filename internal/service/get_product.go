package service

import (
	"context"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/google/uuid"
)

func (s *AdminService) GetProduct(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	product, err := s.storage.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.Product{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price.Float64,
		Description: product.Description.String,
		Quantity:    product.Quantity,
	}, nil
}
