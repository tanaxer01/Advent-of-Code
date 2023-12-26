package main

import (
	"AdventOfCode/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

const DAY = "Day 03"

func PartA() {
	var lines []string = utils.ReadLines("inputs/day03.txt")

	r, _ := regexp.Compile(`[0-9]+|[^.0-9]`)

    var symbols [][]int
    var numbers [][]int

    for y, line := range(lines) {
        items := r.FindAllStringIndex(line, -1)

        for _, item := range items {
            _, err := strconv.Atoi(lines[y][item[0]:item[1]])
            if err != nil {
                symbols = append(symbols, append(item, y))
            } else {
                numbers = append(numbers, append(item, y))
            }
        }
    }


    var total int = 0
    for _, num := range numbers {
        for _, sym := range symbols {
            if math.Abs( float64(sym[2]-num[2]) ) >= 2 {
                continue
            }

            if sym[1] == num[0] || sym[0] == num[1] {
                n, _ := strconv.Atoi(lines[num[2]][num[0]:num[1]])
                total += n
            }

            if sym[0] >= num[0] && sym[1] <= num[1]  {
                n, _ := strconv.Atoi(lines[num[2]][num[0]:num[1]])
                total += n
            }
        }
    }


	fmt.Printf("\t part A: %d\n", total)
}

func PartB() {
	var lines []string = utils.ReadLines("inputs/day03.txt")

	r, _ := regexp.Compile(`[0-9]+|\*`)

    var symbols [][]int
    var numbers [][]int

    for y, line := range(lines) {
        items := r.FindAllStringIndex(line, -1)

        for _, item := range items {
            _, err := strconv.Atoi(lines[y][item[0]:item[1]])
            if err != nil {
                symbols = append(symbols, append(item, y))
            } else {
                numbers = append(numbers, append(item, y))
            }
        }
    }
    
    var total int = 0
    for _, sym := range symbols {
        var gear_nums []int

        for _, num := range numbers {
            if math.Abs( float64(sym[2]-num[2]) ) >= 2 {
                continue
            }

            if sym[1] == num[0] || sym[0] == num[1] {
                n, _ := strconv.Atoi(lines[num[2]][num[0]:num[1]])
                gear_nums = append(gear_nums, n)
            }

            if sym[0] >= num[0] && sym[1] <= num[1] {
                n, _ := strconv.Atoi(lines[num[2]][num[0]:num[1]])
                gear_nums = append(gear_nums, n)
            }
        }

        if len(gear_nums) > 1 {
            var temp int = 1
            for _, num := range gear_nums {
                temp *= num
            }
            total += temp
        }
    }


	fmt.Printf("\t part B: %d\n", total)
}

func main() {
	fmt.Println(">> " + DAY)
	PartA()
	PartB()
}
