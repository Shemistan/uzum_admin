package service

import (
	"context"
	"errors"
	"github.com/Shemistan/uzum_admin/generated/protos/login_v1"
	"github.com/google/uuid"
)

func (s *AdminService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	user, err := s.loginClient.GetData(ctx, &login_v1.GetData_Request{EndpointAddress: ""})
	if err != nil {
		return err
	}
	if user.Role != "admin" {
		return errors.New("access denied")
	}
	return s.storage.DeleteProduct(ctx, id)
}
