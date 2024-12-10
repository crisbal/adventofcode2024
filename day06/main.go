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
	mapp [][]rune
}

var DIRECTIONS = [][]int{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func findGuard(mapp [][]rune) (int, int) {
	for y, row := range mapp {
		for x, cell := range row {
			if cell == '^' {
				return y, x
			}
		}
	}
	return -1, -1
}

// tempWallY and tempWallX are the coordinates of the wall that we will place to create a loop, set to -1 if we don't want to create a loop
// reason why those exists is because we don't want to alter the original mapp, so we can easily run this function in parallel
func walk(mapp [][]rune, startY int, startX int, startDirectionIndex int, tempWallY int, tempWallX int) (map[[2]int]bool, bool) {
	currDirectionIndex := startDirectionIndex
	currY := startY
	currX := startX
	// preallocate the maps to avoid resizing, 6000 is decided based on observations on the input
	posVisited := make(map[[2]int]bool, 6000)
	posVisitedWithDir := make(map[[3]int]bool, 6000)
	for {
		if posVisitedWithDir[[3]int{currY, currX, currDirectionIndex}] {
			return posVisited, true
		}
		posVisited[[2]int{currY, currX}] = true
		posVisitedWithDir[[3]int{currY, currX, currDirectionIndex}] = true
		// check if we can move forward
		nextY := currY + DIRECTIONS[currDirectionIndex][0]
		nextX := currX + DIRECTIONS[currDirectionIndex][1]
		if nextY < 0 || nextY >= len(mapp) || nextX < 0 || nextX >= len(mapp[0]) {
			break
		}
		if mapp[nextY][nextX] == '#' || (nextY == tempWallY && nextX == tempWallX) {
			// turn right
			currDirectionIndex = (currDirectionIndex + 1) % 4
		} else {
			// move forward
			currY = nextY
			currX = nextX
		}
	}
	return posVisited, false
}

func part1(input Input) int {
	currDirectionIndex := 0
	currY, currX := findGuard(input.mapp)
	posVisited, isLoop := walk(input.mapp, currY, currX, currDirectionIndex, -1, -1)
	if isLoop {
		panic("loop should not happen")
	}
	return len(posVisited)
}

func part2(input Input) int {
	startDirectionIndex := 0
	startY, startX := findGuard(input.mapp)
	posVisited, isLoop := walk(input.mapp, startY, startX, startDirectionIndex, -1, -1)
	if isLoop {
		panic("loop should not happen")
	}

	// run each possible wall placement, test in parallel
	results := make(chan bool, len(posVisited))
	for pos := range posVisited {
		go func(pos [2]int) {
			tempWallY, tempWallX := pos[0], pos[1]
			_, isLoop := walk(input.mapp, startY, startX, startDirectionIndex, tempWallY, tempWallX)
			results <- isLoop
		}(pos)
	}

	res := 0
	for range posVisited {
		if <-results {
			res++
		}
	}
	return res
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	mapp := make([][]rune, len(lines))
	for i, line := range lines {
		mapp[i] = []rune(line)
	}
	return Input{
		mapp: mapp,
	}
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
