package model

const TableNameSetting = "settings"

// Setting mapped from table <settings>
type Setting struct {
	Key   string `gorm:"type:character varying(32);primaryKey"`
	Value string `gorm:"type:json;not null"`
}

// TableName Setting's table name
func (*Setting) TableName() string {
	return TableNameSetting
}
