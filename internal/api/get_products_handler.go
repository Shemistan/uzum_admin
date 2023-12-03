package api

import (
	"context"
	"github.com/Shemistan/uzum_admin/generated/protos/admin_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *API) GetProducts(ctx context.Context, _ *emptypb.Empty) (*admin_v1.Products_GetResponse, error) {
	products, err := a.service.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	var response []*admin_v1.ShortProductInfo

	for _, product := range products {
		response = append(response, &admin_v1.ShortProductInfo{
			ProductId: product.ID.String(),
			Name:      product.Name,
			Price:     float32(product.Price),
		})
	}

	return &admin_v1.Products_GetResponse{Products: response}, nil
}
