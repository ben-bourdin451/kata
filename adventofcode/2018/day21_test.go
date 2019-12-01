package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay21Part1(t *testing.T) {
	cases := []struct {
		ip   string
		in   []string
		want int
	}{
		{
			"#ip 0",
			[]string{
				"seti 5 0 1",
				"seti 6 0 2",
				"addi 0 1 0",
				"addr 1 2 3",
				"setr 1 0 0",
				"seti 8 0 4",
				"seti 9 0 5",
			},
			6,
		},
	}

	for _, c := range cases {
		got := day21Part1([]int{0, 0, 0, 0, 0, 0}, c.ip, c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay21Part1Final(t *testing.T) {
	in, err := readStringsPlain("./day21_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day21Part1([]int{0, 0, 0, 0, 0, 0}, in[0], in[1:])
	want := 0
	fmt.Printf("Day 21, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay21Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"",
			},
			0,
		},
	}

	for _, c := range cases {
		got := day21Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay21Part2Final(t *testing.T) {
	in, err := readStrings("./day21_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day21Part2(*in)
	want := 0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 21, part 2 answer: %v\n", got)
}
