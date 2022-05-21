package model

const TableNameAnswer = "answer"

// Answer mapped from table <answer>
type Answer struct {
	ApplicantID uint `gorm:"primaryKey"`
	Applicant   Applicant
	QuestionID  uint `gorm:"primaryKey"`
	Question    Question
	TheAnswer   bool
}

// TableName Answer's table name
func (*Answer) TableName() string {
	return TableNameAnswer
}
