package aocutils

import (
	"math/big"
	"strconv"
)

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
	intParsed, err := strconv.ParseInt(s, 10, 0)
	Check(err)
	return int(intParsed)
}

//Check checks for error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// This file is for stuff related to the chinese remainder theorem!
// borrowed from https://github.com/deanveloper/modmath

// SolveCrt x=a mod m; x=b mod n by using the chinese remainder theorem.
func SolveCrt(a, m, b, n *big.Int) *big.Int {
	gcd := new(big.Int)
	s := new(big.Int)
	t := new(big.Int)
	gcd.GCD(s, t, m, n)

	// let eqn = bsm, eqn2 = ant
	eqn := new(big.Int)
	eqn2 := new(big.Int)
	eqn.Mul(b, s)
	eqn.Mul(eqn, m)
	eqn2.Mul(a, n)
	eqn2.Mul(eqn2, t)

	// now, let eqn = bsm + ant, eqn2 = m * n
	eqn.Add(eqn, eqn2)
	eqn2.Mul(m, n)
	return eqn.Mod(eqn, eqn2)
}

// CrtEntry Represents an entry in the Extended Chinese Remainder Theorem
type CrtEntry struct {
	A, N *big.Int
}

// SolveCrtMany the solution to x=(a1 mod m1); x=(a2 mod m2); x=...
//
// If len(eqs) == 0, it panics.
func SolveCrtMany(eqs []CrtEntry) *big.Int {
	if len(eqs) == 0 {
		panic("cannot have 0 entries to solve")
	}
	if len(eqs) == 1 {
		return new(big.Int).Mod(eqs[0].A, eqs[0].N)
	}
	eqs2 := make([]CrtEntry, len(eqs))
	copy(eqs2, eqs)
	return solveCrtManyIntern(eqs2)
}

func solveCrtManyIntern(eqs []CrtEntry) *big.Int {
	f := eqs[0]
	s := eqs[1]
	x := SolveCrt(f.A, f.N, s.A, s.N)
	if len(eqs) == 2 {
		return x
	}
	eqs[1] = CrtEntry{x, new(big.Int).Mul(f.N, s.N)}
	return solveCrtManyIntern(eqs[1:])
}
