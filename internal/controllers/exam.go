package controllers

import (
	"encoding/json"
	"examination-sys/internal/models"
	"examination-sys/internal/service"
	"examination-sys/internal/util"
	"github.com/astaxie/beego"
	"strconv"
)

type ExamController struct {
	beego.Controller
}

func (this *ExamController) FindExamByPage() {
	pageNum, _ := strconv.Atoi(this.Ctx.Input.Param(":pageNum"))
	pageSize, _ := strconv.Atoi(this.Ctx.Input.Param(":pageSize"))

	res, err := service.QueryExamByPage(pageNum, pageSize)
	if err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}
	util.Json(this.Controller, res, "success", 200)
	return
}

func (this *ExamController) FindExamById() {
	examCode, _ := strconv.Atoi(this.Ctx.Input.Param(":examCode"))
	res, err := service.QueryExamById(examCode)
	if err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}
	util.Json(this.Controller, res, "success", 200)
}

func (this *ExamController) UpdateExam() {

	data := this.Ctx.Input.RequestBody
	exam := models.ExamManage{}

	if err := json.Unmarshal(data, &exam); err != nil {
		util.Json(this.Controller, nil, "err"+err.Error(), 500)
		return
	}
	if err := service.UpdateExam(&exam); err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}

	util.Json(this.Controller, "", "success", 200)
	return
}

// add paper
func (this *ExamController) AddExam() {
	data := this.Ctx.Input.RequestBody
	exam := models.ExamManage{}

	if err := json.Unmarshal(data, &exam); err != nil {
		util.Json(this.Controller, nil, "json unmarshal wrong:"+err.Error(), 500)
		return
	}
	if err := service.AddExam(&exam); err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}

	util.Json(this.Controller, "", "success", 200)
	return
}

func (this *ExamController) DeleteExamById() {
	examCode, _ := strconv.Atoi(this.Ctx.Input.Param(":examCode"))
	if err := service.DeleteExam(examCode); err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}
	util.Json(this.Controller, "", "success", 200)
}

func (this *ExamController) FindLastPaperId() {
	res, err := service.FindLastPaperId()
	if err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}
	util.Json(this.Controller, res, "success", 200)
}

func (this *ExamController) FindAllExams() {
	res, err := service.FindAllExams()
	if err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}
	util.Json(this.Controller, res, "success", 200)
}
