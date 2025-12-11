package main

import "testing"

var example1 = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

var example2 = `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`

func TestPart1(t *testing.T) {
	got := part1(example1)
	want := 5

	if got != want {
		t.Fatalf("part1: got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(example2)
	want := 2

	if got != want {
		t.Fatalf("part2: got %d, want %d", got, want)
	}
}
