package adventofcode

import (
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	regLength = 4
)

type sample struct {
	before []int
	op     []int
	after  []int
}

// Function type which defines operations on a register
type opFunc func(reg []int, in []int) []int
type op struct {
	name      string
	behaviour opFunc
	code      int
}

func day16Part1(samples []*sample) int {
	sum := 0
	ops := initBehaviours()
	for _, s := range samples {
		if c, _ := behaviourCount(s, ops); c >= 3 {
			sum++
		}
	}
	return sum
}

func behaviourCount(s *sample, ops []*op) (int, *op) {
	count := 0
	var last *op
	for _, op := range ops {
		if op.code < 0 && reflect.DeepEqual(s.after, op.behaviour(s.before, s.op[1:])) {
			count++
			last = op
		}
	}
	return count, last
}

func day16Part2(samples []*sample, input []string) int {
	ops := initBehaviours()

	for hasUnknownCode(ops) {
		for _, s := range samples {
			c, o := behaviourCount(s, ops)
			if c == 1 {
				o.code = s.op[0]
			}
		}
	}
	sortOps(ops)

	reg := []int{0, 0, 0, 0}
	for _, in := range input {
		i := parseRegister(in)
		reg = ops[i[0]].behaviour(reg, i[1:])
	}

	return reg[0]
}

func sortOps(ops []*op) {
	sort.Slice(ops, func(i, j int) bool {
		return ops[i].code < ops[j].code
	})
}

func hasUnknownCode(ops []*op) bool {
	for _, op := range ops {
		if op.code < 0 {
			return true
		}
	}

	return false
}

func initBehaviours() []*op {
	ops := []*op{
		&op{"addr", addr, -1},
		&op{"addi", addi, -1},

		&op{"mulr", mulr, -1},
		&op{"muli", muli, -1},

		&op{"banr", banr, -1},
		&op{"bani", bani, -1},

		&op{"borr", borr, -1},
		&op{"bori", bori, -1},

		&op{"setr", setr, -1},
		&op{"seti", seti, -1},

		&op{"gtir", gtir, -1},
		&op{"gtri", gtri, -1},
		&op{"gtrr", gtrr, -1},

		&op{"eqir", eqir, -1},
		&op{"eqri", eqri, -1},
		&op{"eqrr", eqrr, -1},
	}

	return ops
}

func addr(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	out[in[2]] = r[in[0]] + r[in[1]]
	return out
}
func addi(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	out[in[2]] = r[in[0]] + in[1]
	return out
}
func mulr(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	out[in[2]] = r[in[0]] * r[in[1]]
	return out
}
func muli(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	out[in[2]] = r[in[0]] * in[1]
	return out
}
func banr(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	out[in[2]] = r[in[0]] & r[in[1]]
	return out
}
func bani(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	out[in[2]] = r[in[0]] & in[1]
	return out
}
func borr(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	out[in[2]] = r[in[0]] | r[in[1]]
	return out
}
func bori(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	out[in[2]] = r[in[0]] | in[1]
	return out
}
func setr(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	out[in[2]] = r[in[0]]
	return out
}
func seti(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	out[in[2]] = in[0]
	return out
}
func gtir(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	if in[0] > r[in[1]] {
		out[in[2]] = 1
	} else {
		out[in[2]] = 0
	}
	return out
}
func gtri(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	if r[in[0]] > in[1] {
		out[in[2]] = 1
	} else {
		out[in[2]] = 0
	}
	return out
}
func gtrr(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	if r[in[0]] > r[in[1]] {
		out[in[2]] = 1
	} else {
		out[in[2]] = 0
	}
	return out
}
func eqir(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	if in[0] == r[in[1]] {
		out[in[2]] = 1
	} else {
		out[in[2]] = 0
	}
	return out
}
func eqri(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	if r[in[0]] == in[1] {
		out[in[2]] = 1
	} else {
		out[in[2]] = 0
	}
	return out
}
func eqrr(r []int, in []int) []int {
	out := make([]int, len(r))
	copy(out[:], r)

	if r[in[0]] == r[in[1]] {
		out[in[2]] = 1
	} else {
		out[in[2]] = 0
	}
	return out
}

func parseSamples(in []string) []*sample {
	samples := []*sample{}
	var curr *sample
	for _, line := range in {
		if line == "" {
			continue

		} else if strings.HasPrefix(line, "Before") {
			curr = &sample{} // new sample
			curr.before = parseRegister(line)

		} else if strings.HasPrefix(line, "After") {
			curr.after = parseRegister(line)
			samples = append(samples, curr) // end of current sample

		} else {
			curr.op = parseRegister(line)
		}
	}

	return samples
}

func parseRegister(in string) []int {
	r := regexp.MustCompile(`[0-9]+`)
	nums := r.FindAllStringSubmatch(in, regLength)

	reg := []int{}
	for _, n := range nums {
		v, _ := strconv.Atoi(n[0])
		reg = append(reg, v)
	}

	return reg
}
