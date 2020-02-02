package adventofcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Takes initial memory alongside input & output channels
// Modifies memory (intended side effect, non-purity)
func intcode(mem []int, in <-chan int, out chan<- int) {
	for i := 0; i < len(mem); {
		if mem[i] == 99 {
			return
		}
		_, b, c := getDigitAt(mem[i], 4), getDigitAt(mem[i], 3), getDigitAt(mem[i], 2)
		op := getOpCode(mem[i])

		switch op {
		case 1: // add
			// fmt.Println(i, mem[i:i+4])
			mem[mem[i+3]] = getParam(mem, i+1, c) + getParam(mem, i+2, b)
			i += 4
			break

		case 2: // multiply
			// fmt.Println(i, mem[i:i+4])
			mem[mem[i+3]] = getParam(mem, i+1, c) * getParam(mem, i+2, b)
			i += 4
			break

		case 3: // input
			// fmt.Println(i, mem[i:i+2])
			// fmt.Println("waiting for input")
			mem[mem[i+1]] = <-in
			i += 2
			break

		case 4: // output
			// fmt.Println(i, mem[i:i+2])
			// fmt.Println("sending output")
			out <- getParam(mem, i+1, c)
			i += 2
			break

		case 5: // jump-if-true
			// fmt.Println(i, mem[i:i+3])
			if getParam(mem, i+1, c) != 0 {
				// fmt.Println("i -->", getParam(mem, i+2, b))
				i = getParam(mem, i+2, b)
			} else {
				i += 3
			}
			break

		case 6: // jump-if-false
			// fmt.Println(i, mem[i:i+3])
			if getParam(mem, i+1, c) == 0 {
				// fmt.Println("i -->", getParam(mem, i+2, b))
				i = getParam(mem, i+2, b)
			} else {
				i += 3
			}
			break

		case 7: // less-than
			// fmt.Println(i, mem[i:i+4])
			// fmt.Println(getParam(mem, i+1, c), "<", getParam(mem, i+2, b), "-->", mem[mem[i+3]])
			if getParam(mem, i+1, c) < getParam(mem, i+2, b) {
				mem[mem[i+3]] = 1
			} else {
				mem[mem[i+3]] = 0
			}
			i += 4
			break

		case 8: // equals
			// fmt.Println(i, mem[i:i+4])
			// fmt.Println(getParam(mem, i+1, c), "==", getParam(mem, i+2, b), "-->", mem[mem[i+3]])
			if getParam(mem, i+1, c) == getParam(mem, i+2, b) {
				mem[mem[i+3]] = 1
			} else {
				mem[mem[i+3]] = 0
			}
			i += 4
			break

		default:
			fmt.Println("unknown op", op)
			return
		}
	}
}

func initCodes(in string) []int {
	mem := strings.Split(in, ",")
	codes := []int{}
	for _, code := range mem {
		n, _ := strconv.Atoi(code)
		codes = append(codes, n)
	}

	return codes
}

// getDigitAt(9876, 1) --> 7
func getDigitAt(in, d int) int {
	return int(math.Floor(float64(in)/float64(math.Pow10(d)))) % 10
}

func getOpCode(in int) int {
	return getDigitAt(in, 0) + getDigitAt(in, 1)*10
}

func getParam(codes []int, pos, mode int) int {
	if mode == 1 {
		return codes[pos]
	}

	return codes[codes[pos]]
}
