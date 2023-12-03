package api

import (
	"context"
	"github.com/Shemistan/uzum_admin/generated/protos/admin_v1"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *API) UpdateProduct(ctx context.Context, req *admin_v1.Product_UpdateRequest) (*emptypb.Empty, error) {
	err := a.service.UpdateProduct(ctx, &models.Product{
		ID:          uuid.MustParse(req.Product.ProductId),
		Name:        req.Product.Name,
		Price:       float64(req.Product.Price),
		Description: req.Product.Description,
		Quantity:    int32(req.Product.Quantity),
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
