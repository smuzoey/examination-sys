package controllers

import (
	"encoding/json"
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
	"examination-sys/internal/service"
	"examination-sys/internal/util"
	"github.com/astaxie/beego"
	"strconv"
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
	if err := json.Unmarshal(data, &user); err != nil {
		util.Json(this.Controller, "", "error:"+err.Error(), 500)
		return
	}
	if s, err := dao.DB.QueryStudent(user.Name); err == nil && s.StudentId != "" {
		util.Json(this.Controller, s, "success test", 200)
		return
	} else if t, err := dao.DB.QueryTeacher(user.Name); err == nil && t.TeacherId != "" {
		util.Json(this.Controller, t, "success", 200)
		return
	}
	util.Json(this.Controller, "无该用户", "err", 500)
	return
}

func (this *UserController) AddStudent() {
	data := this.Ctx.Input.RequestBody
	student := models.Student{}
	if err := json.Unmarshal(data, &student); err != nil {
		util.Json(this.Controller, nil, "err: "+err.Error(), 500)
		return
	}
	if err := service.AddStudent(&student); err != nil {
		util.Json(this.Controller, nil, "err: "+err.Error(), 500)
		return
	}
	util.Json(this.Controller, student, "success", 200)
	return
}

func (this *UserController) QueryStudentByPage() {
	pageNum, _ := strconv.Atoi(this.Ctx.Input.Param(":pageNum"))
	pageSize, _ := strconv.Atoi(this.Ctx.Input.Param(":pageSize"))
	res, err := service.QueryStudentByPage(pageNum, pageSize)
	if err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}
	util.Json(this.Controller, res, "success", 200)
	return
}

func (this *UserController) UpdateStudentSomeValues() {
	data := this.Ctx.Input.RequestBody
	student := models.Student{}
	if err := json.Unmarshal(data, &student); err != nil {
		util.Json(this.Controller, nil, "err", 500)
		return
	}
	if err := service.UpdateStudentSomeValues(&student); err != nil {
		util.Json(this.Controller, nil, "err: "+err.Error(), 500)
		return
	}
	util.Json(this.Controller, "", "success", 200)
	return
}

func (this *UserController) QueryStudentById() {
	StudentId := this.Ctx.Input.Param(":studentId")
	res, err := service.QueryStudentById(StudentId)
	if err != nil {
		util.Json(this.Controller, nil, "err: "+err.Error(), 500)
		return
	}
	util.Json(this.Controller, res, "success", 200)
	return
}
