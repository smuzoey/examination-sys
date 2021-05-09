package controllers

import (
	"encoding/json"
	"examination-sys/internal/models"
	"examination-sys/internal/service"
	"examination-sys/internal/util"
	"fmt"
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

func (this *PaperController) GroupPaper() {
	data := this.Ctx.Input.RequestBody
	param := new(struct {
		Total        int      `json:"total"`
		Point        []string `json:"point"`
		ChangeNumber int      `json:"changeNumber"`
		FillNumber   int      `json:"fillNumber"`
		JudgeNumber  int      `json:"judgeNumber"`
		PaperId      int      `json:"paperId"`
		Difficult    float64  `json:"difficult"`
	})

	if err := json.Unmarshal(data, param); err != nil {
		log.Errorf("unmarshal err: %v", err)
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}

	fmt.Println(param)

	res := models.PaperLimit{
		Id:         param.PaperId,
		TotalScore: param.Total,
		Difficulty: param.Difficult,
		Point:      param.Point,
	}
	res.EachTypeCount[0], res.EachTypeCount[1], res.EachTypeCount[2] = param.ChangeNumber, param.JudgeNumber, param.FillNumber

	p := service.Inherit(&res)
	var answer []models.Paper
	for _, v := range (*p)[0].ProblemList {
		temp := models.Paper{
			PaperId:      res.Id,
			QuestionId:   v.QuestionId,
			QuestionType: v.Type,
		}
		answer = append(answer, temp)
	}

	if err := service.AddBatchQuetionToPaper(&answer); err != nil {
		util.Json(this.Controller, nil, "err:"+err.Error(), 500)
		return
	}

	util.Json(this.Controller, "", "success", 200)
	return
}

func (this *PaperController) Test() {
	fmt.Println("<<<<<<<<<<")
	util.Json(this.Controller, nil, "success", 200)
	return
}
