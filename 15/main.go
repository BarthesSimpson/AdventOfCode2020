package main

import (
	"fmt"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

var input = []int{8, 0, 17, 4, 1, 12}

func part1() {
	fmt.Println(runTillTarget(2020))
}
func part2() {
	fmt.Println(runTillTarget(30000000))
}

func runTillTarget(target int) int {
	lastSeen := make(map[int]int)
	for i, n := range input {
		lastSeen[n] = i + 1
	}
	num := 0
	start := len(input) + 1
	for i := start; i < target; i++ {
		if idx, ok := lastSeen[num]; ok {
			lastSeen[num] = i
			num = i - idx
			continue
		}
		lastSeen[num] = i
		num = 0
	}
	return num
}
