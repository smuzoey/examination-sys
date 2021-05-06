package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
)

func QueryQuestionsByPage(pageNum, pageSize int) (*models.Page, error) {
	res, err := dao.DB.QueryQuestionsByPage(pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return res, nil
}
