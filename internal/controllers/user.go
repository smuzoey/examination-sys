package controllers

import (
	"encoding/json"
	"examination-sys/internal/dao"
	"examination-sys/internal/util"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type User struct {
	Name string `json:"username"`
	Pwd  string `json:"password"`
}

func (this *UserController) Login() {
	data := this.Ctx.Input.RequestBody
	user := &User{}
	err := json.Unmarshal(data, &user)
	if err != nil {
		util.Json(this.Controller, "", "error", -500)
		return
	}
	if s, err := dao.DB.QueryStudent(user.Name); err == nil {
		util.Json(this.Controller, s, "success test", 200)
		return
	} else if t, err := dao.DB.QueryTeacher(user.Name); err == nil {
		util.Json(this.Controller, t, "success", 200)
		return
	}
	util.Json(this.Controller, nil, "err", -500)
	return
}
