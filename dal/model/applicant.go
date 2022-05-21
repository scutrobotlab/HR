package model

import (
	"gorm.io/gorm"
)

const TableNameApplicant = "applicant"

// Applicant mapped from table <applicant>
type Applicant struct {
	gorm.Model
	Name    string  `gorm:"type:character varying(32);not null"`
	Gender  *bool   `gorm:"type:boolean"`
	Phone   *string `gorm:"type:character varying(16)"`
	Avatar  *string `gorm:"type:character varying(256)"`
	Form    string  `gorm:"type:jsonb;not null"`
	Intents []Intent
}

// TableName Applicant's table name
func (*Applicant) TableName() string {
	return TableNameApplicant
}
