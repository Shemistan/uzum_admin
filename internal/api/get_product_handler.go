package api

import (
	"context"
	"github.com/Shemistan/uzum_admin/generated/protos/admin_v1"
	"github.com/google/uuid"
)

func (a *API) GetProduct(ctx context.Context, req *admin_v1.Product_GetRequest) (*admin_v1.Product_GetResponse, error) {
	product, err := a.service.GetProduct(ctx, uuid.MustParse(req.ProductId))
	if err != nil {
		return nil, err
	}
	return &admin_v1.Product_GetResponse{Product: &admin_v1.ProductInfo{
		ProductId:   product.ID.String(),
		Name:        product.Name,
		Price:       float32(product.Price),
		Description: product.Description,
		Quantity:    uint32(product.Quantity),
	}}, nil
}
