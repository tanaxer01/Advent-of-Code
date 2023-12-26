package main

import (
	"AdventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const DAY string = "Day 1"

func PartA() {
	var lines []string = utils.ReadLines("inputs/day01.txt")

	var total int = 0
	for _, line := range lines {
		total += LineNum(line)
	}

	fmt.Printf("\t part A: %d\n", total)
}

func PartB() {
	var lines []string = utils.ReadLines("inputs/day01.txt")

	var total int = 0
	for _, line := range lines {
		modLine := ReplaceStrNums(line)
        total += LineNum(modLine)
	}

	fmt.Printf("\t part B: %d\n", total)
}

func LineNum(line string) int {
	r, _ := regexp.Compile("[0-9]")
	all := r.FindAllString(line, -1)

	n, _ := strconv.Atoi(all[0] + all[len(all)-1])

	return n
}

func ReplaceStrNums(line string) string {
	nums := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

    for old, new := range nums {
        tempLine := strings.ReplaceAll(line, old, new)
        line = tempLine
    }

	return line
}

func main() {
	fmt.Println(">> " + DAY)
	PartA()
	PartB()
}
