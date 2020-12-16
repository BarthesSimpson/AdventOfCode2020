package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	depart, buses, _ := parseInput()
	bus, wait := getBusAndWait(depart, buses)
	fmt.Println(bus * wait)
}

func part2() {
	_, ibuses, offsets := parseInput()
	buses := make([]int, len(ibuses))
	for i, b := range ibuses {
		buses[i] = int(b)
	}
	ts := findSeq(buses, offsets)
	fmt.Println(ts)
}

func getBusAndWait(depart int, buses []int) (int, int) {
	bus, wait := 0xFFFFFFFF, 0xFFFFFFFF
	for _, b := range buses {
		mod := depart % b
		if mod == 0 {
			return b, 0
		}
		bwait := b - mod
		if bwait < wait {
			bus, wait = b, bwait
		}
	}
	return bus, wait
}

func findSeq(buses, offsets []int) int {
	ts := 0
	var m int = 1
	for i, bus := range buses {
		offset := offsets[i]
		for (ts+offset)%bus != 0 {
			ts += m
		}
		m *= int(bus)
	}
	return ts
}

func parseInput() (depart int, buses []int, offsets []int) {
	lines := utils.SplitFileBy("13/input.txt", "\n")
	depart, _ = strconv.Atoi(lines[0])
	strbus := strings.Split(lines[1], ",")
	var offset = -1
	for _, b := range strbus {
		offset++
		if b == "x" {
			continue
		}
		offsets = append(offsets, int(offset))
		bus, _ := strconv.Atoi(b)
		buses = append(buses, bus)
	}
	return depart, buses, offsets
}
