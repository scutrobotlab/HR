package model

const TableNameAdmin = "admin"

// Admin mapped from table <admin>
type Admin struct {
	ID         uint      `gorm:"primaryKey" json:"-"`
	Name       string    `gorm:"type:character varying(16);not null" json:"-"`
	StandardID *uint     `json:"standard_id"`
	Standard   *Standard `json:"-"`
}

// TableName Admin's table name
func (*Admin) TableName() string {
	return TableNameAdmin
}
