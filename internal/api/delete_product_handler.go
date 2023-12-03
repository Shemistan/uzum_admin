package api

import (
	"context"
	"github.com/Shemistan/uzum_admin/generated/protos/admin_v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *API) DeleteProduct(ctx context.Context, req *admin_v1.Product_DeleteRequest) (*emptypb.Empty, error) {
	err := a.service.DeleteProduct(ctx, uuid.MustParse(req.ProductId))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
