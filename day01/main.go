package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/crisbal/aoc24/util"
)

//go:embed input.txt
var input string

type Input struct {
	leftList  []int
	rightList []int
}

func part1(input Input) int {
	sort.Ints(input.leftList)
	sort.Ints(input.rightList)
	totalDistance := 0
	for i := 0; i < len(input.leftList); i++ {
		totalDistance += util.IntAbs(input.leftList[i] - input.rightList[i])
	}
	return totalDistance
}

func part2(input Input) int {
	counter := make(map[int]int)
	for i := 0; i < len(input.rightList); i++ {
		counter[input.rightList[i]]++
	}
	totalSimilarity := 0
	for i := 0; i < len(input.leftList); i++ {
		totalSimilarity += counter[input.leftList[i]] * input.leftList[i]
	}
	return totalSimilarity
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	leftList := make([]int, len(lines))
	rightList := make([]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "   ")
		leftList[i], _ = strconv.Atoi(parts[0])
		rightList[i], _ = strconv.Atoi(parts[1])
	}
	return Input{leftList: leftList, rightList: rightList}
}

func main() {
	input := parseInput(input)

	res1 := part1(input)
	fmt.Println(res1)

	res2 := part2(input)
	fmt.Println(res2)
}
