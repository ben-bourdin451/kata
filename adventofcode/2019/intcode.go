package adventofcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Takes initial memory alongside input & output channels
// Modifies memory (int64ended side effect, non-purity)
func intcode(mem []int64, in <-chan int64, out chan<- int64) {
	if out != nil {
		defer close(out)
	}

	rbase := int64(0)
	for i := int64(0); i < int64(len(mem)); {
		if mem[i] == 99 {
			return
		}
		a, b, c := getDigitAt(mem[i], 4), getDigitAt(mem[i], 3), getDigitAt(mem[i], 2)
		op := getOpCode(mem[i])

		switch op {
		case 1: // add
			res := getParam(mem, i+1, rbase, c) + getParam(mem, i+2, rbase, b)
			setParam(mem, res, i+3, rbase, a)
			i += 4
			break

		case 2: // multiply
			res := getParam(mem, i+1, rbase, c) * getParam(mem, i+2, rbase, b)
			setParam(mem, res, i+3, rbase, a)
			i += 4
			break

		case 3: // input
			setParam(mem, <-in, i+1, rbase, c)
			i += 2
			break

		case 4: // output
			out <- getParam(mem, i+1, rbase, c)
			i += 2
			break

		case 5: // jump-if-true
			if getParam(mem, i+1, rbase, c) != 0 {
				i = getParam(mem, i+2, rbase, b)
			} else {
				i += 3
			}
			break

		case 6: // jump-if-false
			if getParam(mem, i+1, rbase, c) == 0 {
				i = getParam(mem, i+2, rbase, b)
			} else {
				i += 3
			}
			break

		case 7: // less-than
			if getParam(mem, i+1, rbase, c) < getParam(mem, i+2, rbase, b) {
				setParam(mem, 1, i+3, rbase, a)
			} else {
				setParam(mem, 0, i+3, rbase, a)
			}
			i += 4
			break

		case 8: // equals
			if getParam(mem, i+1, rbase, c) == getParam(mem, i+2, rbase, b) {
				setParam(mem, 1, i+3, rbase, a)
			} else {
				setParam(mem, 0, i+3, rbase, a)
			}
			i += 4
			break

		case 9: // adjust relative base
			rbase += getParam(mem, i+1, rbase, c)
			i += 2
			break

		default:
			fmt.Println("unknown op", op)
			return
		}
	}
}

func initCodes(in string) []int64 {
	mem := strings.Split(in, ",")
	codes := make([]int64, 10000)
	for i, code := range mem {
		n, _ := strconv.ParseInt(code, 10, 64)
		codes[i] = n
	}

	return codes
}

// getDigitAt(9876, 1) --> 7
func getDigitAt(in, d int64) int64 {
	return int64(math.Floor(float64(in)/float64(math.Pow(10, float64(d))))) % 10
}

func getOpCode(in int64) int64 {
	return getDigitAt(in, 0) + getDigitAt(in, 1)*10
}

func setParam(mem []int64, value, pos, rbase, mode int64) {
	param := mem[pos]
	switch mode {
	case 0:
		mem[param] = value
	case 1:
		mem[pos] = value
	case 2:
		mem[rbase+param] = value
	default:
		fmt.Println("unknown mode")
	}
}

func getParam(mem []int64, pos, rbase, mode int64) int64 {
	param := mem[pos]
	switch mode {
	case 0:
		return mem[param]
	case 1:
		return mem[pos]
	case 2:
		return mem[rbase+param]
	default:
		fmt.Println("unknown mode")
		return -1
	}
}
