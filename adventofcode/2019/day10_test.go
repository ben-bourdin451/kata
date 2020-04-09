package adventofcode

import (
	"fmt"
	"testing"
)

func TestDay10Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		// {
		// 	[]string{
		// 		".#..#",
		// 		".....",
		// 		".....",
		// 		".....",
		// 		".....",
		// 	},
		// 	1,
		// },
		// {
		// 	[]string{
		// 		".#..#",
		// 		".....",
		// 		"#####",
		// 		"....#",
		// 		"...##",
		// 	},
		// 	8,
		// },
		// {
		// 	[]string{
		// 		"......#.#.",
		// 		"#..#.#....",
		// 		"..#######.",
		// 		".#.#.###..",
		// 		".#..#.....",
		// 		"..#....#.#",
		// 		"#..#....#.",
		// 		".##.#..###",
		// 		"##...#..#.",
		// 		".#....####",
		// 	},
		// 	33,
		// },
		// {
		// 	[]string{
		// 		"#.#...#.#.",
		// 		".###....#.",
		// 		".#....#...",
		// 		"##.#.#.#.#",
		// 		"....#.#.#.",
		// 		".##..###.#",
		// 		"..#...##..",
		// 		"..##....##",
		// 		"......#...",
		// 		".####.###.",
		// 	},
		// 	35,
		// },
		// {
		// 	[]string{
		// 		".#..#..###",
		// 		"####.###.#",
		// 		"....###.#.",
		// 		"..###.##.#",
		// 		"##.##.#.#.",
		// 		"....###..#",
		// 		"..#.#..#.#",
		// 		"#..#.#.###",
		// 		".##...##.#",
		// 		".....#.#..",
		// 	},
		// 	41,
		// },
		// {
		// 	[]string{
		// 		".#..##.###...#######",
		// 		"##.############..##.",
		// 		".#.######.########.#",
		// 		".###.#######.####.#.",
		// 		"#####.##.#.##.###.##",
		// 		"..#####..#.#########",
		// 		"####################",
		// 		"#.####....###.#.#.##",
		// 		"##.#################",
		// 		"#####.##.###..####..",
		// 		"..######..##.#######",
		// 		"####.##.####...##..#",
		// 		".#####..#.######.###",
		// 		"##...#.##########...",
		// 		"#.##########.#######",
		// 		".####.#.###.###.#.##",
		// 		"....##.##.###..#####",
		// 		".#.#.###########.###",
		// 		"#.#.#.#####.####.###",
		// 		"###.##.####.##.#..##",
		// 	},
		// 	210,
		// },
	}

	for _, c := range cases {
		got := day10Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay10Part1Final(t *testing.T) {
	in, err := readStrings("./day10_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day10Part1(in)
	want := 0
	fmt.Printf("Day 10, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay10Part2(t *testing.T) {
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
		got := day10Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay10Part2Final(t *testing.T) {
	in, err := readStrings("./day10_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day10Part2(in)
	want := 0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 10, part 2 answer: %v\n", got)
}
