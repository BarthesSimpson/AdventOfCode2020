package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	instructions := parseInstructions()
	var nand, or uint64
	mem := make(map[uint64]uint64)
	for _, inst := range instructions {
		if inst.Cmd == "mask" {
			nand, or, _ = parseMask([]rune(inst.Mask))
			continue
		}
		mem[inst.Addr] = or | (inst.Val & (^nand))
	}
	var sum uint64
	for _, n := range mem {
		sum += n
	}
	fmt.Println(sum)
}

func part2() {
	instructions := parseInstructions()
	var or uint64
	flt := map[int]bool{}
	mem := map[uint64]uint64{}
	for _, inst := range instructions {
		if inst.Cmd == "mask" {
			_, or, flt = parseMask([]rune(inst.Mask))
			continue
		}
		addrs := getAddresses(inst.Addr, or, flt)
		for addr := range addrs {
			mem[addr] = inst.Val
		}
	}
	var sum uint64
	for _, n := range mem {
		sum += n
	}
	fmt.Println(sum)
}

func getAddresses(addr uint64, or uint64, flt map[int]bool) map[uint64]bool {
	root := or | addr
	addrs := map[uint64]bool{root: true}
	for i := range flt {
		toAdd := make(map[uint64]bool)
		for addr := range addrs {
			toAdd[addr|1<<i] = true
			toAdd[addr & ^(1<<i)] = true
		}
		for addr := range toAdd {
			addrs[addr] = true
		}
	}
	return addrs
}

func parseMask(mask []rune) (nand, or uint64, flt map[int]bool) {
	flt = map[int]bool{}
	for i := 0; i < len(mask); i++ {
		c := mask[len(mask)-(1+i)]
		if c == '0' {
			nand |= (1 << i)
		} else if c == '1' {
			or |= (1 << i)
		} else {
			flt[i] = true
		}
	}
	return nand, or, flt
}

type instruction struct {
	Cmd  string
	Mask string
	Addr uint64
	Val  uint64
}

var memRegex = regexp.MustCompile(`^mem\[(\d+)\]\s=\s(\d+)$`)

func parseInstructions() []instruction {
	instructions := make([]instruction, 0)
	utils.ApplyToEachLine("14/input.txt", func(line string) {
		memMatch := memRegex.FindStringSubmatch(line)
		var inst instruction
		if len(memMatch) == 0 {
			comp := strings.Split(line, " = ")
			inst = instruction{Cmd: comp[0], Mask: comp[1]}
		} else {
			addr, _ := strconv.Atoi(memMatch[1])
			val, _ := strconv.Atoi(memMatch[2])
			inst = instruction{Cmd: "mem", Addr: uint64(addr), Val: uint64(val)}
		}
		instructions = append(instructions, inst)
	})
	return instructions
}
