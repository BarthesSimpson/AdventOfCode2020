package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run ./template <day number>")
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Usage: go run ./template <day number>")
	}

	if err := os.Mkdir(os.Args[1], 0753); err != nil {
		log.Fatal("Failed to create directory", day)
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%d/main.go", day), templateCode, 0644); err != nil {
		log.Fatalf("Failed to create %d/main.go", day)
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%d/input.txt", day), []byte(""), 0644); err != nil {
		log.Fatalf("Failed to create %d/input.txt", day)
	}

}

var templateCode = []byte(
	`package main

import (
	"github.com/BarthesSimpson/AdventOfCode2020/utils"
)

func main() {
	utils.RouteToPart(part1, part2)
}

func part1() {

}

func part2() {

}
`)
