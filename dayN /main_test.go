package main

import "testing"

var example = ``

func TestPart1(t *testing.T) {
	got := part1(example)
	want := 0

	if got != want {
		t.Fatalf("part1: got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(example)
	want := 0

	if got != want {
		t.Fatalf("part2: got %d, want %d", got, want)
	}
}
