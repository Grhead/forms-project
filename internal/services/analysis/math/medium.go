package math

import (
	"strconv"
	"tusur-forms/internal/domain"
)

type MediumAnalysis interface {
	CalculateDiscipline(data *domain.Form, discipline string) *Medium
	CalculateQuestion(data *domain.Form) *Medium
	CalculateOverall()
}
type Medium struct {
	Value []float64
}

func CalculateQuestion(data *domain.Form) *Medium {
	var medium Medium
	for _, item := range data.Questions {
		var tempMedium float64
		for _, ans := range item.Answers {
			intContent, _ := strconv.ParseFloat(ans.Content, 32)
			tempMedium += intContent
		}
		tempMedium = tempMedium / float64(len(item.Answers))
		medium.Value = append(medium.Value, tempMedium)
	}
	return &medium
}

func CalculateDiscipline(data *domain.Form, discipline string) *Medium {
	var medium Medium
	var tempMedium float64
	var discCount = 0
	for _, item := range data.Questions {
		if item.Description == discipline {
			for _, ans := range item.Answers {
				intContent, _ := strconv.ParseFloat(ans.Content, 32)
				tempMedium += intContent
			}
			discCount++
		}
	}
	tempMedium = tempMedium / float64(discCount)
	medium.Value = append(medium.Value, tempMedium)
	return &medium
}
