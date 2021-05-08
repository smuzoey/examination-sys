package models

type Score struct {
	ExamCode   int    `json:"examCode" gorm:"column:examCode"`
	StudentId  string `json:"studentId" gorm:"column:studentId"`
	Subject    string `json:"subject"`
	PtScore    int    `json:"ptScore" gorm:"column:ptScore"`
	EtScore    int    `json:"etScore" gorm:"column:etScore"`
	Score      int    `json:"score"`
	AnswerDate string `json:"answerDate" gorm:"column:answerDate"`
}
