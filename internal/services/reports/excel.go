package reports

import (
	"fmt"
	"log"
	"strconv"
	"tusur-forms/internal/domain"

	"github.com/xuri/excelize/v2"
)

type Spreadsheet struct {
	file *excelize.File
}

type ExcelReport interface {
	CreateFile() *Spreadsheet
	CreateSpreadsheet(title string) (int, error)
	SetHeader(headers *domain.Form) error
	SetData(data *domain.Form) error
	SaveFile(filename string) error
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
		cell := rune('a' + (i))
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
		cell := rune('a' + (i))
		for j, ans := range item.Answers {
			log.Println(string(cell) + strconv.Itoa(j+2))
			err := s.file.SetCellValue(title, string(cell)+strconv.Itoa(j+2), ans.Content)
			if err != nil {
				return err
			}
		}
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

//func main() {
//	err = f.SetCellValue("Sheet2", "A2", "Hello world.")
//	if err != nil {
//		return
//	}
//	err = f.SetCellValue("Sheet1", "B2", 100)
//	if err != nil {
//		return
//	}
//}
