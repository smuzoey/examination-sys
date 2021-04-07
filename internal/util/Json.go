package util

import (
	"github.com/astaxie/beego"
)

type ReturnData struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Json(c beego.Controller, data interface{}, message string, code int) {
	r := ReturnData{
		Code:    code,
		Data:    data,
		Message: message,
	}
	c.Data["json"] = &r
	c.ServeJSON()
}
