package main

import (
	"github.com/Shemistan/uzum_admin/cmd/api/handlers"
	"github.com/Shemistan/uzum_admin/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	httpPort = ":8080"
)

func main() {
	deliveryService := service.NewService()
	handler := handlers.NewHandler(deliveryService)

	router := mux.NewRouter()

	router.HandleFunc("/healthz", handler.Healthz).Methods(http.MethodGet)
	router.HandleFunc("/deliver/v1/give-delivery/{id:[0-9]+}", handler.GiveOrder).Methods(http.MethodGet)
	router.HandleFunc("/deliver/v1/add-order/", handler.AddOrder).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:        httpPort,
		ReadTimeout: time.Second * 10,
		Handler:     router,
	}

	log.Println("delivery server is running at port:", httpPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
