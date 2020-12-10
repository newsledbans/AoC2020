package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	aocu "github.com/newsledbans/aoc2020/aocutils"
)

var inArr []string
var intArr []int
var part1, part2 int

func day(input []byte) {
	inArr := strings.Split(string(input), "\n")
	intArr := aocu.SArrtoIArr(inArr)

	for i, line := range intArr {
		fmt.Println(i, line)
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func main() {
	dat, err := ioutil.ReadFile("inputs/11.txt")
	aocu.Check(err)
	day(dat)

	// EXAMPLE CASE(S)
	//
	// 	day([]byte(
	// 		``))
}
