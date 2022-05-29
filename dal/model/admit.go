package model

import (
	"time"
)

const TableNameAdmit = "admit"

// Admit mapped from table <admit>
type Admit struct {
	Group       string    `gorm:"type:character varying(16);primaryKey"`
	ApplicantID uint      `gorm:"primaryKey;autoIncrement:false"`
	Applicant   Applicant `json:"-"`
	AdminID     uint      `swaggerignore:"true"`
	Admin       Admin     `json:"-"`
	UpdatedAt   time.Time `swaggerignore:"true"`
}

// TableName Admit's table name
func (*Admit) TableName() string {
	return TableNameAdmit
}
