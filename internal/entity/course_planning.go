package entity

import "github.com/google/uuid"

type CoursePlanning struct {
	UserId  uuid.UUID            `json:"userId,type:uuid"`
	Courses []CourseRelationship `json:"courses"`
}

type CourseRelationship struct {
	DesiredCourse  string `json:"desiredCourse"`
	RequiredCourse string `json:"requiredCourse"`
}

type OrderedCoursePlanning struct {
	UserId  uuid.UUID                   `json:"userId,type:uuid"`
	Courses []OrderedCourseRelationship `json:"courses"`
}

type OrderedCourseRelationship struct {
	CourseName string `json:"courseName"`
	Order      int    `json:"order"`
}

type UserOrderedCourse struct {
	UserId     uuid.UUID `json:"user_id"`
	CourseName string    `json:"course_name"`
	Order      int       `json:"course_order"`
}
