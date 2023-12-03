package main

import (
	"fmt"
	"testing"

	"github.com/xuri/excelize/v2"
)

// func TestCount(t *testing.T) {
// 	c := countWords("vpss.xlsx")
// 	r := sortWords(c)
// 	// res := onlyLength(r, 4)
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(res[i].word, ":", res[i].count)

// 	}
// }

func TestDebug(t *testing.T) {
	f, _ := excelize.OpenFile("vpss.xlsx")
	fmt.Println(getSheetRows(f))
}
