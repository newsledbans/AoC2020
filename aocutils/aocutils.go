package aocutils

import "strconv"

//MinMax returns min and max of an []int
func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

//SArrtoIArr converts []string to []int
func SArrtoIArr(sa []string) []int {
	var iArr []int
	for _, v := range sa {
		iArr = append(iArr, StoI(v))
	}
	return iArr
}

//StoI converts string to an int
func StoI(s string) int {
	intL, err := strconv.ParseInt(s, 10, 0)
	Check(err)
	return int(intL)
}

//Check checks for error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
