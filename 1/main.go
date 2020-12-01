package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part := flag.Int("part", 1, "part 1 or part 2")
	flag.Parse()

	if *part < 1 || *part > 2 {
		log.Fatal("please specify -part=1 or -part=2")
	}

	if *part == 1 {
		part1()
	} else {
		part2()
	}
}

func part1() {
	lookup := make(map[int]interface{})

	applyToEachLine("input.txt", func(line string) {
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
	applyToEachLine("input.txt", func(line string) {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	})

	lookup := make(map[int][]int)
	for i, num1 := range nums {
		for j, num2 := range nums[i+1:] {
			if num1+num2 > 2020 {
				continue
			}
			lookup[num1+num2] = []int{i, 1 + i + j}
		}
	}

	for i, k := range nums {
		if idxs, found := lookup[2020-k]; found {
			if contains(idxs, i) {
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

func contains(arr []int, i int) bool {
	for _, j := range arr {
		if j == i {
			return true
		}
	}
	return false
}

func applyToEachLine(filepath string, op func(string)) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		op(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
