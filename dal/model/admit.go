package model

import (
	"time"
)

const TableNameAdmit = "admit"

// Admit mapped from table <admit>
type Admit struct {
	ApplicantID uint `gorm:"primaryKey"`
	Applicant   Applicant
	Group       string `gorm:"type:character(16);primaryKey"`
	AdminID     uint
	Admin       Admin
	CreatedAt   time.Time
}

// TableName Admit's table name
func (*Admit) TableName() string {
	return TableNameAdmit
}
