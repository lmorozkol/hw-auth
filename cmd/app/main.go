package main

import (
	"context"
	"fmt"
	"log"
	"ms-hw/cmd/init/db"
	"ms-hw/internal/adapters/db/postgres"
	"ms-hw/internal/config"
	"ms-hw/internal/service/authentication"
	"ms-hw/internal/transport/rest"
	"ms-hw/internal/transport/rest/handler"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}

	cfg := config.SetConfig()
	srv := rest.NewServer(&cfg)

	// gc
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		osCall := <-c
		fmt.Printf("system call :%+v", osCall)
		srv.Shutdown()
		wg.Done()
	}(&wg)

	// connections
	conn, err := db.Conn(ctx, &cfg)
	if err != nil {
		log.Println(err)
	}

	// storages
	storage := postgres.NewPostgresRepo(conn)

	// services
	authService := authentication.NewAuthService(storage)

	// handlers
	authHandler := handler.NewAuthHandler(authService)

	routes := rest.Linker(authHandler)

	srv.Serve(routes)

	wg.Wait()
}
