// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMytable = "mytables"

// Mytable mapped from table <mytables>
type Mytable struct {
	ID       int32   `gorm:"column:ID;type:int(11);not null" json:"ID_example"`
	Username *string `gorm:"column:username;type:varchar(16);index:idx_username,priority:1;default:NULL" json:"username_example"`
	Age      int32   `gorm:"column:age;type:int(8);not null" json:"age_example"`
	Phone    string  `gorm:"column:phone;type:varchar(11);not null" json:"phone_example"`
}

// TableName Mytable's table name
func (*Mytable) TableName() string {
	return TableNameMytable
}
