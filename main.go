package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shortly/handler"
	"shortly/repository/inmemory"
	"shortly/service"
	"syscall"
)

func main() {
	addr := flag.String("address", ":8080", "address for running server")
	flag.Parse()

	repo := inmemory.NewInMemoryDB()
	urlService := service.NewURLService(repo)
	urlHandler := handler.NewHandler(urlService)

	router := http.NewServeMux()
	router.HandleFunc("/encode/", urlHandler.EncodeURL)
	router.HandleFunc("/metrics", urlHandler.Metrics)
	router.HandleFunc("/", urlHandler.Redirect)

	s := &http.Server{
		Addr:    *addr,
		Handler: router,
	}

	logger := log.Default()

	go func() {
		logger.Printf("starting server at :8080")
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatalf("server closed with error %v", err)
		} else {
			logger.Println("server closed!")
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	logger.Printf("caught sig: %+v", sig)
	logger.Printf("graceful shutdown for signal %v,", sig)

	if err := s.Shutdown(context.Background()); err != nil {
		logger.Printf("error while shutting down server: %v", err)
	}
}
