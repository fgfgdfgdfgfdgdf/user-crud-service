package httphandlers

import (
	"context"
	"time"

	"github.com/google/uuid"

	db "github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/infra/db/gen"
)

type UserService interface {
	Delete(ctx context.Context, userID uuid.UUID, hard bool) error
	Get(ctx context.Context, userID uuid.UUID) (db.User, error)
	Create(ctx context.Context, name string, password string, email string) (uuid.UUID, error)
	Update(ctx context.Context, userID uuid.UUID, name, email, password string, active bool) error
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type UserCreateInput struct {
	Body struct {
		Name     string `json:"name"`
		Password string `json:"password" minLength:"8" maxLength:"100"`
		Email    string `json:"email" format:"email"`
	}
}

type UserCreateOutput struct {
	Body struct {
		ID uuid.UUID `json:"id"`
	}
}

func (h *UserHandler) Create(ctx context.Context, in *UserCreateInput) (*UserCreateOutput, error) {
	id, err := h.userService.Create(ctx, in.Body.Name, in.Body.Password, in.Body.Email)
	if err != nil {
		return nil, err
	}

	out := &UserCreateOutput{}
	out.Body.ID = id

	return out, nil
}

type UserGetInput struct {
	ID uuid.UUID `path:"id" format:"uuid"`
}

type UserGetOutput struct {
	Body struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email" format:"email"`
		Active    bool      `json:"active"`
		CreatedAt time.Time `json:"created_at"`
	}
}

func (h *UserHandler) Get(ctx context.Context, in *UserGetInput) (*UserGetOutput, error) {
	user, err := h.userService.Get(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	out := &UserGetOutput{}
	out.Body.ID = user.ID
	out.Body.Name = user.Name
	out.Body.Email = user.Email
	out.Body.Active = user.Active
	out.Body.CreatedAt = user.CreatedAt

	return out, nil
}

type UserUpdateInput struct {
	ID   uuid.UUID `path:"id"`
	Body struct {
		Name     string `json:"name"`
		Password string `json:"password" minLength:"8" maxLength:"100"`
		Email    string `json:"email" format:"email"`
		Active   bool   `json:"active"`
	}
}

func (h *UserHandler) Update(ctx context.Context, in *UserUpdateInput) (*struct{}, error) {
	if err := h.userService.Update(ctx, in.ID, in.Body.Name, in.Body.Email, in.Body.Password, in.Body.Active); err != nil {
		return nil, err
	}

	return nil, nil
}

type UserDeleteInput struct {
	ID   uuid.UUID `path:"id"`
	Body struct {
		Hard bool `json:"hard"`
	}
}

type UserDeleteOutput struct {
	Body struct {
		ID uuid.UUID `json:"id"`
	}
}

func (h *UserHandler) Delete(ctx context.Context, in *UserDeleteInput) (*UserDeleteOutput, error) {
	if err := h.userService.Delete(ctx, in.ID, in.Body.Hard); err != nil {
		return nil, err
	}

	out := &UserDeleteOutput{}
	out.Body.ID = in.ID

	return out, nil
}
