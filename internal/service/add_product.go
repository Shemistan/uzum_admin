package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"

	"github.com/Shemistan/uzum_admin/internal/models"
)

func (s *service) AddProduct(ctx context.Context, req *models.Product) (int64, error) {
	req.Price = req.Price * 100

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := s.authClient.Healthz(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("failed to get note by id: %v", err)
	}

	return s.store.AddProduct(ctx, req)
}
