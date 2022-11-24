package usecase

import (
	"context"
	"fmt"
	"strings"

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
	orderedCourses := make([]entity.OrderedCourseRelationship, 0, len(t.Courses)+1)

	if len(t.Courses) == 0 {
		return entity.OrderedCoursePlanning{}, nil
	}
	desiredCourses := make([]string, 0, len(t.Courses)-1)

	courseMap := make(map[string]string)
	for _, v := range t.Courses {
		desiredCourses = append(desiredCourses, v.DesiredCourse)
		if val, ok := courseMap[v.RequiredCourse]; ok {
			courseMap[v.RequiredCourse] = val + "-" + v.DesiredCourse
		} else {
			courseMap[v.RequiredCourse] = v.DesiredCourse
		}
	}

	order := 0
	for k := range courseMap {
		if !contains(desiredCourses, k) {
			orderedCourses = append(orderedCourses, entity.OrderedCourseRelationship{CourseName: k, Order: 0})

			for i := 0; i < len(desiredCourses)-1; i++ {
				if courseMap[orderedCourses[i].CourseName] != "" {
					splitted := strings.Split(courseMap[orderedCourses[i].CourseName], "-")
					if len(splitted) == 1 {
						order += 1
						orderedCourses = append(orderedCourses, entity.OrderedCourseRelationship{CourseName: courseMap[orderedCourses[i].CourseName], Order: order})
					} else {
						for j := 0; j < len(splitted); j++ {
							order += 1
							orderedCourses = append(orderedCourses, entity.OrderedCourseRelationship{CourseName: splitted[j], Order: order})
						}
					}

				}
			}

			break
		}
	}

	if len(orderedCourses) > 0 {
		err := uc.repo.DeleteUserCourses(ctx, t.UserId)
		if err != nil {
			return entity.OrderedCoursePlanning{}, fmt.Errorf("CoursePlanningUseCase - CoursePlanning - s.repo.DeleteUserCourses: %w", err)
		}

		var orderedCourse entity.UserOrderedCourse

		for _, ocr := range orderedCourses {
			orderedCourse = entity.UserOrderedCourse{
				UserId:     t.UserId,
				CourseName: ocr.CourseName,
				Order:      ocr.Order,
			}

			err := uc.repo.Store(ctx, orderedCourse)

			if err != nil {
				return entity.OrderedCoursePlanning{}, fmt.Errorf("CoursePlanningUseCase - CoursePlanning - s.repo.Store: %w", err)
			}
		}

	}

	orderedCourse := entity.OrderedCoursePlanning{
		UserId:  t.UserId,
		Courses: orderedCourses,
	}

	return orderedCourse, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
