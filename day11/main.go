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
	numbers []int
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	numbers := []int{}
	for _, strNum := range strings.Split(input, " ") {
		num, _ := strconv.Atoi(strNum)
		numbers = append(numbers, num)
	}
	return Input{
		numbers: numbers,
	}
}

type Cache struct {
	data map[[2]int]int
}

func (c *Cache) Get(key [2]int) (int, bool) {
	res, ok := c.data[key]
	return res, ok
}
func (c *Cache) Set(key [2]int, value int) {
	c.data[key] = value
}

var cache = Cache{
	data: make(map[[2]int]int),
}

func blinkNum(num int, steps int) int {
	if steps == 0 {
		return 1
	}

	cached, ok := cache.Get([2]int{num, steps})
	if ok {
		return cached
	}

	if num == 0 {
		res := blinkNum(1, steps-1)
		cache.Set([2]int{num, steps}, res)
		return res
	}
	strNum := strconv.Itoa(num)
	if len(strNum)%2 == 0 {
		half := len(strNum) / 2
		left := strNum[:half]
		leftNum, _ := strconv.Atoi(left)
		leftRes := blinkNum(leftNum, steps-1)
		cache.Set([2]int{leftNum, steps - 1}, leftRes)

		right := strNum[half:]
		rightNum, _ := strconv.Atoi(right)
		rightRes := blinkNum(rightNum, steps-1)
		cache.Set([2]int{rightNum, steps - 1}, rightRes)

		return leftRes + rightRes
	}
	res := blinkNum(num*2024, steps-1)
	cache.Set([2]int{num, steps}, res)
	return res
}

func part1(input Input) int {
	numbers := input.numbers
	tot := 0
	for _, num := range numbers {
		tot += blinkNum(num, 25)
	}
	return tot
}

func part2(input Input) int {
	numbers := input.numbers
	tot := 0
	for _, num := range numbers {
		tot += blinkNum(num, 75)
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
