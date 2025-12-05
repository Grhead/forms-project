package reports

import (
	"fmt"
	"strconv"
	"tusur-forms/internal/domain"
	"tusur-forms/internal/services/analysis/math"

	"github.com/xuri/excelize/v2"
)

type Spreadsheet struct {
	file *excelize.File
}

type ExcelReport interface {
	CreateFile() *Spreadsheet
	CreateSpreadsheet(title string) (int, error)
	SetHeader(index int, headers *domain.Form) error
	SetData(index int, data *domain.Form) error
	SetMediumQuestion(index int, data *domain.Form) error
	SaveFile(index int, filename string) error
}

func CreateFile() *Spreadsheet {
	return &Spreadsheet{
		file: excelize.NewFile(),
	}
}
func (s *Spreadsheet) CreateSpreadsheet(title string) (int, error) {
	index, err := s.file.NewSheet(title)
	if err != nil {
		return 0, err
	}
	return index, nil
}
func (s *Spreadsheet) SetHeader(index int, headers *domain.Form) error {
	title := s.file.GetSheetName(index)
	if title == "" {
		return fmt.Errorf("index is not correct")
	}
	for i, item := range headers.Questions {
		cell := rune('b' + (i))
		err := s.file.SetCellValue(title, string(cell)+"1", item.Title)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *Spreadsheet) SetData(index int, data *domain.Form) error {
	title := s.file.GetSheetName(index)
	if title == "" {
		return fmt.Errorf("index is not correct")
	}
	for i, item := range data.Questions {
		cell := rune('b' + (i))
		for j, ans := range item.Answers {
			err := s.file.SetCellValue(title, string('a')+strconv.Itoa(j+2), ans.SubmittedAt)
			intContent, err := strconv.ParseInt(ans.Content, 10, 64)
			if err != nil {
				return err
			}
			err = s.file.SetCellInt(title, string(cell)+strconv.Itoa(j+2), intContent)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (s *Spreadsheet) SetMediumQuestion(index int, data *domain.Form) error {
	title := s.file.GetSheetName(index)
	if title == "" {
		return fmt.Errorf("index is not correct")
	}
	med := math.CalculateQuestion(data)
	for i, item := range data.Questions {
		cell := rune('b' + (i))
		err := s.file.SetCellFloat(title,
			string(cell)+strconv.Itoa(len(item.Answers)+2),
			med.Value[i],
			2,
			32)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *Spreadsheet) SetMediumDiscipline(index int, data *domain.Form) error {
	title := s.file.GetSheetName(index)
	if title == "" {
		return fmt.Errorf("index is not correct")
	}
	prev := ""
	for j, item := range data.Questions {
		if item.Description == prev {
			continue
		}
		cell := rune('b' + len(data.Questions))
		med := math.CalculateDiscipline(data, item.Description)
		err := s.file.SetCellValue(title,
			string(cell)+strconv.Itoa(j+1),
			item.Description)
		if err != nil {
			return err
		}
		cell = rune('b' + len(data.Questions) + 2)

		err = s.file.SetCellFloat(title,
			string(cell)+strconv.Itoa(j+1),
			med.Value[0],
			2,
			32)
		if err != nil {
			return err
		}
		prev = item.Description

	}

	return nil
}
func (s *Spreadsheet) SaveFile(index int, filename string) error {
	s.file.SetActiveSheet(index)
	err := s.file.SaveAs(filename)
	if err != nil {
		return err
	}
	return nil
}
