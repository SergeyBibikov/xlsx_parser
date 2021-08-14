package main

import (
	"fmt"
	"strings"
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
	// re := regexp.MustCompile("(^|[^0-9A-Za-zА-Яа-я])kpi(s|es|$|[^0-9A-Za-zА-Яа-я])")
	// fmt.Println(re.FindStringIndex("  kpis"))
	// st := ".ts.ts.ts   ts ts"
	// fmt.Println(countMatches(&st, "(^|[^0-9A-Za-zА-Яа-я_=+#~])ts(s|es|$|[^0-9A-Za-zА-Яа-я_=+#~])"))
	st := fileToString("senru.xlsx")
	fmt.Println(strings.Contains(*st, "kpi"))
	fmt.Println(countMatches(*st, "(^|[^0-9A-Za-zА-Яа-я_=+#~])preparation(s|es|$|[^0-9A-Za-zА-Яа-я_=+#~])"))
}
