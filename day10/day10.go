package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var inArr []string

func day(input []byte) {

	inArr := strings.Split(string(input), "\n")

	for i, intCh := range inArr {
	}
}

func castSArrtoIArr(sa []string) []int {
	var iArr []int
	for _, v := range sa {
		iArr = append(iArr, castStoI(v))
	}
	return iArr
}
func castStoI(s string) int {
	// converts string to an int, returns 0 if err
	intL, err := strconv.ParseInt(s, 10, 0)
	check(err)
	return int(intL)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("inputs/09.txt")
	check(err)
	day(dat)

	// EXAMPLE CASE(S)
	//
	// day([]byte(
	// 	``))
}
