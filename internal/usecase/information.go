package usecase

import (
	"context"
	"learningSpace/internal/entity"
	"learningSpace/internal/model"
	"learningSpace/internal/repository"
	"time"

	"github.com/google/uuid"
)

type IInformationUsecase interface {
	CreateInformation(ctx context.Context, Information model.CreateInformation) (*model.InformationResponse, error)
	GetInformation(ctx context.Context, pagination model.Pagination) ([]model.InformationResponse, error)
	DeleteInformation(ctx context.Context, id uuid.UUID) error
	EditInformation(ctx context.Context, id uuid.UUID, edit model.EditInformation) error
}

type InformationUsecase struct {
	InformationRepository repository.IInformationRepository
}

func NewInformationUsecase(InformationRepository repository.IInformationRepository) *InformationUsecase {
	return &InformationUsecase{InformationRepository}
}

func (p *InformationUsecase) CreateInformation(ctx context.Context, CreateInformation model.CreateInformation) (*model.InformationResponse, error) {
	Information := entity.Information{
		Id:        uuid.New(),
		Name:      CreateInformation.Name,
		Title:     CreateInformation.Title,
		Content:   CreateInformation.Content,
		CreatedAt: time.Now(),
	}

	err := p.InformationRepository.CreateInformation(ctx, Information)
	if err != nil {
		return nil, err
	}

	response := model.ToInformationResponse(Information)
	return &response, nil
}

func (p *InformationUsecase) GetInformation(ctx context.Context, pagination model.Pagination) ([]model.InformationResponse, error) {
	information, err := p.InformationRepository.GetInformation(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToInformationResponses(information)
	return responses, nil
}

func (p *InformationUsecase) DeleteInformation(ctx context.Context, id uuid.UUID) error {
	return p.InformationRepository.DeleteInformation(ctx, id)
}

func (p *InformationUsecase) EditInformation(ctx context.Context, id uuid.UUID, edit model.EditInformation) error {
	return p.InformationRepository.EditInformation(ctx, id, edit)
}
