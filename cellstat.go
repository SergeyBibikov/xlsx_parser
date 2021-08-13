package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
)

//Преобразует файл с терминами ы мапу
func termFileToMap(filename string) *map[string]int {
	termsMap := make(map[string]int)
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	rows, err := f.GetRows("Лист1")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, cells := range rows {
		for _, cell := range cells {
			cell = strings.ToLower(cell)
			termsMap[cell] = 0
		}
	}
	return &termsMap
}

func CountCells(sourceFile string, termsFile string) {
	m := termFileToMap(termsFile)
	var wg sync.WaitGroup
	ch := make(chan string)
	resultMap := make(map[string]int)

	f, err := excelize.OpenFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("Лист1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, cells := range rows {
		for _, cell := range cells {
			if cell != "" && cell != " " {
				cell = strings.ToLower(cell)
				wg.Add(1)
				go looop(&wg, m, &cell, ch)
			}
		}
	}
	counter := 0
outer:
	for {
		select {
		case key := <-ch:
			resultMap[key]++
			counter = 0
		default:
			if counter > 7 {
				break outer
			}
			counter++
			time.Sleep(time.Second)
		}
	}
	wg.Wait()
	MapToFile(&resultMap, "CellCountResults.xlsx")
}

func MapToFile(m *map[string]int, resultsName string) {
	termIndex := 2
	countIndex := 2
	nf := excelize.NewFile()
	index := nf.NewSheet("Sheet1")
	nf.SetActiveSheet(index)
	nf.SetCellValue("Sheet1", "A1", "Term")
	nf.SetCellValue("Sheet1", "B1", "Count")
	for k, v := range *m {
		termCell := fmt.Sprintf("A%d", termIndex)
		countCell := fmt.Sprintf("B%d", termIndex)
		nf.SetCellValue("Sheet1", termCell, k)
		nf.SetCellValue("Sheet1", countCell, v)
		termIndex++
		countIndex++
	}
	if err := nf.SaveAs(resultsName); err != nil {
		fmt.Println(err)
	}
}

//Выполняет всю грязную работу
func looop(wg *sync.WaitGroup, m *map[string]int, cell *string, ch chan string) {
	allSigns := "[^A-Za-zА-Яа-я-]"
	for k := range *m {
		if strings.Contains(*cell, k) {
			exp := fmt.Sprintf("%s%s%s", allSigns, k, allSigns)
			if b, _ := regexp.MatchString(exp, *cell); b {
				ch <- k
			}
		}
	}
	wg.Done()
}
