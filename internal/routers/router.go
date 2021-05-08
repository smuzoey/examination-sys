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
	beego.Router("student", &controllers.UserController{}, "put:UpdateStudentSomeValues")
	beego.Router("student/:studentId", &controllers.UserController{}, "get:QueryStudentById")

	beego.Router("exams/:pageNum/:pageSize", &controllers.ExamController{}, "get:FindExamByPage")
	beego.Router("exam/:examCode", &controllers.ExamController{}, "get:FindExamById")
	beego.Router("exam/:examCode", &controllers.ExamController{}, "delete:DeleteExamById")
	beego.Router("exam", &controllers.ExamController{}, "post:AddExam")
	beego.Router("exam", &controllers.ExamController{}, "put:UpdateExam")
	beego.Router("examManagePaperId", &controllers.ExamController{}, "get:FindLastPaperId")

	beego.Router("paper/:paperId", &controllers.PaperController{}, "get:FindPaperById")
	beego.Router("paperManage", &controllers.PaperController{}, "post:AddQuestionToPaper")

	beego.Router("messages/:pageNum/:pageSize", &controllers.MessageController{}, "get:FindMessageByPage")
	beego.Router("message", &controllers.MessageController{}, "post:AddMessage")

	beego.Router("score/:pageNum/:pageSize/:studentId", &controllers.ScoreController{}, "get:QueryStudentScoreByPage")
	beego.Router("score/:studentId", &controllers.ScoreController{}, "get:QueryScoreByStudentId")
	beego.Router("scores/:examCode", &controllers.ScoreController{}, "get:QueryScoresByExamCode")
	beego.Router("score", &controllers.ScoreController{}, "post:AddScore")

	beego.Router("answers/:pageNum/:pageSize", &controllers.QuestionController{}, "get:QueryQuestionsByPage")
	beego.Router("fillQuestion", &controllers.QuestionController{}, "post:AddFillQuestion")
	beego.Router("MultiQuestion", &controllers.QuestionController{}, "post:AddSelectQuestion")
	beego.Router("judgeQuestion", &controllers.QuestionController{}, "post:AddJudgeQuestion")
	beego.Router("judgeQuestionId", &controllers.QuestionController{}, "get:QueryLastJudgeQuestion")
	beego.Router("multiQuestionId", &controllers.QuestionController{}, "get:QueryLastSelectQuestion")
	beego.Router("fillQuestionId", &controllers.QuestionController{}, "get:QueryLastFillQuestion")
}

// /api/exam
// /api/student 添加学生
// /api/student   只有改变的title 》》》》 put

// api/examManagePaperId
