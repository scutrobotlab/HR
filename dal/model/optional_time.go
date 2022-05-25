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
	Group       string `gorm:"type:character varying(16)"`
	IntentRank  uint   // 限定面试轮次
	TotalCount  uint   // 人数
}

// TableName OptionalTime's table name
func (*OptionalTime) TableName() string {
	return TableNameOptionalTime
}
