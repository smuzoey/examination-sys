package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
)

func QueryScoreByExamCode(examCode int) (*[]models.Score, error) {
	res, err := dao.DB.QueryScoreByExamCode(examCode)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryScoreByStudentId(studentId string) (*[]models.Score, error) {
	res, err := dao.DB.QueryScoreByStudentId(studentId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryStudentScoreByPage(studentId string, pageNum, pageSize int) (*models.Page, error) {
	res, err := dao.DB.QueryStudentScoreByPage(studentId, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddScore(score *models.Score) error {
	if err := dao.DB.AddScore(score); err != nil {
		return err
	}
	return nil
}
