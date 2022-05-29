package model

const TableNameIntent = "intent"

// Intent mapped from table <intent>
type Intent struct {
	ApplicantID    uint      `gorm:"primaryKey;autoIncrement:false" json:"-"`
	Applicant      Applicant `json:"-"`
	Group          string    `gorm:"type:character varying(16);primaryKey;autoIncrement:false" json:"group"`
	IntentRank     uint      // 0 为平行志愿
	OptionalTimeID uint
	OptionalTime   OptionalTime
}

// TableName Intent's table name
func (*Intent) TableName() string {
	return TableNameIntent
}
