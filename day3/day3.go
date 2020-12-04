package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func day3(input []byte) {

	xInc := [5]int{1, 3, 5, 7, 1}
	yInc := [5]int{1, 1, 1, 1, 2}

	var tree = []byte(`#`)[0]
	var results [5]int

	inArray := bytes.Split(input, []byte("\n"))
	// fmt.Println(inArray[0], tree)
	for i := 0; i < 5; i++ {
		treeCount, xPos := 0, 0
		for yPos, line := range inArray {
			if yPos%yInc[i] != 0 {
				continue
			}
			lnString := string(line)
			if xPos >= len(lnString) {
				xPos %= len(lnString)
			}
			if lnString[xPos] == tree {
				treeCount++
			}
			xPos += xInc[i]
			// fmt.Println("lnString", lnString, "; treeCnt: ", treeCount, "; xpos: ", xPos)

		}
		results[i] = treeCount
	}

	finalResult := 1

	for j, rslt := range results {
		fmt.Println("Case:", j+1, "=", rslt)
		finalResult *= rslt
	}
	fmt.Println("Product:", finalResult)

}

func castBtoI(byteSlice []byte) int {
	byteToInt, err := strconv.Atoi(string(byteSlice))
	check(err)
	return byteToInt
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("03.txt")
	check(err)
	day3(dat)

	// EXAMPLE CASE(S)
	// 	day3([]byte(
	// 		`..##.......
	// #...#...#..
	// .#....#..#.
	// ..#.#...#.#
	// .#...##..#.
	// ..#.##.....
	// .#.#.#....#
	// .#........#
	// #.##...#...
	// #...##....#
	// .#..#...#.#`))
	//
}
