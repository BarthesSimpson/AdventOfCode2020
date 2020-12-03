package main

import (
	"fmt"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	treeMap := parseMap()
	numTrees := countCollisions(treeMap, 3, 1)
	fmt.Println(numTrees)
}

func part2() {
	treeMap := parseMap()
	numTrees := (countCollisions(treeMap, 1, 1) *
		countCollisions(treeMap, 3, 1) *
		countCollisions(treeMap, 5, 1) *
		countCollisions(treeMap, 7, 1) *
		countCollisions(treeMap, 1, 2))
	fmt.Println(numTrees)
}

// TODO: convert to a bitmap if I can be bothered
func parseMap() [][]int {
	treeMap := make([][]int, 0)
	utils.ApplyToEachLine("3/input.txt", func(line string) {
		row := make([]int, 0)
		for _, c := range line {
			if c == '#' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		treeMap = append(treeMap, row)
	})
	return treeMap
}

func countCollisions(treeMap [][]int, x, y int) int {
	mapWidth := len(treeMap[0])
	slope := float32(x) / float32(y)
	numCollisions := 0
	for row, col := range treeMap {
		if (row % y) != 0 {
			continue
		}
		_x := int(slope*float32(row)) % mapWidth
		numCollisions += col[_x]
	}
	return numCollisions
}
