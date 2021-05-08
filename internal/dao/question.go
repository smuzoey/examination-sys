package dao

import (
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
)

const (
	sqlQuerySelectByPaperId = "SELECT multi_question.questionId,multi_question.`subject`,multi_question.question,multi_question.answerA,multi_question.answerB,multi_question.answerC,multi_question.answerD,multi_question.rightAnswer, multi_question.analysis,multi_question.score, multi_question.section, multi_question.`level` FROM multi_question LEFT JOIN paper_manage ON multi_question.questionId = paper_manage.questionId WHERE paper_manage.paperId = ? AND paper_manage.questionType=1"
	sqlQueryFillByPagerId   = "SELECT fill_question.questionId,fill_question.`subject`,fill_question.question,fill_question.answer,fill_question.analysis,fill_question.score,fill_question.score,fill_question.`level`,fill_question.section FROM fill_question LEFT JOIN paper_manage ON fill_question.questionId = paper_manage.questionId WHERE paper_manage.paperId = ? AND paper_manage.questionType=2"
	sqlQueryJudgeByPagerId  = "SELECT judge_question.questionId,judge_question.`subject`,judge_question.question,judge_question.answer,judge_question.analysis,judge_question.score,judge_question.`level`,judge_question.section FROM judge_question LEFT JOIN paper_manage ON judge_question.questionId = paper_manage.questionId WHERE paper_manage.paperId = ? AND paper_manage.questionType=3"

	sqlQuerySelectByQuestionId = "SELECT * FROM multi_question"
	sqlQueryFillByQuestionId   = "SELECT * FROM fill_question"
	sqlQueryJudgeByQuestionId  = "SELECT * FROM judge_question"

	sqlQueryQuestionsByPage = "select question, subject, score, section,level, \"选择题\" as type from multi_question union select  question, subject, score, section,level, \"判断题\" as type  from judge_question union select  question, subject, score, section,level, \"填空题\" as type from fill_question"
	sqlQueryQuestionsCount  = "select sum(x) from (select count(*) as x from multi_question union select count(*) as x from judge_question union select count(*) as x from fill_question) as total"
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
	return &res, nil
}

func (d *dao) QueryQuestionsByPage(pageNum, pageSize int) (*models.Page, error) {
	var (
		res   []models.Question
		count int64
	)
	if err := d.orm.Raw(sqlQueryQuestionsByPage).Scan(&res).Error; err != nil {
		log.Errorf("raw query questions by page err(%v)", err)
		return nil, err
	}
	ans := res[(pageNum-1)*pageSize : pageNum*pageSize]

	if err := d.orm.Raw(sqlQueryQuestionsCount).Scan(&count).Error; err != nil {
		log.Errorf("raw query question count err(%v)", err)
		return nil, err
	}

	pageRes := models.Page{
		Records: ans,
		Total:   count,
		Size:    pageSize,
		Current: pageNum,
	}
	return &pageRes, nil
}

func (d *dao) AddSelectQuestion(question *models.SelectQuestion) (int, error) {
	if err := d.orm.Table("multi_question").Create(question).Error; err != nil {
		log.Errorf("add select question wrong(%v)", err)
		return 0, err
	}
	var questionId int
	if err := d.orm.Raw("select questionId from multi_question order by questionId desc limit 1").Scan(&questionId).Error; err != nil {
		log.Errorf("get questionId wrong(%v)", err)
		return 0, err
	}
	return questionId, nil
}

func (d *dao) AddFillQuestion(question *models.FillQuestion) (int, error) {
	if err := d.orm.Table("fill_question").Create(question).Error; err != nil {
		log.Errorf("add fill question error(%v)", err)
		return 0, err
	}
	var questionId int
	if err := d.orm.Raw("select questionId from fill_question order by questionId desc limit 1").Scan(&questionId).Error; err != nil {
		log.Errorf("get questionId wrong(%v)", err)
		return 0, err
	}
	return questionId, nil
}

func (d *dao) AddJudgeQuestion(question *models.JudgeQuestion) (int, error) {
	if err := d.orm.Table("judge_question").Create(question).Error; err != nil {
		log.Errorf("add judge question error(%v)", err)
		return 0, err
	}
	var questionId int
	if err := d.orm.Raw("select questionId from judge_question order by questionId desc limit 1").Scan(&questionId).Error; err != nil {
		log.Errorf("get questionId wrong(%v)", err)
		return 0, err
	}
	return questionId, nil
}

func (d *dao) QuerySelectQuestionByQuestionId(questionId int) (*models.SelectQuestion, error) {
	res := models.SelectQuestion{}
	if err := d.orm.Raw(sqlQuerySelectByQuestionId, questionId).Scan(&res).Error; err != nil {
		log.Error("raw query select question wrong (%v)", err)
		return nil, err
	}
	return &res, nil
}

func (d *dao) QueryFillQuestionByQuestionId(questionId int) (*models.FillQuestion, error) {
	res := models.FillQuestion{}
	if err := d.orm.Raw(sqlQueryFillByQuestionId, questionId).Scan(&res).Error; err != nil {
		log.Error("raw query fill question wrong (%v)", err)
		return nil, err
	}
	return &res, nil
}

func (d *dao) QueryJudgeQuestionByQuestionId(questionId int) (*models.JudgeQuestion, error) {
	res := models.JudgeQuestion{}
	if err := d.orm.Raw(sqlQueryJudgeByQuestionId, questionId).Scan(&res).Error; err != nil {
		log.Error("raw query judge question wrong (%v)", err)
		return nil, err
	}
	return &res, nil
}

func (d *dao) QueryLastJudgeQuestion() (*models.JudgeQuestion, error) {
	res := models.JudgeQuestion{}
	if err := d.orm.Table("judge_question").Order("questionId desc").Limit(1).Find(&res).Error; err != nil {
		log.Errorf("query last judge question(%v)", err)
		return nil, err
	}
	return &res, nil
}

func (d *dao) QueryLastSelectQuestion() (*models.SelectQuestion, error) {
	res := models.SelectQuestion{}
	if err := d.orm.Table("multi_question").Order("questionId desc").Limit(1).Find(&res).Error; err != nil {
		log.Errorf("query last select question(%v)", err)
		return nil, err
	}
	return &res, nil
}

func (d *dao) QueryLastFillQuestion() (*models.FillQuestion, error) {
	res := models.FillQuestion{}
	if err := d.orm.Table("fill_question").Order("questionId desc").Limit(1).Find(&res).Error; err != nil {
		log.Errorf("query last fill question(%v)", err)
		return nil, err
	}
	return &res, nil
}
