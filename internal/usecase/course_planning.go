package usecase

import (
	"context"
	"fmt"

	"github.com/ferralucho/digital-course-api/internal/entity"
	"github.com/google/uuid"
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
func (uc *CoursePlanningUseCase) CoursePlanning(ctx context.Context, userId uuid.UUID) (entity.OrderedCoursePlanning, error) {
	courses, err := uc.repo.GetCoursePlanning(ctx, userId)
	if err != nil {
		return entity.OrderedCoursePlanning{}, fmt.Errorf("CoursePlanningUseCase - CoursePlanning - s.repo.GetCoursePlanning: %w", err)
	}

	relationships := make([]entity.OrderedCourseRelationship, len(courses))
	for i, c := range courses {
		relationships[i] = entity.OrderedCourseRelationship{
			CourseName: c.CourseName,
			Order:      c.Order,
		}
	}

	ordered := entity.OrderedCoursePlanning{
		UserId:  userId,
		Courses: relationships,
	}

	return ordered, nil
}

// OrderCoursePlanning -.
func (uc *CoursePlanningUseCase) OrderCoursePlanning(ctx context.Context, t entity.CoursePlanning) (entity.OrderedCoursePlanning, error) {
	orderedCourse := entity.OrderedCoursePlanning{}

	return orderedCourse, nil
}
