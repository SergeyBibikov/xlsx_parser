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
	fmt.Println(trimPunctuation("\"Hey,"))
}
