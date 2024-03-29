package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
)

// Преобразует файл с терминами в мапу
func termFileToMap(filename string) *map[string]int {
	termsMap := make(map[string]int)
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	rows := getSheetRows(f)
	for _, cells := range rows {
		for _, cell := range cells {
			cell = strings.ToLower(cell)
			termsMap[cell] = 0
		}
	}
	return &termsMap
}

// Считает, в скольких ячейках есть тот или иной термин
// Выводит результат в файл
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
	rows := getSheetRows(f)
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
	mapToFile(&resultMap, "CellCountResults.xlsx")
}

// Считает количество ячеек с терминами.
// Выводит результат в терминале
func CountCellsWithTerms(sourceFile string, termsFile string) {
	m := termFileToMap(termsFile)

	f, err := excelize.OpenFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := getSheetRows(f)

	ch := make(chan struct{}, 2000)
	for _, row := range rows {
		for _, cell := range row {
			go findTerm(m, cell, ch)
		}
	}
	resultsCount := 0
	firstRead := false
	selectCount := 0
outer:
	for {
		select {
		case <-ch:
			firstRead = true
			resultsCount++
		default:
			if firstRead && selectCount > 10 {
				break outer
			} else if !firstRead {

			} else {
				selectCount++
				time.Sleep(time.Second)
			}
		}
	}

	fmt.Println("Количество ячеек, в которых есть хотя бы один термин:", resultsCount)
}

// Пишет в канал ключи, который найдены в ячейке
func looop(wg *sync.WaitGroup, m *map[string]int, cell *string, ch chan string) {
	startSigns := "(^|[^0-9A-Za-zА-Яа-я_=+#~])"
	endSigns := "(s|es|$|[^0-9A-Za-zА-Яа-я_=+#~])"
	for k := range *m {
		if strings.Contains(*cell, k) {
			exp := fmt.Sprintf("%s%s%s", startSigns, k, endSigns)
			if b, _ := regexp.MatchString(exp, *cell); b {
				ch <- k
			}
		}
	}
	wg.Done()
}

// Пишет в канал 1, после нахождения первого термина
func findTerm(m *map[string]int, cell string, ch chan struct{}) {
	startSigns := "(^|[^0-9A-Za-zА-Яа-я_=+#~])"
	endSigns := "(s|es|$|[^0-9A-Za-zА-Яа-я_=+#~])"
	for k := range *m {
		if strings.Contains(cell, k) {
			exp := fmt.Sprintf("%s%s%s", startSigns, k, endSigns)
			if b, _ := regexp.MatchString(exp, cell); b {
				ch <- struct{}{}
				break
			}
		}
	}
}
