package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"

	aocu "github.com/newsledbans/aoc2020/aocutils"
)

var inArr []string
var part1, part2 int

func day(input []byte) {
	inArr = strings.Split(string(input), "\n")
	busArr := strings.Split(inArr[1], ",")

	// PART 1
	earliestD, _ := strconv.Atoi(inArr[0])
	nextD, thisD, busID := earliestD, 0, 1
	for _, bus := range busArr {
		id, nobus := strconv.Atoi(bus)
		if nobus != nil {
			// fmt.Println("error", bus, nobus)
			continue
		}
		thisD = (earliestD/id+1)*id - earliestD
		// fmt.Println("bus", id, thisD, nextD, earliestD/id)
		if thisD < nextD {
			nextD = thisD
			busID = id
		}
	}

	// PART 2
	var busses []aocu.CrtEntry
	for offset, bus := range busArr {
		busInt, x := strconv.Atoi(bus)
		if x == nil {
			busses = append(busses, aocu.CrtEntry{A: big.NewInt(int64(busInt - offset)), N: big.NewInt(int64(busInt))})
		}
	}
	timestamp := aocu.SolveCrtMany(busses).String()

	// timestamp := 523 + 547 + 47 + 37 + 29 + 23 + 19 + 17 + 13
	// for {
	// 	skip := 1
	// 	valid := true
	// 	for offset, bus := range busses {
	// 		if (timestamp+offset)%bus != 0 {
	// 			valid = false
	// 			break
	// 		}
	// 		skip *= bus
	// 	}
	// 	if valid {
	// 		break
	// 	}
	// 	timestamp += skip
	// }
	// fmt.Println("CRT method", crt(busses))

	// FAILED PART TWO, WORKS FOR TEST CASES NOT FOR LARGE INPUT
	// for n := 1; n < 100000000000000; {
	// 	targetTime := buses[biggestBus]*n - biggestBus
	// 	counter := 0
	// 	for offset, bus := range buses {
	// 		if (targetTime+offset)%bus != 0 {break}
	// 		counter++}
	// 	if counter == len(buses) {
	// 		fmt.Println("SUCCESS!", targetTime)
	// 		break}
	// 	n++ }

	fmt.Println("part1:", nextD*busID)
	fmt.Println("part2:", timestamp)
}

// func crt(arr map[int]int) int {
// 	// Compute product of all numbers
// 	var prod, result = 1, 0
// 	for _, val := range arr {
// 		prod = prod * val
// 	}
// 	for i, val := range arr {
// 		pp := prod / val
// 		result = result + i*modInverse(pp, val)*pp
// 		// fmt.Println("modInverse", prod, pp, val)
// 	}
// 	return result % prod
// }
// func modInverse(a, m int) int {
// 	m0, x0, x1 := m, 0, 1

// 	if m == 1 {
// 		return 0
// 	}
// 	for a > 1 {
// 		q := a / m
// 		t := m

// 		m = a % m
// 		a = t
// 		t = x0
// 		x0 = x1 - q*x0
// 		x1 = t
// 	}
// 	if x1 < 0 {
// 		x1 = x1 + m0
// 	}
// 	return x1
// }

func main() {
	dat, _ := ioutil.ReadFile("inputs/13.txt")
	day(dat)

	// EXAMPLE CASE(S)

	// 	day([]byte(
	// 		`939
	// 7,13,x,x,59,x,31,19`)) // part two answer is 1068781
	// 	day([]byte(
	// 		`3417
	// 17,x,13,19`))
	// 	day([]byte(
	// 		`754018
	// 67,7,59,61`))
	// 	day([]byte(
	// 		`779210
	// 67,x,7,59,61`))
	// 	day([]byte(
	// 		`1261476
	// 67,7,x,59,61`))
	// 	day([]byte(
	// 		`1202161486
	// 1789,37,47,1889`))
}
