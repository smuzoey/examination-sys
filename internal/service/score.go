package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
)

func QueryScoreByExamCode(examCode string) (*[]models.Score, error) {
	res, err := dao.DB.QueryScoreByExamCode(examCode)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryScoreByStudentId(studentId int) (*[]models.Score, error) {
	res, err := dao.DB.QueryScoreByStudentId(studentId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryStudentScoreByPage(studentId, pageNum, pageSize int) (*models.Page, error) {
	res, err := dao.DB.QueryStudentScoreByPage(studentId, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return res, nil
}
