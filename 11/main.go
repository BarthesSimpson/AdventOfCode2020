package main

import (
	"fmt"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	seats := parseSeats()
	prev, curr := -1, 0
	rule := Rule{4, countOccupiedAdjacent}
	for prev != curr {
		prev = curr
		curr = applyRules(seats, rule, prev)
	}
	fmt.Println(curr)
}

func part2() {
	seats := parseSeats()
	prev, curr := -1, 0
	rule := Rule{5, countOccupiedAdjacentV2}
	for prev != curr {
		prev = curr
		curr = applyRules(seats, rule, prev)
	}
	fmt.Println(curr)
}

// Rule is a rule for counting adjacent occupied seats and a max
// number of adjacent occupied seats that a seated person will tolerate
type Rule struct {
	MaxAdj   int
	CountAdj func(int, int, *[][]int) int
}

func applyRules(seats *[][]int, rule Rule, bal int) int {
	nxt := makeCopy(seats)
	for r, row := range *seats {
		for c, s := range row {
			if s == -1 {
				continue
			}
			adj := rule.CountAdj(r, c, seats)
			if s == 0 && adj == 0 {
				bal++
				(*nxt)[r][c] = 1
			} else if s == 1 && adj >= rule.MaxAdj {
				bal--
				(*nxt)[r][c] = 0
			}
		}
	}
	copy(*seats, *nxt)
	return bal
}

var neighbors = [][]int{[]int{-1, -1}, []int{-1, 0}, []int{-1, 1},
	[]int{0, 1}, []int{1, 1}, []int{1, 0},
	[]int{1, -1}, []int{0, -1}}

func countOccupiedAdjacent(r, c int, seats *[][]int) int {
	R, C := len(*seats), len((*seats)[0])
	var count int
	for _, n := range neighbors {
		_r, _c := n[0], n[1]
		rr, cc := r+_r, c+_c
		if (rr > -1) && (rr < R) && (cc > -1) && (cc < C) && (*seats)[rr][cc] == 1 {
			count++
		}
	}
	return count
}

func countOccupiedAdjacentV2(r, c int, seats *[][]int) int {
	R, C := len(*seats), len((*seats)[0])
	var count int
	for _, n := range neighbors {
		_r, _c := n[0], n[1]
		rr, cc := r+_r, c+_c
		for (rr > -1) && (rr < R) && (cc > -1) && (cc < C) {
			s := (*seats)[rr][cc]
			if s == 0 {
				break
			} else if s == 1 {
				count++
				break
			}
			rr, cc = rr+_r, cc+_c
			continue
		}
	}
	return count
}

func makeCopy(seats *[][]int) *[][]int {
	nxt := make([][]int, len(*seats))
	for r, row := range *seats {
		_row := make([]int, len(row))
		copy(_row, row)
		nxt[r] = _row
	}
	return &nxt
}

func parseSeats() *[][]int {
	seats := make([][]int, 0)
	utils.ApplyToEachLine("11/input.txt", func(line string) {
		row := make([]int, len(line))
		for i, c := range line {
			if c == '.' {
				row[i] = 2
			}
		}
		seats = append(seats, row)
	})
	return &seats
}
