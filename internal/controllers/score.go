package controllers

import (
	"examination-sys/internal/service"
	"examination-sys/internal/util"
	"github.com/astaxie/beego"
	"strconv"
)

type ScoreController struct {
	beego.Controller
}

func (c *ScoreController) QueryStudentScoreByPage() {
	StudentId, _ := strconv.Atoi(c.Ctx.Input.Param(":studentId"))
	PageNum, _ := strconv.Atoi(c.Ctx.Input.Param(":pageNum"))
	PageSize, _ := strconv.Atoi(c.Ctx.Input.Param(":pageSize"))

	res, err := service.QueryStudentScoreByPage(StudentId, PageNum, PageSize)
	if err != nil {
		util.Json(c.Controller, res, "err", 500)
		return
	}
	util.Json(c.Controller, res, "success", 200)
	return
}

func (c *ScoreController) QueryScoresByExamCode() {
	examCode := c.Ctx.Input.Param(":examCode")
	res, err := service.QueryScoreByExamCode(examCode)
	if err != nil {
		util.Json(c.Controller, nil, "err", 500)
		return
	}
	util.Json(c.Controller, res, "success", 200)

}
