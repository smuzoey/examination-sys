package dao

import (
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
)

const (
	sqlQuerySelectByPaperId = "SELECT multi_question.questionId,multi_question.`subject`,multi_question.question,multi_question.answerA,multi_question.answerB,multi_question.answerC,multi_question.answerD,multi_question.rightAnswer, multi_question.analysis,multi_question.score, multi_question.section, multi_question.`level` FROM multi_question LEFT JOIN paper_manage ON multi_question.questionId = paper_manage.questionId WHERE paper_manage.paperId = ?"
	sqlQueryFillByPageId    = "SELECT fill_question.questionId,fill_question.`subject`,fill_question.question,fill_question.answer,fill_question.analysis,fill_question.score,fill_question.score,fill_question.`level`,fill_question.section FROM fill_question LEFT JOIN paper_manage ON fill_question.questionId = paper_manage.questionId WHERE paper_manage.paperId = ?"
)

func (d *dao) QuerySelectQuestionByPaperId(paperId int) (*[]models.SelectQuestion, error) {
	res := []models.SelectQuestion{}
	if err := d.orm.Raw(sqlQuerySelectByPaperId, paperId).Scan(&res).Error; err != nil {
		log.Error("raw query select question wrong (%v)", err)
		return nil, err
	}
	return &res, nil
}

func (d *dao) QueryFillQuestionByPaperId(paperId int) (*[]models.FillQuestion, error) {
	res := []models.FillQuestion{}
	if err := d.orm.Raw(sqlQueryFillByPageId, paperId).Scan(&res).Error; err != nil {
		log.Error("raw query fill question wrong (%v)", err)
		return nil, err
	}
	return &res, nil
}
