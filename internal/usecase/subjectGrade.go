package usecase

import (
	"context"
	"learningSpace/internal/entity"
	"learningSpace/internal/model"
	"learningSpace/internal/repository"

	"github.com/google/uuid"
)

type ISubjectGradeUsecase interface {
	CreateSubjectGrade(ctx context.Context, SubjectGrade model.CreateSubjectGrade) (*model.SubjectGradeResponse, error)
	GetSubjectGrade(ctx context.Context, pagination model.Pagination) ([]model.SubjectGradeResponse, error)
	DeleteSubjectGrade(ctx context.Context, id uuid.UUID) error
	EditSubjectGrade(ctx context.Context, id uuid.UUID, edit model.EditSubjectGrade) error
}

type SubjectGradeUsecase struct {
	SubjectGradeRepository repository.ISubjectGradeRepository
}

func NewSubjectGradeUsecase(SubjectGradeRepository repository.ISubjectGradeRepository) *SubjectGradeUsecase {
	return &SubjectGradeUsecase{SubjectGradeRepository}
}

func (p *SubjectGradeUsecase) CreateSubjectGrade(ctx context.Context, CreateSubjectGrade model.CreateSubjectGrade) (*model.SubjectGradeResponse, error) {
	SubjectGrade := entity.SubjectGrade{
		Id:        uuid.New(),
		Grade:      CreateSubjectGrade.Grade,
	}

	err := p.SubjectGradeRepository.CreateSubjectGrade(ctx, SubjectGrade)
	if err != nil {
		return nil, err
	}

	response := model.ToSubjectGradeResponse(SubjectGrade)
	return &response, nil
}

func (p *SubjectGradeUsecase) GetSubjectGrade(ctx context.Context, pagination model.Pagination) ([]model.SubjectGradeResponse, error) {
	SubjectGrade, err := p.SubjectGradeRepository.GetSubjectGrade(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToSubjectGradeResponses(SubjectGrade)
	return responses, nil
}

func (p *SubjectGradeUsecase) DeleteSubjectGrade(ctx context.Context, id uuid.UUID) error {
	return p.SubjectGradeRepository.DeleteSubjectGrade(ctx, id)
}

func (p *SubjectGradeUsecase) EditSubjectGrade(ctx context.Context, id uuid.UUID, edit model.EditSubjectGrade) error {
	return p.SubjectGradeRepository.EditSubjectGrade(ctx, id, edit)
}
