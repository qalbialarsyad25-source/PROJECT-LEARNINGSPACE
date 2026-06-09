package usecase

import (
	"context"
	"learningSpace/internal/entity"
	"learningSpace/internal/model"
	"learningSpace/internal/repository"

	"github.com/google/uuid"
)

type IReportCardUsecase interface {
	CreateReportCard(ctx context.Context, ReportCard model.CreateReportCard) (*model.ReportCardResponse, error)
	GetReportCard(ctx context.Context, pagination model.Pagination) ([]model.ReportCardResponse, error)
	DeleteReportCard(ctx context.Context, id uuid.UUID) error
	EditReportCard(ctx context.Context, id uuid.UUID, edit model.EditReportCard) error
}

type ReportCardUsecase struct {
	ReportCardRepository repository.IReportCardRepository
}

func NewReportCardUsecase(ReportCardRepository repository.IReportCardRepository) *ReportCardUsecase {
	return &ReportCardUsecase{ReportCardRepository}
}

func (p *ReportCardUsecase) CreateReportCard(ctx context.Context, CreateReportCard model.CreateReportCard) (*model.ReportCardResponse, error) {
	ReportCard := entity.ReportCard{
		Id:        uuid.New(),
		Average:      CreateReportCard.Average,
	}

	err := p.ReportCardRepository.CreateReportCard(ctx, ReportCard)
	if err != nil {
		return nil, err
	}

	response := model.ToReportCardResponse(ReportCard)
	return &response, nil
}

func (p *ReportCardUsecase) GetReportCard(ctx context.Context, pagination model.Pagination) ([]model.ReportCardResponse, error) {
	ReportCard, err := p.ReportCardRepository.GetReportCard(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToReportCardResponses(ReportCard)
	return responses, nil
}

func (p *ReportCardUsecase) DeleteReportCard(ctx context.Context, id uuid.UUID) error {
	return p.ReportCardRepository.DeleteReportCard(ctx, id)
}

func (p *ReportCardUsecase) EditReportCard(ctx context.Context, id uuid.UUID, edit model.EditReportCard) error {
	return p.ReportCardRepository.EditReportCard(ctx, id, edit)
}
