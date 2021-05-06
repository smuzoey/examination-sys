package routers

import (
	"examination-sys/internal/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("login", &controllers.UserController{}, "post:Login")
	beego.Router("students/:pageNum/:pageSize", &controllers.UserController{}, "get:QueryStudentByPage")
	beego.Router("student", &controllers.UserController{}, "post:AddStudent")

	beego.Router("exams/:pageNum/:pageSize", &controllers.ExamController{}, "get:FindExamByPage")
	beego.Router("exam/:examCode", &controllers.ExamController{}, "get:FindExamById")
	beego.Router("exam", &controllers.ExamController{}, "post:UpdateExam")

	beego.Router("paper/:paperId", &controllers.PaperController{}, "get:FindPaperById")

	beego.Router("messages/:pageNum/:pageSize", &controllers.MessageController{}, "get:FindMessageByPage")
	beego.Router("message", &controllers.MessageController{}, "post:AddMessage")

	beego.Router("score/:pageNum/:pageSize/:studentId", &controllers.ScoreController{}, "get:QueryStudentScoreByPage")
	beego.Router("score/:studentId", &controllers.ScoreController{}, "get:QueryScoreByStudentId")
	beego.Router("scores/:examCode", &controllers.ScoreController{}, "get:QueryScoresByExamCode")

	beego.Router("answers/:pageNum/:pageSize", &controllers.QuestionController{}, "get:QueryQuestionsByPage")
}

// /api/exam
// /api/student 添加学生
// /api/student   只有改变的title 》》》》 put
