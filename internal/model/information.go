package model

import (
	"learningSpace/internal/entity"

	"github.com/google/uuid"
)

type InformationResponse struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

type CreateInformation struct {
	Name    string `json:"name" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type EditInformation struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ToInformationResponse(information entity.Information) InformationResponse {
	return InformationResponse{
		Id:      information.Id,
		Name:    information.Name,
		Title:   information.Title,
		Content: information.Content,
	}
}
