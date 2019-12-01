package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay3ParseClaim(t *testing.T) {
	cases := []struct {
		in   []string
		want []claim
	}{
		{[]string{
			"#1 @ 1,3: 4x4",
			"#2 @ 3,1: 4x4",
			"#3 @ 5,5: 2x2",
		}, []claim{
			{id: 1, y: 1, x: 3, w: 4, h: 4},
			{id: 1, y: 3, x: 1, w: 4, h: 4},
			{id: 1, y: 3, x: 1, w: 4, h: 4},
		}},
	}
	for _, c := range cases {
		got := day3Part1(&c.in)
		if reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay3Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{
			"#1 @ 1,3: 4x4",
			"#2 @ 3,1: 4x4",
			"#3 @ 5,5: 2x2",
		}, 4},
	}
	for _, c := range cases {
		got := day3Part1(&c.in)
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
	want := 119572
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
		{[]string{
			"#1 @ 1,3: 4x4",
			"#2 @ 3,1: 4x4",
			"#3 @ 5,5: 2x2",
		}, 3},
	}
	for _, c := range cases {
		got := day3Part2(&c.in)
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
	want := 775
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 3, part 2 answer: %v\n", got)
}
