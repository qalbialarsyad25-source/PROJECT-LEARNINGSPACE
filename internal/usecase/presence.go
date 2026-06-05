package usecase

import (
	"context"
	"learningSpace/internal/entity"
	"learningSpace/internal/model"
	"learningSpace/internal/repository"
	"time"

	"github.com/google/uuid"
)

type IPresenceUsecase interface {
	CreatePresence(ctx context.Context, CreatePresence model.CreatePresence) (*model.PresenceResponse, error)
	ConfirmPresence(ctx context.Context, userId uuid.UUID, ConfirmPresence model.ConfirmPresence) (*model.PresenceResponse, error)
	GetPresence(ctx context.Context, pagination model.Pagination) ([]model.PresenceResponse, error)
	DeletePresence(ctx context.Context, id uuid.UUID) error
	EditPresence(ctx context.Context, id uuid.UUID, edit model.EditPresence) error
}

type PresenceUsecase struct {
	PresenceRepository repository.IPresenceRepository
}

func NewPresenceUsecase(PresenceRepository repository.IPresenceRepository) *PresenceUsecase {
	return &PresenceUsecase{PresenceRepository}
}

func (p *PresenceUsecase) CreatePresence(ctx context.Context, CreatePresence model.CreatePresence) (*model.PresenceResponse, error) {
	Presence := entity.Presence{
		Id:        uuid.New(),
		Presence:  false,
		Title:     CreatePresence.Title,
		Content:   CreatePresence.Content,
		CreatedAt: time.Now(),
	}

	err := p.PresenceRepository.CreatePresence(ctx, Presence)
	if err != nil {
		return nil, err
	}

	response := model.ToPresenceResponse(Presence)
	return &response, nil
}

func (p *PresenceUsecase) ConfirmPresence(ctx context.Context, userId uuid.UUID, ConfirmPresence model.ConfirmPresence) (*model.PresenceResponse, error){
	Presence := entity.Presence{
		Presence: ConfirmPresence.Presence,
	}

	err := p.PresenceRepository.ConfirmPresence(ctx, userId, Presence)
	if err != nil{
		return nil, err
	}

	response := model.ToPresenceResponse(Presence)
	return &response, nil
}

func (p *PresenceUsecase) GetPresence(ctx context.Context, pagination model.Pagination) ([]model.PresenceResponse, error) {
	Presence, err := p.PresenceRepository.GetPresence(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToPresenceResponses(Presence)
	return responses, nil
}

func (p *PresenceUsecase) DeletePresence(ctx context.Context, id uuid.UUID) error {
	return p.PresenceRepository.DeletePresence(ctx, id)
}

func (p *PresenceUsecase) EditPresence(ctx context.Context, id uuid.UUID, edit model.EditPresence) error {
	return p.PresenceRepository.EditPresence(ctx, id, edit)
}
