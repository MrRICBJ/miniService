package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"miniService/internal/config"
	v1 "miniService/internal/controllers/v1"
	"miniService/internal/db"
	"miniService/internal/repositories"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s\n", err.Error())
	}

	cfg := config.New()

	postgres, err := db.NewDatabase(cfg.Postgres)
	if err != nil {
		log.Fatalf("error connection db: %s\n", err.Error())
	}

	repo := repositories.New(postgres.Client)

	handlers := v1.New(repo)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: handlers.Register(),
	}

	go func() {
		fmt.Println("server started")
		_ = srv.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Server Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := postgres.Client.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}
