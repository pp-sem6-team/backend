package service

import (
	"context"

	"github.com/google/uuid"
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

func (s *UserService) GetByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		BirthDate: user.BirthDate,
		Gender:    user.Gender,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *UserService) Update(ctx context.Context, userID uuid.UUID, input domain.UpdateUserInput) error {
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

func (s *UserService) Delete(ctx context.Context, userID uuid.UUID) error {
	return s.repo.Delete(ctx, userID)
}
