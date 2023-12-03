package service

import (
	"context"
	"github.com/Shemistan/uzum_admin/generated/protos/login_v1"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func (s *AdminService) GetUser(ctx context.Context) (models.User, error) {
	emp := &login_v1.GetData_Request{EndpointAddress: ""}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	res, err := s.loginClient.GetData(ctx, emp)
	if err != nil {
		return models.User{}, err
	}

	return models.User{ID: uuid.MustParse(res.UserId), Role: res.Role}, err
}
