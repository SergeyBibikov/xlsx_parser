package main

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

func isAcronym(word string) bool {
	// Проверяет, является ли слово аббревиатурой
	return len(word) > 1 && strings.ToUpper(word) == word
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
