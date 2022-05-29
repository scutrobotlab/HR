package model

const TableNameRemark = "remarks"

// Remark mapped from table <remarks>
type Remark struct {
	AdminID     uint      `gorm:"primaryKey;autoIncrement:false" json:"-"`
	Admin       Admin     `json:"-"`
	ApplicantID uint      `gorm:"primaryKey;autoIncrement:false"`
	Applicant   Applicant `json:"-"`
	TheRemark   string    `gorm:"type:text"`
}

// TableName Remark's table name
func (*Remark) TableName() string {
	return TableNameRemark
}
