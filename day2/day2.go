package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func day2(input []byte) int {

	var goodPass int
	reNum := regexp.MustCompile(`\d+`)
	rePass := regexp.MustCompile(`[a-z]+`)

	inArray := bytes.Split(input, []byte("\n"))
	for _, value := range inArray {
		minmax := reNum.FindAll(value, -1)

		password := rePass.FindAll(value, -1)
		count := bytes.Count(password[1], password[0])

		if count >= castBtoI(minmax[0]) && count <= castBtoI(minmax[1]) {
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
	dat, err := ioutil.ReadFile("inputs/02.txt")
	check(err)
	answer := day2(dat)
	fmt.Println(answer)

	// EXAMPLE CASE(S)
	// answer := day2([]byte(
	// 	`1-3 a: abcde
	// 	1-3 b: cdefg
	// 	2-9 c: ccccccccc`))
	// fmt.Println(answer)

}
