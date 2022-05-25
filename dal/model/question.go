package model

import "gorm.io/gorm"

const TableNameQuestion = "question"

// Question mapped from table <question>
type Question struct {
	gorm.Model  `swaggerignore:"true"`
	TheQuestion string `gorm:"type:text" json:"question"`
	Group       string `gorm:"type:character varying(16)" json:"group"`
}

// TableName Question's table name
func (*Question) TableName() string {
	return TableNameQuestion
}
