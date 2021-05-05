package models

type Score struct {
	ExamCode   string `json:"examCode" gorm:"column:examCode"`
	StudentId  int    `json:"studentId" gorm:"column:studentId"`
	Subject    string `json:"subject"`
	PtScore    int    `json:"ptScore" gorm:"column:ptScore"`
	EtScore    int    `json:"etScore" grom:"column:etScore"`
	Score      int    `json:"score"`
	AnswerDate string `json:"answerDate" gorm:"column:answerDate"`
}
