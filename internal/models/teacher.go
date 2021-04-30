package models

type Teacher struct {
	TeacherId   string `json:"teacherId"`
	TeacherName string `json:"teacherName"`
	Institute   string `json:"institute"`
	Tel         string `json:"tel"`
	Email       string `json:"email"`
	Pwd         string `json:"pwd"`
	CardId      string `json:"cardId"`
	Type        string `json:"type"`
	Sex         string `json:"sex"`
	Role        string `json:"role"`
}
