package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day(input []byte) {

	inArray := strings.Split(string(input), "\n")
	result := 0
	for _, item := range inArray {
		data := strings.Fields(item)

		result += data
	}
	fmt.Println("Result:", result)
}

func castStoI(s string) int {
	int64, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int(int64)
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
	// ``))

}
