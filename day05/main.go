package main

import (
	"AdventOfCode/utils"
	"fmt"
	"strconv"
	"strings"
)

const DAY string = "Day 5"

func PartA() {
	var lines []string = utils.ReadLines("inputs/day05.txt")

    seedsStr := strings.Split(lines[0][7:], " ")
    seedsInt := make([]int, len(seedsStr)) 

    for i := 0; i < len(seedsStr); i++ {
        n, _ := strconv.Atoi(seedsStr[i])
        seedsInt[i] = n
    }

    var idx int = 1
    var options [][]int

    for idx < len(lines)-1 {
        idx, options = ParseMap(lines, idx+2)
        seedsInt = ApplyMap(options, seedsInt)
    }

    var min int = seedsInt[0] 
    for _, i := range seedsInt {
        if min > i {
            min = i
        }
    }

	fmt.Printf("\t part A: %d\n", min)
}

func PartB() {
	var lines []string = utils.ReadLines("inputs/day05.txt")

    seedsStr := strings.Split(lines[0][7:], " ")
    var seedsInt [][]int

    for i := 0; i < len(seedsStr); i += 2 {
        n1, _ := strconv.Atoi(seedsStr[i])
        n2, _ := strconv.Atoi(seedsStr[i+1])

        seedsInt = append(seedsInt, []int{n1, n2})
    }

    var idx int = 1
    var options [][]int
    for idx < len(lines)-1 {
        idx, options = ParseMap(lines, idx+2)
        seedsInt = ApplyMapRange(options, seedsInt)
    }


    var min int = seedsInt[0][0]
    for _, i := range seedsInt {
        if min > i[0] {
            min = i[0]
        }
    }

    fmt.Printf("\t part B: %d\n", min)

}

func ParseMap(lines []string, offset int) (int, [][]int) {
    var idx int = offset
    var options [][]int

    for ;lines[idx] != "" && idx < len(lines) - 1; idx++ {
        partsStr := strings.Split(lines[idx], " ")
        partsInt := make([]int, len(partsStr)) 

        for i := 0; i < len(partsStr); i++ {
	        n, _ := strconv.Atoi(partsStr[i])
            partsInt[i] = n
        }

        options = append(options, partsInt)
    }

    return idx, options 
}

func ApplyMap(options [][]int, input []int) []int {
    output := make([]int, len(input))

    for i := 0; i < len(input); i++ {
        var changed bool = false
        for _, opt := range(options) {
            if input[i] >= opt[1] && input[i] < opt[1] + opt[2] {
                output[i] = input[i] - opt[1] + opt[0]
                changed = true
            }
        }

        if !changed {
            output[i] = input[i]
        }
    }

    return output
}

func ApplyMapRange(options [][]int, input [][]int) [][]int {
    var output [][]int

    for ;len(input) != 0; input = input[1:]  {
        curr := input[0]

        var changed bool = false
        for _, option := range options {
            if curr[0] >= option[1] && curr[0] < option[1] + option[2] {
                a := curr[0] - option[1] + option[0]
                b := curr[1]
                changed = true

                if a + b > option[0] + option[2] {
                    extra := (curr[0] + curr[1]) - (option[1] + option[2])
                    b -= extra

                    input = append(input, []int{curr[0] + (curr[1] - extra), extra})
                }

                output = append(output, []int{a,b})
            }
        }

        if !changed {
            output = append(output, curr)
        }
    }
    
    return output
}

func main() {
	fmt.Println(">> " + DAY)
	PartA()
	PartB()
}
