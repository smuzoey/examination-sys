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
