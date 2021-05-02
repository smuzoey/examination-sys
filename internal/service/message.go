package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
)

func QueryMessageByPage(pageNum, pageSize int) (*models.Page, error) {
	res, err := dao.DB.QueryMessageByPage(pageNum, pageSize)
	if err != nil {
		log.Error("query message by page err")
		return nil, err
	}
	return res, nil
}

func AddMessage(message *models.Message) error {
	if err := dao.DB.AddMessage(message); err != nil {
		log.Error("add message wrong ")
		return err
	}
	return nil
}
