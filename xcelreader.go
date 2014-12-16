package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "excel.xlsx"
	xlFile, error := xlsx.OpenFile(excelFileName)
	if error == nil {
		for _, sheet := range xlFile.Sheets {
			for _, row := range sheet.Rows {
				for _, cell := range row.Cells {
					fmt.Printf("%s ", cell.String())
				}
				fmt.Printf("\n")
			}
		}
	}

}
