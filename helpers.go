package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/xuri/excelize/v2"
)

func hasLetters(word string) bool {
	for _, v := range word {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}

func isAcronym(word string) bool {
	// Проверяет, является ли слово аббревиатурой
	if len([]rune(word)) < 2 {
		return false
	}
	return strings.ToUpper(word) == word
}

func collectString(filename string, makeLower bool) *string {
	strArray := make([]string, 10000)
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
	for _, row := range rows {
		for _, cell := range row {
			if cell != "" && cell != " " {
				if makeLower {
					strArray = append(strArray, strings.ToLower(cell))
				} else {
					strArray = append(strArray, cell)
				}
			}
		}
	}
	finst := strings.Join(strArray, " ")
	return &finst
}

func fileToString(filename string) *string {
	//Собирает xlsx файл в одну строку
	return collectString(filename, false)
}

func fileToStringLowercase(filename string) *string {
	//Собирает xlsx файл в одну строку, делая слова lowercase
	return collectString(filename, true)
}

func _setupTheFile() *excelize.File {
	resultFile := excelize.NewFile()
	index := resultFile.NewSheet("Sheet1")
	resultFile.SetActiveSheet(index)
	resultFile.SetCellValue("Sheet1", "A1", "Term")
	resultFile.SetCellValue("Sheet1", "B1", "Count")
	return resultFile
}
func mapToFile(m *map[string]int, resultsName string) {
	termIndex := 2
	countIndex := 2
	resultFile := _setupTheFile()
	for k, v := range *m {
		termCell := fmt.Sprintf("A%d", termIndex)
		countCell := fmt.Sprintf("B%d", termIndex)
		resultFile.SetCellValue("Sheet1", termCell, k)
		resultFile.SetCellValue("Sheet1", countCell, v)
		termIndex++
		countIndex++
	}
	if err := resultFile.SaveAs(resultsName); err != nil {
		fmt.Println(err)
	}
}

func sliceToFile(data []tuple, resultsName string) {
	termIndex := 2
	countIndex := 2
	resultFile := _setupTheFile()
	for _, v := range data {
		termCell := fmt.Sprintf("A%d", termIndex)
		countCell := fmt.Sprintf("B%d", termIndex)
		resultFile.SetCellValue("Sheet1", termCell, v.word)
		resultFile.SetCellValue("Sheet1", countCell, v.count)
		termIndex++
		countIndex++
	}
	if err := resultFile.SaveAs(resultsName); err != nil {
		fmt.Println(err)
	}
}

func parseFlags() (string, string, int, int) {
	help := flag.Bool("help", false, "Показать помощь")

	modesHelp := strings.Trim(`
Режим работы команды
Доступные режимы:
	countCells: full, cell, total
	countWords: minLen, minCount
Для команды countCells обязательно нужно указать один режим работы,
Для команды countWords можно указывать один или два параметра, а также не указывать ни одного
    `, "\t \n")
	mode := flag.String("mode", "", modesHelp)
	command := flag.String("command", "", "Команда, которую нужно выполнить. Доступные команды: countCells, countWords")
	minLen := flag.Int("minLen", -1, "Минимальная длина слов, которые нужно выводить в отчёте")
	minCount := flag.Int("minCount", -1, "Минимальное количество повторений слова, которые нужно выводить в отчёте")

	flag.Parse()

	if *help {
		fmt.Println("Использование: xlsx_parser.exe [файлы] [команда] [режим работы команды]")
		fmt.Println("Пример: xlsx_parser.exe -command=countCells -mode=full analyzeMe.xlsx terms.xlsx")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *command == "" {
		fmt.Printf("Не выбрана ни одна команда, для вывода помощи используйте флаг -help\n")
		os.Exit(1)
	}

	countCellModes, ok := possibleWorkModes[*command]
	if !ok {
		fmt.Printf("Указанная команда не поддерживается, для вывода помощи используйте флаг -help\n")
		os.Exit(1)
	}

	if *command == "countCells" {
		_, ok = countCellModes[*mode]
		if !ok {
			fmt.Printf("Указанный режим работы не поддерживается, для вывода помощи используйте флаг -help\n")
			os.Exit(1)
		}
	}

	return *command, *mode, *minLen, *minCount
}
