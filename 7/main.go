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
	_, canBeIn := parseGraphs()
	visited := make(map[Bag]bool)
	startNode := Bag{Modifier: "shiny", Color: "gold"}
	stack := []Bag{startNode}
	var nxt Bag
	for len(stack) != 0 {
		nxt, stack = stack[len(stack)-1], stack[:len(stack)-1]
		visited[nxt] = true
		candidates, _ := canBeIn[nxt]
		for c := range candidates {
			if _, seen := visited[c]; !seen {
				stack = append(stack, c)
			}
		}
	}
	fmt.Println(len(visited) - 1)
}

func part2() {
	mustContain, _ := parseGraphs()
	startNode := Bag{Modifier: "shiny", Color: "gold"}
	stack := []tail{tail{Bag: startNode, Num: 1}}
	var totalBags int
	var nxt tail
	for len(stack) != 0 {
		nxt, stack = stack[len(stack)-1], stack[:len(stack)-1]
		totalBags += nxt.Num
		bags := mustContain[nxt.Bag]
		for b, n := range bags {
			num, _ := n.(int)
			newNode := tail{Bag: b, Num: num * nxt.Num}
			stack = append(stack, newNode)
		}
	}
	fmt.Println(totalBags - 1)
}

// Bag contains the info relevant to each bag
type Bag struct {
	Modifier string
	Color    string
}

func (b Bag) String() string {
	return fmt.Sprintf("%s_%s", b.Modifier, b.Color)
}

// AdjacencyList is an adjacency list representation of a graph
// mustContain: {"drab_gold": {"bright_red": 1, "dull_blue": 2}, "bright_red": {"dull_blue": 1}}
// canBeIn: {"drab_gold": {"bright_red": true, "dull_blue": true}, "dull_blue": {"drab_gold": true}}
type AdjacencyList = map[Bag]map[Bag]interface{}

func parseGraphs() (AdjacencyList, AdjacencyList) {
	mustContain := make(AdjacencyList)
	canBeIn := make(AdjacencyList)
	utils.ApplyToEachLine("7/input.txt", func(line string) {
		head, tails := parseLine(line)
		if _, ok := mustContain[head]; !ok {
			mustContain[head] = make(map[Bag]interface{})
		}
		for _, t := range tails {
			if _, ok := canBeIn[t.Bag]; !ok {
				canBeIn[t.Bag] = make(map[Bag]interface{})
			}
			mustContain[head][t.Bag] = t.Num
			canBeIn[t.Bag][head] = true

		}
	})
	return mustContain, canBeIn
}

type tail struct {
	Bag Bag
	Num int
}

var headRegex = regexp.MustCompile(`(\w+)\s(\w+)\sbags\scontain`)
var tailRegex = regexp.MustCompile(`(\d+)\s(\w+)\s(\w+)\sbag`)

func parseLine(line string) (Bag, []tail) {
	headMatch := headRegex.FindStringSubmatch(line)
	head := Bag{Modifier: headMatch[1], Color: headMatch[2]}
	tailMatches := tailRegex.FindAllStringSubmatch(line, -1)

	tails := make([]tail, len(tailMatches))
	for i, t := range tailMatches {
		tailBag := Bag{Modifier: t[2], Color: t[3]}
		num, _ := strconv.Atoi(t[1])
		tails[i] = tail{Bag: tailBag, Num: num}
	}
	return head, tails
}
