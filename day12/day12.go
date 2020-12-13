package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inArr []string
var part1, part2 int
var xpos, ypos int
var waypoint = complex(10, 1)

func day(input []byte) {
	inArr = strings.Split(string(input), "\n")
	// dirMap := map[int]string{0: "E", 1: "S", 2: "W", 3: "N"}
	rotation := map[int]complex128{0: 1 + 0i, 1: 0 + 1i, 2: -1 + 0i, 3: 0 - 1i}
	// ydir = []int{0, -1, 0, 1}

	for _, line := range inArr {
		facing := 0
		amt, _ := strconv.Atoi(line[1:])
		switch line[0] {
		case 'E', 'S', 'W', 'N':
			dirInc(line)
		case 'F':
			// curDir := dirMap[facing] + line[1:]
			// dirInc(curDir)
			xpos += amt * int(real(waypoint))
			ypos += amt * int(imag(waypoint))
			fmt.Println(line, "new pos", xpos, ypos)
		case 'L':
			facing += amt / 90
			facing %= 4
			if facing < 0 {
				facing += 4
			}
			waypoint *= rotation[facing]
			fmt.Println(line, "new waypoint", waypoint)
		case 'R':
			facing -= amt / 90
			facing %= 4
			if facing < 0 {
				facing += 4
			}
			waypoint *= rotation[facing]
			fmt.Println(line, "new waypoint", waypoint, amt, facing, rotation[facing])
		}
	}
	part1 = intAbs(xpos) + intAbs(ypos)

	fmt.Println("part1:", part1, xpos, ypos)
	fmt.Println("part2:", part2)
}

func dirInc(op string) {
	inc, _ := strconv.ParseFloat(op[1:], 64)
	switch op[0] {
	case 'N':
		waypoint += complex(0, inc)
	case 'S':
		waypoint -= complex(0, inc)
	case 'E':
		waypoint += complex(inc, 0)
	case 'W':
		waypoint -= complex(inc, 0)
	}
	fmt.Println(op, "new waypoint", waypoint)
}

func intAbs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func main() {
	dat, _ := ioutil.ReadFile("inputs/12.txt")
	day(dat)

	// EXAMPLE CASE(S)

	// day([]byte(
	// 		`F10
	// N3
	// F7
	// R90
	// F11`))

}
