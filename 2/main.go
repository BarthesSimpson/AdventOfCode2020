package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// PasswordRule contains the parameters for validating a password.
// The validation rules are different for part 1 and part 2
type PasswordRule struct {
	letter rune
	min    int
	max    int
}

func (pr *PasswordRule) isValidPwdV1(pwd string) bool {
	count := 0
	for _, c := range pwd {
		if c != pr.letter {
			continue
		}
		count++
		if count > pr.max {
			return false
		}
	}
	return count >= pr.min
}

func (pr *PasswordRule) isValidPwdV2(pwd string) bool {
	count := 0
	if pr.min <= len(pwd) && rune(pwd[pr.min-1]) == pr.letter {
		count++
	}
	if pr.max <= len(pwd) && rune(pwd[pr.max-1]) == pr.letter {
		count++
	}
	return count == 1
}

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
	count := 0
	applyToEachLine("input.txt", func(line string) {
		rule, pwd := parseLine(line)
		if rule.isValidPwdV1(pwd) {
			count++
		}
	})
	fmt.Println(count)
}

func part2() {
	count := 0
	applyToEachLine("input.txt", func(line string) {
		rule, pwd := parseLine(line)
		if rule.isValidPwdV2(pwd) {
			count++
		}
	})
	fmt.Println(count)
}

var lineParser = regexp.MustCompile(`(?P<Min>\d+)-(?P<Max>\d+)\s(?P<Letter>\w):\s(?P<Password>\w+)`)

func parseLine(line string) (PasswordRule, string) {
	match := lineParser.FindStringSubmatch(line)
	min, _ := strconv.Atoi(match[1])
	max, _ := strconv.Atoi(match[2])
	letter := []rune(match[3])[0]
	pwd := match[4]
	return PasswordRule{min: min, max: max, letter: letter}, pwd
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
