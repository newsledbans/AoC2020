package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	reHgt := regexp.MustCompile(`(^[\d]{2,3})[ic][nm]$`)
	test := "150in"
	answer := reHgt.MatchString(test)
	units := test[len(test)-2:]
	measure, _ := strconv.Atoi(test[:len(test)-2])
	if units == "in" {
		fmt.Println(answer, measure, units)
	}
}
