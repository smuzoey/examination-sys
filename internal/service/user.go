package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
)

func QueryStudentByPage(pageNum, pageSize int) (*models.Page, error) {
	res, err := dao.DB.QueryStudentByPage(pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddStudent(s *models.Student) error {
	if err := dao.DB.AddStudent(s); err != nil {
		return err
	}
	return nil
}

func UpdateStudentSomeValues(s *models.Student) error {
	if err := dao.DB.UpdateStudentSomeValues(s); err != nil {
		return err
	}
	return nil
}

func QueryStudentById(id string) (*models.Student, error) {
	res, err := dao.DB.QueryStudent(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
