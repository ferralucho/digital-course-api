// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/ferralucho/digital-course-api/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// CoursePlanning -.
	CoursePlanning interface {
		OrderCoursePlanning(context.Context, entity.CoursePlanning) (entity.UserOrderedCourse, error)
		CoursePlanning(context.Context) ([]entity.UserOrderedCourse, error)
	}

	// CoursePlanningRepo -.
	CoursePlanningRepo interface {
		Store(context.Context, entity.UserOrderedCourse) error
		GetCoursePlanning(context.Context) ([]entity.UserOrderedCourse, error)
	}
)
