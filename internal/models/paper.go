package models

type Paper struct {
	PaperId      int `json:"paperId" gorm:"column:paperId"`
	QuestionId   int `json:"questionId" gorm:"column:questionId"`
	QuestionType int `json:"questionType" gorm:"column:questionType"`
}
