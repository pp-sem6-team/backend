package mapper

import (
	"github.com/pp-sem6-team/backend/internal/db/model"
	"github.com/pp-sem6-team/backend/internal/domain"
	"github.com/pp-sem6-team/backend/internal/dto"
)

func ToDomain(m *model.User) *domain.User {
	return &domain.User{
		ID:        m.ID,
		Email:     m.Email,
		Name:      m.Name,
		BirthDate: m.BirthDate,
		Gender:    (*domain.Gender)(m.Gender),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToResponse(u *domain.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		BirthDate: u.BirthDate,
		Gender:    u.Gender,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
