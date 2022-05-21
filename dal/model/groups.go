package model

import "gorm.io/gorm"

const TableNameGroup = "groups"

// Group mapped from table <groups>
type Group struct {
	ID        uint           `gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

// TableName Group's table name
func (*Group) TableName() string {
	return TableNameGroup
}
