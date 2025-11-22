package reports

import (
	"fmt"
	"tusur-forms/internal/domain"

	"github.com/xuri/excelize/v2"
)

type Spreadsheet struct {
	File *excelize.File
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
		File: excelize.NewFile(),
	}
}

func (s *Spreadsheet) CreateSpreadsheet(title string) (int, error) {
	return 0, nil
}
func (s *Spreadsheet) SetHeader(headers *domain.Form) error {
	return nil
}
func (s *Spreadsheet) SetData(data *domain.Form) error {
	return nil
}
func (s *Spreadsheet) SaveFile(filename string) error {
	return nil
}

func main() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set value of a cell.
	err = f.SetCellValue("Sheet2", "A2", "Hello world.")
	if err != nil {
		return
	}
	err = f.SetCellValue("Sheet1", "B2", 100)
	if err != nil {
		return
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
