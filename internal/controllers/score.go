package controllers

import (
	"encoding/json"
	"examination-sys/internal/models"
	"examination-sys/internal/service"
	"examination-sys/internal/util"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type ScoreController struct {
	beego.Controller
}

func (c *ScoreController) QueryStudentScoreByPage() {
	StudentId := c.Ctx.Input.Param(":studentId")
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
	examCode, _ := strconv.Atoi(c.Ctx.Input.Param(":examCode"))
	res, err := service.QueryScoreByExamCode(examCode)
	if err != nil {
		util.Json(c.Controller, nil, "err", 500)
		return
	}
	util.Json(c.Controller, res, "success", 200)
}

func (c *ScoreController) QueryScoreByStudentId() {
	studentId := c.Ctx.Input.Param(":studentId")
	res, err := service.QueryScoreByStudentId(studentId)
	if err != nil {
		util.Json(c.Controller, nil, "err:"+err.Error(), 500)
		return
	}
	util.Json(c.Controller, res, "success", 200)
	return
}

func (c *ScoreController) AddScore() {
	data := c.Ctx.Input.RequestBody
	score := models.Score{}

	fmt.Println("<<<<<<", score)
	if err := json.Unmarshal(data, &score); err != nil {
		fmt.Println("<<<<<<<<<")
		util.Json(c.Controller, nil, "err:"+err.Error(), 500)
		return
	}
	fmt.Println("<<<<<<", score)
	if err := service.AddScore(&score); err != nil {
		util.Json(c.Controller, nil, "err:"+err.Error(), 500)
		return
	}
	util.Json(c.Controller, "", "success", 200)
	return
}
