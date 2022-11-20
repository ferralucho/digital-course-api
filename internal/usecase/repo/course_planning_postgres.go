package repo

import (
	"context"
	"fmt"

	"github.com/ferralucho/digital-course-api/internal/entity"
	"github.com/ferralucho/digital-course-api/pkg/postgres"
)

const _defaultEntityCap = 64

// CoursePlanningRepo -.
type CoursePlanningRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *CoursePlanningRepo {
	return &CoursePlanningRepo{pg}
}

// GetCoursePlanning -.
func (r *CoursePlanningRepo) GetCoursePlanning(ctx context.Context, userId int) ([]entity.UserOrderedCourse, error) {
	sql, _, err := r.Builder.
		Select("user_id, course_name, course_order").
		From("user_course_planning").
		Where("user_id = ?", userId).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("CoursePlanningRepo - GetCoursePlanning - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("CoursePlanningRepo - GetCoursePlanning - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.UserOrderedCourse, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.UserOrderedCourse{}

		err = rows.Scan(&e.UserId, &e.CourseName, &e.Order)
		if err != nil {
			return nil, fmt.Errorf("CoursePlanningRepo - GetCoursePlanning - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// Store -.
func (r *CoursePlanningRepo) Store(ctx context.Context, t entity.UserOrderedCourse) error {
	sql, args, err := r.Builder.
		Insert("user_course_planning").
		Columns("user_id, course_order, course_name, ").
		Values(t.UserId, t.Order, t.CourseName).
		ToSql()
	if err != nil {
		return fmt.Errorf("CoursePlanningRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("CoursePlanningRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteUserCourses -.
func (r *CoursePlanningRepo) DeleteUserCourses(ctx context.Context, t entity.UserOrderedCourse) error {
	sql, args, err := r.Builder.
		Delete("user_course_planning").
		Where("user_id = ?", t.UserId).
		ToSql()
	if err != nil {
		return fmt.Errorf("CoursePlanningRepo - DeleteUserCourses - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("CoursePlanningRepo - DeleteUserCourses - r.Pool.Exec: %w", err)
	}

	return nil
}
