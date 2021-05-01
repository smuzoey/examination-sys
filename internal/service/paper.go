package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
)

func QueryPaperById(id int) (*[]models.Paper, error) {
	paper, err := dao.DB.QueryPaperById(id)
	if err != nil {
		return nil, err
	}
	return paper, nil
}
