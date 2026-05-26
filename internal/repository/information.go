package repository 

import(
	"learningSpace/internal/entity"
	"learningSpace/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IInformationRepository interface{
	CreateInformation(ctx context.Context, information entity.Information) error
	GetInformation(ctx context.Context, pagination model.Pagination) ([]entity.Information, error)
	DeleteInformation(ctx context.Context, id uuid.UUID) error
	EditInformation(ctx context.Context, id uuid.UUID, edit model.EditInformation) error
}

type InformationRepository struct{
	db *gorm.DB
}

func NewInformationRepository(db *gorm.DB) *InformationRepository{
	return &InformationRepository{db}
}

func (p *InformationRepository) CreateInformation(ctx context.Context, information entity.Information) error {
	err := gorm.G[entity.Information](p.db).Create(ctx, &information)
	if err != nil {
		return err
	}

	return nil
}

func (p *InformationRepository) GetInformation(ctx context.Context, pagination model.Pagination) ([]entity.Information, error) {
	information, err := gorm.G[entity.Information](p.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("Created_At DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return information, nil
}

func (p *InformationRepository) DeleteInformation(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Information](p.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p *InformationRepository) EditInformation(ctx context.Context, id uuid.UUID, edit model.EditInformation) error {
	result := p.db.WithContext(ctx).Model(&entity.Information{}).
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