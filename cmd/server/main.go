package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Neurasita/restful-api/pkg/config"
)

func main() {

	// Register handler routing here
	mux := http.NewServeMux()

	// Enable global middleware here
	var h http.Handler = mux

	// Create and run server
	s := http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.AppHost, config.AppPort),
		Handler: h,
	}
	go func(s *http.Server) {
		log.Printf("Server listening on %s\n", s.Addr)
		if err := s.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalln(err)
			}
			return
		}
	}(&s)

	// Handle gracefull shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdownRelease()
	if err := s.Shutdown(shutdownCtx); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
			return
		}
	}
	log.Println("Server shutdown successfully")
}
