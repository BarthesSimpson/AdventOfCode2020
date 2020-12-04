package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

// Passport represents a passport.
// All fields are represented as strings for convenience.
type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func parsePassport(s string) *Passport {
	p := Passport{}
	queue := make([]rune, 0)
	var label string
	for _, c := range s {
		if c == ' ' || c == '\n' {
			p.setField(label, string(queue))
			queue = nil
			continue
		}
		if c == ':' {
			label = string(queue)
			queue = nil
			continue
		}
		queue = append(queue, c)
	}
	if len(queue) != 0 {
		p.setField(label, string(queue))
	}
	return &p
}

func (p *Passport) setField(k, v string) {
	switch k {
	case "byr":
		p.byr = v
	case "iyr":
		p.iyr = v
	case "eyr":
		p.eyr = v
	case "hgt":
		p.hgt = v
	case "hcl":
		p.hcl = v
	case "ecl":
		p.ecl = v
	case "pid":
		p.pid = v
	case "cid":
		p.cid = v
	default:
		log.Fatal(fmt.Sprintf("encountered invalid key %s", k))
	}
}

func (p *Passport) isValid() bool {
	return (p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != "")
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
func (p *Passport) isValidV2() bool {
	if !validateYear(p.byr, 1920, 2002) {
		return false
	}
	if !validateYear(p.iyr, 2010, 2020) {
		return false
	}
	if !validateYear(p.eyr, 2020, 2030) {
		return false
	}
	if !validateHeight(p.hgt) {
		return false
	}
	if !validateHairColor(p.hcl) {
		return false
	}
	if !validateEyeColor(p.ecl) {
		return false
	}
	if !validatePassportID(p.pid) {
		return false
	}
	return true
}

func validateYear(year string, min, max int) bool {
	if len(year) != 4 {
		return false
	}
	v, _ := strconv.Atoi(year)
	return v >= min && v <= max
}

var hgtRegex = regexp.MustCompile(`^(\d+)(\w+)$`)

func validateHeight(hgt string) bool {
	match := hgtRegex.FindStringSubmatch(hgt)
	if len(match) < 2 {
		return false
	}
	num, _ := strconv.Atoi(match[1])
	unit := match[2]
	if unit == "cm" {
		return num >= 150 && num <= 193
	}
	if unit == "in" {
		return num >= 59 && num <= 76
	}
	return false
}

var hclRegex = regexp.MustCompile(`^#[0-9,a-f]{6}$`)

func validateHairColor(hcl string) bool {
	return hclRegex.Match([]byte(hcl))
}

var validEcls map[string]bool = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func validateEyeColor(ecl string) bool {
	_, ok := validEcls[ecl]
	return ok
}

var pidRegex = regexp.MustCompile(`^\d{9}$`)

func validatePassportID(pid string) bool {
	return pidRegex.Match([]byte(pid))
}

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	passports := parsePassports()
	count := 0
	for _, p := range passports {
		if p.isValid() {
			count++
		}
	}
	fmt.Println(count)
}

func part2() {
	passports := parsePassports()
	count := 0
	for _, p := range passports {
		if p.isValidV2() {
			count++
		}
	}
	fmt.Println(count)
}

func parsePassports() []*Passport {
	passportStrings := utils.SplitFileBy("4/input.txt", "\n\n")
	parsedPassports := make([]*Passport, 0)

	for _, ps := range passportStrings {
		parsedPassports = append(parsedPassports, parsePassport(ps))
	}
	return parsedPassports
}
