package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed example.txt
var example string

func TestPart1(t *testing.T) {
	lines := strings.Split(example, "\n")
	input := parseInput(lines[0])
	want := 161
	got := part1(input)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines := strings.Split(example, "\n")
	input := parseInput(lines[1])
	want := 48
	got := part2(input)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := parseInput(example)
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := parseInput(example)
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
