package api

import (
	"context"
	"github.com/Shemistan/uzum_admin/generated/protos/admin_v1"
	"github.com/Shemistan/uzum_admin/internal/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *API) AddProduct(ctx context.Context, req *admin_v1.Product_CreateRequest) (*emptypb.Empty, error) {
	err := a.service.AddProduct(ctx, &models.Product{
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
