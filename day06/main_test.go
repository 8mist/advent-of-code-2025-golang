package main

import "testing"

var example = `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `

func TestParts(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string) int
		want int
	}{
		{"part1 example", part1, 4277556},
		{"part2 example", part2, 3263827},
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
