package app

import (
	"context"
	"github.com/Shemistan/uzum_admin/generated/protos/login_v1"
	"github.com/Shemistan/uzum_admin/internal/api"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type App struct {
	dbConn *sqlx.DB

	adminAPI     api.API
	loginClient login_v1.LoginV1Client
	grpcServer  *grpc.Server
	httpMux     *runtime.ServeMux

	config *models.Config
}

func NewApp(ctx context.Context) (*App, error) {
	var app App
	err := app.loadConfig()
	if err != nil {
		return nil, err
	}
	err = app.initAPI()
	if err != nil {
		return nil, err
	}
	app.InitGRPC()
	err = app.InitHTTP(ctx)
	if err != nil {
		return nil, err
	}

	return &app, err
}
