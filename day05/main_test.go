package main

import "testing"

var example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestParts(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string) int64
		want int64
	}{
		{"part1 example", part1, 3},
		{"part2 example", part2, 14},
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
