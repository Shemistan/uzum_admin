package service

import (
	"context"
	"errors"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/google/uuid"
)

func (s *AdminService) GetStatistics(ctx context.Context) (*models.Statistics, error) {
	user, err := s.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	if user.Role != "admin" {
		return nil, errors.New("access denied")
	}
	sales, err := s.storage.GetBaskets(ctx)
	if err != nil {
		return nil, err
	}

	var response models.Statistics
	stats := make(map[uuid.UUID]int32)

	for _, sale := range sales {
		stats[sale.ProductID] += sale.Quantity
	}

	for id, quantity := range stats {
		product, err := s.storage.GetProduct(ctx, id)
		if err != nil {
			continue
		}
		response.Sales = append(response.Sales, &models.Sales{
			ProductName: product.Name,
			Quantity:    quantity,
		})
		response.TotalIncome += product.Price.Float64 * float64(quantity)
	}

	return &response, nil
}
