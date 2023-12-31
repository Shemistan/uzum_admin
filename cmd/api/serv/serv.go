package serv

import (
	"github.com/Shemistan/uzum_admin/cmd/api/handlers"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/Shemistan/uzum_admin/internal/service"
	"github.com/Shemistan/uzum_admin/internal/storage/postgres"
	desc "github.com/Shemistan/uzum_admin/pkg/auth_v1"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

func GetServ(cnf models.Config, db *sqlx.DB, c desc.AuthV1Client) (*http.Server, error) {

	store := postgres.NewRepoPostgres(db)
	deliveryService := service.NewService(store, c)
	handler := handlers.NewHandler(deliveryService)

	router := mux.NewRouter()

	router.HandleFunc("/healthz", handler.Healthz).Methods(http.MethodGet)
	router.HandleFunc("/product/add", handler.AddProduct).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:        cnf.App.Port,
		ReadTimeout: time.Second * 10,
		Handler:     router,
	}

	return srv, nil
}
