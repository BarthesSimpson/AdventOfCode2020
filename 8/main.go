package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	inst := parseInstructions()
	sum, _ := runTillExit(inst)
	fmt.Println(sum)
}

func part2() {
	// Brute force. Can try a smarter solution later.
	inst := parseInstructions()
	for idx, i := range inst {
		if i.Cmd == acc {
			continue
		}
		inst2 := make([]instruction, len(inst))
		copy(inst2, inst)
		if i.Cmd == jmp {
			inst2[idx].Cmd = nop
		} else {
			inst2[idx].Cmd = jmp
		}
		sum, didLoop := runTillExit(inst2)
		if !didLoop {
			fmt.Println(sum)
			return
		}
	}
}

// Runs till the instruction set terminates (either because of a loop or EOF)
func runTillExit(inst []instruction) (sum int, didLoop bool) {
	didRun := make([]bool, len(inst))
	var i int
	for {
		if i == len(inst) {
			fmt.Println("reached EOF")
			return sum, false
		}
		if didRun[i] {
			fmt.Println("detected loop at instruction ", i)
			return sum, true
		}
		didRun[i] = true
		nxt := inst[i]
		if nxt.Cmd == jmp {
			i += nxt.Num
			continue
		}
		if nxt.Cmd == acc {
			sum += nxt.Num
		}
		i++
	}
}

type cmd = int

const (
	acc cmd = iota
	jmp
	nop
)

func toCommand(s string) cmd {
	switch s {
	case "acc":
		return acc
	case "jmp":
		return jmp
	case "nop":
		return nop
	default:
		return nop
	}
}

func toString(c cmd) string {
	switch c {
	case acc:
		return "acc"
	case jmp:
		return "jmp"
	case nop:
		return "nop"
	default:
		return "nop"
	}
}

type instruction struct {
	Cmd cmd
	Num int
}

func (i instruction) String() string {
	return fmt.Sprintf("%s %d", toString(i.Cmd), i.Num)
}

var cmdRegex = regexp.MustCompile(`^(\w+)\s(.*)$`)

func parseInstructions() []instruction {
	inst := make([]instruction, 0)
	utils.ApplyToEachLine("8/input.txt", func(line string) {
		match := cmdRegex.FindStringSubmatch(line)
		num, _ := strconv.Atoi(match[2])
		i := instruction{Cmd: toCommand(match[1]), Num: num}
		inst = append(inst, i)
	})
	return inst
}
