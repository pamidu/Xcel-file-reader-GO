package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {

	//initliz
	excelFileName := "excel.xlsx"
	rowcount := 0
	colunmcount := 0
	var exceldata []map[string]string
	var colunName []string

	//file read
	xlFile, error := xlsx.OpenFile(excelFileName)

	if error == nil {
		for _, sheet := range xlFile.Sheets {
			//max colunm count and row count
			rowcount = sheet.MaxRow
			colunmcount = sheet.MaxCol
			//make colunm titile  to arry
			colunName = make([]string, colunmcount)
			for _, row := range sheet.Rows {
				for j, cel := range row.Cells {
					colunName[j] = cel.String()
				}
				break
			}
		}
	}
	fmt.Println(rowcount, "ghjkl")
	exceldata = make(([]map[string]string), rowcount)

	if error == nil {
		for _, sheet := range xlFile.Sheets {
			fmt.Println(sheet.MaxCol, sheet.MaxRow)
			for rownumber, row := range sheet.Rows {
				currentRow := make(map[string]string)
				exceldata[rownumber] = currentRow
				for cellnumber, cell := range row.Cells {
					fmt.Println(rownumber + 1)
					fmt.Println(colunName[cellnumber])
					fmt.Println(cell.String())

					exceldata[rownumber][colunName[cellnumber]] = cell.String()
				}
				fmt.Printf("\n")

			}
		}
	}
	fmt.Println(rowcount, colunmcount, colunName, exceldata)

}
