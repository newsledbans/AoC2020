package main

import (
	"fmt"
	"strings"
)

var inArr []string
var part1, part2 int

func day(input []byte) {

	//Parsing Input
	inArr = strings.Split(string(input), "\n")
	parseArr := strings.Split(inArr[1], ",")

	// PART 1
	for i, val := range parseArr {
		fmt.Println(i, val)
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func main() {
	// dat, _ := ioutil.ReadFile("inputs/17.txt")
	// day(dat)

	// EXAMPLE CASE(S)

	day([]byte(``))

}
