package controllers

import (
	"examination-sys/internal/service"
	"examination-sys/internal/util"
	"github.com/astaxie/beego"
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
	util.Json(this.Controller, res, "success", 200)
	return
}
