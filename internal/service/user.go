package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/entity"
	db "github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/infra/db/gen"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	q *db.Queries
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	return &UserService{
		q: db.New(pool),
	}
}

func (s *UserService) Create(ctx context.Context, name, password, email string) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}

	user, err := s.q.CreateUser(ctx, db.CreateUserParams{
		Name:     name,
		Active:   true,
		Password: string(hash),
		Email:    email,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func (s *UserService) Get(ctx context.Context, userID uuid.UUID) (db.User, error) {
	user, err := s.q.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return db.User{}, entity.ErrUserNotFound.WithCause(err)
		}
		return db.User{}, err
	}

	return user, nil
}

func (s *UserService) Update(ctx context.Context, userID uuid.UUID, name, email, password string, active bool) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.q.UpdateUserByID(ctx, db.UpdateUserByIDParams{
		Name:     name,
		Email:    email,
		Password: string(hash),
		Active:   active,
		ID:       userID,
	})
}

func (s *UserService) Delete(ctx context.Context, userID uuid.UUID, hard bool) error {
	if hard {
		return s.q.HardDeleteUser(ctx, userID)
	}

	return s.q.DeleteUser(ctx, userID)
}

func (s *UserService) RenameUser(ctx context.Context, userID uuid.UUID, newName string) error {
	return s.q.UpdateUserNameByID(ctx, db.UpdateUserNameByIDParams{
		Name: newName,
		ID:   userID,
	})
}

func (s *UserService) ChangePassword(ctx context.Context, userID uuid.UUID, newPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.q.UpdateUserPasswordByID(ctx, db.UpdateUserPasswordByIDParams{
		Password: string(hash),
		ID:       userID,
	})
}
