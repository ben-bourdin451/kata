package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay19Part1(t *testing.T) {
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
		got := day19([]int{0, 0, 0, 0, 0, 0}, c.ip, c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay19ParseOpInstruction(t *testing.T) {
	cases := []struct {
		in   string
		want opInstruction
	}{
		{
			"addi 5 16 5",
			opInstruction{"addi", addi, []int{5, 16, 5}},
		},
		{
			"seti 1 7 3",
			opInstruction{"seti", seti, []int{1, 7, 3}},
		},
	}

	for _, c := range cases {
		got := parseOpInstruction(c.in)

		if got.name != c.want.name || !reflect.DeepEqual(got.vals, c.want.vals) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay19Part1Final(t *testing.T) {
	in, err := readStringsPlain("./day19_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day19([]int{0, 0, 0, 0, 0, 0}, in[0], in[1:])
	want := 878
	fmt.Printf("Day 19, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay19Part2Final(t *testing.T) {
	in, err := readStringsPlain("./day19_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	// got := day19([]int{1, 0, 0, 0, 0, 0}, in[0], in[1:])
	got := day19([]int{10551278, 10551276, 10551277, 10551277, 0, 9}, in[0], in[1:])
	want := 10551278
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 19, part 2 answer: %v\n", got)
}
