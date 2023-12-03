package app

import (
	"context"
	build "github.com/Shemistan/uzum_admin/dev"
	"github.com/Shemistan/uzum_admin/generated/protos/admin_v1"
	"github.com/Shemistan/uzum_admin/generated/protos/login_v1"
	"github.com/Shemistan/uzum_admin/internal/api"
	"github.com/Shemistan/uzum_admin/internal/db"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/Shemistan/uzum_admin/internal/service"
	"github.com/Shemistan/uzum_admin/internal/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (app *App) InitGRPC() {
	app.grpcServer = grpc.NewServer()

	admin_v1.RegisterShopServiceServer(app.grpcServer, &app.adminAPI)
}

func (app *App) InitHTTP(ctx context.Context) error {
	app.httpMux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := admin_v1.RegisterShopServiceHandlerFromEndpoint(ctx, app.httpMux, app.config.Host+app.config.GrpcPort, opts)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) initAPI() (err error) {
	app.dbConn, err = storage.NewDatabase(app.config.DB_CONN_STR)
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(app.config.Host+app.config.LoginClientPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	app.loginClient = login_v1.NewLoginV1Client(conn)

	shopService := service.NewAdminService(db.New(app.dbConn), app.loginClient)
	app.adminAPI = api.NewAPI(shopService)

	return err
}

func (app *App) loadConfig() error {
	var conf models.Config

	if build.DEBUG {
		err := build.SetConfig()
		if err != nil {
			return err
		}
	}
	err := envconfig.Process("", &conf)

	app.config = &conf
	return err
}
