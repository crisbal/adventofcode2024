package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Coord struct {
	y int
	x int
}
type Input struct {
	width    int
	height   int
	antennas map[rune][]Coord
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	antennas := make(map[rune][]Coord)
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	for y, line := range lines {
		for x, char := range line {
			if char == '.' || char == '#' {
				continue
			}
			antennas[char] = append(antennas[char], Coord{y: y, x: x})
		}
	}
	return Input{antennas: antennas, width: width, height: height}
}

func part1(input Input) int {
	antinodes := make(map[Coord]bool)
	for _, coords := range input.antennas {
		for i, coord1 := range coords {
			for j, coord2 := range coords {
				if i == j {
					continue
				}
				dx := coord1.x - coord2.x
				dy := coord1.y - coord2.y
				ax := coord1.x + dx
				ay := coord1.y + dy
				if ax < 0 || ax >= input.width || ay < 0 || ay >= input.height {
					continue
				}
				antinodes[Coord{y: ay, x: ax}] = true
			}
		}
	}
	return len(antinodes)
}

func part2(input Input) int {
	antinodes := make(map[Coord]bool)
	for _, coords := range input.antennas {
		for i, coord1 := range coords {
			for j, coord2 := range coords {
				if i == j {
					continue
				}
				c := 0
				for {
					dx := coord1.x - coord2.x
					dy := coord1.y - coord2.y
					ax := coord1.x + dx*c
					ay := coord1.y + dy*c
					if ax < 0 || ax >= input.width || ay < 0 || ay >= input.height {
						break
					}
					antinodes[Coord{y: ay, x: ax}] = true
					c += 1
				}
			}
		}
	}
	return len(antinodes)
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
