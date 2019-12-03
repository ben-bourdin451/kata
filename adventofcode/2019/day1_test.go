package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{"12"}, 2},
		{[]string{"14"}, 2},
		{[]string{"1969"}, 654},
		{[]string{"100756"}, 33583},
	}

	for _, c := range cases {
		got := day1Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay1Part1Final(t *testing.T) {
	in, err := readStrings("./day1_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day1Part1(in)
	want := 3147032
	fmt.Printf("Day 1, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay1Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{"14"}, 2},
		{[]string{"1969"}, 966},
		{[]string{"100756"}, 50346},
	}

	for _, c := range cases {
		got := day1Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay1Part2Final(t *testing.T) {
	in, err := readStrings("./day1_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day1Part2(in)
	want := 4717699
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 1, part 2 answer: %v\n", got)
}
