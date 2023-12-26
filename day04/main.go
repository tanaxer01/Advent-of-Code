package main

import (
	"AdventOfCode/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const DAY = "Day 04"

func PartA() {
	var lines []string = utils.ReadLines("inputs/day04.txt")

    var total int = 0
    for _, line := range lines {
        games := strings.Split(line,  ": ")[1]
        parts := strings.Split(games, " | ")

        winNums := ParseNums(parts[0])
        ownNums := ParseNums(parts[1])

        sort.Sort( sort.IntSlice(winNums) )
        sort.Sort( sort.IntSlice(ownNums) )

        curr, _ := GetScore(winNums, ownNums)
        total += curr
    }

	fmt.Printf("\t part A: %d\n", total)
}

func PartB() {
	var lines []string = utils.ReadLines("inputs/day04.txt")

    var ocurrences = make([]int, len(lines))

    for y, line := range lines {
        games := strings.Split(line,  ": ")[1]
        parts := strings.Split(games, " | ")
        
        winNums := ParseNums(parts[0])
        ownNums := ParseNums(parts[1])

        sort.Sort( sort.IntSlice(winNums) )
        sort.Sort( sort.IntSlice(ownNums) )

        _, cant := GetScore(winNums, ownNums)
        ocurrences[y]++
        for i := 1; i < cant+1 && y+i < len(lines); i++ {
            ocurrences[y+i] += ocurrences[y]
        }
    }

    var total int = 0
    for i := 0; i < len(lines); i++ {
        total += ocurrences[i] 
    }

	fmt.Printf("\t part B: %d\n", total)
}

func ParseNums(str string) []int {
    var out []int

    for _, i := range strings.Split(str, " ") {
        if len(i) != 0 {
            n, _ := strconv.Atoi(i)
            out = append(out, n)
        }
    }

    return out
}

func GetScore(winNums []int, ownNums []int) (int, int) {
    var curr int = 0
    var cant int = 0

    var i int = 0
    var j int = 0
    for i < len(winNums) && j < len(ownNums) {
        if winNums[i] == ownNums[j] {
            if curr == 0 {
                curr += 1
            } else {
                curr *= 2
            }

            cant += 1

            i += 1
            j += 1
        } else if winNums[i] > ownNums[j] {
            j += 1
        } else {
            i += 1
        }
    }

    return curr, cant
}

func main() {
	fmt.Println(">> " + DAY)
	PartA()
	PartB()
}
