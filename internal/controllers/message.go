package controllers

import (
	"examination-sys/internal/service"
	"examination-sys/internal/util"
	"strconv"
)

func (this *ExamController) FindMessageByPage() {
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
