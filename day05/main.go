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

type Rules map[int][]int
type Update = []int

type Input struct {
	rules   Rules
	updates []Update
}

func isValid(update Update, rules Rules) (bool, int, int) {
	numToPos := make(map[int]int)
	for i, num := range update {
		numToPos[num] = i
	}
	for _, num := range update {
		numRules := rules[num]
		for _, otherNum := range numRules {
			if _, ok := numToPos[otherNum]; !ok {
				continue
			}
			if numToPos[otherNum] < numToPos[num] {
				return false, numToPos[num], numToPos[otherNum]
			}
		}
	}
	return true, -1, -1
}

func part1(input Input) int {
	res := 0
	for _, update := range input.updates {
		validUpdate, _, _ := isValid(update, input.rules)
		if validUpdate {
			if len(update)%2 == 0 {
				panic("update has even length?")
			}
			halfIndex := len(update) / 2
			res += update[halfIndex]
		}
	}
	return res
}

func part2(input Input) int {
	// First find all the invalid updates
	invalidUpdates := make([]Update, 0)
	for _, update := range input.updates {
		validUpdate, _, _ := isValid(update, input.rules)
		if !validUpdate {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	// Fix the invalid updates
	// Use a very simple algorithm, just swap the first two numbers that are invalid, and check again
	for _, update := range invalidUpdates {
		for {
			validUpdate, index1, index2 := isValid(update, input.rules)
			if validUpdate {
				break
			}
			update[index1], update[index2] = update[index2], update[index1]
		}
	}

	res := 0
	for _, update := range invalidUpdates {
		if len(update)%2 == 0 {
			panic("update has even length?")
		}
		halfIndex := len(update) / 2
		res += update[halfIndex]
	}
	return res
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, "\n\n")

	rulesStrings := strings.Split(strings.TrimSpace(parts[0]), "\n")
	rules := make(Rules, len(rulesStrings))
	for _, ruleString := range rulesStrings {
		ruleParts := strings.Split(ruleString, "|")
		ruleLeft, _ := strconv.Atoi(ruleParts[0])
		ruleRight, _ := strconv.Atoi(ruleParts[1])
		if _, ok := rules[ruleLeft]; !ok {
			rules[ruleLeft] = []int{}
		}
		rules[ruleLeft] = append(rules[ruleLeft], ruleRight)
	}

	updatesString := strings.Split(strings.TrimSpace(parts[1]), "\n")
	updates := make([]Update, len(updatesString))
	for i, updateString := range updatesString {
		update := Update{}
		for _, updatePart := range strings.Split(updateString, ",") {
			num, _ := strconv.Atoi(updatePart)
			update = append(update, num)
		}
		updates[i] = update
	}
	return Input{
		rules:   rules,
		updates: updates,
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
