package dao

import (
	"examination-sys/internal/models"
	"fmt"
	"github.com/prometheus/common/log"
)

func (d *dao) QueryStudent(id string) (*models.Student, error) {
	s := &models.Student{}
	var o = d.orm.Model(models.Student{})
	err := o.Table("student").Where("studentId = ?", id).Find(s).Error
	fmt.Println(s)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return s, nil
}

func (d *dao) QueryStudentByPage(pageNum, pageSize int) (*models.Page, error) {
	var (
		res   []models.Student
		count int64
	)
	if err := d.orm.Table("student").Order("studentId desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&res).Error; err != nil {
		log.Errorf("query student by page err(%v)", err)
		return nil, err
	}
	if err := d.orm.Table("student").Count(&count).Error; err != nil {
		log.Errorf("query count(student) err(%v)", err)
		return nil, err
	}

	resPage := models.Page{
		Records: res,
		Total:   count,
		Size:    pageSize,
		Current: pageNum,
	}
	return &resPage, nil
}

func (d *dao) AddStudent(student *models.Student) error {
	if err := d.orm.Table("student").Create(student).Error; err != nil {
		log.Errorf("add studnet err(%v)", err)
		return err
	}
	return nil
}

func (d *dao) DeleteStudent(studentId int) error {
	if err := d.orm.Table("student").Where("studentId = ?", studentId).Delete(&models.Student{}).Error; err != nil {
		log.Errorf("delete student err(%v)", err)
		return err
	}
	return nil
}

func (d *dao) StudentChangePwd(studentId int, pwd string) error {
	if err := d.orm.Table("student").Where("studentId=?", studentId).Update("pwd", pwd).Error; err != nil {
		log.Errorf("student change pwd err(%v)", err)
		return err
	}
	return nil
}

func (d *dao) UpdateStudent(student *models.Student) error {
	if err := d.orm.Table("student").Where("studentId=?", student.StudentId).Updates(student).Error; err != nil {
		log.Errorf("student update information err(%v)", err)
		return err
	}
	return nil
}
