package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inArr []string

func day(input []byte) {

	inArr := strings.Fields(string(input))

	p2Target := 0

	for i, intCh := range inArr {
		ok := false
		if i < 25 {
			continue
		}
		hi := i - 1
		lo := i - 25
		intTarget := castStoI(intCh)
		for j := lo; j < hi; j++ {
			for k := j + 1; k <= hi; k++ {
				if intTarget == castStoI(inArr[j])+castStoI(inArr[k]) {
					ok = true
				}
			}
		}
		if !ok {
			p2Target = castStoI(intCh)
			fmt.Println("part1:", intCh)
			break
		}

	}

	// var working []int
	var sum int
	for j := 0; j < len(inArr); j++ {
		for k := j + 1; k <= len(inArr); k++ {
			sum = 0
			for _, element := range inArr[j:k] {

				sum += castStoI(element)
			}
			if sum == p2Target {
				tMin, tMax := minMax(castSArrtoIArr(inArr[j:k]))
				fmt.Println("target", sum, "Part 2:", tMin+tMax)
				break
			} else if sum > p2Target {
				continue
			}
		}
		if sum == p2Target {
			break
		}
		if sum > p2Target {
			continue
		}
	}
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
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
	dat, err := ioutil.ReadFile("09.txt")
	check(err)
	day(dat)

	// EXAMPLE CASE(S)
	//
	// day([]byte(
	// 	`35
	// 20
	// 15
	// 25
	// 47
	// 40
	// 62
	// 55
	// 65
	// 95
	// 102
	// 117
	// 150
	// 182
	// 127
	// 219
	// 299
	// 277
	// 309
	// 576`))
}
