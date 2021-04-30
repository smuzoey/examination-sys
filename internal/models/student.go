package models

type Student struct {
	StudentId   string `json:"studentId" gorm:"column:studentId"`
	StudentName string `json:"studentName" gorm:"column:studentName"`
	Grade       string `json:"grade"`
	Major       string `json:"major"`
	Clazz       string `json:"clazz"`
	Institute   string `json:"institute"`
	Tel         string `json:"tel"`
	Email       string `json:"email"`
	Pwd         string `json:"pwd"`
	CardId      string `json:"cardId"`
	Sex         string `json:"sex"`
	Role        string `json:"role"`
}
