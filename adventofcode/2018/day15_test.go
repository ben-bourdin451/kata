package adventofcode

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDay15Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"#######",
				"#.G...#",
				"#...EG#",
				"#.#.#G#",
				"#..G#E#",
				"#.....#",
				"#######",
			},
			27730,
		},
		{
			[]string{
				"#######",
				"#G..#E#",
				"#E#E.E#",
				"#G.##.#",
				"#...#E#",
				"#...E.#",
				"#######",
			},
			36334,
		},
		{
			[]string{
				"#######",
				"#E..EG#",
				"#.#G.E#",
				"#E.##E#",
				"#G..#.#",
				"#..E#.#",
				"#######",
			},
			39514,
		},
		{
			[]string{
				"#######",
				"#E.G#.#",
				"#.#G..#",
				"#G.#.G#",
				"#G..#.#",
				"#...E.#",
				"#######",
			},
			27755,
		},
		{
			[]string{
				"#######",
				"#.E...#",
				"#.#..G#",
				"#.###.#",
				"#E#G#G#",
				"#...#G#",
				"#######",
			},
			28944,
		},
		{
			[]string{
				"#########",
				"#G......#",
				"#.E.#...#",
				"#..##..G#",
				"#...##..#",
				"#...#...#",
				"#.G...G.#",
				"#.....G.#",
				"#########",
			},
			18740,
		},
	}

	for _, c := range cases {
		got := day15Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay15FindReachablePoints(t *testing.T) {
	cases := []struct {
		cave     [][]string
		u        *unit
		e        []*unit
		want     *point
		distance int
	}{
		{
			[][]string{ //   0123456
				strings.Split("#######", ""), // 0
				strings.Split("#.G...#", ""), // 1
				strings.Split("#...EG#", ""), // 2
				strings.Split("#.#.#G#", ""), // 3
				strings.Split("#..G#E#", ""), // 4
				strings.Split("#.....#", ""), // 5
				strings.Split("#######", ""), // 6
			},
			&unit{&point{2, 1}, 200, "G"},
			[]*unit{
				&unit{&point{4, 2}, 200, "E"},
				&unit{&point{5, 4}, 200, "E"},
			},
			&point{3, 1},
			2,
		},
		{
			[][]string{ //   0123456
				strings.Split("#######", ""), // 0
				strings.Split("#.G...#", ""), // 1
				strings.Split("#...EG#", ""), // 2
				strings.Split("#.#.#G#", ""), // 3
				strings.Split("#..G#E#", ""), // 4
				strings.Split("#.....#", ""), // 5
				strings.Split("#######", ""), // 6
			},
			&unit{&point{2, 1}, 200, "G"},
			[]*unit{
				&unit{&point{5, 4}, 200, "E"},
			},
			&point{1, 1},
			9,
		},
		{
			[][]string{ //   0123456
				strings.Split("#######", ""), // 0
				strings.Split("#.G...#", ""), // 1
				strings.Split("#...EG#", ""), // 2
				strings.Split("#.#.#G#", ""), // 3
				strings.Split("#...#E#", ""), // 4
				strings.Split("#.....#", ""), // 5
				strings.Split("#######", ""), // 6
			},
			&unit{&point{2, 1}, 200, "G"},
			[]*unit{
				&unit{&point{5, 4}, 200, "E"},
			},
			&point{3, 1},
			7,
		},
		{
			[][]string{ //   0123456
				strings.Split("#######", ""), // 0
				strings.Split("#.G...#", ""), // 1
				strings.Split("#...EG#", ""), // 2
				strings.Split("#.#.#G#", ""), // 3
				strings.Split("#..G#E#", ""), // 4
				strings.Split("#.....#", ""), // 5
				strings.Split("#######", ""), // 6
			},
			&unit{&point{3, 4}, 200, "G"},
			[]*unit{
				&unit{&point{4, 2}, 200, "E"},
				&unit{&point{5, 4}, 200, "E"},
			},
			&point{3, 3},
			2,
		},
		{
			[][]string{ //   0123456
				strings.Split("#######", ""), // 0
				strings.Split("#.G...#", ""), // 1
				strings.Split("#...EG#", ""), // 2
				strings.Split("#.#.#G#", ""), // 3
				strings.Split("#..G#E#", ""), // 4
				strings.Split("#.....#", ""), // 5
				strings.Split("#######", ""), // 6
			},
			&unit{&point{3, 4}, 200, "G"},
			[]*unit{
				&unit{&point{5, 4}, 200, "E"},
			},
			&point{3, 5},
			3,
		},
		{
			[][]string{ //   0123456
				strings.Split("#######", ""), // 0
				strings.Split("#.G...#", ""), // 1
				strings.Split("#...EG#", ""), // 2
				strings.Split("#.#.#G#", ""), // 3
				strings.Split("#...#E#", ""), // 4
				strings.Split("#....G#", ""), // 5
				strings.Split("#######", ""), // 6
			},
			&unit{&point{5, 3}, 200, "G"},
			[]*unit{
				&unit{&point{4, 2}, 200, "E"},
				&unit{&point{5, 4}, 200, "E"},
			},
			&point{5, 3},
			0,
		},
		{
			[][]string{ //   0123456
				strings.Split("#######", ""), // 0
				strings.Split("#.G...#", ""), // 1
				strings.Split("#..G.G#", ""), // 2
				strings.Split("#...#E#", ""), // 3
				strings.Split("#.....#", ""), // 4
				strings.Split("#E....#", ""), // 5
				strings.Split("#######", ""), // 6
			},
			&unit{&point{3, 2}, 200, "G"},
			[]*unit{
				&unit{&point{5, 3}, 200, "E"},
				&unit{&point{1, 5}, 200, "E"},
			},
			&point{2, 2},
			4,
		},
	}

	for _, c := range cases {
		got, d := c.u.findMoveTowardsEnemy(c.cave, c.e)
		if !reflect.DeepEqual(got, c.want) || d != c.distance {
			t.Errorf("got %v:%v, want %v:%v", got, d, c.want, c.distance)
		}
	}
}

func TestDay15Part1Final(t *testing.T) {
	in, err := readStrings("./day15_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day15Part1(*in)
	want := 198744
	fmt.Printf("Day 15, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay15Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"#######",
				"#.G...#",
				"#...EG#",
				"#.#.#G#",
				"#..G#E#",
				"#.....#",
				"#######",
			},
			4988,
		},
		{
			[]string{
				"#######",
				"#E..EG#",
				"#.#G.E#",
				"#E.##E#",
				"#G..#.#",
				"#..E#.#",
				"#######",
			},
			31284,
		},
		{
			[]string{
				"#######",
				"#E.G#.#",
				"#.#G..#",
				"#G.#.G#",
				"#G..#.#",
				"#...E.#",
				"#######",
			},
			3478,
		},
		{
			[]string{
				"#######",
				"#.E...#",
				"#.#..G#",
				"#.###.#",
				"#E#G#G#",
				"#...#G#",
				"#######",
			},
			6474,
		},
		{
			[]string{
				"#########",
				"#G......#",
				"#.E.#...#",
				"#..##..G#",
				"#...##..#",
				"#...#...#",
				"#.G...G.#",
				"#.....G.#",
				"#########",
			},
			1140,
		},
	}

	for _, c := range cases {
		got := day15Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay15Part2Final(t *testing.T) {
	in, err := readStrings("./day15_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day15Part2(*in)
	want := 0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 15, part 2 answer: %v\n", got)
}
