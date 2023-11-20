package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Переданы не все аргументы")
		fmt.Println("Необходимо передать название source файла, файла-глоссария и режим подсчёта(full, cell или total)")
		return
	}
	sourceFile := os.Args[1]
	termsFile := os.Args[2]
	mode := os.Args[3]

	if mode != "full" && mode != "cell" && mode != "total" {
		fmt.Println("Возможны три режима работы: full, cell или total.\nНеобходимо указать один из них в качестве последнего аргумента")
		return
	}
	fmt.Println("Начинаю подсчёт. Для выхода нажмите Ctrl+C\n===============")
	start := time.Now().Unix()
	if mode == "full" {
		st := fileToStringLowercase(sourceFile)
		CountFull(termsFile, st)
	}
	if mode == "cell" {
		CountCells(sourceFile, termsFile)
	}
	if mode == "total" {
		CountCellsWithTerms(sourceFile, termsFile)
	}
	finish := time.Now().Unix()
	fmt.Println("Готово\n", "Отчёт сформирован за", finish-start, "секунд")
}

func CountFull(termsFilename string, st *string) {
	// Считает, сколько раз в тексте содержится каждое из слов глоссария
	f, err := excelize.OpenFile(termsFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("Лист1")
	if err != nil {
		fmt.Println(err)
		return
	}
	resultsMap := make(map[string]int)

	ch := make(chan []interface{}, 200000)
	for _, row := range rows {
		for _, cell := range row {
			if cell != "" && cell != " " {
				cell = strings.ToLower(cell)
				go temp(*st, cell, ch)
			}
		}
	}
	counter := 0
	firstRead := false
outer:
	for {
		select {
		case slice := <-ch:
			firstRead = true
			key := slice[0].(string)
			v := slice[1].(int)
			resultsMap[key] = v
			counter = 0
		default:
			if firstRead && counter > 10 {
				break outer
			} else if !firstRead {

			} else {
				counter++
				time.Sleep(time.Second)
			}
		}
	}
	MapToFile(&resultsMap, "FullTextCountResults.xlsx")
}

func countMatches(st string, regex string) int {
	startingIndex := 0
	counter := 0
	re := regexp.MustCompile(regex)
	for {
		if res := re.FindStringIndex(st[startingIndex:]); res != nil {
			counter++
			startingIndex = startingIndex + res[0] + 2
		} else {
			break
		}
	}
	return counter
}

func temp(st string, cell string, ch chan []interface{}) {
	if strings.Contains(st, cell) {
		regex := fmt.Sprintf("(^|[^0-9A-Za-zА-Яа-я_=+#~])%s(s|es|$|[^0-9A-Za-zА-Яа-я_=+#~])", cell)
		count := countMatches(st, regex)
		if count > 0 {
			ch <- []interface{}{cell, count}
		}
	}
}
