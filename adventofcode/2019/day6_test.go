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
				"COM)B",
				"B)C",
				"C)D",
				"D)E",
				"E)F",
				"B)G",
				"G)H",
				"D)I",
				"E)J",
				"J)K",
				"K)L",
			},
			42,
		},
	}

	for _, c := range cases {
		got := day6Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay6Part1Final(t *testing.T) {
	in, err := readStrings("./day6_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day6Part1(in)
	want := 110190
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
				"COM)B",
				"B)C",
				"C)D",
				"D)E",
				"E)F",
				"B)G",
				"G)H",
				"D)I",
				"E)J",
				"J)K",
				"K)L",
				"K)YOU",
				"I)SAN",
			},
			4,
		},
	}

	for _, c := range cases {
		got := day6Part2(c.in)
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

	got := day6Part2(in)
	want := 343
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 6, part 2 answer: %v\n", got)
}
