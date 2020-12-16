package main

import (
	"fmt"
	"strconv"
	"strings"
)

var inArr []string
var part1, part2 int

func day(input []byte) {
	inArr = strings.Split(string(input), ",")
	// order := make([]int, 0)
	spoken := make(map[int]*int)
	consider, zero := 0, 0
	var newConsider *int

	// PART 1
	for i, val := range inArr {
		valint, _ := strconv.Atoi(val)
		if i == len(inArr)-1 {
			consider = valint
			continue
		}
		// order = append(order, valint)
		i2 := i + 1
		spoken[valint] = &i2
	}

	finalNum := 30000000
	for j := len(spoken) + 1; j < finalNum+1; j++ {
		if index, ok := spoken[consider]; ok {
			j2 := j - *index
			newConsider = &j2
		} else {
			newConsider = &zero
		}
		// order = append(order, consider)
		if j == finalNum {
			fmt.Println(consider)
			break
		}
		newIndex := j
		spoken[consider] = &newIndex
		consider = *newConsider
	}

	fmt.Println("part1:")
	fmt.Println("part2:", part2)
}

func main() {
	// dat, _ := ioutil.ReadFile("inputs/13.txt")
	// day(dat)

	// EXAMPLE CASE(S)
	// day([]byte(`0,3,6`))
	day([]byte(`8,0,17,4,1,12`))

}
