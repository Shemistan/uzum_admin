package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/Shemistan/uzum_admin/internal/service"
	"io"
	"log"
	"net/http"
)

type IHandlers interface {
	Healthz(w http.ResponseWriter, r *http.Request)
	GiveOrder(w http.ResponseWriter, r *http.Request)
	AddOrder(w http.ResponseWriter, r *http.Request)
}

func NewHandler(serv service.IService) IHandlers {
	return &handler{serv: serv}
}

type handler struct {
	serv service.IService
}

func (h *handler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *handler) AddOrder(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Body.Read", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	coordinate := &models.Coordinate{}
	order := &models.Order{
		Coordinate: coordinate,
	}

	err = json.Unmarshal(b, &order)
	if err != nil {
		log.Println("Unmarshal", err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(fmt.Printf("%+v\n", order)) // Записали в БД

	w.WriteHeader(http.StatusCreated)

}
