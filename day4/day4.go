package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func day(input []byte) {

	inArray := strings.Split(string(input), "\n\n")
	passCount := 0
	reHgt := regexp.MustCompile(`(^[\d]{2,3})[ic][nm]$`)
	reHcl := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	colors := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	rePid := regexp.MustCompile(`^[0-9]{9}$`)

	for _, passport := range inArray {
		data := strings.Fields(passport)
		dataMap := make(map[string]string)
		dataCount := 0

		// skip passports that don't contain enough data
		if len(data) < 6 {
			fmt.Println("skipped < 6", "\n")
			continue
		}

		for _, item := range data {
			data := strings.Split(item, ":")
			dataMap[data[0]] = data[1]
		}
		// cid (Country ID) - ignored, missing or not.
		delete(dataMap, "cid")
		fmt.Println(dataMap)

		// byr (Birth Year) - four digits; at least 1920 and at most 2002.
		if val, ok := dataMap["byr"]; ok {
			intVal := castStoI(val)
			if intVal >= 1920 && intVal <= 2004 {
				dataCount++
				fmt.Println("byr True", val)
			}
		}
		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		if val, ok := dataMap["iyr"]; ok {
			intVal := castStoI(val)
			if intVal >= 2010 && intVal <= 2020 {
				dataCount++
				fmt.Println("iyr True", val)

			}
		}
		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		if val, ok := dataMap["eyr"]; ok {
			intVal := castStoI(val)
			if intVal >= 2020 && intVal <= 2030 {
				dataCount++
				fmt.Println("eyr True", val)

			}
		}
		// hgt (Height) - a number followed by either cm or in:
		// If cm, the number must be at least 150 and at most 193.
		// If in, the number must be at least 59 and at most 76.
		if val, ok := dataMap["hgt"]; ok {
			if reHgt.MatchString(val) {
				units := val[len(val)-2:]
				measure := castStoI(val[:len(val)-2])
				if units == "in" {
					if measure >= 59 && measure <= 76 {
						dataCount++
						fmt.Println("hgt(in) True", val)
					}
				} else if units == "cm" {
					if measure >= 150 && measure <= 193 {
						dataCount++
						fmt.Println("hgt(cm) True", val)

					}
				}

			}
		}
		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		if val, ok := dataMap["hcl"]; ok {
			if reHcl.MatchString(val) {
				dataCount++
				fmt.Println("hcl True", val)

			}
		}
		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		if val, ok := dataMap["ecl"]; ok {
			for _, color := range colors {
				if color == val {
					dataCount++
					fmt.Println("ecl True", val, color)

				}
			}
		}
		// pid (Passport ID) - a nine-digit number, including leading zeroes.
		if val, ok := dataMap["pid"]; ok {
			if rePid.MatchString(val) {
				dataCount++
				fmt.Println("pid True", val)

			}
		}
		if dataCount > 6 {
			passCount++
			fmt.Println("passport accepted", dataCount)

		}
	}
	fmt.Println("Result:", passCount)
}

func castStoI(s string) int {
	int64, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int(int64)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("04.txt")
	check(err)
	day(dat)

	// EXAMPLE CASE(S)
	// 	day([]byte(
	// 		`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
	// byr:1937 iyr:2017 cid:147 hgt:183cm

	// iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
	// hcl:#cfa07d byr:1929

	// hcl:#ae17e1 iyr:2013
	// eyr:2024
	// ecl:brn pid:760753108 byr:1931
	// hgt:179cm

	// hcl:#cfa07d eyr:2025 pid:166559648
	// iyr:2011 ecl:brn hgt:59in`))

}
