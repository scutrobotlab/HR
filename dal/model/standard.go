package model

import "gorm.io/gorm"

const TableNameStandard = "standard"

// Standard mapped from table <standard>
type Standard struct {
	gorm.Model
	Name    string
	Content string `gorm:"type:jsonb;not null"`
	AdminID uint   // 上次修改的管理员
}

// TableName Standard's table name
func (*Standard) TableName() string {
	return TableNameStandard
}
