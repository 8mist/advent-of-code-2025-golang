package main

import "testing"

var example = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestParts(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string) int
		want int
	}{
		{"part1 example", part1, 3},
		{"part2 example", part2, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn(example)
			if got != tt.want {
				t.Fatalf("got %d, want %d", got, tt.want)
			}
		})
	}
}
