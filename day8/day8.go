package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inMap = make(map[int][]string)
var i, accumulator, nextOp int
var beenRun []int

func day(input []byte) {

	inArr := bytes.Split(input, []byte("\n"))
	for i, instruction := range inArr {
		inMap[i] = strings.Fields(string(instruction))
	}

	for iCh, strCh := range inMap {
		// loop through the instructions, swapping nop/jmp until we hit a solution.
		i, nextOp, accumulator = 0, 0, 0
		beenRun = []int{}

		if strCh[0] == "acc" {
			continue
		}
		for i >= 0 {
			if nextOp == len(inMap) {
				fmt.Println("part2:", accumulator)
				break
			}
			if alreadyRan(nextOp) {
				break
			}

			beenRun = append(beenRun, nextOp)
			if nextOp == iCh {
				tempStr := swapNJ(strCh)
				runOperation(tempStr)
			} else {
				runOperation(inMap[nextOp])
			}
			i++
		}
	}
}

func swapNJ(instr []string) []string {
	switch instr[0] {
	case "jmp":
		newStr := []string{"nop", instr[1]}
		return newStr
	case "nop":
		newStr := []string{"jmp", instr[1]}
		return newStr

	}
	return instr
}

func alreadyRan(opNum int) bool {
	for _, opNum := range beenRun {
		if opNum == nextOp {
			return true
		}
	}
	return false
}

func runOperation(instr []string) {
	// performs operation and returns int for index of next operation
	switch instr[0] {
	case "acc":
		accumulator += castStoI(instr[1])
		nextOp++
		return
	case "jmp":
		nextOp += castStoI(instr[1])
		return
	case "nop":
		nextOp++
		return
	}
}
func castStoI(s string) int {
	// converts string to an int, returns 0 if err
	int64, err := strconv.Atoi(s)
	check(err)
	return int(int64)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("08.txt")
	check(err)
	day(dat)

	// EXAMPLE CASE(S)
	//
	// day([]byte(
	// 	`nop +0
	// acc +1
	// jmp +4
	// acc +3
	// jmp -3
	// acc -99
	// acc +1
	// jmp -4
	// acc +6`))
}
