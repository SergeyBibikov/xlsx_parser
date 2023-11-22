package main

import (
	"fmt"
	"testing"
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
	words := getWordsAfterFiltering("vpss.xlsx", 5, 30)
	for _, v := range words {
		fmt.Printf("%s:%d\n", v.word, v.count)
	}
}
