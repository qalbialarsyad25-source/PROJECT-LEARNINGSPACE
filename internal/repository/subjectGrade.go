package repository

import (
	"context"
	"learningSpace/internal/entity"
	"learningSpace/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ISubjectGradeRepository interface {
	CreateSubjectGrade(ctx context.Context, grade entity.SubjectGrade) error
	GetSubjectGrade(ctx context.Context, pagination model.Pagination) ([]entity.SubjectGrade, error)
	DeleteSubjectGrade(ctx context.Context, id uuid.UUID) error
	EditSubjectGrade(ctx context.Context, id uuid.UUID, edit model.EditSubjectGrade) error
}

type SubjectGradeRepository struct {
	db *gorm.DB
}

func NewSubjectGradeRepository(db *gorm.DB) *SubjectGradeRepository {
	return &SubjectGradeRepository{db}
}

func (p *SubjectGradeRepository) CreateSubjectGrade(ctx context.Context, grade entity.SubjectGrade) error {
	err := gorm.G[entity.SubjectGrade](p.db).Create(ctx, &grade)
	if err != nil {
		return err
	}

	return nil
}

func (p *SubjectGradeRepository) GetSubjectGrade(ctx context.Context, pagination model.Pagination) ([]entity.SubjectGrade, error) {
	SubjectGrade, err := gorm.G[entity.SubjectGrade](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("Created_At DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return SubjectGrade, nil
}

func (p *SubjectGradeRepository) DeleteSubjectGrade(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.SubjectGrade](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p *SubjectGradeRepository) EditSubjectGrade(ctx context.Context, id uuid.UUID, edit model.EditSubjectGrade) error {
	result := p.db.WithContext(ctx).Model(&entity.SubjectGrade{}).
		Where("id = ?", id).
		Updates(edit.ToMap())

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
