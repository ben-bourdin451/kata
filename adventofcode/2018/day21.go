package adventofcode

import (
	"regexp"
	"strconv"
)

func day21Part1(reg []int, ip string, in []string) int {
	count := 0
	halted := false
	for count < 100 && !halted {
		reg[0] = count
		if c := doWork(reg, ip, in); c < 10000 {
			return c
		}
		count++
	}

	return -1
}

func doWork(reg []int, ip string, in []string) int {
	ops := []opInstruction{}
	for _, i := range in {
		ops = append(ops, parseOpInstruction(i))
	}

	ip = regexp.MustCompile(`[0-9]+`).FindStringSubmatch(ip)[0]
	boundReg, _ := strconv.Atoi(ip)

	ipVal := reg[boundReg]
	count := 0
	for ipVal >= 0 && ipVal < len(ops) && count < 100000 {
		reg[boundReg] = ipVal
		reg = ops[ipVal].behaviour(reg, ops[ipVal].vals)
		ipVal = reg[boundReg]
		ipVal++
		count++
		// fmt.Printf("ip=%v\t%v\t%v\n", ipVal, reg, ops[ipVal].vals)
	}

	return count
}

func day21Part2(in []string) int {
	return 0
}
