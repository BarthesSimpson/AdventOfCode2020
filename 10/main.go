package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	adapters := parseAdapters()
	sort.Ints(adapters)
	var d1 int
	var d3 int
	for i, a1 := range adapters[1:] {
		a2 := adapters[i]
		if a1-a2 == 1 {
			d1++
		} else {
			d3++
		}
	}
	fmt.Println(d1 * (d3 + 1))
}

func part2() {
	adapters := parseAdapters()
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	lookup := make(map[int]int)
	for i, n := range adapters {
		lookup[n] = i
	}
	dp := make([]int, len(adapters))
	dp[0] = 1
	for i, a1 := range adapters[1:] {
		if j, ok := lookup[a1-1]; ok && j < (i+1) {
			dp[i+1] += dp[j]
		}
		if j, ok := lookup[a1-2]; ok && j < (i+1) {
			dp[i+1] += dp[j]
		}
		if j, ok := lookup[a1-3]; ok && j < (i+1) {
			dp[i+1] += dp[j]
		}
	}
	fmt.Println(dp[len(dp)-1])
}

func parseAdapters() []int {
	strs := utils.SplitFileBy("10/input.txt", "\n")
	nums := make([]int, len(strs))
	for i, str := range strs {
		num, _ := strconv.Atoi(str)
		nums[i] = num
	}
	return nums
}
