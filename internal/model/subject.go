package model

import (
	"learningSpace/internal/entity"

	"github.com/google/uuid"
)

type CreateSubject struct {
	Name string `json:"name" validate:"required"`
}

type SubjectResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type EditSubject struct {
	Name string `json:"name"`
}

func ToSubjectResponse(subject entity.Subject) SubjectResponse {
	return SubjectResponse{
		Id:   subject.Id,
		Name: subject.Name,
	}
}

func ToSubjectResponses(subject []entity.Subject) []SubjectResponse {
	var responses []SubjectResponse
	for _, subject := range subject{
		responses = append(responses, ToSubjectResponse(subject))
	}

	return responses
}

func (p *EditSubject) ToMap() map[string]any {
	Update := map[string]any{}

	if p.Name != "" {
		Update["name"] = p.Name
	}

	return Update
}
