package main

import (
	"fmt"
	"testing"
)

// func TestF2M(t *testing.T) {
// 	FileToMap("glosru.xlsx")
// }

// func TestInc(t *testing.T) {
// 	m := make(map[string]int)
// 	m["one"] = 1
// 	IncreaseCount(&m, "one")
// 	fmt.Println(m["one"])
// }

func TestCount(t *testing.T) {
	m := make(map[string]int)
	m["hey"]++
	fmt.Println(m)
}
