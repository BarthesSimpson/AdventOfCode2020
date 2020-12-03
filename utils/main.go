package utils

import (
	"bufio"
	"flag"
	"log"
	"os"
)

// ApplyToEachLine applies the provided lambda to each line of the specified file
func ApplyToEachLine(filepath string, op func(string)) {
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

// RouteToPart routes to the handler for Part1 (default) or Part2
// based on the --part command line flag
func RouteToPart(part1Handler, part2Handler func()) {
	part := flag.Int("part", 1, "part 1 or part 2")
	flag.Parse()

	if *part < 1 || *part > 2 {
		log.Fatal("please specify -part=1 or -part=2")
	}

	if *part == 1 {
		part1Handler()
	} else {
		part2Handler()
	}
}

// IntArrayContains checks if the int i exists in the int array (slice) arr
func IntArrayContains(arr []int, i int) bool {
	for _, j := range arr {
		if j == i {
			return true
		}
	}
	return false
}
