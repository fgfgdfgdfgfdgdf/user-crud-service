package main

import (
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	httphandlers "github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/api/http/handlers"
	httproutes "github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/api/http/routers"
	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	api := humachi.New(router, huma.DefaultConfig("User CRUD service", "1.0.0"))

	userHandler := httphandlers.NewUserHandler(nil)

	httproutes.AddUserRoutes(api, userHandler)

	spec, err := api.OpenAPI().YAML()
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("openapi.yaml", spec, 0644); err != nil {
		panic(err)
	}
}
