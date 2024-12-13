package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Input struct {
	field [][]rune
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	field := make([][]rune, len(lines))
	for i, line := range lines {
		field[i] = []rune(line)
	}
	return Input{field}
}

type Perimeter = int
type Area = int
type Region = map[[2]int]bool

func visit(field [][]rune, y int, x int, current rune, visited map[[2]int]bool) (Perimeter, Area, Region) {
	if y < 0 || y >= len(field) || x < 0 || x >= len(field[y]) {
		return 1, 0, nil
	}
	if current != '.' && field[y][x] != current {
		return 1, 0, nil
	}
	if visited[[2]int{y, x}] {
		return 0, 0, nil
	}
	visited[[2]int{y, x}] = true
	perimeter := 0
	area := 1
	region := make(map[[2]int]bool)
	region[[2]int{y, x}] = true
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, direction := range directions {
		dy, dx := direction[0], direction[1]
		p, a, r := visit(field, y+dy, x+dx, field[y][x], visited)
		perimeter += p
		area += a
		for k, v := range r {
			region[k] = v
		}
	}
	return perimeter, area, region
}

func part1(input Input) int {
	visitedPlants := make(map[[2]int]bool)
	tot := 0
	for y, row := range input.field {
		for x, _ := range row {
			if visitedPlants[[2]int{y, x}] {
				continue
			}
			perimeter, area, _ := visit(input.field, y, x, '.', visitedPlants)
			tot += perimeter * area
		}
	}
	return tot
}

func countSides(region Region) int {
	sides := 0
	for point, _ := range region {
		y, x := point[0], point[1]
		// bottom left corner
		if !region[[2]int{y, x - 1}] && !region[[2]int{y - 1, x}] {
			sides++
		}
		// bottom right corner
		if !region[[2]int{y, x + 1}] && !region[[2]int{y - 1, x}] {
			sides++
		}
		// top left corner
		if !region[[2]int{y, x - 1}] && !region[[2]int{y + 1, x}] {
			sides++
		}
		// top right corner
		if !region[[2]int{y, x + 1}] && !region[[2]int{y + 1, x}] {
			sides++
		}

		if region[[2]int{y, x - 1}] && region[[2]int{y - 1, x}] && !region[[2]int{y - 1, x - 1}] {
			sides++
		}
		if region[[2]int{y, x + 1}] && region[[2]int{y - 1, x}] && !region[[2]int{y - 1, x + 1}] {
			sides++
		}
		if region[[2]int{y, x - 1}] && region[[2]int{y + 1, x}] && !region[[2]int{y + 1, x - 1}] {
			sides++
		}
		if region[[2]int{y, x + 1}] && region[[2]int{y + 1, x}] && !region[[2]int{y + 1, x + 1}] {
			sides++
		}
	}
	return sides
}

func part2(input Input) int {
	visitedPlants := make(map[[2]int]bool)
	tot := 0
	for y, row := range input.field {
		for x, _ := range row {
			if visitedPlants[[2]int{y, x}] {
				continue
			}
			_, area, region := visit(input.field, y, x, '.', visitedPlants)
			sides := countSides(region)
			tot += area * sides
		}
	}
	return tot
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
