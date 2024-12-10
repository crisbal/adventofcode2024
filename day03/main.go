package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Input struct {
	lines []string
}

func part1(input Input) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	result := 0
	for _, line := range input.lines {
		matches := re.FindAllStringSubmatch(line, -1)
		if matches == nil {
			continue
		}
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			result += (a * b)
		}
	}
	return result
}

func part2(input Input) int {
	result := 0
	enabled := true
	re := regexp.MustCompile(`(do|don't|mul)\((.*?)\)`)
	for _, line := range input.lines {
		matches := re.FindAllStringSubmatch(line, -1)
		if matches == nil {
			continue
		}
		for _, match := range matches {
			switch match[1] {
			case "do":
				enabled = true
			case "don't":
				enabled = false
			default:
				if enabled {
					args := strings.Split(match[2], ",")
					a, _ := strconv.Atoi(args[0])
					b, _ := strconv.Atoi(args[1])
					result += (a * b)
				}
			}
		}
	}
	return result
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	return Input{lines: lines}
}

func main() {
	input := parseInput(input)

	res1 := part1(input)
	fmt.Println(res1)

	res2 := part2(input)
	fmt.Println(res2)
}
