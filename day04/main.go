package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Input struct {
	puzzle [][]rune
}

func searchPuzzle(puzzle [][]rune, word string, startx, starty int, dirx, diry int) bool {
	for i := 0; i < len(word); i++ {
		x := startx + i*dirx
		y := starty + i*diry
		if y < 0 || y >= len(puzzle) || x < 0 || x >= len(puzzle[0]) {
			return false
		}
		if puzzle[y][x] != rune(word[i]) {
			return false
		}
	}
	return true
}

func part1(input Input) int {
	word := "XMAS"
	nFound := 0
	directions := [][]int{
		{1, 0},   // right
		{0, 1},   // down
		{-1, 0},  // left
		{0, -1},  // up
		{1, 1},   // down-right
		{1, -1},  // up-right
		{-1, 1},  // down-left
		{-1, -1}, // up-left
	}
	for y := 0; y < len(input.puzzle); y++ {
		for x := 0; x < len(input.puzzle[y]); x++ {
			for _, dir := range directions {
				if searchPuzzle(input.puzzle, word, x, y, dir[0], dir[1]) {
					nFound++
				}
			}
		}
	}
	return nFound
}

func part2(input Input) int {
	nFound := 0
	for y := 0; y < len(input.puzzle); y++ {
		for x := 0; x < len(input.puzzle[y]); x++ {
			if input.puzzle[y][x] == 'A' {
				if y-1 < 0 || y+1 >= len(input.puzzle) || x-1 < 0 || x+1 >= len(input.puzzle[y]) {
					continue
				}
				isMas1 := (input.puzzle[y-1][x-1] == 'M' && input.puzzle[y+1][x+1] == 'S') || (input.puzzle[y-1][x-1] == 'S' && input.puzzle[y+1][x+1] == 'M')
				isMas2 := (input.puzzle[y-1][x+1] == 'M' && input.puzzle[y+1][x-1] == 'S') || (input.puzzle[y-1][x+1] == 'S' && input.puzzle[y+1][x-1] == 'M')
				if isMas1 && isMas2 {
					nFound++
				}
			}
		}
	}
	return nFound
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	puzzle := make([][]rune, len(lines))
	for i, line := range lines {
		puzzle[i] = []rune(line)
	}
	return Input{puzzle}
}

func main() {
	input := parseInput(input)

	res1 := part1(input)
	fmt.Println(res1)

	res2 := part2(input)
	fmt.Println(res2)
}
