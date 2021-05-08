package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
)

func QueryQuestionsByPage(pageNum, pageSize int) (*models.Page, error) {
	res, err := dao.DB.QueryQuestionsByPage(pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddSelectQuestion(question *models.SelectQuestion) error {
	if err := dao.DB.AddSelectQuestion(question); err != nil {
		return err
	}
	return nil
}

func AddFillQuestion(question *models.FillQuestion) error {
	if err := dao.DB.AddFillQuestion(question); err != nil {
		return err
	}
	return nil
}

func AddJudgeQuestion(question *models.JudgeQuestion) error {
	if err := dao.DB.AddJudgeQuestion(question); err != nil {
		return err
	}
	return nil
}
