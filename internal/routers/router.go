package routers

import (
	"examination-sys/internal/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("login", &controllers.UserController{}, "post:Login")
	beego.Router("exams/:pageNum/:pageSize", &controllers.ExamController{}, "get:FindExamByPage")
	beego.Router("exam/:examCode", &controllers.ExamController{}, "get:FindExamById")
	beego.Router("paper/:paperId", &controllers.PaperController{}, "get:FindPaperById")
}

/*
http://localhost:8088/api/score/1/10/20154084
http://localhost:8088/api/messages/3/8 message分页
http://localhost:8088/api/message post
content: "testteest"
time: "2021-05-02T16:38:36.059Z"
title: "test"
*/
