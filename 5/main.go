package main

import (
	"fmt"
	"math"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

// Pass is a boarding pass
type Pass struct {
	row int
	col int
}

func (p *Pass) seatID() int {
	return (8 * p.row) + p.col
}

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	passes := parsePasses()
	var maxID int
	for _, p := range passes {
		seatID := p.seatID()
		if seatID > maxID {
			maxID = seatID
		}
	}
	fmt.Println(maxID)
}

func part2() {
	passes := parsePasses()
	minID, maxID := 1024, 0
	var sum int
	for _, p := range passes {
		seatID := p.seatID()
		sum += seatID
		if seatID > maxID {
			maxID = seatID
		} else if seatID < minID {
			minID = seatID
		}
	}
	// like a GAUSS
	expectedSum := (minID + maxID) * (maxID - minID + 1) / 2
	fmt.Println(expectedSum - sum)
}

func parsePasses() []*Pass {
	passes := make([]*Pass, 0)
	utils.ApplyToEachLine("5/input.txt", func(line string) {
		passes = append(passes, parsePass(line))
	})
	return passes
}

func parsePass(s string) *Pass {
	p := Pass{}
	p.row = binarySearch([]rune(s), 0, 7)
	p.col = binarySearch([]rune(s), 7, 10)
	return &p
}

func binarySearch(s []rune, start, end int) int {
	lo, hi := 0, int(math.Pow(2, float64(end-start)))
	var m int
	var d Direction
	for i := start; i < end; i++ {
		m = lo + ((hi - lo) / 2)
		d = inferDirection(s[i])
		if d == up {
			lo = m
		} else {
			hi = m
		}
	}
	return lo
}

// Direction is an enum that means go higher or lower
type Direction int

const (
	up Direction = iota
	down
)

func inferDirection(r rune) Direction {
	if r == 'B' || r == 'R' {
		return up
	}
	return down
}
