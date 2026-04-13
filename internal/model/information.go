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

func (p *EditInformation) ToMap() map[string]any {
	Update := map[string]any{}

	if p.Name != ""{
		Update["name"] = p.Name
	}

	if p.Title != ""{
		Update["title"] = p.Title
	}

	if p.Content != ""{
		Update["content"] = p.Content
	}

	return Update
}
