package dao

import (
	"examination-sys/internal/models"
	"fmt"
	"github.com/prometheus/common/log"
)

func (d *dao) QueryStudent(id string) (*models.Student, error) {
	s := &models.Student{}
	var o = d.orm.Model(models.Student{})
	err := o.Table("student").Where("studentId = ?", id).Find(s).Error
	fmt.Println(s)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return s, nil
}

func (d *dao) CreateStudent() {}
