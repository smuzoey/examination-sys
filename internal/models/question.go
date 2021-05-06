package models

// question 目前有三类：填空题、选择题、判断题

// table: multi_question 1
type SelectQuestion struct {
	QuestionId  int    `json:"question_id" gorm:"column:questionId"`
	Subject     string `json:"subject"`
	Question    string `json:"question"`
	AnswerA     string `json:"answerA" gorm:"column:answerA"`
	AnswerB     string `json:"answerB" gorm:"column:answerB"`
	AnswerC     string `json:"answerC" gorm:"column:answerC"`
	AnswerD     string `json:"answerD" gorm:"column:answerD"`
	RightAnswer string `json:"rightAnswer" gorm:"column:rightAnswer"`
	Analysis    string `json:"analysis"`
	Score       int    `json:"score"`
	Level       string `json:"level"`
	Section     string `json:"section"`
}

// table: file_question 2
type FillQuestion struct {
	QuestionId int    `json:"question_id" gorm:"column:questionId"`
	Subject    string `json:"subject"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	Analysis   string `json:"analysis"`
	Score      int    `json:"score"`
	Level      string `json:"level"`
	Section    string `json:"section"`
}

// table: judge_question 3
type JudgeQuestion struct {
	QuestionId int    `json:"question_id" gorm:"column:questionId"`
	Subject    string `json:"subject"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	Analysis   string `json:"analysis"`
	Score      int    `json:"score"`
	Level      string `json:"level"`
	Section    string `json:"section"`
}

type Question struct {
	Question string `json:"question"`
	Subject  string `json:"subject"`
	Score    string `json:"score"`
	Section  string `json:"section"`
	Type     string `json:"type"`
}
