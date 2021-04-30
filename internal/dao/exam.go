package dao

import (
	"examination-sys/internal/models"
	"fmt"
	"github.com/prometheus/common/log"
)

func (d *dao) QueryExamById(id string) (*models.ExamManage, error) {
	res := &models.ExamManage{}
	var o = d.orm.Model(models.ExamManage{})
	err := o.Table("exam_manage").Where("examCode=?", id).Find(res).Error
	if err != nil {
		log.Error("exam_manage query exam(%v)", err)
		return nil, err
	}

	return res, nil
}

func (d *dao) AddExam(e *models.ExamManage) error {
	if err := d.orm.Model(models.ExamManage{}).Create(e).Error; err != nil {
		log.Error("exam insert exam error(%v)", err)
		return err
	}
	return nil
}

func (d *dao) QueryExamByPage(pageNum, pageSize int) (*models.Page, error) {

	var res = &[]models.ExamManage{}
	var count int64

	var o = d.orm.Model(models.ExamManage{})
	if err := o.Table("exam_manage").Order("examCode desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(res).Error; err != nil {
		log.Error("query exam_manage page err(%v)", err)
		return nil, err
	}
	if err := o.Table("exam_manage").Count(&count).Error; err != nil {
		log.Error("query exam_manage count(%v)", err)
		return nil, err
	}

	fmt.Println("<<<<<<<", count)

	pageRes := &models.Page{
		Records: new([]models.ExamManage),
		Size:    pageSize,
		Current: pageNum,
	}
	pageRes.Records = res
	pageRes.Total = count

	return pageRes, nil
}
