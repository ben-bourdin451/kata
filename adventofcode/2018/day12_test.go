package adventofcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay12Part1(t *testing.T) {
	cases := []struct {
		in     string
		combos []string
		want   int
	}{
		{
			"#..#.#..##......###...###",
			[]string{
				"...## => #",
				"..#.. => #",
				".#... => #",
				".#.#. => #",
				".#.## => #",
				".##.. => #",
				".#### => #",
				"#.#.# => #",
				"#.### => #",
				"##.#. => #",
				"##.## => #",
				"###.. => #",
				"###.# => #",
				"####. => #",
			},
			325,
		},
	}

	for _, c := range cases {
		got := day12Part1(c.in, c.combos, 20)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay12FindMatch(t *testing.T) {
	combos := []string{
		"...## => #",
		"..#.. => #",
		".#... => #",
		".#.#. => #",
		".#.## => #",
		".##.. => #",
		".#### => #",
		"#.#.# => #",
		"#.### => #",
		"##.#. => #",
		"##.## => #",
		"###.. => #",
		"###.# => #",
		"####. => #",
	}

	cases := []struct {
		in     string
		combos []string
		want   string
	}{
		{"...##", combos, "#"},
		{"#####", combos, "."},
		{"####.", combos, "#"},
		{"..#..", combos, "#"},
		{"..###", combos, "."},
	}

	for _, c := range cases {
		got := findMatch(c.in, c.combos)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay12TunnelValue(t *testing.T) {
	cases := []struct {
		in   *plantTunnel
		want int
	}{
		{
			&plantTunnel{
				[]string{},
				strings.Split(".#....##....#####...#######....#.#..##.", ""),
				3,
			}, 325,
		},
	}

	for _, c := range cases {
		got := c.in.getValue()
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay12Part1Final(t *testing.T) {
	combos, err := readStrings("./day12_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}
	init := "###..###....####.###...#..#...##...#..#....#.##.##.#..#.#..##.#####..######....#....##..#...#...#.#"
	got := day12Part1(init, *combos, 20)
	want := 6201
	fmt.Printf("Day 12, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay12Part2Final(t *testing.T) {
	combos, err := readStrings("./day12_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}
	init := "###..###....####.###...#..#...##...#..#....#.##.##.#..#.#..##.#####..######....#....##..#...#...#.#"
	got := day12Part2(init, *combos, 50000000000)
	want := 9300000001023
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 12, part 2 answer: %v\n", got)
}
