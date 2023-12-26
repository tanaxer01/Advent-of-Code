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
    // Seed to soil
    idx, options := ParseMap(lines, idx+2)
    seedsInt = ApplyMap(options, seedsInt)

    // Soil to fertilizer
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMap(options, seedsInt)

    // Fertilizer to water
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMap(options, seedsInt)

    // Water to light
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMap(options, seedsInt)

    // Light to temperature
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMap(options, seedsInt)

    // Temperature to humidity 
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMap(options, seedsInt)

    // Humidity to location
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMap(options, seedsInt)

    var min int = seedsInt[0] 
    for _, i := range seedsInt {
        if min > i {
            min = i
        }
    }

	fmt.Printf("\t part A: %d\n", min)
}

func PartB() {
	var lines []string = utils.ReadLines("inputs/test05.txt")

    seedsStr := strings.Split(lines[0][7:], " ")
    var seedsInt [][]int

    for i := 0; i < len(seedsStr); i += 2 {
        n1, _ := strconv.Atoi(seedsStr[i])
        n2, _ := strconv.Atoi(seedsStr[i+1])

        //for i := 0; i < n2; i++ {
        //    seedsInt = append(seedsInt, n1+i)
        //}
        seedsInt = append(seedsInt, []int{n1, n2})
    }

    var idx int = 1
    // Seed to soil
    idx, options := ParseMap(lines, idx+2)
    //seedsInt = ApplyMapRange(options, seedsInt)
    seedsInt = ApplyMapRange(options, seedsInt)
 
    // Soil to fertilizer
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMapRange(options, seedsInt)

    // Fertilizer to water
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMapRange(options, seedsInt)

    // Water to light
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMapRange(options, seedsInt)

    // Light to temperature
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMapRange(options, seedsInt)

    // Temperature to humidity 
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMapRange(options, seedsInt)

    // Humidity to location
    idx, options = ParseMap(lines, idx+2)
    seedsInt = ApplyMapRange(options, seedsInt)

    /*
    var min int = seedsInt[0] 
    for _, i := range seedsInt {
        if min > i {
            min = i
        }
    }
    */

    fmt.Printf("\t part B: %d\n", 0)

}

func ParseMap(lines []string, offset int) (int, [][]int) {// {{{
    var idx int = offset
    var options [][]int

    for ;lines[idx] != "" && idx < len(lines) - 1;idx++ {
        partsStr := strings.Split(lines[idx], " ")
        partsInt := make([]int, len(partsStr)) 

        for i := 0; i < len(partsStr); i++ {
	        n, _ := strconv.Atoi(partsStr[i])
            partsInt[i] = n
        }

        options = append(options, partsInt)
    }

    return idx, options 
}// }}}

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

    for _, seedRange := range input {
        fmt.Print(seedRange, ": ")

        for _, option := range options {
            if seedRange[0] >= option[1] && seedRange[0] + seedRange[1] <= option[1] + option[2] {
                var a int = seedRange[0] - option[1] + option[0]
                var b int = seedRange[0] 

                if a + b > option[0] + option[2] {

                }
                fmt.Print(option, "--", a, "-", b, " ")
                output = append(output, []int{a, b})
            }
        }
        fmt.Println()
    }
    fmt.Println("-->", output)

    return output
}

func main() {
	fmt.Println(">> " + DAY)
	PartA()
	PartB()
}
