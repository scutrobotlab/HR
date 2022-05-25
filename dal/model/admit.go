package model

import (
	"time"
)

const TableNameAdmit = "admit"

// Admit mapped from table <admit>
type Admit struct {
	ApplicantID uint      `gorm:"primaryKey"`
	Applicant   Applicant `json:"-"`
	AdminID     uint      `swaggerignore:"true"`
	Admin       Admin     `json:"-"`
	Group       string    `gorm:"type:character varying(16);primaryKey"`
	UpdatedAt   time.Time `swaggerignore:"true"`
}

// TableName Admit's table name
func (*Admit) TableName() string {
	return TableNameAdmit
}
