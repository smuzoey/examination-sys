package dao

import (
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
)

func (d *dao) QueryPaperById(paperId int) (*[]models.Paper, error) {
	res := []models.Paper{}
	if err := d.orm.Table("paper_manage").Where("paperId = ?", paperId).Find(&res).Error; err != nil {
		log.Errorf("paper_manage query paper by Id (%v)", err)
		return nil, err
	}
	return &res, nil
}
