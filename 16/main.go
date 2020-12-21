package main

import (
	"math"
	"math/bits"
	"regexp"
	"sort"
	"strconv"

	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	_, otherTickets, rules := parseInput()
	validRanges := mergeAllRanges(rules)
	errorRate, _ := filterValidTickets(otherTickets, validRanges)
	fmt.Println(errorRate)
}

func part2() {
	ticket, otherTickets, rules := parseInput()
	validRanges := mergeAllRanges(rules)
	_, validTickets := filterValidTickets(otherTickets, validRanges)
	fieldMap := makeFieldMap(rules)
	filterRules(rules, validTickets, fieldMap)
	res := 1
	for _, r := range rules {
		if !strings.HasPrefix(r.Field, "departure") {
			continue
		}
		idx := math.Log2(float64(fieldMap[r.Field]))
		res *= ticket[int(idx)]
	}
	fmt.Println(res)
}

func filterRules(rules []rule, tickets [][]int, fieldMap map[string]uint32) {
	for _, t := range tickets {
		for i, n := range t {
			for _, r := range rules {
				if n < r.Min1 || (n > r.Max1 && n < r.Min2) || n > r.Max2 {
					(fieldMap)[r.Field] &^= uint32(1) << i
					if getHammingWeight((fieldMap)[r.Field]) == 1 {
						clearOthers(fieldMap, r.Field)
					}
				}
			}
		}
	}
}

func printMap(fieldMap map[string]uint32) {
	for r, n := range fieldMap {
		fmt.Println(r)
		fmt.Println(strconv.FormatUint(uint64(n), 2))
	}
}

func clearOthers(fieldMap map[string]uint32, key string) {
	for k := range fieldMap {
		if k != key {
			pre := getHammingWeight((fieldMap)[k])
			(fieldMap)[k] &^= (fieldMap)[key]
			post := getHammingWeight((fieldMap)[k])
			if pre != 1 && post == 1 {
				clearOthers(fieldMap, k)
			}
		}
	}
}

func getHammingWeight(i uint32) int {
	return bits.OnesCount32(i)
}

func makeFieldMap(rules []rule) map[string]uint32 {
	fieldMap := make(map[string]uint32)
	for _, r := range rules {
		fieldMap[r.Field] = (uint32(1) << len(rules)) - 1
	}
	return fieldMap
}

func mergeAllRanges(rules []rule) [][]int {
	ranges := make([][]int, 0)
	for _, r := range rules {
		insertRange(&ranges, []int{r.Min1, r.Max1})
		insertRange(&ranges, []int{r.Min2, r.Max2})
	}
	return ranges
}

func insertRange(ranges *[][]int, r []int) {
	i := sort.Search(len(*ranges), func(i int) bool { return (*ranges)[i][0] >= r[0] })
	// r is the first member of ranges
	if len(*ranges) == 0 {
		*ranges = append(*ranges, r)
		return
	}

	overlapLeft := i > 0 && (*ranges)[i-1][1] > r[0]
	overlapRight := i < len(*ranges)-1 && (*ranges)[i+1][0] < r[1]

	if !overlapLeft && !overlapRight {
		*ranges = append(*ranges, []int{})
		copy((*ranges)[i+1:], (*ranges)[i:])
		(*ranges)[i] = r
		return
	}

	if overlapLeft && overlapRight {
		left := (*ranges)[i-1][0]
		if r[0] < left {
			left = r[0]
		}
		right := (*ranges)[i][1]
		if r[1] > right {
			right = r[1]
		}
		(*ranges)[i-1][0] = left
		(*ranges)[i-1][1] = right
		copy((*ranges)[i:], (*ranges)[i+1:])
		(*ranges) = (*ranges)[:len(*ranges)-1]
		return
	}

	if overlapLeft {
		(*ranges)[i-1][1] = r[1]
	}

	if overlapRight {
		(*ranges)[i][1] = r[1]
	}
}

func filterValidTickets(tickets, validRanges [][]int) (errorRate int, validTickets [][]int) {
	validTickets = make([][]int, 0)
	for _, t := range tickets {
		isValidTicket := true
		for _, n := range t {
			isValid := false
			for _, r := range validRanges {
				if n >= r[0] && n <= r[1] {
					isValid = true
					break
				}
			}
			if !isValid {
				errorRate += n
				isValidTicket = false
			}
		}
		if isValidTicket {
			validTickets = append(validTickets, t)
		}
	}
	return errorRate, validTickets
}

var fieldRegex = regexp.MustCompile(`(?m)^([\w|\s]*):\s(\d+)-(\d+)\sor\s(\d+)-(\d+)`)
var ticketsRegex = regexp.MustCompile(`(?m)^([\d|,]+)`)

type rule struct {
	Field string
	Min1  int
	Max1  int
	Min2  int
	Max2  int
}

func parseInput() ([]int, [][]int, []rule) {
	fileBytes, err := ioutil.ReadFile("16/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	file := string(fileBytes)
	ticketsMatch := ticketsRegex.FindAllStringSubmatch(file, -1)
	ticket := makeTicket(ticketsMatch[0][1])
	otherTickets := make([][]int, len(ticketsMatch)-1)
	for i := 1; i < len(ticketsMatch); i++ {
		otherTickets[i-1] = makeTicket(ticketsMatch[i][1])
	}
	fieldmatch := fieldRegex.FindAllStringSubmatch(file, -1)
	rules := make([]rule, len(fieldmatch))
	for i := 0; i < len(fieldmatch); i++ {
		m := fieldmatch[i]
		min1, _ := strconv.Atoi(m[2])
		max1, _ := strconv.Atoi(m[3])
		min2, _ := strconv.Atoi(m[4])
		max2, _ := strconv.Atoi(m[5])
		rules[i] = rule{m[1], min1, max1, min2, max2}
	}
	return ticket, otherTickets, rules
}

func makeTicket(line string) []int {
	str := strings.Split(line, ",")
	ticket := make([]int, len(str))
	for i, s := range str {
		n, _ := strconv.Atoi(s)
		ticket[i] = n
	}
	return ticket
}
