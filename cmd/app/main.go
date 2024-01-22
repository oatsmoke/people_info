package main

import (
	"context"
	"fmt"
	"github.com/gofor-little/env"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oatsmoke/people_info/internal/app/handler"
	"github.com/oatsmoke/people_info/internal/app/repository"
	"github.com/oatsmoke/people_info/internal/app/service"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	if err := env.Load(".env"); err != nil {
		log.Fatalf(".env initialization error: %s", err.Error())
		return
	}
	log.Println(".env initialized")
	connectionDB, err := newConnectionDB(env.Get("DB_STRING", ""))
	if err != nil {
		log.Fatalf("DB connection initialization error: %s", err.Error())
		return
	}
	defer connectionDB.Close()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	log.Println("DB connected")
	repositories := repository.NewRepository(connectionDB)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	server := &http.Server{
		Addr:    env.Get("PORT", ""),
		Handler: handlers.InitRoutes(),
	}
	log.Println("starting api server")
	go func() {
		<-ctx.Done()
		if err := server.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server startup error: %s", err.Error())
		return
	}
}

func newConnectionDB(connectionString string) (*pgxpool.Pool, error) {
	connectDB, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("main newConnectionDB: %s", err.Error())
	}
	err = connectDB.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("main newConnectionDB: %s", err.Error())
	}
	return connectDB, nil
}
