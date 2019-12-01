package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay11Part1(t *testing.T) {
	cases := []struct {
		in   int
		want *point
	}{
		{18, &point{33, 45}},
		{42, &point{21, 61}},
	}

	for _, c := range cases {
		got := day11Part1(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay11FuelCellPower(t *testing.T) {
	cases := []struct {
		x, y int
		s    int
		want int
	}{
		{3, 5, 8, 4},
		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
	}

	for _, c := range cases {
		got := fuelCellPower(c.x, c.y, c.s)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay11Part1Final(t *testing.T) {
	got := day11Part1(1309)
	want := &point{20, 43}
	fmt.Printf("Day 11, part 1 answer: %v\n", got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay11Part2(t *testing.T) {
	cases := []struct {
		in   int
		p    *point
		size int
	}{
		{18, &point{90, 269}, 16},
		{42, &point{232, 251}, 12},
	}

	for _, c := range cases {
		p, s := day11Part2(c.in)
		if !reflect.DeepEqual(p, c.p) || s != c.size {
			t.Errorf("got %v,%v want %v,%v", p, s, c.p, c.size)
		}
	}
}

func TestDay11Part2Final(t *testing.T) {
	p, s := day11Part2(1309)
	want, size := &point{233, 271}, 13
	if !reflect.DeepEqual(p, want) || s != size {
		t.Errorf("got %v,%v want %v,%v", p, s, want, size)
	}
	fmt.Printf("Day 11, part 2 answer: %v,%v\n", p, s)
}
