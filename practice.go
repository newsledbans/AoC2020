package main

import (
	"fmt"
)

func main() {

	var bagMap = make(map[string]map[string]int)
	bagMap["shinygold"] = map[string]int{}
	bagMap["shinygold"]["darkred"] = 2
	bagMap["darkred"] = nil

	if v, _ := bagMap["darkred"]; v != nil {
		fmt.Println("not ok", bagMap["darkred"])
	} else {
		fmt.Println("is ok", bagMap["darkred"])
	}

	fmt.Println(bagMap)

	// reHgt := regexp.MustCompile(`(^[\d]{2,3})[ic][nm]$`)
	// test := "150in"
	// answer := reHgt.MatchString(test)
	// units := test[len(test)-2:]
	// measure, _ := strconv.Atoi(test[:len(test)-2])
	// if units == "in" {
	// 	fmt.Println(answer, measure, units)
	// }

}
