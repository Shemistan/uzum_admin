package api

import (
	"github.com/Shemistan/uzum_admin/generated/protos/admin_v1"
	"github.com/Shemistan/uzum_admin/internal/service"
)

type API struct {
	admin_v1.UnimplementedShopServiceServer
	service *service.AdminService
}

func NewAPI(service *service.AdminService) API {
	return API{
		service: service,
	}
}
