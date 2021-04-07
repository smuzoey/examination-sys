package models

import (
	"github.com/prometheus/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB = New()

type Dao interface {
	Ping() error
	Close()

	QueryStudent(id string) (*Student, error)
	QueryTeacher(id string) (*Teacher, error)
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
