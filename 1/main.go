package main

import (
	"fmt"
	"strconv"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	lookup := make(map[int]interface{})

	utils.ApplyToEachLine("1/input.txt", func(line string) {
		num, _ := strconv.Atoi(line)
		lookup[num] = 1
	})

	for k := range lookup {
		if _, found := lookup[2020-k]; found {
			fmt.Println((2020 - k) * k)
			return
		}
	}
}

func part2() {
	nums := []int{}
	utils.ApplyToEachLine("1/input.txt", func(line string) {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	})

	lookup := make(map[int][]int)
	for i, num1 := range nums {
		for j, num2 := range nums[i+1:] {
			if num1+num2 <= 2020 {
				lookup[num1+num2] = []int{i, 1 + i + j}
			}
		}
	}

	for i, k := range nums {
		if idxs, found := lookup[2020-k]; found {
			if utils.IntArrayContains(idxs, i) {
				continue
			}
			ans := k
			for _, idx := range idxs {
				ans *= nums[idx]
			}
			fmt.Println(ans)
			return
		}
	}
}
