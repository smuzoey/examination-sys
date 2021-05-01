package dao

import (
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
)

func (d *dao) QueryExamById(id int) (*models.ExamManage, error) {
	res := models.ExamManage{}
	if err := d.orm.Table("exam_manage").Where("examCode= ?", id).Find(&res).Error; err != nil {
		log.Error("exam_manage query exam(%v)", err)
		return nil, err
	}

	return &res, nil
}

func (d *dao) AddExam(e *models.ExamManage) error {
	//if err := d.orm.Model(models.ExamManage{}).Create(e).Error; err != nil {
	if err := d.orm.Table("exam_manage").Create(e).Error; err != nil {
		log.Error("exam insert exam error(%v)", err)
		return err
	}
	return nil
}

func (d *dao) QueryExamByPage(pageNum, pageSize int) (*models.Page, error) {

	var (
		res   []models.ExamManage
		count int64
	)

	if err := d.orm.Table("exam_manage").Order("examCode desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&res).Error; err != nil {
		log.Error("query exam_manage page err(%v)", err)
		return nil, err
	}

	if err := d.orm.Table("exam_manage").Count(&count).Error; err != nil {
		log.Error("query exam_manage count(%v)", err)
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
