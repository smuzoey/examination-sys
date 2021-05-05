package controllers

import (
	"encoding/json"
	"examination-sys/internal/models"
	"examination-sys/internal/service"
	"examination-sys/internal/util"
	"github.com/astaxie/beego"
	"strconv"
)

type MessageController struct {
	beego.Controller
}

func (this *MessageController) FindMessageByPage() {
	pageNum, _ := strconv.Atoi(this.Ctx.Input.Param(":pageNum"))
	pageSize, _ := strconv.Atoi(this.Ctx.Input.Param(":pageSize"))

	res, err := service.QueryMessageByPage(pageNum, pageSize)
	if err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}
	util.Json(this.Controller, res, "success", 200)
	return
}

func (this *MessageController) AddMessage() {
	data := this.Ctx.Input.RequestBody
	message := models.Message{}
	if err := json.Unmarshal(data, &message); err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}

	if err := service.AddMessage(&message); err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}
	util.Json(this.Controller, "", "success", 200)
	return
}
