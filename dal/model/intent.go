package model

import "gorm.io/gorm"

const TableNameIntent = "intent"

// Intent mapped from table <intent>
type Intent struct {
	gorm.Model
	ApplicantID    uint
	Applicant      Applicant
	Group          string `gorm:"type:character(16)"`
	IntentRank     *int16 // nil 为平行志愿
	OptionalTimeID uint
	OptionalTime   OptionalTime
}

// TableName Intent's table name
func (*Intent) TableName() string {
	return TableNameIntent
}