package entity

import "github.com/google/uuid"

type CoursePlanning struct {
	UserId  uuid.UUID            `json:"userId,type:uuid,default:uuid_generate_v4()"`
	Courses []CourseRelationship `json:"courses"`
}

type CourseRelationship struct {
	DesiredCourse  string `json:"desiredCourse"`
	RequiredCourse string `json:"requiredCourse"`
}

type OrderedCoursePlanning struct {
	UserId  uuid.UUID                   `json:"userId,type:uuid,default:uuid_generate_v4()"`
	Courses []OrderedCourseRelationship `json:"courses"`
}

type OrderedCourseRelationship struct {
	CourseName string `json:"courseName"`
	Order      int    `json:"order"`
}
