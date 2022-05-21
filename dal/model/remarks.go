package model

const TableNameRemark = "remarks"

// Remark mapped from table <remarks>
type Remark struct {
	AdminID     int32  `gorm:"primaryKey"`
	ApplicantID int32  `gorm:"primaryKey"`
	TheRemark   string `gorm:"type:text"`
}

// TableName Remark's table name
func (*Remark) TableName() string {
	return TableNameRemark
}
