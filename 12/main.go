package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {
	insts := parseInstructions()
	x, y := runInstructions(insts)
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func part2() {
	insts := parseInstructions()
	x, y := runInstructionsV2(insts)
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func runInstructions(insts []instruction) (x, y int) {
	dir := 0
	for _, inst := range insts {
		handleInst(inst, &x, &y, &dir)
	}
	return x, y
}

func runInstructionsV2(insts []instruction) (x, y int) {
	wp := waypoint{instruction{east, 10}, instruction{north, 1}}
	for _, inst := range insts {
		handleInstV2(inst, &x, &y, &wp)
	}
	return x, y
}

func handleInst(inst instruction, x, y, dir *int) {
	if inst.Cmd == right || inst.Cmd == left {
		rotateDir(inst, dir)
		return
	}
	if inst.Cmd == forward {
		handleInst(instruction{command(*dir), inst.Mag}, x, y, dir)
		return
	}
	applyMove(inst, x, y)
}

func handleInstV2(inst instruction, x, y *int, wp *waypoint) {
	if inst.Cmd == east || inst.Cmd == west {
		wp.CmdX.add(inst)
		return
	}
	if inst.Cmd == north || inst.Cmd == south {
		wp.CmdY.add(inst)
		return
	}
	if inst.Cmd == right || inst.Cmd == left {
		rotateDir(inst, &(wp.CmdX.Cmd))
		rotateDir(inst, &(wp.CmdY.Cmd))
		if wp.CmdX.Cmd == north || wp.CmdX.Cmd == south {
			wp.CmdX, wp.CmdY = wp.CmdY, wp.CmdX
		}
	}
	if inst.Cmd == forward {
		for moves := inst.Mag; moves > 0; moves-- {
			applyMove(wp.CmdX, x, y)
			applyMove(wp.CmdY, x, y)
		}
	}
}

func applyMove(inst instruction, x, y *int) {
	if inst.Cmd == east {
		*x += inst.Mag
		return
	}
	if inst.Cmd == west {
		*x -= inst.Mag
		return
	}
	if inst.Cmd == north {
		*y += inst.Mag
		return
	}
	if inst.Cmd == south {
		*y -= inst.Mag
		return
	}
}

func rotateDir(inst instruction, dir *int) {
	if inst.Cmd == right {
		*dir = (*dir + (inst.Mag / 90)) % 4
		return
	}
	*dir -= (inst.Mag / 90)
	if *dir < 0 {
		*dir = 4 + *dir
	}
	*dir %= 4
}

type command = int

const (
	east command = iota
	south
	west
	north
	forward
	left
	right
	unknown
)

var commands = []string{"E", "S", "W", "N", "F", "L", "R", "nil"}

type instruction struct {
	Cmd command
	Mag int
}

func (inst instruction) String() string {
	return fmt.Sprintf("%s%d", commands[inst.Cmd], inst.Mag)
}

func (inst *instruction) add(inst2 instruction) {
	if inst.Cmd == inst2.Cmd {
		inst.Mag += inst2.Mag
		return
	}
	if math.Abs(float64(inst.Cmd)-float64(inst2.Cmd)) == 2 {
		inst.Mag -= inst2.Mag
	}
}

type waypoint struct {
	CmdX instruction
	CmdY instruction
}

func toCmd(r rune) command {
	switch r {
	case 'E':
		return east
	case 'S':
		return south
	case 'W':
		return west
	case 'N':
		return north
	case 'F':
		return forward
	case 'L':
		return left
	case 'R':
		return right
	default:
		return unknown
	}
}

func parseInstructions() []instruction {
	insts := make([]instruction, 0)
	utils.ApplyToEachLine("12/input.txt", func(line string) {
		num, _ := strconv.Atoi(line[1:])
		inst := instruction{toCmd([]rune(line)[0]), num}
		insts = append(insts, inst)
	})
	return insts
}
