package main

import "testing"

var example = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func TestParts(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, int) int
		want int
	}{
		{"part1 example", part1, 40},
		{"part2 example", part2, 25272},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn(example, 10)
			if got != tt.want {
				t.Fatalf("got %d, want %d", got, tt.want)
			}
		})
	}
}
