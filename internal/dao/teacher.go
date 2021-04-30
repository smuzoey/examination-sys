package dao

import (
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
)

func (d *dao) QueryTeacher(id string) (*models.Teacher, error) {
	t := &models.Teacher{}
	var o = d.orm.Model(models.Teacher{})
	err := o.Table("teacher").Where("teacherId = ?", id).Find(t).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return t, nil
}
