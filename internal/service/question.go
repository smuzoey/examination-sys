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

func AddSelectQuestion(question *models.SelectQuestion) (int, error) {
	id, err := dao.DB.AddSelectQuestion(question)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func AddFillQuestion(question *models.FillQuestion) (int, error) {
	id, err := dao.DB.AddFillQuestion(question)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func AddJudgeQuestion(question *models.JudgeQuestion) (int, error) {
	id, err := dao.DB.AddJudgeQuestion(question)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func QuerySelectQuestionByQuestionId(id int) (*models.SelectQuestion, error) {
	res, err := dao.DB.QuerySelectQuestionByQuestionId(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryFillQuestionByQuestionId(id int) (*models.FillQuestion, error) {
	res, err := dao.DB.QueryFillQuestionByQuestionId(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryJudgeQuestionByQuestionId(id int) (*models.JudgeQuestion, error) {
	res, err := dao.DB.QueryJudgeQuestionByQuestionId(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryLastSelectQuestion() (*models.SelectQuestion, error) {
	res, err := dao.DB.QueryLastSelectQuestion()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryLastFillQuestion() (*models.FillQuestion, error) {
	res, err := dao.DB.QueryLastFillQuestion()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryLastJudgeQuestion() (*models.JudgeQuestion, error) {
	res, err := dao.DB.QueryLastJudgeQuestion()
	if err != nil {
		return nil, err
	}
	return res, nil
}
