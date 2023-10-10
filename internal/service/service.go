package service

import (
	"github.com/Shemistan/uzum_admin/internal/models"
	"time"
)

type IService interface {
	GetOrder(id int64) (*models.Order, error)
}

func NewService() IService {
	return &service{}
}

type service struct {
	//Store storage.IStorage
}

func (s *service) GetOrder(id int64) (*models.Order, error) {
	return &models.Order{
		ID:      id,
		Name:    "test Name",
		Phone:   "Phone",
		Address: "Address",
		Coordinate: &models.Coordinate{
			Longitude: 54.234,
			Latitude:  37.454,
		},
		Meta:         "Address",
		Status:       "Address",
		DeliveryTime: time.Time{},
	}, nil
}
