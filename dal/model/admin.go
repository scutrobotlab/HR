package model

const TableNameAdmin = "admin"

// Admin mapped from table <admin>
type Admin struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:character varying(32);not null"`
	StandardID *uint
	Standard   *Standard
}

// TableName Admin's table name
func (*Admin) TableName() string {
	return TableNameAdmin
}
