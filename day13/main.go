package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Problem struct {
	x1, y1, x2, y2 int
	tx, ty         int
}

type Input struct {
	problems []Problem
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	problems := []Problem{}
	problemsStr := strings.Split(input, "\n\n")
	for _, problemStr := range problemsStr {
		problem := Problem{}
		lines := strings.Split(problemStr, "\n")
		// Button A: X+94, Y+34
		// Button B: X+22, Y+67
		// Prize: X=8400, Y=5400
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &problem.x1, &problem.y1)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &problem.x2, &problem.y2)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &problem.tx, &problem.ty)
		problems = append(problems, problem)
	}
	return Input{problems}
}

func part1(input Input) int {
	res := 0
	const A_COST = 3
	const B_COST = 1
	for _, problem := range input.problems {
		bPresses := (problem.x1*problem.ty - problem.y1*problem.tx) / (problem.y2*problem.x1 - problem.x2*problem.y1)
		aPresses := (problem.tx - problem.x2*bPresses) / (problem.x1)
		// check if the integer results are correct, otherwise it means that the solutions are actually float
		if problem.tx != problem.x1*aPresses+problem.x2*bPresses || problem.ty != problem.y1*aPresses+problem.y2*bPresses {
			continue
		}
		res += aPresses*A_COST + bPresses*B_COST
	}
	return res
}

func part2(input Input) int {
	res := 0
	const A_COST = 3
	const B_COST = 1
	for _, problem := range input.problems {
		problem.tx += 10000000000000
		problem.ty += 10000000000000
		bPresses := (problem.x1*problem.ty - problem.y1*problem.tx) / (problem.y2*problem.x1 - problem.x2*problem.y1)
		aPresses := (problem.tx - problem.x2*bPresses) / (problem.x1)
		// check if the integer results are correct, otherwise it means that the solutions are actually float
		if problem.tx != problem.x1*aPresses+problem.x2*bPresses || problem.ty != problem.y1*aPresses+problem.y2*bPresses {
			continue
		}
		res += aPresses*A_COST + bPresses*B_COST
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
