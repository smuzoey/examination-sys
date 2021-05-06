package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
)

func QueryExamByPage(pageNum, pageSize int) (*models.Page, error) {
	examPage, err := dao.DB.QueryExamByPage(pageNum, pageSize)
	if err != nil {
		log.Error("query Exam by page wrong")
		return nil, err
	}
	return examPage, err
}

func QueryExamById(id int) (*models.ExamManage, error) {
	exam, err := dao.DB.QueryExamById(id)
	if err != nil {
		log.Error("query Exam by id wrong")
		return nil, err
	}
	return exam, nil
}

func AddExam(exam *models.ExamManage) error {
	if err := dao.DB.AddExam(exam); err != nil {
		log.Error("add exam wrong")
		return err
	}
	return nil
}

func UpdateExam(exam *models.ExamManage) error {
	if err := dao.DB.UpdateExam(exam); err != nil {
		return err
	}
	return nil
}
