package controllers

import (
	"examination-sys/internal/service"
	"examination-sys/internal/util"
	"github.com/astaxie/beego"
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
