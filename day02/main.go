package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/crisbal/aoc24/util"
)

//go:embed input.txt
var input string

type Input struct {
	reports [][]int
}

func isReportValid(report []int) bool {
	// A report is valid if:
	// All numbers are either all increasing or all decreasing
	// Any two adjacent numbers differ by at least one and at most three
	increasing := true
	for i := 1; i < len(report); i++ {
		if report[i] < report[i-1] {
			increasing = false
			break
		}
	}
	decreasing := true
	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1] {
			decreasing = false
			break
		}
	}
	if !increasing && !decreasing {
		return false
	}

	for i := 1; i < len(report); i++ {
		diff := util.IntAbs(report[i] - report[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func part1(input Input) int {
	validReports := 0
	for _, report := range input.reports {
		if isReportValid(report) {
			validReports++
		}
	}
	return validReports
}

func part2(input Input) int {
	validReports := 0
	invalidReports := make([][]int, 0, len(input.reports))
	for _, report := range input.reports {
		if isReportValid(report) {
			validReports++
		} else {
			invalidReports = append(invalidReports, report)
		}
	}
	// For each invalid report see if we can fix it by removing one number
	for _, invalidReport := range invalidReports {
		for i := 0; i < len(invalidReport); i++ {
			newReport := make([]int, 0, len(invalidReport)-1)
			for j := 0; j < len(invalidReport); j++ {
				if i == j {
					continue
				}
				newReport = append(newReport, invalidReport[j])
			}
			if isReportValid(newReport) {
				validReports++
				break
			}
		}
	}
	return validReports
}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	reports := make([][]int, 0, len(lines))
	for _, line := range lines {
		var report []int
		for _, stringNum := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(stringNum)
			report = append(report, num)
		}
		reports = append(reports, report)
	}
	return Input{
		reports: reports,
	}
}

func main() {
	input := parseInput(input)

	res1 := part1(input)
	fmt.Println(res1)

	res2 := part2(input)
	fmt.Println(res2)
}
