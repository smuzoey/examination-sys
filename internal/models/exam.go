package models

import "github.com/prometheus/common/log"

type ExamManage struct {
	ExamCode    int    `json:"examCode"`
	Description string `json:"description"`
	Source      string `json:"source"`
	PaperId     int    `json:"paperId"`
	ExamDate    string `json:"examDate"`
	TotalTime   int    `json:"totalTime"`
	Grade       string `json:"grade"`
	Term        string `json:"term"`
	Major       string `json:"major"`
	Institute   string `json:"institute"`
	TotalSource int    `json:"totalSource"`
	Type        string `json:"type"`
	Tips        string `json:"tips"`
}

func (d *dao) QueryExamById(id string) (*ExamManage, error) {
	res := &ExamManage{}
	var o = d.orm.Model(ExamManage{})
	err := o.Table("exam_manage").Where("examCode=?", id).Find(res).Error
	if err != nil {
		log.Error("exam_manage query exam(%v)", err)
		return nil, err
	}
	return res, nil
}

func (d *dao) AddExam(e *ExamManage) error {
	if err := d.orm.Model(ExamManage{}).Create(e).Error; err != nil {
		log.Error("exam insert exam error(%v)", err)
		return err
	}
	return nil
}

func (d *dao) QueryExamByPage(pageNum, pageSize int) (*[]ExamManage, error) {
	var res = &[]ExamManage{}
	var o = d.orm.Model(ExamManage{})
	if err := o.Table("exam_manage").Order("examCode desc").Offset(pageNum).Limit(pageSize).Find(res).Error; err != nil {
		log.Error("query exam_manage err(%v)", err)
		return nil, err
	}
	return res, nil
}
