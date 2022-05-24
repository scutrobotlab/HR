package model

import (
	"time"
)

const TableNameScore = "scores"

// Score mapped from table <scores>
type Score struct {
	AdminID           uint      `gorm:"primaryKey"`
	Admin             Admin     `json:"-"`
	ApplicantID       uint      `gorm:"primaryKey"`
	Applicant         Applicant `json:"-"`
	Group             string    `gorm:"type:character(16);primaryKey"`
	Score             float64
	StandardID        uint
	Standard          Standard `json:"-"`
	EvaluationDetails string   `gorm:"type:jsonb"`
	UpdatedAt         time.Time
}

// TableName Score's table name
func (*Score) TableName() string {
	return TableNameScore
}
