package dao

import (
	"examination-sys/internal/models"
	"github.com/prometheus/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB = New()

type Dao interface {
	Ping() error
	Close()

	// user
	QueryStudent(id string) (*models.Student, error)
	QueryStudentByPage(pageNum, pageSize int) (*models.Page, error)
	AddStudent(student *models.Student) error
	DeleteStudent(studentId string) error
	StudentChangePwd(studentId string, pwd string) error
	UpdateStudent(student *models.Student) error
	UpdateStudentSomeValues(s *models.Student) error

	QueryTeacher(id string) (*models.Teacher, error)

	// exam
	QueryExamByPage(pageNum, pageSize int) (*models.Page, error)
	QueryExamById(id int) (*models.ExamManage, error)
	AddExam(e *models.ExamManage) error
	DeleteExam(examCode int) error
	UpdateExam(exam *models.ExamManage) error
	FindLastPaperId() (*models.ExamManage, error)
	FindAllExams() (*[]models.ExamManage, error)

	// paper
	QueryPaperById(paperId int) (*[]models.Paper, error)
	AddPaper(p *models.Paper) error
	AddBatchPaper(p *[]models.Paper) error

	// question
	QuerySelectQuestionByPaperId(paperId int) (*[]models.SelectQuestion, error)
	QueryFillQuestionByPaperId(paperId int) (*[]models.FillQuestion, error)
	QueryJudgeQuestionByPaperId(paperId int) (*[]models.JudgeQuestion, error)
	QueryQuestionsByPage(pageNum, pageSize int) (*models.Page, error)
	AddSelectQuestion(question *models.SelectQuestion) (int, error)
	AddFillQuestion(question *models.FillQuestion) (int, error)
	AddJudgeQuestion(question *models.JudgeQuestion) (int, error)
	QuerySelectQuestionByQuestionId(questionId int) (*models.SelectQuestion, error)
	QueryFillQuestionByQuestionId(questionId int) (*models.FillQuestion, error)
	QueryJudgeQuestionByQuestionId(questionId int) (*models.JudgeQuestion, error)
	QueryLastJudgeQuestion() (*models.JudgeQuestion, error)
	QueryLastSelectQuestion() (*models.SelectQuestion, error)
	QueryLastFillQuestion() (*models.FillQuestion, error)

	QueryQuestionByInherit() (*[]models.Problem, error)

	// message
	QueryMessageByPage(pageNum, pageSize int) (*models.Page, error)
	AddMessage(m *models.Message) error

	// score
	QueryAllScores() (*[]models.Score, error)                                              // ????????????????????????
	QueryScoreByExamCode(examCode int) (*[]models.Score, error)                            // ??????????????????????????????????????????
	QueryScoreByStudentId(studentId string) (*[]models.Score, error)                       // ????????????????????????????????? ?????????
	QueryStudentScoreByPage(studentId string, pageNum, pageSize int) (*models.Page, error) // ????????????????????????????????? ??????
	AddScore(score *models.Score) error                                                    // ????????????
}

type dao struct {
	orm *gorm.DB
}

func New() Dao {
	return &dao{
		orm: NewMySQL(),
	}
}

// Close close the resource.
func (d *dao) Close() {
	sqlDB, _ := d.orm.DB()
	sqlDB.Close()
}

// Ping ping the resource.
func (d *dao) Ping() (err error) {
	sqlDB, _ := d.orm.DB()
	return sqlDB.Ping()
}

func NewMySQL() (db *gorm.DB) {
	dsn := "root:root@tcp(139.196.218.71:3306)/examination?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("orm: open error(%v)", err)
		panic(err)
	}

	sqlDB, err := db.DB()
	// SetMaxIdleConns ?????????????????????????????????????????????
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns ?????????????????????????????????????????????
	sqlDB.SetMaxOpenConns(20)
	// SetConnMaxLifetime ??????????????????????????????????????????
	sqlDB.SetConnMaxLifetime(time.Hour * 4)

	return
}
