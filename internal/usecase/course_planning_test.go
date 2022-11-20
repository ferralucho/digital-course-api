package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ferralucho/digital-course-api/internal/entity"
	"github.com/ferralucho/digital-course-api/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func coursePlanning(t *testing.T) (*usecase.CoursePlanningUseCase, *MockCoursePlanningRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockCoursePlanningRepo(mockCtl)

	coursePlanning := usecase.New(repo)

	return coursePlanning, repo
}

func TestCoursePlanning(t *testing.T) {

	t.Parallel()

	coursePlanning, repo := coursePlanning(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().GetCoursePlanning(context.Background(), 1).Return(nil, nil)
			},
			res: []entity.UserOrderedCourse(nil),
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().GetCoursePlanning(context.Background(), 1).Return(nil, errInternalServErr)
			},
			res: []entity.UserOrderedCourse(nil),
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := coursePlanning.CoursePlanning(context.Background(), 1)

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
