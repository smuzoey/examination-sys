package models

import (
	"strconv"
)

type PaperLimit struct {
	Id            int      `json:"id"`            // 试卷编号
	TotalScore    int      `json:"totalScore"`    // 试卷总分
	Difficulty    float64  `json:"difficulty"`    // 难度
	Point         []string `json:"point"`         // 类型
	EachTypeCount [3]int   `json:"eachTypeCount"` // 三种题型数量
}

type Unit struct {
	Id               int
	AdaptationDegree float64
	KPCoverage       float64
	ProblemList      []Problem
}

type Problem struct {
	QuestionId int    `json:"questionId" gorm:"column:questionId"`
	Score      int    `json:"score"`
	Section    string `json:"section"`
	Type       int    `json:"type"`
	Level      string `json:"level"`
}

func (u Unit) GetSumScore() int {
	sum := 0
	for _, v := range u.ProblemList {
		sum = sum + v.Score
	}
	return sum
}

func (u Unit) Difficulty() float64 {
	diff := 0.00
	for _, v := range u.ProblemList {
		tmp, _ := strconv.Atoi(v.Level)
		diff += float64(tmp) * 0.2 * float64(v.Score)
	}
	return diff / float64(u.GetSumScore())
}
