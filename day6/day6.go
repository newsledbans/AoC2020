package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day(input []byte) {

	inArr := strings.Split(string(input), "\n\n")
	part1, part2 := 0, 0

	for _, data := range inArr { // for each group
		charMap := make(map[rune]int)
		dataArr := strings.Split(string(data), "\n")

		for _, report := range dataArr { // for person
			for _, char := range report {
				charMap[char]++
			}
		}
		for _, v := range charMap {
			if len(dataArr) == v {
				part2++
			}
		}

		part1 += len(charMap)
	}
	fmt.Println("part1", part1)
	fmt.Println("part2", part2)
}

func castStoI(s string) int {
	// converts string to an int, returns 0 if err
	int64, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int(int64)
}
func intPow(a, n int) int {
	// return a^n
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
	dat, err := ioutil.ReadFile("06.txt")
	check(err)
	day(dat)

	// EXAMPLE CASE(S)
	// 	day([]byte(
	// 		`abc

	// a
	// b
	// c

	// ab
	// ac

	// a
	// a
	// a
	// a

	// b`))

}
