package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var inArr []string
var part1, part2 int

func day(input []byte) {
	inArr = strings.Split(string(input), "\n")
	busArr := strings.Split(inArr[1], ",")

	// PART 1

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func main() {
	dat, _ := ioutil.ReadFile("inputs/13.txt")
	day(dat)

	// EXAMPLE CASE(S)

	day([]byte(
		`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`))

}
