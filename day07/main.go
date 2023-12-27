package main

import (
	"AdventOfCode/utils"
	"fmt"
	"sort"
	"strings"
)


const DAY string = "Day 7"

func PartA() {
	var lines []string = utils.ReadLines("inputs/test07.txt")

    for _, line := range lines {
        parts := strings.Split(line, " ")
        score := RateHand(parts[0])

        fmt.Println(parts, score)

    }


	fmt.Printf("\t part A: %d\n", 0)
}

func PartB() {
	//var lines []string = utils.ReadLines("inputs/day06.txt")
	fmt.Printf("\t part B: %d\n", 0)
}

func RateHand(hand string) int {
    freqMap := make(map[rune]int)
    for _, card := range hand {
        freqMap[card]++
    }

    var freqArr []int
    for _, val := range freqMap {
        freqArr = append(freqArr, val)
    }

    sort.Sort(sort.Reverse( sort.IntSlice(freqArr) ) )

    switch freqArr[0] {
    case 5:
        return 7
    case 4:
        return 6
    case 3:
        if freqArr[1] == 2 {
            return 5
        } else {
            return 4
        }
    case 2:
        if freqArr[1] == 2 {
            return 3
        } else {
            return 2
        }
    default:
        return 0
    }
}


func main() {
	fmt.Println(">> " + DAY)
	PartA()
	PartB()
}
