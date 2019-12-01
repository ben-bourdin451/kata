package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay10Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"position=< 9,  1> velocity=< 0,  2>",
				"position=< 7,  0> velocity=<-1,  0>",
				"position=< 3, -2> velocity=<-1,  1>",
				"position=< 6, 10> velocity=<-2, -1>",
				"position=< 2, -4> velocity=< 2,  2>",
				"position=<-6, 10> velocity=< 2, -2>",
				"position=< 1,  8> velocity=< 1, -1>",
				"position=< 1,  7> velocity=< 1,  0>",
				"position=<-3, 11> velocity=< 1, -2>",
				"position=< 7,  6> velocity=<-1, -1>",
				"position=<-2,  3> velocity=< 1,  0>",
				"position=<-4,  3> velocity=< 2,  0>",
				"position=<10, -3> velocity=<-1,  1>",
				"position=< 5, 11> velocity=< 1, -2>",
				"position=< 4,  7> velocity=< 0, -1>",
				"position=< 8, -2> velocity=< 0,  1>",
				"position=<15,  0> velocity=<-2,  0>",
				"position=< 1,  6> velocity=< 1,  0>",
				"position=< 8,  9> velocity=< 0, -1>",
				"position=< 3,  3> velocity=<-1,  1>",
				"position=< 0,  5> velocity=< 0, -1>",
				"position=<-2,  2> velocity=< 2,  0>",
				"position=< 5, -2> velocity=< 1,  2>",
				"position=< 1,  4> velocity=< 2,  1>",
				"position=<-2,  7> velocity=< 2, -2>",
				"position=< 3,  6> velocity=<-1, -1>",
				"position=< 5,  0> velocity=< 1,  0>",
				"position=<-6,  0> velocity=< 2,  0>",
				"position=< 5,  9> velocity=< 1, -2>",
				"position=<14,  7> velocity=<-2,  0>",
				"position=<-3,  6> velocity=< 2, -1>",
			},
			3,
		},
	}

	for _, c := range cases {
		got := day10(c.in, 80)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay10ParseInput(t *testing.T) {
	cases := []struct {
		in   string
		want *light
	}{
		{"position=< 9,  1> velocity=< 0,  2>", &light{&point{9, 1}, &point{0, 2}}},
		{"position=< 7,  0> velocity=<-1,  0>", &light{&point{7, 0}, &point{-1, 0}}},
		{"position=< 3, -2> velocity=<-1,  1>", &light{&point{3, -2}, &point{-1, 1}}},
	}

	for _, c := range cases {
		got := parseLight(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay10Move(t *testing.T) {
	cases := []struct {
		in   *light
		want *light
	}{
		{&light{&point{9, 1}, &point{0, 2}}, &light{&point{9, 3}, &point{0, 2}}},
		{&light{&point{5, 2}, &point{3, 2}}, &light{&point{8, 4}, &point{3, 2}}},
		{&light{&point{7, 0}, &point{-1, 0}}, &light{&point{6, 0}, &point{-1, 0}}},
		{&light{&point{7, 0}, &point{3, 0}}, &light{&point{10, 0}, &point{3, 0}}},
		{&light{&point{3, -2}, &point{-1, -1}}, &light{&point{2, -3}, &point{-1, -1}}},
	}

	for _, c := range cases {
		c.in.move()
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("got %v, want %v", c.in, c.want)
		}
	}
}

func TestDay10(t *testing.T) {
	in, err := readStrings("./day10_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	fmt.Println("Day 10, part 1 answer:")
	got := day10(*in, 1000)
	want := 10659
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 10, part 2 answer: %v\n", got)
}
