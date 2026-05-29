package repository

import (
	"context"
	"learningSpace/internal/entity"
	"learningSpace/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ISubjectRepository interface {
	CreateSubject(ctx context.Context, subject entity.Subject) error
	GetSubject(ctx context.Context, pagination model.Pagination) ([]entity.Subject, error)
	DeleteSubject(ctx context.Context, id uuid.UUID) error
	EditSubject(ctx context.Context, id uuid.UUID, edit model.EditSubject) error
}

type SubjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{db}
}

func (p *SubjectRepository) CreateSubject(ctx context.Context, subject entity.Subject) error {
	err := gorm.G[entity.Subject](p.db).Create(ctx, &subject)
	if err != nil {
		return err
	}

	return nil
}

func (p *SubjectRepository) GetSubject(ctx context.Context, pagination model.Pagination) ([]entity.Subject, error) {
	subject, err := gorm.G[entity.Subject](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("Created_At DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (p *SubjectRepository) DeleteSubject(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Subject](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p *SubjectRepository) EditSubject(ctx context.Context, id uuid.UUID, edit model.EditSubject) error {
	result := p.db.WithContext(ctx).Model(&entity.Subject{}).
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
