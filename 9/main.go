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
	nums := parseNums()
	fmt.Println(findFirstBadNumber(nums))
}

func part2() {
	nums := parseNums()
	target := findFirstBadNumber(nums)
	set := findContiguousSum(nums, target)
	var sum int
	for _, n := range set {
		sum += n
	}
	sort.Ints(set)
	fmt.Println(set[0] + set[len(set)-1])
}

func parseNums() []int {
	lines := utils.SplitFileBy("9/input.txt", "\n")
	nums := make([]int, len(lines))
	for i, line := range lines {
		n, _ := strconv.Atoi(line)
		nums[i] = n
	}
	return nums
}

func findFirstBadNumber(nums []int) int {
	preamble := nums[:25]
	sums := make(map[int]bool)
	for i, n1 := range preamble {
		for _, n2 := range preamble[i+1:] {
			sums[n1+n2] = true
		}
	}

	for i, n := range nums[25:] {
		if _, ok := sums[n]; !ok {
			return n
		}
		for _, n1 := range nums[:25+i] {
			sums[n+n1] = true
		}
	}
	return -1 // lulz
}

func findContiguousSum(nums []int, target int) []int {
	i, j := 0, 1
	sum := nums[i] + nums[j]
	for {
		if sum == target && j > i {
			return nums[i : j+1]
		}
		for sum < target {
			j++
			sum += nums[j]
		}
		for sum > target {
			sum -= nums[i]
			i++
		}
	}
}
