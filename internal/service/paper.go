package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
)

func QueryPaperQuestionById(id int) (*[]models.SelectQuestion, error) {
	//paper, err := dao.DB.QueryPaperById(id)
	selectQuestion, err := dao.DB.QuerySelectQuestionByPaperId(id)
	if err != nil {
		return nil, err
	}
	return selectQuestion, nil
}
