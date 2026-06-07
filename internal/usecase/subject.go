package usecase

import (
	"context"
	"learningSpace/internal/entity"
	"learningSpace/internal/model"
	"learningSpace/internal/repository"

	"github.com/google/uuid"
)

type ISubjectUsecase interface {
	CreateSubject(ctx context.Context, Subject model.CreateSubject) (*model.SubjectResponse, error)
	GetSubject(ctx context.Context, pagination model.Pagination) ([]model.SubjectResponse, error)
	DeleteSubject(ctx context.Context, id uuid.UUID) error
	EditSubject(ctx context.Context, id uuid.UUID, edit model.EditSubject) error
}

type SubjectUsecase struct {
	SubjectRepository repository.ISubjectRepository
}

func NewSubjectUsecase(SubjectRepository repository.ISubjectRepository) *SubjectUsecase {
	return &SubjectUsecase{SubjectRepository}
}

func (p *SubjectUsecase) CreateSubject(ctx context.Context, CreateSubject model.CreateSubject) (*model.SubjectResponse, error) {
	Subject := entity.Subject{
		Id:        uuid.New(),
		Name:      CreateSubject.Name,
	}

	err := p.SubjectRepository.CreateSubject(ctx, Subject)
	if err != nil {
		return nil, err
	}

	response := model.ToSubjectResponse(Subject)
	return &response, nil
}

func (p *SubjectUsecase) GetSubject(ctx context.Context, pagination model.Pagination) ([]model.SubjectResponse, error) {
	Subject, err := p.SubjectRepository.GetSubject(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToSubjectResponses(Subject)
	return responses, nil
}

func (p *SubjectUsecase) DeleteSubject(ctx context.Context, id uuid.UUID) error {
	return p.SubjectRepository.DeleteSubject(ctx, id)
}

func (p *SubjectUsecase) EditSubject(ctx context.Context, id uuid.UUID, edit model.EditSubject) error {
	return p.SubjectRepository.EditSubject(ctx, id, edit)
}
