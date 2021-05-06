package dao

import (
	"examination-sys/internal/models"
	"fmt"
	"github.com/prometheus/common/log"
)

const (
	sqlQuerySelectByPaperId = "SELECT multi_question.questionId,multi_question.`subject`,multi_question.question,multi_question.answerA,multi_question.answerB,multi_question.answerC,multi_question.answerD,multi_question.rightAnswer, multi_question.analysis,multi_question.score, multi_question.section, multi_question.`level` FROM multi_question LEFT JOIN paper_manage ON multi_question.questionId = paper_manage.questionId WHERE paper_manage.paperId = ? AND paper_manage.questionType=1"
	sqlQueryFillByPagerId   = "SELECT fill_question.questionId,fill_question.`subject`,fill_question.question,fill_question.answer,fill_question.analysis,fill_question.score,fill_question.score,fill_question.`level`,fill_question.section FROM fill_question LEFT JOIN paper_manage ON fill_question.questionId = paper_manage.questionId WHERE paper_manage.paperId = ? AND paper_manage.questionType=2"
	sqlQueryJudgeByPagerId  = "SELECT judge_question.questionId,judge_question.`subject`,judge_question.question,judge_question.answer,judge_question.analysis,judge_question.score,judge_question.`level`,judge_question.section FROM judge_question LEFT JOIN paper_manage ON judge_question.questionId = paper_manage.questionId WHERE paper_manage.paperId = ? AND paper_manage.questionType=3"

	sqlQueryQuestionsByPage = "select question, subject, score, section,level, \"选择题\" as type from multi_question union select  question, subject, score, section,level, \"判断题\" as type  from judge_question union select  question, subject, score, section,level, \"填空题\" as type from fill_question offset ? limit ?"
	sqlQueryQuestionsCount  = "select question, subject, score, section,level, \"选择题\" as type from multi_question union select  question, subject, score, section,level, \"判断题\" as type  from judge_question union select  question, subject, score, section,level, \"填空题\" as type from fill_question"
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
	if err := d.orm.Raw(sqlQueryFillByPagerId, paperId).Scan(&res).Error; err != nil {
		log.Error("raw query fill question wrong (%v)", err)
		return nil, err
	}
	return &res, nil
}

func (d *dao) QueryJudgeQuestionByPaperId(paperId int) (*[]models.JudgeQuestion, error) {
	res := []models.JudgeQuestion{}
	if err := d.orm.Raw(sqlQueryJudgeByPagerId, paperId).Scan(&res).Error; err != nil {
		log.Error("raw query judge question wrong (%v)", err)
		return nil, err
	}
	fmt.Println()
	return &res, nil
}

func (d *dao) QueryQuestionsByPage(pageNum, pageSize int) (*models.Page, error) {
	var (
		res   []models.Question
		count int64
	)
	if err := d.orm.Raw(sqlQueryQuestionsByPage, (pageNum-1)*pageSize, pageSize).Scan(&res).Error; err != nil {
		log.Errorf("raw query questions by page err(%v)", err)
		return nil, err
	}

	if err := d.orm.Raw(sqlQueryQuestionsCount).Count(&count).Error; err != nil {
		log.Errorf("raw query question count err(%v)", err)
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
