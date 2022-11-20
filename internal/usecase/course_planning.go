package usecase

import (
	"context"
	"fmt"

	"github.com/ferralucho/digital-course-api/internal/entity"
)

// CoursePlanningUseCase -.
type CoursePlanningUseCase struct {
	repo CoursePlanningRepo
}

// New -.
func New(r CoursePlanningRepo) *CoursePlanningUseCase {
	return &CoursePlanningUseCase{
		repo: r,
	}
}

// CoursePlanning - getting course planning for a user from store.
func (uc *CoursePlanningUseCase) CoursePlanning(ctx context.Context, userId int) ([]entity.UserOrderedCourse, error) {
	translations, err := uc.repo.GetCoursePlanning(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("CoursePlanningUseCase - CoursePlanning - s.repo.GetCoursePlanning: %w", err)
	}

	return translations, nil
}

// OrderCoursePlanning -.
func (uc *CoursePlanningUseCase) OrderCoursePlanning(ctx context.Context, t entity.CoursePlanning) (entity.UserOrderedCourse, error) {
	orderedCourse := entity.UserOrderedCourse{}

	return orderedCourse, nil
}
