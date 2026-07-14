package httproutes

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	httphandlers "github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/api/http/handlers"
)

type UserHandler interface {
	Create(ctx context.Context, in *httphandlers.UserCreateInput) (*httphandlers.UserCreateOutput, error)
	Get(ctx context.Context, in *httphandlers.UserGetInput) (*httphandlers.UserGetOutput, error)
	Update(ctx context.Context, in *httphandlers.UserUpdateInput) (*struct{}, error)
	Delete(ctx context.Context, in *httphandlers.UserDeleteInput) (*httphandlers.UserDeleteOutput, error)
}

func AddUserRoutes(api huma.API, handler UserHandler) {
	huma.Register(api, huma.Operation{Method: http.MethodPost, Path: "/users"}, handler.Create)
	huma.Register(api, huma.Operation{Method: http.MethodGet, Path: "/users/{id}"}, handler.Get)
	huma.Register(api, huma.Operation{Method: http.MethodPut, Path: "/users/{id}"}, handler.Update)
	huma.Register(api, huma.Operation{Method: http.MethodDelete, Path: "/users/{id}"}, handler.Delete)
}
