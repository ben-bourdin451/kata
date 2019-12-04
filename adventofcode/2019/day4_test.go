package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay4Part1(t *testing.T) {
	cases := []struct {
		min  int
		max  int
		want int
	}{
		{
			111111, 111113,
			3,
		},
		{
			111111, 111119,
			9,
		},
	}

	for _, c := range cases {
		got := day4(c.min, c.max, part1)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay4PasswordMatchPart1(t *testing.T) {
	cases := map[int]bool{
		111111: true,
		122345: true,
		111123: true,
		335679: true,
		223450: false,
		123789: false,
	}

	for k, v := range cases {
		got := passwordMatch(k, part1)
		if got != v {
			t.Errorf("%v got %v, want %v", k, got, v)
		}
	}
}

func TestDay4PasswordMatchPart2(t *testing.T) {
	cases := map[int]bool{
		112233: true,
		111122: true,
		778999: true,
		123444: false,
		111111: false,
		111123: false,
		123241: false,
	}

	for k, v := range cases {
		got := passwordMatch(k, part2)
		if got != v {
			t.Errorf("%v got %v, want %v", k, got, v)
		}
	}
}

func TestDay4Part1Final(t *testing.T) {
	got := day4(382345, 843167, part1)
	want := 460
	fmt.Printf("Day 4, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay4Part2Final(t *testing.T) {
	got := day4(382345, 843167, part2)
	want := 290
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 4, part 2 answer: %v\n", got)
}
