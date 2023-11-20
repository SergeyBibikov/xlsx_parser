package main

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	c := countWords("vpss.xlsx")
	r := sortWords(c)
	for _, v := range r {
		fmt.Println(v.word, ":", v.count)
		break
	}
	for i := 0; i < 10; i++ {
		fmt.Println(r[i].word, ":", r[i].count)

	}
}
