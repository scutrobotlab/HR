package model

import "gorm.io/gorm"

const TableNameQuestion = "question"

// Question mapped from table <question>
type Question struct {
	gorm.Model  `swaggerignore:"true"`
	TheQuestion string `gorm:"type:text"`
	Group       string `gorm:"type:character(16)"`
}

// TableName Question's table name
func (*Question) TableName() string {
	return TableNameQuestion
}
