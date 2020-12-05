package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func inRange(val, min, max int) bool {
	return min <= val && val <= max
}

func main() {
	buffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	passportChunks := strings.Split(string(buffer), "\n\n")
	validPassports := 0
	for _, passportChunk := range passportChunks {
		chunkLines := strings.Split(passportChunk, "\n")
		var fields []string

		for _, chunkLine := range chunkLines {
			fields = append(fields, strings.Split(chunkLine, " ")...)
		}

		passport := map[string]string{
			"byr": "",
			"iyr": "",
			"eyr": "",
			"hgt": "",
			"hcl": "",
			"ecl": "",
			"pid": "",
			"cid": "ignore",
		}

		for _, field := range fields {
			if field == "" {
				continue
			}
			fs := strings.Split(field, ":")
			passport[fs[0]] = fs[1]
		}

		valid := true
		for _, v := range passport {
			if v == "" {
				valid = false
				break
			}
		}

		if !valid {
			continue
		}

		byr, _ := strconv.Atoi(passport["byr"])
		valid = valid && inRange(byr, 1920, 2002)
		iyr, _ := strconv.Atoi(passport["iyr"])
		valid = valid && inRange(iyr, 2010, 2020)
		eyr, _ := strconv.Atoi(passport["eyr"])
		valid = valid && inRange(eyr, 2020, 2030)

		// hgt
		if strings.HasSuffix(passport["hgt"], "cm") {
			hgt, _ := strconv.Atoi(strings.TrimSuffix(passport["hgt"], "cm"))
			valid = valid && inRange(hgt, 150, 193)
		} else if strings.HasSuffix(passport["hgt"], "in") {
			hgt, _ := strconv.Atoi(strings.TrimSuffix(passport["hgt"], "in"))
			valid = valid && inRange(hgt, 59, 76)
		} else {
			valid = false
		}

		// hcl
		matched, _ := regexp.MatchString("^#[0-9a-f]{6}$", passport["hcl"])
		valid = valid && matched

		// ecl
		matched, _ = regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", passport["ecl"])
		valid = valid && matched

		// pid
		matched, _ = regexp.MatchString("^[0-9]{9}$", passport["pid"])
		valid = valid && matched

		if valid {
			validPassports++
			fmt.Println(passport)
		} else {
		}
	}
	fmt.Println(validPassports)
}
