package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/model"
	"github.com/pp-sem6-team/backend/internal/domain"
	"github.com/pp-sem6-team/backend/internal/repository"
	"github.com/pp-sem6-team/backend/internal/security"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// TODO: заменить model.User на DTO
func (s *UserService) GetMe(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return s.repo.GetByID(ctx, userID)
}

func (s *UserService) UpdateMe(ctx context.Context, userID uuid.UUID, input domain.UpdateUserInput) error {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	if input.BirthDate != nil {
		user.BirthDate = input.BirthDate
	}

	if input.Gender != nil {
		user.Gender = input.Gender
	}

	return s.repo.Update(ctx, user)
}

func (s *UserService) UpdatePassword(ctx context.Context, userID uuid.UUID, password string) error {
	hash, err := security.HashPassword(password)
	if err != nil {
		return err
	}

	return s.repo.UpdatePassword(ctx, userID, string(hash))
}

func (s *UserService) DeleteMe(ctx context.Context, userID uuid.UUID) error {
	return s.repo.Delete(ctx, userID)
}
