package model

import (
	"time"
)

const TableNameAdmit = "admit"

// Admit mapped from table <admit>
type Admit struct {
	ApplicantID int32 `gorm:"primaryKey"`
	GroupID     int32 `gorm:"primaryKey"`
	AdminID     int32
	CreatedAt   time.Time
}

// TableName Admit's table name
func (*Admit) TableName() string {
	return TableNameAdmit
}
