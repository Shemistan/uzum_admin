package models

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
}

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Quantity    int32     `json:"quantity"`
}

type Sales struct {
	ProductName string
	Quantity    int32
}

type Statistics struct {
	TotalIncome float64
	Sales       []*Sales
}

type Config struct {
	Host            string `envconfig:"HOST"`
	DB_CONN_STR     string `envconfig:"POSTGRES_DSN"`
	GrpcPort        string `envconfig:"GRPC_PORT"`
	HttpPort        string `envconfig:"HTTP_PORT"`
	LoginClientPort string `envconfig:"LOGIN_CLIENT_PORT"`
}
