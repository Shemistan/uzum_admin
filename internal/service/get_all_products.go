package service

import (
	"context"
	"github.com/Shemistan/uzum_admin/internal/models"
)

func (s *AdminService) GetAllProducts(ctx context.Context) ([]*models.Product, error) {
	products, err := s.storage.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	var res []*models.Product
	for _, product := range products {
		res = append(res, &models.Product{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price.Float64,
			Description: product.Description.String,
			Quantity:    product.Quantity,
		})
	}
	return res, nil
}
