package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Shemistan/uzum_admin/internal/models"
	"github.com/Shemistan/uzum_admin/internal/service"
)

type IHandlers interface {
	Healthz(w http.ResponseWriter, r *http.Request)

	AddProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

func NewHandler(serv service.IService) IHandlers {
	return &handler{serv: serv}
}

type handler struct {
	serv service.IService
}

func (h *handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := getContext()
	defer cancel()

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Body.Read", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := &models.Product{}

	err = json.Unmarshal(b, &product)
	if err != nil {
		log.Println("Unmarshal", err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.serv.AddProduct(ctx, product)
	if err != nil {
		log.Println("failed to add product", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(strconv.FormatInt(res, 10)))
	if err != nil {
		log.Println("failed to write body", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx.Value("test")

	w.WriteHeader(http.StatusOK)
}
func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx.Value("test")

	w.WriteHeader(http.StatusOK)
}
func (h *handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx.Value("test")

	w.WriteHeader(http.StatusOK)
}
func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx.Value("test")

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Healthz(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx.Value("test")

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
