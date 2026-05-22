package model

import (
	"learningSpace/internal/entity"

	"github.com/google/uuid"
)

type ReportCardResponse struct {
	Id      uuid.UUID `json:"id"`
	Average float64   `json:"average"`
}

type CreateReportCard struct {
	Average float64 `json:"average"`
}

type EditReportCard struct {
	Average float64 `json:"average"`
}

func ToReportCardResponse(reportCard entity.ReportCard) ReportCardResponse {
	return ReportCardResponse{
		Id:      reportCard.Id,
		Average: reportCard.Average,
	}
}

func (p *EditReportCard) ToMap() map[string]any {
	Update := map[string]any{}

	if p.Average != 0 {
		Update["average"] = p.Average
	}

	return Update
}
