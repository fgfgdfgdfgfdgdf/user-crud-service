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

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	httphandlers "github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/api/http/handlers"
	httproutes "github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/api/http/routers"
	"github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/config"
	"github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/infra/db"
	"github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {

	if err := config.Init(); err != nil {
		panic("Failed to initialize config")
	}

	router := chi.NewRouter()

	api := humachi.New(router, huma.DefaultConfig("User CRUD service", "1.0.0"))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pgPool, err := db.NewPGPool(ctx)
	if err != nil {
		panic(err)
	}

	userService := service.NewUserService(pgPool)

	userHandler := httphandlers.NewUserHandler(userService)

	httproutes.AddUserRoutes(api, userHandler)

	addr := fmt.Sprintf("%s:%d", config.API.API_LISTEN_ADDR, config.API.API_PORT)
	srv := &http.Server{
		Addr:    addr,
		Handler: api.Adapter(),
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listen error: %v", err)
		}
	}()

	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	pgPool.Close()

}
