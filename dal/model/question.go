package model

import "gorm.io/gorm"

const TableNameQuestion = "question"

// Question mapped from table <question>
type Question struct {
	ID          int32          `gorm:"primaryKey"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	TheQuestion *string        `gorm:"type:text"`
	GroupID     int32
}

// TableName Question's table name
func (*Question) TableName() string {
	return TableNameQuestion
}
