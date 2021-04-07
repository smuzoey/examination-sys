package models

import "github.com/prometheus/common/log"

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

func (d *dao) QueryTeacher(id string) (*Teacher, error) {
	t := &Teacher{}
	var o = d.orm.Model(Teacher{})
	err := o.Table("teacher").Where("teacherId = ?", id).Find(t).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return t, nil
}
