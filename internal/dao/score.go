package dao

import (
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
)

func (d *dao) QueryAllScores() (*[]models.Score, error) {
	var res []models.Score
	if err := d.orm.Table("score").Find(&res).Error; err != nil {
		log.Errorf("query all scores err(%v)", err)
		return nil, err
	}
	return &res, nil
}

func (d *dao) QueryScoreByExamCode(examCode int) (*[]models.Score, error) {
	var res []models.Score
	if err := d.orm.Table("score").Where("examCode=? ", examCode).Find(&res).Error; err != nil {
		log.Errorf("score query by examCode err(%v)", err)
		return nil, err
	}
	return &res, nil
}

// 不分页 查询 student 全部成绩
func (d *dao) QueryScoreByStudentId(studentId string) (*[]models.Score, error) {
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

// 分页查询 student score
func (d *dao) QueryStudentScoreByPage(studentId string, pageNum, pageSize int) (*models.Page, error) {
	var (
		res   []models.Score
		count int64
	)

	if err := d.orm.Table("score").Order("answerDate desc").Offset((pageNum-1)*pageSize).Limit(pageSize).Where("studentId=?", studentId).Find(&res).Error; err != nil {
		log.Errorf("query student score by page err(%v)", err)
		return nil, err
	}

	if err := d.orm.Table("score").Where("studentId=?", studentId).Count(&count).Error; err != nil {
		log.Errorf("query student score count err (%v)", err)
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
