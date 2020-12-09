package main

import (
	"fmt"
	"strconv"
)

func castStoI(s string) int {
	// converts string to an int, returns 0 if err
	int64, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int(int64)
}

func main() {

	instr := "acc -1"
	inInt := castStoI(instr[4:])
	fmt.Println(inInt)

	// var bagMap = make(map[string]map[string]int)
	// bagMap["shinygold"] = map[string]int{}
	// bagMap["shinygold"]["darkred"] = 2
	// bagMap["darkred"] = nil

	// if v, _ := bagMap["darkred"]; v != nil {
	// 	fmt.Println("not ok", bagMap["darkred"])
	// } else {
	// 	fmt.Println("is ok", bagMap["darkred"])
	// }

	// fmt.Println(bagMap)

	// reHgt := regexp.MustCompile(`(^[\d]{2,3})[ic][nm]$`)
	// test := "150in"
	// answer := reHgt.MatchString(test)
	// units := test[len(test)-2:]
	// measure, _ := strconv.Atoi(test[:len(test)-2])
	// if units == "in" {
	// 	fmt.Println(answer, measure, units)
	// }

}
