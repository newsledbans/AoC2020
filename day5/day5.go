package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day(input []byte) {

	inArray := strings.Split(string(input), "\n")
	// BFFFBBFRRR: row 70, column 7, seat ID 567.
	// FFFBBBFRRR: row 14, column 7, seat ID 119.
	// BBFFBBFRLL: row 102, column 4, seat ID 820.
	//   0  1  2  3 4 5 6 7  8 9 10
	// 128 64 32 16 8 4 2 1  4 2 1
	result := make(map[int]string)
	highest := 0

	for _, seat := range inArray {
		row, col := 0, 0
		for i, sByte := range seat {
			if sByte == rune(`B`[0]) {
				row += intPow(2, (6 - i))
			} else if sByte == rune(`R`[0]) {
				col += intPow(2, (9 - i))
			}
		}
		seatID := row*8 + col
		result[seatID] = "row:" + fmt.Sprint(row) + " col:" + fmt.Sprint(col)
		if seatID > highest {
			highest = seatID
		}
	}
	fmt.Println("Part1: Id", highest, result[highest])
	for k := range result {
		if _, ok := result[k+1]; !ok {
			if _, ok2 := result[k+2]; ok2 {
				fmt.Println("Part2: id", k+1)
			}
		}
	}
}

func castStoI(s string) int {
	int64, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int(int64)
}

// return a^n
func intPow(a, n int) int {
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
	dat, err := ioutil.ReadFile("05.txt")
	check(err)
	day(dat)

	// EXAMPLE CASE(S)
	// 	day([]byte(
	// 		`BFFFBBFRRR
	// FFFBBBFRRR
	// BBFFBBFRLL`))

}
