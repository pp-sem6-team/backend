package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/domain"
	"github.com/pp-sem6-team/backend/internal/repository"
	"github.com/pp-sem6-team/backend/internal/security"
	"github.com/pp-sem6-team/backend/internal/service/mapper"
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

	return mapper.ToUserDomain(user), nil
}

func (s *UserService) Update(
	ctx context.Context,
	userID uuid.UUID,
	email *string,
	name *string,
	birthDate *time.Time,
	gender *domain.Gender,
) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if email != nil {
		user.Email = *email
	}

	if name != nil {
		user.Name = *name
	}

	if birthDate != nil {
		user.BirthDate = birthDate
	}

	if gender != nil {
		user.Gender = gender
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return mapper.ToUserDomain(user), nil
}

func (s *UserService) UpdatePassword(ctx context.Context, userID uuid.UUID, currentPassword, newPassword string) error {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if err := security.CheckPassword(currentPassword, user.PasswordHash); err != nil {
		return err
	}

	hash, err := security.HashPassword(newPassword)
	if err != nil {
		return err
	}

	return s.repo.UpdatePassword(ctx, userID, string(hash))
}

func (s *UserService) Delete(ctx context.Context, userID uuid.UUID) error {
	return s.repo.Delete(ctx, userID)
}
