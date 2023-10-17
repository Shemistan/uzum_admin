package main

import (
	"github.com/Shemistan/uzum_admin/cmd/api/serv"
	"github.com/Shemistan/uzum_admin/cmd/conf"
	"github.com/Shemistan/uzum_admin/internal/models"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	httpPort = ":8080"
)

func main() {
	cfg, err := conf.NewConfig()
	if err != nil {
		log.Fatal("failed to get congigs", err.Error())
	}

	db, err := initDB(cfg)
	if err != nil {
		log.Fatal("failed to init DB", err.Error())
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Println("failed to close connection to DV:", err.Error())
		}
	}()

	srv, err := serv.GetServ(cfg, db)
	if err != nil {
		log.Fatal("failed to get serv", err.Error())
	}

	log.Println("delivery server is running at port:", httpPort)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func initDB(cnf models.Config) (*sqlx.DB, error) {
	sqlConnectionString := conf.GetSqlConnectionString(cnf)

	db, err := sqlx.Open("postgres", sqlConnectionString)
	if err != nil {
		return nil, err
	}

	// Проверка доступности БД
	if err = db.Ping(); err != nil {
		log.Println("failed to ping DB")
		return nil, err
	}

	log.Println("Connection to DB success")

	return db, nil
}
