package model

import (
	"learningSpace/internal/entity"
	"time"

	"github.com/google/uuid"
)

type CreatePresence struct {
	Presence bool   `json:"presence" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Title    string `json:"title" validate:"required"`
}

type ConfirmPresence struct {
	Presence bool `json:"presence"`
}

type PresenceResponse struct {
	Id        uuid.UUID `json:"id"`
	Presence  bool      `json:"presence"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"Created_At"`
}

type EditPresence struct {
	Presence bool   `json:"presence"`
	Content  string `json:"content"`
	Title    string `json:"title"`
}

func ToPresenceResponse(presence entity.Presence) PresenceResponse {
	return PresenceResponse{
		Id:       presence.Id,
		Presence: presence.Presence,
		Content:  presence.Content,
		Title:    presence.Title,
	}
}

func ToPresenceResponses(presence []entity.Presence) []PresenceResponse{
	var response []PresenceResponse
	for _, presence := range presence {
		response = append(response, ToPresenceResponse(presence))
	}

	return response
}

func (p *EditPresence) ToMap() map[string]any {
	Update := map[string]any{}

	if p.Presence != false {
		Update["presence"] = p.Presence
	}

	if p.Content != "" {
		Update["content"] = p.Content
	}

	if p.Title != "" {
		Update["title"] = p.Title
	}

	return Update
}
