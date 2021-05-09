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

func (d *dao) AddPaper(p *models.Paper) error {
	if err := d.orm.Table("paper_manage").Create(p).Error; err != nil {
		log.Errorf("add paper error(%v)", err)
		return err
	}
	return nil
}

func (d *dao) AddBatchPaper(p *[]models.Paper) error {
	if err := d.orm.Table("paper_manage").CreateInBatches(p, len(*p)).Error; err != nil {
		log.Errorf("add batch paper error(%v)", err)
		return err
	}
	return nil
}
