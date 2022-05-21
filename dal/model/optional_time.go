package model

import (
	"time"
)

const TableNameOptionalTime = "optional_time"

// OptionalTime mapped from table <optional_time>
type OptionalTime struct {
	ID          uint      `gorm:"primaryKey"`
	TheDate     time.Time `gorm:"type:date;not null"`
	TheTime     time.Time `gorm:"type:time;not null"`
	TheLocation *string   `gorm:"index"`
	GroupID     int32
	IntentRank  *int16 // 限定面试轮次
}

// TableName OptionalTime's table name
func (*OptionalTime) TableName() string {
	return TableNameOptionalTime
}
