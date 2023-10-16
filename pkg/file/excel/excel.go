package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-boss/config"
	"go-boss/pkg/crypt"
	"go-boss/pkg/util"
)

// 创建excel
func Create() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	sheet := "Sheet1"
	index, err := f.NewSheet(sheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set value of a cell.
	f.SetCellValue(sheet, "A1", "姓名")
	f.SetCellValue(sheet, "B1", "年龄")

	f.SetCellValue(sheet, "A2", "golang")
	f.SetCellValue(sheet, "B2", "15")

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	c := config.NewConfig()
	viper := c.Viper
	dir := viper.Get("file.excel.export").(string)
	filename := crypt.Md5(util.RandStr(20))
	if err := f.SaveAs(dir + "/" + filename + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}

// 读取excel
func Read() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
