package main

import (
	"fmt"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	grid := parseGrid(false)
	for i := 0; i < 6; i++ {
		grid = runCycle(grid)
	}
	ans := 0
	for range grid {
		ans++
	}
	fmt.Println(ans)
}

func part2() {
	grid := parseGrid(true)
	for i := 0; i < 6; i++ {
		grid = runCycle(grid)
	}
	ans := 0
	for range grid {
		ans++
	}
	fmt.Println(ans)
}

func runCycle(grid map[cuboid]bool) map[cuboid]bool {
	visited := make(map[cuboid]bool)
	toVisit := make([]cuboid, 0)
	next := make(map[cuboid]bool)
	for c := range grid {
		toVisit = append(toVisit, c)
	}
	for len(toVisit) > 0 {
		curr := toVisit[len(toVisit)-1]
		toVisit = toVisit[:len(toVisit)-1]
		cnt := 0
		for _, n := range curr.getNeighbors() {
			if _, neighborActive := grid[n]; neighborActive {
				cnt++
			}
			if _, currActive := grid[curr]; currActive {
				if _, neighborSeen := visited[n]; !neighborSeen {
					toVisit = append(toVisit, n)
				}
			}
		}
		visited[curr] = true
		if _, ok := grid[curr]; ok {
			if cnt == 2 || cnt == 3 {
				next[curr] = true
			}
		} else {
			if cnt == 3 {
				next[curr] = true
			}
		}
	}
	return next
}

type cube struct {
	X int
	Y int
	Z int
}

type hypercube struct {
	X int
	Y int
	Z int
	W int
}

type cuboid interface {
	getNeighbors() []cuboid
}

func (c cube) getNeighbors() []cuboid {
	neigh := make([]cuboid, 0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				neigh = append(neigh, cube{c.X + i, c.Y + j, c.Z + k})
			}
		}
	}
	return neigh
}

func (h hypercube) getNeighbors() []cuboid {
	neigh := make([]cuboid, 0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				for m := -1; m < 2; m++ {
					if i == 0 && j == 0 && k == 0 && m == 0 {
						continue
					}
					neigh = append(neigh, hypercube{h.X + i, h.Y + j, h.Z + k, h.W + m})
				}
			}
		}
	}
	return neigh
}

func parseGrid(use4D bool) map[cuboid]bool {
	grid := make(map[cuboid]bool)
	var y int
	utils.ApplyToEachLine("17/input.txt", func(line string) {
		for x, c := range line {
			if c == '#' {
				if use4D {
					grid[hypercube{x, y, 0, 0}] = true
				} else {
					grid[cube{x, y, 0}] = true
				}

			}
		}
		y++
	})
	return grid
}
