package model

import (
	"time"
)

const TableNameScore = "scores"

// Score mapped from table <scores>
type Score struct {
	AdminID           int32 `gorm:"primaryKey"`
	ApplicantID       int32 `gorm:"primaryKey"`
	GroupID           int32 `gorm:"primaryKey"`
	Score             float64
	StandardID        *int32
	EvaluationDetails *string `gorm:"type:jsonb"`
	UpdatedAt         time.Time
}

// TableName Score's table name
func (*Score) TableName() string {
	return TableNameScore
}
