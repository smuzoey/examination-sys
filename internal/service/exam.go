package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
)

func QueryExamByPage(pageNum, pageSize int) (*models.Page, error) {

	examPage, err := dao.DB.QueryExamByPage(pageNum, pageSize)
	if err != nil {
		log.Errorf("query Exam by page wrong %v", err)
		return nil, err
	}

	return examPage, err
}
