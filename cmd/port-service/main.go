package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dkhrunov/hexagonal-architecture/internal/config"
	"github.com/dkhrunov/hexagonal-architecture/internal/repository/inmem"
	"github.com/dkhrunov/hexagonal-architecture/internal/services"
	"github.com/dkhrunov/hexagonal-architecture/transport"
	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	//  read config from env
	cfg := config.Read()

	// create port repository
	portStore := inmem.NewPortInmemStore()

	// create port service
	portService := services.NewPortService(portStore)

	// create http server with application injecterd
	httpServer := transport.NewHttpServer(portService)

	// create http router
	router := mux.NewRouter()
	router.HandleFunc("/port", httpServer.GetPort).Methods(http.MethodGet)
	router.HandleFunc("/count", httpServer.CountPorts).Methods(http.MethodGet)
	router.HandleFunc("/port", httpServer.UploadPorts).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    cfg.PortService.HTTPAddr,
		Handler: router,
	}

	// listen to OS signals and gracefully shutfown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP Server on %v", cfg.PortService.HTTPAddr)

	// start HTTP server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP Server ListenAndServe Error: %v", err)
	}

	<-stopped

	log.Printf("Have a nice day!")

	return nil
}
