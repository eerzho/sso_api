package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/handler"
	"sso/internal/srvc"
	"syscall"
	"time"
)

func main() {
	app := app.MustNew()
	server := setup(app)

	errChan := make(chan error, 1)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go start(server, app, errChan)
	wait(errChan, stopChan)
	shutdown(server)
}

func setup(app *app.App) *http.Server {
	userSrvc := srvc.NewUser()
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Cfg.HTTP.Port),
		Handler: handler.New(app,  userSrvc),
	}
}
 
func start(server *http.Server, app *app.App, errChan chan<- error) {
	log.Printf("starting app on port: %s\n", app.Cfg.HTTP.Port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		errChan <- err
	}
}

func wait(errChan <-chan error, stopChan <-chan os.Signal) {
	select {
	case <-stopChan:
		log.Println("stopping app...")
	case err := <-errChan:
		log.Fatalf("server: %v", err)
	}
}

func shutdown(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("shutdown: %v", err)
	}

	log.Println("server stopped gracefully")
}