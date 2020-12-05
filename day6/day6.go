package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day(input []byte) {

	inArray := strings.Split(string(input), "\n")
	result := make(map[int]string)
	part1 := 0

	for _, data := range inArray {

	}
	fmt.Println(part1)
}

// converts string to an int, returns 0 if err
func castStoI(s string) int {
	int64, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int(int64)
}

// return a^n
func intPow(a, n int) int {
	var i, result int
	result = 1
	for i = 0; i < n; i++ {
		result *= a
	}
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("05.txt")
	check(err)
	day(dat)

	// EXAMPLE CASE(S)
	// 	day([]byte(
	// 		``))

}
