package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

var termIndex = 2
var countIndex = 2

func main() {
	start := time.Now().Unix()
	st := fileToString(os.Args[1])
	count(os.Args[2], &st)
	finish := time.Now().Unix()
	fmt.Println("Done\n", "It took ", finish-start, "seconds")
}

func fileToString(filename string) string {
	strArray := make([]string, 10000)
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	rows, err := f.GetRows("Лист1")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, row := range rows {
		for _, cell := range row {
			if cell != "" && cell != " " {
				strArray = append(strArray, strings.ToLower(cell))
			}
		}
	}
	finst := strings.Join(strArray, " ")
	return finst
}

func count(filename string, st *string) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("Лист1")
	if err != nil {
		fmt.Println(err)
		return
	}
	nf := excelize.NewFile()
	index := nf.NewSheet("Sheet1")
	nf.SetActiveSheet(index)
	nf.SetCellValue("Sheet1", "A1", "Term")
	nf.SetCellValue("Sheet1", "B1", "Count")
	for _, row := range rows {
		for _, cell := range row {
			if cell != "" && cell != " " {
				if count := strings.Count(*st, strings.ToLower(cell)); count > 0 {
					termCell := fmt.Sprintf("A%d", termIndex)
					countCell := fmt.Sprintf("B%d", termIndex)
					nf.SetCellValue("Sheet1", termCell, cell)
					nf.SetCellValue("Sheet1", countCell, count)
					termIndex++
					countIndex++
				}
			}
		}
	}
	if err := nf.SaveAs("Results.xlsx"); err != nil {
		fmt.Println(err)
	}
}
