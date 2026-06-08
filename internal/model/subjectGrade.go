package model

import (
	"learningSpace/internal/entity"

	"github.com/google/uuid"
)

type SubjectGradeResponse struct {
	Id    uuid.UUID `json:"id"`
	Grade float64   `json:"grade"`
}

type CreateSubjectGrade struct {
	Grade float64 `json:"grade"`
}

type EditSubjectGrade struct {
	Grade float64 `json:"grade"`
}

func ToSubjectGradeResponse(subjectGrade entity.SubjectGrade) SubjectGradeResponse {
	return SubjectGradeResponse{
		Id:    subjectGrade.Id,
		Grade: subjectGrade.Grade,
	}
}

func ToSubjectGradeResponses(subjectGrade []entity.SubjectGrade)  []SubjectGradeResponse{
	var responses []SubjectGradeResponse
	for _, subjectGrade := range subjectGrade{
		responses = append(responses, ToSubjectGradeResponse(subjectGrade))
	}

	return responses
}

func (p *EditSubjectGrade) ToMap() map[string]any {
	Update := map[string]any{}

	if p.Grade != 0 {
		Update["grade"] = p.Grade
	}

	return Update
}
