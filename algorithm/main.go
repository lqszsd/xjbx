package main

import (
	"fmt"
	"github.com/xuri/excelize"
)

func main() {
	xlsx := excelize.NewFile()

	index := xlsx.NewSheet("Sheet1")
	xlsx.SetCellValue("Sheet1", "A1", "姓名")
	xlsx.SetCellValue("Sheet1", "B1", "年龄")
	xlsx.SetCellValue("Sheet1", "A2", "狗子")
	xlsx.SetCellValue("Sheet1", "B2", "18")
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	test,_:=excelize.ColumnNumberToName(2)
	fmt.Println(test)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("test_write.xlsx")

	if err != nil {
		fmt.Println(err)
	}
}
