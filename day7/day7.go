package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var bagMap = make(map[string]map[string]int)

func day(input []byte) {

	inArr := bytes.Split(input, []byte(".\n"))
	part1, part2 := 0, 0

	// fill bagmap with the names of bags and directions.
	for _, line := range inArr { // for each instruction
		lineArr := strings.Split(string(line), " ")
		bag0 := lineArr[0] + lineArr[1]
		lineLen := len(lineArr)
		if lineLen < 8 {
			bagMap[bag0] = nil
			continue
		}
		bagMap[bag0] = make(map[string]int)
		for i := 4; i < lineLen; i += 4 {
			bagMap[bag0][lineArr[i+1]+lineArr[i+2]] = castStoI(lineArr[i])
		}
	}
	// fmt.Println(bagMap)
	part2 = getTotal("shinygold")

	for eachBag := range bagMap {
		if getShiny(eachBag) {
			part1++
		}
	}

	fmt.Println("part1", part1)
	fmt.Println("part2", part2)

	// fmt.Println(bagMap)
}

func getTotal(bag string) int {
	subtotal := 0
	if firstKey := bagMap[bag]; firstKey == nil {
		return 0
	}
	for k, v := range bagMap[bag] {
		subtotal += v * (1 + getTotal(k))
	}
	return subtotal
}

func getShiny(bag string) bool {
	//recursive function to get whether it contains a shiny
	for inside := range bagMap[bag] {
		if strings.Contains("shinygold", inside) {
			return true
		} else if getShiny(inside) {
			return true
		}
	}
	return false
}

func castStoI(s string) int {
	// converts string to an int, returns 0 if err
	int64, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int(int64)
}
func intPow(a, n int) int {
	// return a^n
	var i, result int
	result = 1
	for i = 0; i < n; i++ {
		result *= a
	}
	return result
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("07.txt")
	check(err)
	day(dat)

	// EXAMPLE CASE(S)
	// 	day([]byte(
	// 		`shiny gold bags contain 2 dark red bags.
	// dark red bags contain 2 dark orange bags.
	// dark orange bags contain 2 dark yellow bags.
	// dark yellow bags contain 2 dark green bags.
	// dark green bags contain 2 dark blue bags.
	// dark blue bags contain 2 dark violet bags.
	// dark violet bags contain no other bags.`))

}
