package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameOptionalTime = "optional_time"

// OptionalTime mapped from table <optional_time>
type OptionalTime struct {
	gorm.Model  `swaggerignore:"true"`
	TheDate     time.Time `gorm:"type:date"`
	TheTime     time.Time `gorm:"type:time with time zone"`
	TheLocation string
	Group       string `gorm:"type:character(16)"`
	IntentRank  *int16 // 限定面试轮次
}

// TableName OptionalTime's table name
func (*OptionalTime) TableName() string {
	return TableNameOptionalTime
}
