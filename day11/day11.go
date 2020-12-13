package main

import (
	"fmt"
	"strings"
)

var inArr []string
var part1, part2 int
var floor, empty, occupied rune = '.', 'L', '#' // 46="." 76="L" 35="#"
var sb strings.Builder

func day(input []byte) {
	inArr = strings.Split(string(input), "\n")
	// size := len(inArr[0]) - 1

	count := 0
	nexCount := 1
	for count != nexCount {
		count = nexCount
		nexCount = 0
		nexArr := make([]string, len(inArr))
		for i, row := range inArr {
			for j, seat := range row {
				switch seat {
				case floor:
					sb.WriteRune(floor)
				case empty:
					//Rule 1: if empty (L) and no occupied seats addjacent, => occupy
					if seatsTaken(i, j) == 0 {
						sb.WriteRune(occupied)
						nexCount++
					} else {
						sb.WriteRune(empty)
					}
				case occupied:
					//Rule 2: if occupied (#) and four or more adjacent sear are also # => empty
					if seatsTaken(i, j) > 4 {
						sb.WriteRune(empty)
					} else {
						sb.WriteRune(occupied)
						nexCount++
					}
				}
			}
			nexArr[i] = sb.String()
			sb.Reset()
		}
		inArr = nexArr
	}
	fmt.Println("part1:", nexCount)
	fmt.Println("part2:", part2)
}
func seatsSeen(r, c int) int {
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	seatsSeenCount := 0
	for _, xinc, yinc := range directions {
		for x := r + xinc; x > -1 || x < len(inArr); x += xinc {
			for y := c + yinc; y > -1 || len(inArr[x]); y += yinc {
				if inArr[x][y] == byte(occupied) {
					seatsSeenCount++
				}
			}
		}
	}
	fmt.Println(directions)
	return seatsSeenCount
}

func seatsTaken(r, c int) int {
	takenSeatCount := 0
	for x := r - 1; x < r+2; x++ {
		if x == -1 || x == len(inArr) {
			continue
		}
		for y := c - 1; y < c+2; y++ {
			if y == -1 || y == len(inArr[x]) {
				continue
			}
			if inArr[x][y] == byte(occupied) {
				takenSeatCount++
			}
		}
	}
	return takenSeatCount
}

func main() {
	// dat, _ := ioutil.ReadFile("inputs/11.txt")
	// day(dat)

	// EXAMPLE CASE(S)
	//
	day([]byte(
		`L.LL.LL.LL
	LLLLLLL.LL
	L.L.L..L..
	LLLL.LL.LL
	L.LL.LL.LL
	L.LLLLL.LL
	..L.L.....
	LLLLLLLLLL
	L.LLLLLL.L
	L.LLLLL.LL`))
}
