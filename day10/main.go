package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Input struct {
	topomap [][]int
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	topomap := make([][]int, len(lines))
	for i, line := range lines {
		topomap[i] = make([]int, len(line))
		for j, c := range line {
			if c == '.' {
				topomap[i][j] = -1
				continue
			}
			h, _ := strconv.Atoi(string(c))
			topomap[i][j] = h
		}
	}
	return Input{topomap}
}

func explore(topomap [][]int, prev int, y, x int, visited map[[2]int]bool, visitedVal int) int {
	if y < 0 || y >= len(topomap) || x < 0 || x >= len(topomap[0]) {
		return 0
	}
	if topomap[y][x] != prev+1 {
		return 0
	}
	if topomap[y][x] == 9 {
		if visited[[2]int{y, x}] {
			return visitedVal
		}
		visited[[2]int{y, x}] = true
		return 1
	}
	// for each direction, explore and return the sum of the results
	return explore(topomap, topomap[y][x], y+1, x, visited, visitedVal) +
		explore(topomap, topomap[y][x], y-1, x, visited, visitedVal) +
		explore(topomap, topomap[y][x], y, x+1, visited, visitedVal) +
		explore(topomap, topomap[y][x], y, x-1, visited, visitedVal)
}

func part1(input Input) int {
	res := 0
	for y, row := range input.topomap {
		for x, h := range row {
			if h == 0 {
				visited := make(map[[2]int]bool)
				res += explore(input.topomap, -1, y, x, visited, 0)
			}
		}
	}
	return res
}

func part2(input Input) int {
	res := 0
	for y, row := range input.topomap {
		for x, h := range row {
			if h == 0 {
				visited := make(map[[2]int]bool)
				res += explore(input.topomap, -1, y, x, visited, 1)
			}
		}
	}
	return res
}

func main() {
	startTime := time.Now()
	input := parseInput(input)
	fmt.Println("Parsed input in", time.Since(startTime))

	startTime = time.Now()
	res1 := part1(input)
	fmt.Println("Part 1:", res1, "in", time.Since(startTime))

	startTime = time.Now()
	res2 := part2(input)
	fmt.Println("Part 2:", res2, "in", time.Since(startTime))
}
