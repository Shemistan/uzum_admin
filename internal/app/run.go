package app

import (
	"log"
	"net"
	"net/http"
	"runtime"
	"sync"
)

func (app *App) RunGRPC() error {
	lis, err := net.Listen("tcp", app.config.GrpcPort)
	if err != nil {
		return err
	}

	log.Printf("Starting gRPC server on port %s", app.config.GrpcPort)
	return app.grpcServer.Serve(lis)
}

func (app *App) RunHTTP() error {
	log.Printf("Starting HTTP server on port %s", app.config.HttpPort)
	return http.ListenAndServe(app.config.HttpPort, app.httpMux)
}

func (app *App) Run() error {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := app.RunGRPC()
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		defer wg.Done()
		err := app.RunHTTP()
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()

	return nil
}
