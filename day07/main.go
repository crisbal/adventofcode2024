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

type Equation struct {
	result   int
	operands []int
}

type Input struct {
	equations []Equation
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	equations := make([]Equation, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ": ")
		result, _ := strconv.Atoi(parts[0])

		operandsStrings := strings.Split(parts[1], " ")
		operands := make([]int, len(operandsStrings))
		for j, operand := range operandsStrings {
			operandNum, _ := strconv.Atoi(operand)
			operands[j] = operandNum
		}
		equations[i] = Equation{
			result:   result,
			operands: operands,
		}
	}
	return Input{
		equations: equations,
	}
}

func canCalculateRecurse(currentResult int, expectedResult int, operands []int, index int) bool {
	if index >= len(operands) {
		return currentResult == expectedResult
	}
	if currentResult > expectedResult {
		return false
	}
	return canCalculateRecurse(currentResult+operands[index], expectedResult, operands, index+1) || canCalculateRecurse(currentResult*operands[index], expectedResult, operands, index+1)
}

func part1(input Input) int {
	res := 0
	for _, equation := range input.equations {
		if canCalculateRecurse(equation.operands[0], equation.result, equation.operands, 1) {
			res += equation.result
		}
	}
	return res
}

func combine(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a * 10
	}
	multiplier := 1
	for temp := b; temp > 0; temp /= 10 {
		multiplier *= 10
	}
	return a*multiplier + b
}

func canCalculateRecurseP2(currentResult int, expectedResult int, operands []int, index int) bool {
	if index >= len(operands) {
		return currentResult == expectedResult
	}
	if currentResult > expectedResult {
		return false
	}
	plusResult := canCalculateRecurseP2(currentResult+operands[index], expectedResult, operands, index+1)
	if plusResult {
		return true
	}
	multResult := canCalculateRecurseP2(currentResult*operands[index], expectedResult, operands, index+1)
	if multResult {
		return true
	}
	concatResult := canCalculateRecurseP2(combine(currentResult, operands[index]), expectedResult, operands, index+1)
	if concatResult {
		return true
	}
	return false
}

func part2(input Input) int {
	res := 0
	for _, equation := range input.equations {
		result := canCalculateRecurseP2(equation.operands[0], equation.result, equation.operands, 1)
		if result {
			res += equation.result
			continue
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
