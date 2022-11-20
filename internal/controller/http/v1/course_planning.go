package v1

import (
	"net/http"

	"github.com/ferralucho/digital-course-api/internal/entity"
	"github.com/ferralucho/digital-course-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ferralucho/digital-course-api/pkg/logger"
)

type coursePlanningRoutes struct {
	t usecase.CoursePlanning
	l logger.Interface
}

func newCoursePlanningRoutes(handler *gin.RouterGroup, t usecase.CoursePlanning, l logger.Interface) {
	r := &coursePlanningRoutes{t, l}

	h := handler.Group("/course")
	{
		h.GET("/planning", r.coursePlanning)
		h.POST("/order", r.doOrderCourses)
	}
}

type userOrderedResponse struct {
	CoursePlanning entity.OrderedCoursePlanning `json:"course_planning"`
}

// @Summary     Show ordered courses
// @Description Show all ordered courses for the user
// @ID          coursePlanning
// @Tags  	    coursePlanning
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.OrderedCoursePlanning
// @Failure     500 {object} response
// @Router      /course/planning [get]
func (r *coursePlanningRoutes) coursePlanning(c *gin.Context) {
	courses, err := r.t.CoursePlanning(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - course planning")
		errorResponse(c, http.StatusInternalServerError, "server error")

		return
	}

	c.JSON(http.StatusOK, userOrderedResponse{courses})
}

type doOrderCoursesRequest struct {
	UserId  uuid.UUID                   `json:"userId" binding:"required"`
	Courses []entity.CourseRelationship `json:"courses" binding:"required"`
}

// @Description Order user courses
// @ID          do-order
// @Tags  	    coursePlanning
// @Accept      json
// @Produce     json
// @Param       request body doOrderCoursesRequest true "Order courses for the user"
// @Success     200 {object} entity.OrderedCoursePlanning
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /course/order [post]
func (r *coursePlanningRoutes) doOrderCourses(c *gin.Context) {
	var request doOrderCoursesRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doOrderCourses")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	courses, err := r.t.OrderCoursePlanning(
		c.Request.Context(),
		entity.CoursePlanning{
			UserId:  request.UserId,
			Courses: request.Courses,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doOrderCourses")
		errorResponse(c, http.StatusInternalServerError, "courses service problems")

		return
	}

	c.JSON(http.StatusOK, courses)
}
