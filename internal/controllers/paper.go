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

type PaperController struct {
	beego.Controller
}

func (this *PaperController) FindPaperById() {
	paperId, _ := strconv.Atoi(this.Ctx.Input.Param(":paperId"))
	res, err := service.QueryPaperQuestionById(paperId)
	if err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}
	//util.Json(this.Controller, res, "success", 200)
	this.Data["json"] = res
	this.ServeJSON()
	return
}

func (this *PaperController) AddQuestionToPaper() {
	data := this.Ctx.Input.RequestBody
	res := models.Paper{}

	if err := json.Unmarshal(data, &res); err != nil {
		log.Errorf("unmarshal err: %v", err)
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}

	if err := service.AddQuestionToPaper(&res); err != nil {
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}
	util.Json(this.Controller, "", "success", 200)
	return
}
