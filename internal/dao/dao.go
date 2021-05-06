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
	DeleteStudent(studentId int) error
	StudentChangePwd(studentId int, pwd string) error
	UpdateStudent(student *models.Student) error

	QueryTeacher(id string) (*models.Teacher, error)

	// exam
	QueryExamByPage(pageNum, pageSize int) (*models.Page, error)
	QueryExamById(id int) (*models.ExamManage, error)
	AddExam(e *models.ExamManage) error

	// paper
	QueryPaperById(paperId int) (*[]models.Paper, error)

	// question
	QuerySelectQuestionByPaperId(paperId int) (*[]models.SelectQuestion, error)
	QueryFillQuestionByPaperId(paperId int) (*[]models.FillQuestion, error)
	QueryJudgeQuestionByPaperId(paperId int) (*[]models.JudgeQuestion, error)
	QueryQuestionsByPage(pageNum, pageSize int) (*models.Page, error)

	// message
	QueryMessageByPage(pageNum, pageSize int) (*models.Page, error)
	AddMessage(m *models.Message) error

	// score
	QueryAllScores() (*[]models.Score, error)                                       // 所有学生所有成绩
	QueryScoreByExamCode(examCode string) (*[]models.Score, error)                  // 查询某个考试的所有学生的成绩
	QueryScoreByStudentId(studentId int) (*[]models.Score, error)                   // 查询某个学生的所有成绩 不分页
	QueryStudentScoreByPage(studentId, pageNum, pageSize int) (*models.Page, error) // 查询某个学生的所有成绩 分页
	AddScore(score *models.Score) error                                             // 添加成绩
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
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(20)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour * 4)

	return
}
