package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
	"sync"
)

func QueryPaperQuestionById(id int) (map[int]interface{}, error) {

	mp := make(map[int]interface{}, 0)
	//res := make([]interface{}, 3)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func(id int) {
		mp[1], _ = dao.DB.QuerySelectQuestionByPaperId(id)
		wg.Done()
	}(id)
	go func(id int) {
		mp[2], _ = dao.DB.QueryFillQuestionByPaperId(id)
		wg.Done()
	}(id)
	go func(id int) {
		mp[3], _ = dao.DB.QueryJudgeQuestionByPaperId(id)
		wg.Done()
	}(id)

	wg.Wait()

	return mp, nil
}

func AddQuestionToPaper(paper *models.Paper) error {
	if err := dao.DB.AddPaper(paper); err != nil {
		return err
	}
	return nil
}
