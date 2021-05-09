package service

import (
	"examination-sys/internal/dao"
	"examination-sys/internal/models"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 计算unit知识点覆盖率
func GetKPCoverage(unitList *[]models.Unit, paper *models.PaperLimit) *[]models.Unit {
	for i := 0; i < len(*unitList); i++ {
		mp := make(map[string]int, 0)
		for _, v := range (*unitList)[i].ProblemList {
			mp[v.Section] = 1
		}

		containCnt := 0
		for _, v := range paper.Point {
			if mp[v] == 1 {
				containCnt += 1
			}
		}
		(*unitList)[i].KPCoverage = float64(containCnt) / float64(len(paper.Point))
	}
	return unitList
}

// f=1-(1-M/N)*f1-|EP-P|*f2 自适应度
func GetAdaptationDegree(unitList *[]models.Unit, KPCoverage float64, Difficulty float64, paper *models.PaperLimit) *[]models.Unit {
	for i := 0; i < len(*unitList); i++ {
		(*unitList)[i].AdaptationDegree = 1 - (1-(*unitList)[i].KPCoverage)*KPCoverage - math.Abs((*unitList)[i].Difficulty()-paper.Difficulty)*Difficulty
	}
	return unitList
}

// 判断题目知识点是否在试卷中
func PointIsContain(paper *models.PaperLimit, problem *models.Problem) bool {
	for _, v := range paper.Point {
		if v == problem.Section {
			return true
		}
	}
	return false
}

func InitUnit(count int, paper *models.PaperLimit, problemList *[]models.Problem) *[]models.Unit {
	var unitList []models.Unit

	// count：初始化 N 个种群   种群即试卷
	for i := 0; i < count; i++ {
		var unit models.Unit
		unit.Id = i
		unit.AdaptationDegree = 0

		for paper.TotalScore != unit.GetSumScore() {

			unit.ProblemList = []models.Problem{}

			// 三种题型：选择题、判断题、填空题
			for j := 0; j < 3; j++ {
				var tempProblemList []models.Problem

				for _, p := range *problemList {
					if p.Type == j+1 && PointIsContain(paper, &p) { // 知识点是否匹配paper
						tempProblemList = append(tempProblemList, p)
					}
				}

				// eachTypeCount[i] 为第i种题型的数量
				// 随机选择 cnt 道
				for cnt := 0; cnt < paper.EachTypeCount[j]; cnt++ {
					index := rand.Intn(len(tempProblemList) - cnt)
					unit.ProblemList = append(unit.ProblemList, tempProblemList[index])

					// 避免选到重复的，tempProblemList 末尾和 index 调换
					tmp := tempProblemList[len(tempProblemList)-cnt-1]
					tempProblemList[index] = tmp

				}
			}
		}
		unitList = append(unitList, unit)
	}
	// 计算种族的自适应度，用于选择算子进行选择
	unitList = *(GetKPCoverage(&unitList, paper))
	unitList = *(GetAdaptationDegree(&unitList, 0.5, 0.5, paper))

	return &unitList
}

// 选择算子
func Select(unitList *[]models.Unit, cnt int) *[]models.Unit {

	var adapterSum float64
	var selectUnit []models.Unit
	mp := make(map[int]int, 0) // 判断是否重复选择unit
	for _, v := range *unitList {
		adapterSum += v.AdaptationDegree
	}

	for len(selectUnit) != cnt {
		degree := 0.00
		rangeDegree := float64(rand.Intn(100)) * 0.01 * adapterSum

		for k, v := range *unitList {
			degree = degree + v.AdaptationDegree
			if degree >= rangeDegree && mp[k] == 0 {
				selectUnit = append(selectUnit, v)
				mp[k] = 1
				break
			}
		}
	}
	return &selectUnit
}

func Cross(unitList *[]models.Unit, cnt int, paper *models.PaperLimit) *[]models.Unit {

	times := int64(time.Now().Nanosecond())
	rand.Seed(times)

	var crossUnit []models.Unit
	mp := make(map[int]int, 0)

	for len(crossUnit) != cnt {

		// 随机选择两个个体
		index1 := rand.Intn(len(*unitList))
		index2 := rand.Intn(len(*unitList))
		unit1 := (*unitList)[index1]
		unit2 := (*unitList)[index2]

		if unit1.Id != unit2.Id && mp[index1] == 0 && mp[index2] == 0 {
			crossPosition := rand.Intn(len(unit1.ProblemList) - 2) // 交叉点
			score1 := unit1.ProblemList[crossPosition].Score + unit1.ProblemList[crossPosition+1].Score
			score2 := unit2.ProblemList[crossPosition].Score + unit2.ProblemList[crossPosition+1].Score

			if score1 == score2 {
				unitNew1, unitNew2 := models.Unit{}, models.Unit{}
				unitNew1.ProblemList, unitNew2.ProblemList = unit1.ProblemList, unit2.ProblemList

				// 交换交叉点后的两道题
				for i := crossPosition; i < crossPosition+2; i++ {
					unitNew1.ProblemList[i] = unit2.ProblemList[i]
					unitNew2.ProblemList[i] = unit1.ProblemList[i]
				}

				// 添加到新种群集合中
				unitNew1.Id = len(crossUnit)
				unitNew2.Id = unitNew1.Id + 1
				mp[index1], mp[index2] = 1, 1
				if len(crossUnit) < cnt {
					crossUnit = append(crossUnit, unitNew1)
				}
				if len(crossUnit) < cnt {
					crossUnit = append(crossUnit, unitNew2)
				}

			}
		}

	}

	crossUnit = *(GetKPCoverage(&crossUnit, paper))
	crossUnit = *(GetAdaptationDegree(&crossUnit, 0.5, 0.5, paper))
	return &crossUnit
}

func Inherit(limit *models.PaperLimit) *[]models.Unit {
	var unit *[]models.Unit
	problemList, _ := dao.DB.QueryQuestionByInherit()

	unit = InitUnit(8, limit, problemList)
	for _, v := range *unit {
		fmt.Println(v)
	}
	unit = Select(unit, 4)
	unit = Cross(unit, 2, limit)
	return unit
}
