package model

const TableNameRemark = "remarks"

// Remark mapped from table <remarks>
type Remark struct {
	AdminID     uint `gorm:"primaryKey"`
	Admin       Admin
	ApplicantID uint `gorm:"primaryKey"`
	Applicant   Applicant
	TheRemark   string `gorm:"type:text"`
}

// TableName Remark's table name
func (*Remark) TableName() string {
	return TableNameRemark
}
