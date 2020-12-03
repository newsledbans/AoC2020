package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func part2(input []byte) int {

	var goodPass int
	reNum := regexp.MustCompile(`\d+`)
	rePass := regexp.MustCompile(`[a-z]+`)

	inArray := bytes.Split(input, []byte("\n"))
	for _, value := range inArray {
		minmax := reNum.FindAll(value, -1)
		pos1 := castBtoI(minmax[0]) - 1
		pos2 := castBtoI(minmax[1]) - 1
		password := rePass.FindAll(value, -1)

		X := password[0][0] == password[1][pos1]
		// if len(password[1]) <= pos2+1 &&
		Y := password[0][0] == password[1][pos2]
		if (X || Y) && (X != Y) {
			goodPass++
		}
	}
	return goodPass
}

func castBtoI(byteSlice []byte) int {
	byteToInt, err := strconv.Atoi(string(byteSlice))
	check(err)
	return byteToInt
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("02.txt")
	check(err)
	answer := part2(dat)
	fmt.Println(answer)

	// EXAMPLE CASE(S)
	// answer := part2([]byte(
	// 	`1-3 a: abcde
	// 	1-3 b: cdefg
	// 	2-9 c: ccccccccc`))
	// fmt.Println(answer)

}
