package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var inArr []string
var intArr []int
var dp = map[int]int{0: 1}
var count int

func day(input []byte) {
	inArr := strings.Split(string(input), "\n")
	intArr := castSArrtoIArr(inArr)
	intArr = append(intArr, 0)
	sort.Ints(intArr)
	max := intArr[len(intArr)-1]
	intArr = append(intArr, max+3)
	jDif := make([]int, len(intArr))

	for i := 0; i < len(intArr)-1; i++ {
		dif := intArr[i+1] - intArr[i]
		jDif[dif]++
	}

	fmt.Println("part1:", jDif[1]*jDif[3])

	for _, v := range intArr[1:] {
		dp[v] = dp[v-1] + dp[v-2] + dp[v-3]
	}
	part2 := dp[intArr[len(intArr)-1]]
	fmt.Println("part2:", part2)
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
	dat, err := ioutil.ReadFile("inputs/10.txt")
	check(err)
	day(dat)

	// EXAMPLE CASE(S)
	//
	// 	day([]byte(
	// 		`28
	// 33
	// 18
	// 42
	// 31
	// 14
	// 46
	// 20
	// 48
	// 47
	// 24
	// 23
	// 49
	// 45
	// 19
	// 38
	// 39
	// 11
	// 1
	// 32
	// 25
	// 35
	// 8
	// 17
	// 7
	// 9
	// 4
	// 2
	// 34
	// 10
	// 3`))
}
