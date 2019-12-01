package adventofcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay5Part1(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"dabAcCaCBAcCcaDA", 10},
		{"dabAaBAaDA", 4},
	}

	for _, c := range cases {
		got := day5Part1(strings.Split(c.in, ""))
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestAreOpposite(t *testing.T) {
	cases := []struct {
		a    string
		b    string
		want bool
	}{
		{"a", "A", true},
		{"A", "a", true},
		{"c", "C", true},
		{"C", "c", true},
		{"d", "e", false},
		{"K", "f", false},
		{"J", "J", false},
		{"J", "P", false},
	}

	for _, c := range cases {
		got := areOpposite(c.a, c.b)
		if got != c.want {
			t.Errorf("for %v, %v: got %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

func TestDay5Part1Final(t *testing.T) {
	in, err := readStrings("./day5_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day5Part1(strings.Split((*in)[0], ""))
	want := 10250
	fmt.Printf("Day 5, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay5Part2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{
			"dabAcCaCBAcCcaDA",
			4,
		},
	}
	for _, c := range cases {
		got := day5Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay5Part2Final(t *testing.T) {
	in, err := readStrings("./day5_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day5Part2((*in)[0])
	want := 6188
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 5, part 2 answer: %v\n", got)
}
