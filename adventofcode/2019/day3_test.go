package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay3Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"R8,U5,L5,D3",
				"U7,R6,D4,L4",
			},
			6,
		},
		{
			[]string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			},
			159,
		},
		{
			[]string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			},
			135,
		},
	}

	for _, c := range cases {
		got := day3Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay3Part1Final(t *testing.T) {
	in, err := readStrings("./day3_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day3Part1(in)
	want := 260
	fmt.Printf("Day 3, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay3Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"R8,U5,L5,D3",
				"U7,R6,D4,L4",
			},
			30,
		},
		{
			[]string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			},
			610,
		},
		{
			[]string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			},
			410,
		},
	}

	for _, c := range cases {
		got := day3Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay3Part2Final(t *testing.T) {
	in, err := readStrings("./day3_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day3Part2(in)
	want := 15612
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 3, part 2 answer: %v\n", got)
}
