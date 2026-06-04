package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository         IUserRepository
	InformationRepository  IInformationRepository
	PresenceRepository     IPresenceRepository
	ReportCardRepository   IReportCardRepository
	SubjectGradeRepository ISubjectGradeRepository
	SubjectRepository      ISubjectRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:         NewUserRepository(db),
		InformationRepository:  NewInformationRepository(db),
		PresenceRepository:     NewPresenceRepository(db),
		ReportCardRepository:   NewReportCardRepository(db),
		SubjectGradeRepository: NewSubjectGradeRepository(db),
		SubjectRepository:      NewSubjectRepository(db),
	}
}
