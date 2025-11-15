package database

import (
	"tusur-forms/internal/domain"
)

type dbQuestion struct {
	ID           string `gorm:"primaryKey"`
	Title        string
	Description  string
	TypeID       string
	QuestionType dbQuestionType `gorm:"foreignKey:TypeID;references:ID"`
	IsRequired   bool
}

type dbQuestionType struct {
	ID    string `gorm:"primaryKey"`
	Title string
}

func (g *GormRepository) CreateQuestion(q *domain.Question) error {
	qts, err := g.getQuestionType(q.Type.ID)
	if err != nil {
		return err
	}
	exists, err := g.checkQuestionType(qts)
	if err != nil {
		return err
	}
	if !exists {
		g.createQuestionType(&q.Type)
	}
	dbQ := dbQuestion{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		IsRequired:  false,
		TypeID:      q.Type.ID,
	}

	err = g.db.Create(&dbQ).Error
	if err != nil {
		return err
	}
	if q.Type.Title == domain.TypeCheckbox || q.Type.Title == domain.TypeRadio {
		for _, item := range q.PossibleAnswers {
			paID, err := g.getPossibleAnswer(item)
			if err != nil {
				return err
			}
			if paID == nil {
				_, err = g.CreatePossibleAnswer(item, q)
				if err != nil {
					return err
				}
			}
		}
	}

	return g.db.Save(&dbQ).Error
}

func (g *GormRepository) createQuestionType(qt *domain.QuestionType) error {
	dbQt := dbQuestionType{
		ID:    qt.ID,
		Title: string(qt.Title),
	}

	err := g.db.Create(&dbQt).Error
	if err != nil {
		return err
	}
	return g.db.Save(&dbQt).Error
}

func (g *GormRepository) getQuestionType(qtID string) ([]*domain.QuestionType, error) {
	var qts []*dbQuestionType
	err := g.db.Where("id = ?", qtID).Find(&qts)
	if err != nil {
		return nil, err.Error
	}
	var result []*domain.QuestionType
	for i := range qts {
		result = append(result, &domain.QuestionType{
			ID:    qts[i].ID,
			Title: domain.QuestionTypeTitles(qts[i].Title),
		})
	}
	return result, nil
}

func (g *GormRepository) checkQuestionType(qts []*domain.QuestionType) (bool, error) {
	var qtss []string
	var types []*dbQuestionType
	for _, item := range qts {
		qtss = append(qtss, item.ID)
	}
	err := g.db.
		Where("id IN ? ", qtss).
		Find(&types)
	if err != nil {
		return false, err.Error
	}
	return true, nil
}
