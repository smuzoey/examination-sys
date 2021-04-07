package models

import (
	"fmt"
	"github.com/prometheus/common/log"
)

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

func (d *dao) QueryStudent(id string) (*Student, error) {
	s := &Student{}
	var o = d.orm.Model(Student{})
	err := o.Table("student").Where("studentId = ?", id).Find(s).Error
	fmt.Println(s)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return s, nil
}

func (d *dao) CreateStudent() {}
