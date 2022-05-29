package model

import (
	"gorm.io/gorm"
)

const TableNameApplicant = "applicant"

// Applicant mapped from table <applicant>
type Applicant struct {
	gorm.Model `swaggerignore:"true"`
	OpenID     string            `gorm:"type:character varying(64);unique;not null"`
	Name       string            `gorm:"type:character varying(16);not null"`
	Gender     *bool             `gorm:"type:boolean"`
	Phone      *string           `gorm:"type:character varying(16)"`
	Avatar     *string           `gorm:"type:character varying(256)"`
	Form       map[string]string `gorm:"type:jsonb;not null"`
	Intents    []Intent          `gorm:"foreignkey:ApplicantID"`
	Answers    []Answer          `gorm:"foreignkey:ApplicantID"`
}

// TableName Applicant's table name
func (*Applicant) TableName() string {
	return TableNameApplicant
}
