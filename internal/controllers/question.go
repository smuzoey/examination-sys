package controllers

import (
	"encoding/json"
	"examination-sys/internal/models"
	"examination-sys/internal/service"
	"examination-sys/internal/util"
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
	"strconv"
)

type QuestionController struct {
	beego.Controller
}

func (this *QuestionController) QueryQuestionsByPage() {
	pageNum, _ := strconv.Atoi(this.Ctx.Input.Param(":pageNum"))
	pageSize, _ := strconv.Atoi(this.Ctx.Input.Param(":pageSize"))
	res, err := service.QueryQuestionsByPage(pageNum, pageSize)
	if err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}
	util.Json(this.Controller, res, "success", 200)
	return
}

func (this *QuestionController) AddSelectQuestion() {
	data := this.Ctx.Input.RequestBody
	res := models.SelectQuestion{}

	if err := json.Unmarshal(data, &res); err != nil {
		log.Errorf("unmarshal err: %v", err)
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}

	id, err := service.AddSelectQuestion(&res)
	if err != nil {
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}
	util.Json(this.Controller, id, "success", 200)
	return
}

func (this *QuestionController) AddFillQuestion() {
	data := this.Ctx.Input.RequestBody
	res := models.FillQuestion{}

	if err := json.Unmarshal(data, &res); err != nil {
		log.Errorf("unmarshal err: %v", err)
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}

	id, err := service.AddFillQuestion(&res)
	if err != nil {
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}
	util.Json(this.Controller, id, "success", 200)
	return
}

func (this *QuestionController) AddJudgeQuestion() {
	data := this.Ctx.Input.RequestBody
	res := models.JudgeQuestion{}

	if err := json.Unmarshal(data, &res); err != nil {
		log.Errorf("unmarshal err: %v", err)
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}

	id, err := service.AddJudgeQuestion(&res)
	if err != nil {
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}
	util.Json(this.Controller, id, "success", 200)
	return
}

func (this *QuestionController) QueryLastJudgeQuestion() {
	res, err := service.QueryLastJudgeQuestion()
	if err != nil {
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}
	util.Json(this.Controller, res, "success", 200)
	return
}

func (this *QuestionController) QueryLastFillQuestion() {
	res, err := service.QueryLastFillQuestion()
	if err != nil {
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}
	util.Json(this.Controller, res, "success", 200)
	return
}

func (this *QuestionController) QueryLastSelectQuestion() {
	res, err := service.QueryLastSelectQuestion()
	if err != nil {
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}
	util.Json(this.Controller, res, "success", 200)
	return
}
