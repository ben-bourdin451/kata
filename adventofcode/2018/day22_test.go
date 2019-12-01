package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay22Part1(t *testing.T) {
	cases := []struct {
		depth  int
		target point
		want   int
	}{
		{510, point{10, 10}, 114},
	}

	for _, c := range cases {
		got := day22Part1(c.depth, c.target)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay22Part1Final(t *testing.T) {
	got := day22Part1(11541, point{14, 778})
	want := 11575
	fmt.Printf("Day 22, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay22Part2(t *testing.T) {
	cases := []struct {
		depth  int
		target point
		want   int
	}{
		{510, point{10, 10}, 45},
	}

	for _, c := range cases {
		got := day22Part2(c.depth, c.target)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay22Part2Final(t *testing.T) {
	got := day22Part2(11541, point{14, 778})
	want := 1068
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 22, part 2 answer: %v\n", got)
}
