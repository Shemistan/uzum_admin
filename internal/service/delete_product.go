package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

func (s *AdminService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	user, err := s.GetUser(ctx)
	if err != nil {
		return err
	}
	if user.Role != "admin" {
		return errors.New("access denied")
	}
	return s.storage.DeleteProduct(ctx, id)
}
