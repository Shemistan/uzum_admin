package storage

import "github.com/Shemistan/uzum_admin/internal/models"

type IStorage interface {
	GetOrder(id int64) *models.Order
}
