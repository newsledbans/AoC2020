package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var inArr []string
var part1, part2 int

type cube struct {
	x, y, z int
}

type void struct{}

var on = void{}

func newCube(x, y, z int) cube {
	c := cube{x, y, z}
	return c
}

func day(input []byte) {

	//Parsing Input
	inArr = strings.Split(string(input), "\n")
	cubes := make(map[cube]void)
	for i, row := range inArr {
		for j, val := range row {
			if val == '#' {
				cubes[newCube(j, i, 0)] = on
			}
		}
	}

	// PART 1
	for iter := 0; iter < 6; iter++ {
		nextcubes := make(map[cube]void)

		for x := -6; x < 14; x++ {
			for y := -6; y < 14; y++ {
				for z := -6; z < 7; z++ {
					activeCount := 0
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							for dz := -1; dz <= 1; dz++ {
								if (dx == 0) && (dy == 0) && (dz == 0) {
									continue
								}
								if _, ok := cubes[newCube(x+dx, y+dy, z+dz)]; ok {
									activeCount++
								}
							}
						}
					}
					if activeCount == 3 || activeCount == 2 {
						if _, ok := cubes[newCube(x, y, z)]; ok {
							nextcubes[newCube(x, y, z)] = on
							continue
						}
					}
					if activeCount == 3 {
						nextcubes[newCube(x, y, z)] = on
					}
				}
			}
		}
		cubes = nextcubes
	}
	part1 = len(cubes)
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func main() {
	dat, _ := ioutil.ReadFile("inputs/17.txt")
	day(dat)

	// EXAMPLE CASE(S)
	// answer: 112
	// 	day([]byte(`.#.
	// ..#
	// ###`))
}
