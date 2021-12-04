package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	instance = &http.Server{

	}
)

// GracefulRun Stop server in specified timeout (in seconds)
func GracefulRun(srv http.Server, timeout int) {
	// init server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server cannot be started: %s", err)
		}
	}()
	// wait timeout
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutting down server")

	// inform server to finish current job in specified time
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("Server force to shutdown", err)

	}
	log.Println("Server exiting")
}
