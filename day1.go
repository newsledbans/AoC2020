package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day1(input []byte) (a, b, c, d int) {

	var intArray []int
	inArray := strings.Fields(string(input))
	for _, value := range inArray {
		intVal, err := strconv.Atoi(value)
		check(err)
		intArray = append(intArray, intVal)
	}

	for i, value := range intArray {
		target1 := 2020 - value
		for j := i + 1; j < len(intArray); j++ {
			for k := j + 1; k < len(intArray); k++ {
				if target1 == intArray[j]+intArray[k] {
					return value, intArray[j], intArray[k], value * intArray[j] * intArray[k]
				}
			}
		}
	}
	return 0, 1, 2, 3
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("01.txt")
	check(err)
	fmt.Println(day1(dat))

	// EXAMPLE CASE(S)
	// 	a, b, c, d := day1([]byte(`1721
	// 979
	// 366
	// 299
	// 675
	// 1456`))
	// 	fmt.Println(a, b, c, d)

}
