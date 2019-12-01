package adventofcode

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDay4Part1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{
			"[1518-11-01 00:05] falls asleep",
			"[1518-11-02 00:40] falls asleep",
			"[1518-11-01 00:00] Guard #10 begins shift",
			"[1518-11-01 00:30] falls asleep",
			"[1518-11-01 00:25] wakes up",
			"[1518-11-01 23:58] Guard #99 begins shift",
			"[1518-11-05 00:55] wakes up",
			"[1518-11-03 00:24] falls asleep",
			"[1518-11-04 00:02] Guard #99 begins shift",
			"[1518-11-02 00:50] wakes up",
			"[1518-11-03 00:05] Guard #10 begins shift",
			"[1518-11-05 00:03] Guard #99 begins shift",
			"[1518-11-04 00:36] falls asleep",
			"[1518-11-04 00:46] wakes up",
			"[1518-11-03 00:29] wakes up",
			"[1518-11-01 00:55] wakes up",
			"[1518-11-05 00:45] falls asleep",
		}, 240},
	}
	for _, c := range cases {
		got := day4Part1(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestGetTime(t *testing.T) {
	cases := []struct {
		in   string
		want time.Time
	}{
		{"[1518-11-01 00:05] falls asleep", time.Date(1518, 11, 01, 0, 5, 0, 0, time.UTC)},
		{"[1518-11-02 00:40] falls asleep", time.Date(1518, 11, 02, 0, 40, 0, 0, time.UTC)},
		{"[1518-11-01 00:00] Guard #10 begins shift", time.Date(1518, 11, 01, 0, 0, 0, 0, time.UTC)},
		{"[1518-11-01 00:30] falls asleep", time.Date(1518, 11, 01, 0, 30, 0, 0, time.UTC)},
	}
	for _, c := range cases {
		got, err := getTime(c.in)
		if err != nil || !got.Equal(c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestGetGuardID(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"[1518-11-01 00:00] Guard #10 begins shift", 10},
		{"[1518-11-02 23:50] Guard #12 begins shift", 12},
		{"Guard #15 begins shift", 15},
	}
	for _, c := range cases {
		got, err := getGuardID(c.in)
		if err != nil || got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestSort(t *testing.T) {
	cases := []struct {
		in   []string
		want []string
	}{{
		[]string{
			"[1518-11-01 00:05] falls asleep",
			"[1518-11-02 00:40] falls asleep",
			"[1518-11-01 00:00] Guard #10 begins shift",
			"[1518-11-01 00:30] falls asleep",
			"[1518-11-01 00:25] wakes up",
			"[1518-11-01 23:58] Guard #99 begins shift",
			"[1518-11-05 00:55] wakes up",
			"[1518-11-03 00:24] falls asleep",
			"[1518-11-04 00:02] Guard #99 begins shift",
			"[1518-11-02 00:50] wakes up",
			"[1518-11-03 00:05] Guard #10 begins shift",
			"[1518-11-05 00:03] Guard #99 begins shift",
			"[1518-11-04 00:36] falls asleep",
			"[1518-11-04 00:46] wakes up",
			"[1518-11-03 00:29] wakes up",
			"[1518-11-01 00:55] wakes up",
			"[1518-11-05 00:45] falls asleep",
		},
		[]string{
			"[1518-11-01 00:00] Guard #10 begins shift",
			"[1518-11-01 00:05] falls asleep",
			"[1518-11-01 00:25] wakes up",
			"[1518-11-01 00:30] falls asleep",
			"[1518-11-01 00:55] wakes up",
			"[1518-11-01 23:58] Guard #99 begins shift",
			"[1518-11-02 00:40] falls asleep",
			"[1518-11-02 00:50] wakes up",
			"[1518-11-03 00:05] Guard #10 begins shift",
			"[1518-11-03 00:24] falls asleep",
			"[1518-11-03 00:29] wakes up",
			"[1518-11-04 00:02] Guard #99 begins shift",
			"[1518-11-04 00:36] falls asleep",
			"[1518-11-04 00:46] wakes up",
			"[1518-11-05 00:03] Guard #99 begins shift",
			"[1518-11-05 00:45] falls asleep",
			"[1518-11-05 00:55] wakes up",
		}},
	}
	for _, c := range cases {
		sortInput(c.in)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Error("Array not sorted!")
		}
	}
}

func TestDay4Part1Final(t *testing.T) {
	in, err := readStrings("./day4_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day4Part1(*in)
	want := 12169
	fmt.Printf("Day 4, part 1 answer: %v\n", got)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDay4Part2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{[]string{
			"[1518-11-01 00:05] falls asleep",
			"[1518-11-02 00:40] falls asleep",
			"[1518-11-01 00:00] Guard #10 begins shift",
			"[1518-11-01 00:30] falls asleep",
			"[1518-11-01 00:25] wakes up",
			"[1518-11-01 23:58] Guard #99 begins shift",
			"[1518-11-05 00:55] wakes up",
			"[1518-11-03 00:24] falls asleep",
			"[1518-11-04 00:02] Guard #99 begins shift",
			"[1518-11-02 00:50] wakes up",
			"[1518-11-03 00:05] Guard #10 begins shift",
			"[1518-11-05 00:03] Guard #99 begins shift",
			"[1518-11-04 00:36] falls asleep",
			"[1518-11-04 00:46] wakes up",
			"[1518-11-03 00:29] wakes up",
			"[1518-11-01 00:55] wakes up",
			"[1518-11-05 00:45] falls asleep",
		}, 4455},
	}
	for _, c := range cases {
		got := day4Part2(c.in)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func TestDay4Part2Final(t *testing.T) {
	in, err := readStrings("./day4_input.txt")
	if err != nil {
		t.Error("Error while reading input", err)
	}

	got := day4Part2(*in)
	want := 775
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	fmt.Printf("Day 4, part 2 answer: %v\n", got)
}
