package service

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/db/model"
	"github.com/pp-sem6-team/backend/internal/domain"
	"github.com/pp-sem6-team/backend/internal/integration/ml"
	"github.com/pp-sem6-team/backend/internal/integration/storage"
	"github.com/pp-sem6-team/backend/internal/repository"
	"gorm.io/datatypes"
)

type AnalysisService struct {
	photoRepo              repository.PhotoRepository
	analysisRepo           repository.AnalysisRepository
	recommendationRepo     repository.RecommendationRepository
	skinTypeIngredientRepo repository.SkinTypeIngredientRepository
	ingredientRepo         repository.IngredientRepository
	mlClient               ml.Client
	storage                storage.Storage
	presignTTL             time.Duration
}

func NewAnalysisService(
	photoRepo repository.PhotoRepository,
	analysisRepo repository.AnalysisRepository,
	recommendationRepo repository.RecommendationRepository,
	skinTypeIngredientRepo repository.SkinTypeIngredientRepository,
	ingredientRepo repository.IngredientRepository,
	mlClient ml.Client,
	storage storage.Storage,
	presignTTL time.Duration,
) *AnalysisService {
	return &AnalysisService{
		photoRepo:              photoRepo,
		analysisRepo:           analysisRepo,
		recommendationRepo:     recommendationRepo,
		skinTypeIngredientRepo: skinTypeIngredientRepo,
		ingredientRepo:         ingredientRepo,
		mlClient:               mlClient,
		storage:                storage,
		presignTTL:             presignTTL,
	}
}

func (s *AnalysisService) Create(
	ctx context.Context,
	userID uuid.UUID,
	content io.Reader,
	size int64,
	fileName string,
) (*domain.AnalysisCreated, error) {

	ext := strings.ToLower(filepath.Ext(fileName))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return nil, domain.ErrInvalidFileType
	}

	objectKey := fmt.Sprintf("users/%s/photos/%s%s", userID, uuid.New().String(), ext)

	err := s.storage.Upload(
		ctx,
		objectKey,
		content,
		size,
		"image/jpeg",
	)

	if err != nil {
		return nil, err
	}

	photo := &model.Photo{
		ID:        uuid.New(),
		UserID:    userID,
		ObjectKey: objectKey,
	}

	if err := s.photoRepo.Create(ctx, photo); err != nil {
		return nil, err
	}

	analysis := &model.Analysis{
		ID:           uuid.New(),
		PhotoID:      photo.ID,
		Status:       domain.Processing,
		AnalysisData: datatypes.JSON([]byte(`{}`)),
	}

	if err := s.analysisRepo.Create(ctx, analysis); err != nil {
		return nil, err
	}

	go func(parentCtx context.Context, analysisID, photoID uuid.UUID, objectKey string) {
		ctx, cancel := context.WithTimeout(parentCtx, 30*time.Second)
		defer cancel()

		result, err := s.mlClient.Analyze(ctx, objectKey)
		if err != nil || result == nil {
			_ = s.analysisRepo.UpdateStatus(ctx, analysisID, domain.Failed)
			return
		}

		_ = s.analysisRepo.Update(ctx, &model.Analysis{
			ID:           analysisID,
			PhotoID:      photoID,
			Status:       domain.Completed,
			SkinType:     &result.SkinType,
			AnalysisData: result.Data,
		})
	}(ctx, analysis.ID, photo.ID, objectKey)

	url, err := s.storage.GetPresignedURL(ctx, objectKey, s.presignTTL)
	if err != nil {
		return nil, err
	}

	return &domain.AnalysisCreated{
		ID:         analysis.ID,
		PhotoID:    photo.ID,
		FileURL:    url,
		Status:     analysis.Status,
		UploadedAt: photo.UploadedAt,
	}, nil
}

func (s *AnalysisService) ListByUserID(
	ctx context.Context,
	userID uuid.UUID,
	offset int,
	limit int,
) ([]*domain.AnalysisListItem, error) {

	analyses, _, err := s.analysisRepo.ListByUserID(ctx, userID, offset, limit)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.AnalysisListItem, 0, len(analyses))

	for _, a := range analyses {

		photo, err := s.photoRepo.GetByID(ctx, a.PhotoID)
		if err != nil {
			return nil, err
		}

		url, err := s.storage.GetPresignedURL(ctx, photo.ObjectKey, s.presignTTL)
		if err != nil {
			return nil, err
		}

		item := &domain.AnalysisListItem{
			ID:        a.ID,
			PhotoID:   a.PhotoID,
			FileURL:   url,
			Status:    a.Status,
			SkinType:  a.SkinType,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
		}

		result = append(result, item)
	}

	return result, nil
}

func (s *AnalysisService) GetByID(
	ctx context.Context,
	userID, analysisID uuid.UUID,
	recsLimit, ingsLimit int,
) (*domain.AnalysisDetail, error) {

	analysis, err := s.analysisRepo.GetByID(ctx, analysisID)
	if err != nil {
		return nil, err
	}

	photo, err := s.photoRepo.GetByID(ctx, analysis.PhotoID)
	if err != nil {
		return nil, err
	}

	if photo.UserID != userID {
		return nil, domain.ErrForbidden
	}

	url, err := s.storage.GetPresignedURL(ctx, photo.ObjectKey, s.presignTTL)
	if err != nil {
		return nil, err
	}

	response := &domain.AnalysisDetail{
		ID:           analysis.ID,
		PhotoID:      analysis.PhotoID,
		FileURL:      url,
		Status:       analysis.Status,
		SkinType:     analysis.SkinType,
		AnalysisData: analysis.AnalysisData,
		CreatedAt:    analysis.CreatedAt,
		UpdatedAt:    analysis.UpdatedAt,
	}

	if analysis.Status == domain.Completed && analysis.SkinType != nil {

		recs, _, err := s.recommendationRepo.ListBySkinType(ctx, *analysis.SkinType, 0, recsLimit)
		if err == nil {
			response.Recommendations = make([]domain.Recommendation, 0, len(recs))
			for _, r := range recs {
				response.Recommendations = append(response.Recommendations, domain.Recommendation{
					ID:          r.ID,
					Title:       r.Title,
					Description: r.Description,
				})
			}
		}

		ings, _, err := s.skinTypeIngredientRepo.ListBySkinType(
			ctx,
			*analysis.SkinType,
			0,
			ingsLimit,
		)

		if err == nil {
			response.Ingredients = make([]domain.Ingredient, 0, len(ings))

			for _, si := range ings {
				ingredient, err := s.ingredientRepo.GetByID(ctx, si.IngredientID)
				if err != nil {
					continue
				}

				response.Ingredients = append(response.Ingredients, domain.Ingredient{
					ID:          ingredient.ID,
					Name:        ingredient.Name,
					Description: ingredient.Description,
				})
			}
		}
	}

	return response, nil
}

func (s *AnalysisService) Delete(
	ctx context.Context,
	userID, analysisID uuid.UUID,
) error {

	analysis, err := s.analysisRepo.GetByID(ctx, analysisID)
	if err != nil {
		return err
	}

	photo, err := s.photoRepo.GetByID(ctx, analysis.PhotoID)
	if err != nil {
		return err
	}

	if photo.UserID != userID {
		return domain.ErrForbidden
	}

	if err := s.storage.Delete(ctx, photo.ObjectKey); err != nil {
		return err
	}

	if err := s.analysisRepo.Delete(ctx, analysisID); err != nil {
		return err
	}

	if err := s.photoRepo.Delete(ctx, photo.ID); err != nil {
		return err
	}

	return nil
}
