package model

const TableNameIntent = "intent"

// Intent mapped from table <intent>
type Intent struct {
	ID          int32 `gorm:"primaryKey"`
	ApplicantID int32
	GroupID     int32
	IntentRank  *int16 // nil 为平行志愿
}

// TableName Intent's table name
func (*Intent) TableName() string {
	return TableNameIntent
}
