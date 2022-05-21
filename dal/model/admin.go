package model

import "gorm.io/gorm"

const TableNameAdmin = "admin"

// Admin mapped from table <admin>
type Admin struct {
	gorm.Model
	Name       string `gorm:"type:character varying(32);not null"`
	StandardID *uint
	Standard   *Standard
}

// TableName Admin's table name
func (*Admin) TableName() string {
	return TableNameAdmin
}
