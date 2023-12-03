package api

import (
	"context"
	"github.com/Shemistan/uzum_admin/generated/protos/admin_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *API) GetStatistics(ctx context.Context, _ *emptypb.Empty) (*admin_v1.Statistics_GetResponse, error) {
	statistics, err := a.service.GetStatistics(ctx)
	if err != nil {
		return nil, err
	}

	var response []*admin_v1.Sales

	for _, item := range statistics.Sales {
		response = append(response, &admin_v1.Sales{
			ProductName:  item.ProductName,
			ProductSales: uint32(item.Quantity),
		})
	}

	return &admin_v1.Statistics_GetResponse{TotalIncome: float32(statistics.TotalIncome), Sales: response}, nil
}
