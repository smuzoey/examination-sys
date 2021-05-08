package models

type Teacher struct {
	TeacherId   string `json:"teacherId" gorm:"column:teacherId"`
	TeacherName string `json:"teacherName" gorm:"column:teacherName"`
	Institute   string `json:"institute"`
	Tel         string `json:"tel"`
	Email       string `json:"email"`
	Pwd         string `json:"pwd"`
	CardId      string `json:"cardId" gorm:"column:cardId"`
	Type        string `json:"type"`
	Sex         string `json:"sex"`
	Role        int    `json:"role"`
}
