package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay7Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{
			[]string{
				"Step C must be finished before step A can begin.",
				"Step C must be finished before step F can begin.",
				"Step A must be finished before step B can begin.",
				"Step A must be finished before step D can begin.",
				"Step B must be finished before step E can begin.",
				"Step D must be finished before step E can begin.",
				"Step F must be finished before step E can begin.",
			},
			"CABDFE",
		},
		{
			[]string{
				"Step C must be finished before step A can begin.",
				"Step B must be finished before step E can begin.",
				"Step A must be finished before step D can begin.",
				"Step C must be finished before step B can begin.",
				"Step A must be finished before step F can begin.",
				"Step D must be finished before step E can begin.",
				"Step F must be finished before step E can begin.",
			},
			"CABDFE",
		},
		{
			[]string{
				"Step C must be finished before step A can begin.",
				"Step F must be finished before step E can begin.",
				"Step A must be finished before step D can begin.",
				"Step C must be finished before step F can begin.",
				"Step A must be finished before step B can begin.",
				"Step D must be finished before step E can begin.",
				"Step B must be finished before step E can begin.",
				"Step E must be finished before step G can begin.",
				"Step B must be finished before step H can begin.",
				"Step G must be finished before step I can begin.",
				"Step H must be finished before step I can begin.",
			},
			"CABDFEGHI",
		},
	}

	for _, c := range cases {
		got := day7Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestParseLine(t *testing.T) {
	cases := []struct {
		in    string
		want1 string
		want2 string
	}{
		{"Step C must be finished before step A can begin.", "C", "A"},
		{"Step A must be finished before step B can begin.", "A", "B"},
		{"Step B must be finished before step E can begin.", "B", "E"},
	}

	for _, c := range cases {
		g1, g2 := parseLine(c.in)
		if g1 != c.want1 || g2 != c.want2 {
			t.Errorf("got %v:%v, want %v:%v", g1, g2, c.want1, c.want2)
		}
	}
}

func TestDay7FindNodesWithNoDeps(t *testing.T) {
	a := &task{id: "A", deps: []*task{}}
	b := &task{id: "B", deps: []*task{a}}
	c := &task{id: "C", deps: []*task{}}

	cases := []struct {
		in   map[string]*task
		want []*task
	}{
		{
			map[string]*task{
				"A": a,
				"B": b,
				"C": c,
			},
			[]*task{a, c},
		},
		{
			map[string]*task{
				"B": b,
				"C": c,
			},
			[]*task{c},
		},
	}

	for _, c := range cases {
		got := findNodesWithNoDeps(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay7Part1Final(t *testing.T) {
	in, err := readStrings("./day7_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day7Part1(*in)
	want := "CFGHAEMNBPRDISVWQUZJYTKLOX"
	fmt.Printf("Day 7, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay7Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"Step C must be finished before step A can begin.",
				"Step F must be finished before step E can begin.",
				"Step A must be finished before step D can begin.",
				"Step C must be finished before step F can begin.",
				"Step A must be finished before step B can begin.",
				"Step D must be finished before step E can begin.",
				"Step B must be finished before step E can begin.",
			},
			15,
		},
	}

	for _, c := range cases {
		got := day7Part2(c.in, 2, 0)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay7Part2Final(t *testing.T) {
	in, err := readStrings("./day7_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day7Part2(*in, 5, 60)
	want := 828
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 7, part 2 answer: %v\n", got)
}
