package models

type ExamManage struct {
	ExamCode    int    `json:"examCode" gorm:"column:examCode"`
	Description string `json:"description"`
	Source      string `json:"source"`
	PaperId     int    `json:"paperId" gorm:"column:paperId"`
	ExamDate    string `json:"examDate" gorm:"column:examDate"`
	TotalTime   int    `json:"totalTime" gorm:"column:totalTime"`
	Grade       string `json:"grade"`
	Term        string `json:"term"`
	Major       string `json:"major"`
	Institute   string `json:"institute"`
	TotalScore  int    `json:"totalScore" gorm:"column:totalScore"`
	Type        string `json:"type"`
	Tips        string `json:"tips"`
}
