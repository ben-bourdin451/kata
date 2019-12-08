package adventofcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

func intcode(codes []int, input int) ([]int, []int) {
	output := []int{}
	for i := 0; i < len(codes); {
		if codes[i] == 99 {
			return codes, output
		}
		_, b, c := getDigitAt(codes[i], 4), getDigitAt(codes[i], 3), getDigitAt(codes[i], 2)
		op := getOpCode(codes[i])

		switch op {
		case 1: // add
			// fmt.Println(i, codes[i:i+4])
			codes[codes[i+3]] = getParam(codes, i+1, c) + getParam(codes, i+2, b)
			i += 4
			break

		case 2: // multiply
			// fmt.Println(i, codes[i:i+4])
			codes[codes[i+3]] = getParam(codes, i+1, c) * getParam(codes, i+2, b)
			i += 4
			break

		case 3: // input
			// fmt.Println(i, codes[i:i+2])
			codes[codes[i+1]] = input
			i += 2
			break

		case 4: // output
			// fmt.Println(i, codes[i:i+2])
			output = append(output, getParam(codes, i+1, c))
			i += 2
			break

		case 5: // jump-if-true
			// fmt.Println(i, codes[i:i+3])
			if getParam(codes, i+1, c) != 0 {
				// fmt.Println("i -->", getParam(codes, i+2, b))
				i = getParam(codes, i+2, b)
			} else {
				i += 3
			}
			break

		case 6: // jump-if-false
			// fmt.Println(i, codes[i:i+3])
			if getParam(codes, i+1, c) == 0 {
				// fmt.Println("i -->", getParam(codes, i+2, b))
				i = getParam(codes, i+2, b)
			} else {
				i += 3
			}
			break

		case 7: // less-than
			// fmt.Println(i, codes[i:i+4])
			// fmt.Println(getParam(codes, i+1, c), "<", getParam(codes, i+2, b), "-->", codes[codes[i+3]])
			if getParam(codes, i+1, c) < getParam(codes, i+2, b) {
				codes[codes[i+3]] = 1
			} else {
				codes[codes[i+3]] = 0
			}
			i += 4
			break

		case 8: // equals
			// fmt.Println(i, codes[i:i+4])
			// fmt.Println(getParam(codes, i+1, c), "==", getParam(codes, i+2, b), "-->", codes[codes[i+3]])
			if getParam(codes, i+1, c) == getParam(codes, i+2, b) {
				codes[codes[i+3]] = 1
			} else {
				codes[codes[i+3]] = 0
			}
			i += 4
			break

		default:
			fmt.Println("unknown op", op)
			return codes, output
		}
	}

	return codes, output
}
