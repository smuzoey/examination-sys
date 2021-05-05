package dao

import (
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
)

func (d *dao) QueryScoreByExamCode(examCode string) (*[]models.Score, error) {
	var res []models.Score
	if err := d.orm.Table("score").Where("examCode=? ", examCode).Find(&res).Error; err != nil {
		log.Error("score query by examCode err(%v)", err)
		return nil, err
	}
	return &res, nil
}

func (d *dao) QueryScoreByStudentId(studentId int) (*[]models.Score, error) {
	var res []models.Score
	if err := d.orm.Table("score").Where("studentId=?", studentId).Find(&res).Error; err != nil {
		log.Error("score query by studentId err(%v)", err)
		return nil, err
	}
	return &res, nil
}

func (d *dao) AddScore(score *models.Score) error {
	if err := d.orm.Table("score").Create(score).Error; err != nil {
		log.Error("score add err (%v)", err)
		return err
	}
	return nil
}
