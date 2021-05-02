package dao

import (
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
)

func (d *dao) QueryMessageByPage(pageNum, pageSize int) (*models.Page, error) {

	var (
		res   []models.Message
		count int64
	)

	if err := d.orm.Table("message").Order("time desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&res).Error; err != nil {
		log.Error("query messagee page err (%v)", err)
		return nil, err
	}
	if err := d.orm.Table("message").Count(&count).Error; err != nil {
		log.Error("query message count(%v)", err)
		return nil, err
	}

	pageRes := models.Page{
		Records: res,
		Total:   count,
		Size:    pageSize,
		Current: pageNum,
	}
	return &pageRes, nil
}

func (d *dao) AddMessage(m *models.Message) error {
	if err := d.orm.Table("message").Create(m).Error; err != nil {
		log.Error("message insert message error(%v)", err)
		return err
	}
	return nil
}
