package main

import (
	"fmt"
	"strings"
)

type cube struct {
	x, y, z int
}

func newCube(x, y, z int) cube {
	c := cube{x, y, z}
	return c
}

type void struct{}

var on = void{}

func main() {

	// dat, _ := ioutil.ReadFile("inputs/13.txt")
	inArr := strings.Split(
		`.#.
..#
###`, "\n")

	cubes := make(map[cube]void)
	for x, row := range inArr {
		for y, val := range row {
			if val == '#' {
				cubes[newCube(x, y, 0)] = on
			}
		}
	}
	fmt.Println(cubes)

	if _, ok := cubes[newCube(2, 1+1, 0)]; ok {
		fmt.Println("success")
	}
	cubes = make(map[cube]void)
	fmt.Println("after reinstantiating", cubes)

	// var busses []aocu.CrtEntry
	// for i, bus := range strings.Split(busArr[1], ",") {
	// 	busNum, err := strconv.Atoi(bus)
	// 	if err != nil {
	// 		// fmt.Println(i, bus)
	// 		continue
	// 	}
	// 	busses = append(busses, aocu.CrtEntry{A: big.NewInt(int64(busNum - i)), N: big.NewInt(int64(busNum))})

	// }
	// fmt.Println(aocu.SolveCrtMany(busses).String())
	// var waypoint = complex(10, 1)
	// rotation := map[int]complex128{0: 1 + 0i, 1: 0 - 1i, 2: -1 + 0i, 3: 0 + 1i}

	// dirMap := map[int]string{0: "E", 1: "S", 2: "W", 3: "N"}
	// facing := 0
	// line := "F99"
	// if line[0] == 'F' {
	// 	fmt.Println(dirMap[facing] + line[1:])
	// }

	// directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	// for _, dir := range directions {
	// 	for x, y := range dir {
	// 		fmt.Println(x, y)
	// 	}
	// }

	// instr := "acc -1"
	// inInt := castStoI(instr[4:])
	// fmt.Println(inInt)

	// var bagMap = make(map[string]map[string]int)
	// bagMap["shinygold"] = map[string]int{}
	// bagMap["shinygold"]["darkred"] = 2
	// bagMap["darkred"] = nil

	// if v, _ := bagMap["darkred"]; v != nil {
	// 	fmt.Println("not ok", bagMap["darkred"])
	// } else {
	// 	fmt.Println("is ok", bagMap["darkred"])
	// }

	// fmt.Println(bagMap)

	// reHgt := regexp.MustCompile(`(^[\d]{2,3})[ic][nm]$`)
	// test := "150in"
	// answer := reHgt.MatchString(test)
	// units := test[len(test)-2:]
	// measure, _ := strconv.Atoi(test[:len(test)-2])
	// if units == "in" {
	// 	fmt.Println(answer, measure, units)
	// }

}
