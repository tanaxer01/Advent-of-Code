package main

import (
	"AdventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const DAY = "Day 02"

func PartA() {
	var lines []string = utils.ReadLines("inputs/day02.txt")

	r, _ := regexp.Compile("[0-9]+ (red|green|blue)")

    var total int = 0
	for id, line := range lines {
		cubes := r.FindAllString(line, -1)

        var valid bool = true 
		for _, num := range cubes {
            parts := strings.Split(num, " ") 
	        num, _ := strconv.Atoi(parts[0])

			switch parts[1] {
			case "red":
                valid = valid && (num <= 12)
			case "green":
                valid = valid && (num <= 13)
			case "blue":
                valid = valid && (num <= 14)
			}
		}

        if valid == true {
            total += id+1
        }
	}

	fmt.Printf("\t part A: %d\n", total)
}

func PartB() {
	var lines []string = utils.ReadLines("inputs/day02.txt")

	r, _ := regexp.Compile("[0-9]+ (red|green|blue)")

    var total int = 0
    for _, line := range lines {
		cubes := r.FindAllString(line, -1)

        var red int = 0
        var green int = 0
        var blue int = 0

		for _, num := range cubes {
            parts := strings.Split(num, " ") 
	        num, _ := strconv.Atoi(parts[0])

			switch parts[1] {
			case "red":
                if red < num {
                    red = num
                }

			case "green":
                if green < num {
                    green = num
                }

			case "blue":
                if blue < num {
                    blue = num
                }

			}
		}

        total += red * green * blue
    }

	fmt.Printf("\t part B: %d\n", total)
}

func main() {
	fmt.Println(">> " + DAY)
	PartA()
	PartB()
}
