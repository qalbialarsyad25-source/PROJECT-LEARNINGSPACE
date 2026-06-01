package repository

import (
	"context"
	"learningSpace/internal/entity"
	"learningSpace/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IReportCardRepository interface {
	CreateReportCard(ctx context.Context, report entity.ReportCard) error
	GetReportCard(ctx context.Context, pagination model.Pagination) ([]entity.ReportCard, error)
	DeleteReportCard(ctx context.Context, id uuid.UUID) error
	EditReportCard(ctx context.Context, id uuid.UUID, edit model.EditReportCard) error
}

type ReportCardRepository struct {
	db *gorm.DB
}

func NewReportCardRepository(db *gorm.DB) *ReportCardRepository {
	return &ReportCardRepository{db}
}

func (p *ReportCardRepository) CreateReportCard(ctx context.Context, report entity.ReportCard) error {
	err := gorm.G[entity.ReportCard](p.db).Create(ctx, &report)
	if err != nil {
		return err
	}

	return nil
}

func (p *ReportCardRepository) GetReportCard(ctx context.Context, pagination model.Pagination) ([]entity.ReportCard, error) {
	ReportCard, err := gorm.G[entity.ReportCard](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("Created_At DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return ReportCard, nil
}

func (p *ReportCardRepository) DeleteReportCard(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.ReportCard](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p *ReportCardRepository) EditReportCard(ctx context.Context, id uuid.UUID, edit model.EditReportCard) error {
	result := p.db.WithContext(ctx).Model(&entity.ReportCard{}).
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
