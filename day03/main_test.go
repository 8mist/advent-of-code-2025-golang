package main

import "testing"

var example = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestParts(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string) int64
		want int64
	}{
		{"part1 example", part1, 357},
		{"part2 example", part2, 3121910778619},
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
