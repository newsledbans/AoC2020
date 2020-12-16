package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var inArr []string
var part1, part2 int

func day(input []byte) {

	// Parsing
	inArr = strings.Split(string(input), "\n\n")
	reClass := regexp.MustCompile(`[\d]+`)
	classArr := reClass.FindAllString(inArr[0], -1)

	// silent function splits on mulitple runes without multiple calls to split.
	nearbyArr := strings.FieldsFunc(inArr[2], func(r rune) bool { return strings.ContainsRune("\n,", r) })
	nearby, class, invalid := []int{}, []int{}, []int{}
	for _, n := range nearbyArr[1:] {
		num, _ := strconv.Atoi(n)
		nearby = append(nearby, num)
	}
	for _, n := range classArr {
		num, _ := strconv.Atoi(n)
		class = append(class, num)
	}
	// PART 1
	for _, val := range nearby {
		ok := false
		for j := 0; j < len(class)/4; j++ {
			if !((val >= class[j*4] && val <= class[j*4+1]) || (val >= class[j*4+2] && val <= class[4*j+3])) {
				continue
			}
			ok = true
			break
		}
		if !ok {
			invalid = append(invalid, val)
		}
	}

	for _, values := range invalid {
		part1 += values
	}

	fmt.Println("part1:", invalid, part1)
	fmt.Println("part2:", part2)
}

func main() {
	dat, _ := ioutil.ReadFile("inputs/16.txt")
	day(dat)

	// EXAMPLE CASE(S)

	// 	day([]byte(
	// 		`class: 1-3 or 5-7
	// row: 6-11 or 33-44
	// seat: 13-40 or 45-50

	// your ticket:
	// 7,1,14

	// nearby tickets:
	// 7,3,47
	// 40,4,50
	// 55,2,20
	// 38,6,12`))

}
