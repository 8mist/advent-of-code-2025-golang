package main

import "testing"

var example = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224," +
	"1698522-1698528,446443-446449,38593856-38593862,565653-565659," +
	"824824821-824824827,2121212118-2121212124"

func TestParts(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string) int64
		want int64
	}{
		{"part1 example", part1, 1227775554},
		{"part2 example", part2, 4174379265},
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
