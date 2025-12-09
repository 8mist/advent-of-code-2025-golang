package main

import "testing"

var example = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestParts(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string) int
		want int
	}{
		{"part1 example", part1, 50},
		{"part2 example", part2, 24},
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
