package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay6Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"1, 1",
				"1, 6",
				"8, 3",
				"3, 4",
				"5, 5",
				"8, 9",
			},
			17,
		},
	}

	for _, c := range cases {
		got := day6Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay6BoundsContains(t *testing.T) {
	cases := []struct {
		b    bounds
		p    *point
		want bool
	}{
		{
			bounds{0, 0, 19, 19},
			&point{0, 1},
			false,
		},
		{
			bounds{0, 0, 19, 19},
			&point{1, 0},
			false,
		},
		{
			bounds{-1, -1, 19, 19},
			&point{0, 1},
			true,
		},
		{
			bounds{-1, -1, 19, 19},
			&point{1, 0},
			true,
		},
		{
			bounds{-1, -1, 20, 20},
			&point{19, 19},
			true,
		},
		{
			bounds{-1, -1, 20, 20},
			&point{19, 0},
			true,
		},
	}

	for _, c := range cases {
		got := c.b.contains(c.p)
		if got != c.want {
			t.Errorf("bounds %v, point %v, got %v, want %v", c.b, *c.p, got, c.want)
		}
	}
}

func TestDay6Part1Final(t *testing.T) {
	in, err := readStrings("./day6_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day6Part1(*in)
	want := 3401
	fmt.Printf("Day 6, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay6Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"1, 1",
				"1, 6",
				"8, 3",
				"3, 4",
				"5, 5",
				"8, 9",
			},
			16,
		},
	}

	for _, c := range cases {
		got := day6Part2(c.in, 32)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay6Part2Final(t *testing.T) {
	in, err := readStrings("./day6_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day6Part2(*in, 10000)
	want := 49327
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 6, part 2 answer: %v\n", got)
}
