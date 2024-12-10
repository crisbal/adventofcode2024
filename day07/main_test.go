package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

func TestPart1(t *testing.T) {
	input := parseInput(example)
	want := 3749
	got := part1(input)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCombine(t *testing.T) {
	cases := []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 12},
		{1, 3, 13},
		{20, 40, 2040},
		{2, 50, 250},
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 10},
	}
	for _, c := range cases {
		got := combine(c.a, c.b)
		if got != c.want {
			t.Errorf("got %d, want %d", got, c.want)
		}
	}
}

func TestPart2(t *testing.T) {
	input := parseInput(example)
	want := 11387
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
