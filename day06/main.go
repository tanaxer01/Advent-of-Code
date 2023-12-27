package main

import (
	"AdventOfCode/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const DAY string = "Day 6"

func PartA() {
	var lines []string = utils.ReadLines("inputs/day06.txt")

	r, _ := regexp.Compile("[0-9]+")

    timesStr := r.FindAllString(lines[0], -1)
    distancesStr := r.FindAllString(lines[1], -1)

    var total int = 1
    for i := 0; i < len(timesStr); i++ {
        n1, _ := strconv.Atoi(timesStr[i])
        n2, _ := strconv.Atoi(distancesStr[i])

        total *= calcBetterTimes(n1, n2)
    }

	fmt.Printf("\t part A: %d\n", total)
}

func PartB() {
	var lines []string = utils.ReadLines("inputs/day06.txt")

	r, _ := regexp.Compile("[0-9]+")

    timeStr := r.FindAllString(lines[0], -1)
    distanceStr := r.FindAllString(lines[1], -1)
    
    timeInt, _ := strconv.Atoi(strings.Join(timeStr, ""))
    distanceInt, _ := strconv.Atoi(strings.Join(distanceStr, ""))

    total := calcBetterTimesFaster(float64(timeInt), float64(distanceInt)+1)


	fmt.Printf("\t part B: %d\n", total)
}
	
func calcBetterTimes(duration int, best int) int {
    var total int = 0
    for i := 1; i < duration; i++ {
        var distance int = i * (duration - i)

        if distance > best {
            total++
        }
    }

    return total
}

func calcBetterTimesFaster(time float64, distance float64) int {
    b1 := math.Ceil( (time - math.Sqrt(math.Pow(time,2) - 4*distance))/2 )
    b2 := math.Floor( (time + math.Sqrt(math.Pow(time,2) - 4*distance))/2 )

    return int(b2 - b1 + 1)
}

func main() {
	fmt.Println(">> " + DAY)
	PartA()
	PartB()
}
