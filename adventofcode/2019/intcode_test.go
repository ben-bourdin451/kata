package adventofcode

import (
	"testing"
)

func TestIntcodeGetDigitAt(t *testing.T) {
	cases := map[string]struct {
		in   int64
		d    int64
		want int64
	}{
		"2:0":     {2, 0, 2},
		"1002:0":  {1002, 0, 2},
		"1002:3":  {1002, 3, 1},
		"1002:4":  {1002, 4, 0},
		"51002:4": {51002, 4, 5},
	}

	for desc, c := range cases {
		got := getDigitAt(c.in, c.d)
		if got != c.want {
			t.Errorf("%v got %v, want %v", desc, got, c.want)
		}
	}
}

func TestIntcodeGetOpCode(t *testing.T) {
	cases := map[int64]int64{
		2:    2,
		51:   51,
		141:  41,
		1002: 2,
		1032: 32,
		1012: 12,
	}

	for k, v := range cases {
		got := getOpCode(k)
		if got != v {
			t.Errorf("%v got %v, want %v", k, got, v)
		}
	}
}

func TestIntcode(t *testing.T) {
	cases := map[string]struct {
		mem   []int64
		input int64
		want  int64
	}{
		"position equal to 8 ? 1 : 0": {
			[]int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			8,
			1,
		},
		"immediate equal to 8 ? 1 : 0": {
			[]int64{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			8,
			1,
		},
		"position less than 8 ? 1 : 0": {
			[]int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			7,
			1,
		},
		"immediate less than 8 ? 1 : 0": {
			[]int64{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			7,
			1,
		},
		"position jump-if-false if in == 0 ? 0 : 1": {
			[]int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			21,
			1,
		},
		"immediate jump-if-true if in == 0 ? 0 : 1": {
			[]int64{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			0,
			0,
		},
		"in < 8 --> 999": {
			[]int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			5,
			999,
		},
		"in == 8 --> 1000": {
			[]int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			8,
			1000,
		},
		"in > 8 --> 1001": {
			[]int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			34,
			1001,
		},
	}

	for k, c := range cases {
		in, out := make(chan int64, 1), make(chan int64)
		go intcode(c.mem, in, out)
		in <- c.input
		if output := <-out; output != c.want {
			t.Errorf("%v got %v, want %v", k, output, c.want)
		}
	}
}
