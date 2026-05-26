package repository

import (
	"context"
	"learningSpace/internal/entity"
	"learningSpace/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IPresenceRepository interface {
	CreatePresence(ctx context.Context, presence entity.Presence) error
	GetPresence(ctx context.Context, pagination model.Pagination) ([]entity.Presence, error)
	DeletePresence(ctx context.Context, id uuid.UUID) error
	EditPresence(ctx context.Context, id uuid.UUID, edit model.EditPresence) error
}

type PresenceRepository struct {
	db *gorm.DB
}

func NewPresenceRepository(db *gorm.DB) *PresenceRepository {
	return &PresenceRepository{db}
}

func (p *PresenceRepository) CreatePresence(ctx context.Context, presence entity.Presence) error {
	err := gorm.G[entity.Presence](p.db).Create(ctx, &presence)
	if err != nil {
		return err
	}

	return nil
}

func (p *PresenceRepository) GetPresence(ctx context.Context, pagination model.Pagination) ([]entity.Presence, error) {
	presence, err := gorm.G[entity.Presence](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("Created_At DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return presence, nil
}

func (p *PresenceRepository) DeletePresence(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Presence](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p *PresenceRepository) EditPresence(ctx context.Context, id uuid.UUID, edit model.EditPresence) error {
	result := p.db.WithContext(ctx).Model(&entity.Presence{}).
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
