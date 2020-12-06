package main

import (
	"fmt"
	"math/bits"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	groups := parseGroups()
	count := 0
	for _, g := range groups {
		count += getHammingWeight(combineGroupV1(g))
	}
	fmt.Println(count)
}

func part2() {
	groups := parseGroups()
	count := 0
	for _, g := range groups {
		count += getHammingWeight(combineGroupV2(g))
	}
	fmt.Println(count)
}

func getHammingWeight(i uint32) int {
	return bits.OnesCount32(i)
}

func combineGroupV1(g []uint32) uint32 {
	var c uint32
	for _, p := range g {
		c |= p
	}
	return c
}

func combineGroupV2(g []uint32) uint32 {
	c := g[0]
	for _, p := range g {
		c &= p
	}
	return c
}

// Each person is represented as an unsigned 32 bit integer, with one bit for each question.
// Each group is represented as a slice of these integers
func parseGroups() [][]uint32 {
	groups := make([][]uint32, 0)
	var group []uint32
	utils.ApplyToEachLine("6/input.txt", func(line string) {
		if len(line) == 0 {
			groups = append(groups, group)
			group = nil
			return
		}
		group = append(group, parsePerson(line))
	})
	if len(group) != 0 {
		groups = append(groups, group)
	}
	return groups
}

// A person is a 32 bit u_int. The first 6 bits (left to right) are ignored,
// and then the remaining bytes correspond to a-z in the questions.
// A zero bit means the answer was no, and a 1 bit means yes
func parsePerson(s string) uint32 {
	var person uint32
	for _, c := range s {
		person |= 1 << int('z'-c)
	}
	return person
}
