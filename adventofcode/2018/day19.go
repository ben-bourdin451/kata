package adventofcode

import (
	"regexp"
	"strconv"
)

type opInstruction struct {
	name      string
	behaviour opFunc
	vals      []int
}

func day19(reg []int, ip string, in []string) int {

	ops := []opInstruction{}
	for _, i := range in {
		ops = append(ops, parseOpInstruction(i))
	}

	ip = regexp.MustCompile(`[0-9]+`).FindStringSubmatch(ip)[0]
	boundReg, _ := strconv.Atoi(ip)

	ipVal := reg[boundReg]
	for ipVal >= 0 && ipVal < len(ops) {
		// fmt.Printf("ip=%v\t%v\t%v", ipVal, reg, ops[ipVal].vals)
		reg[boundReg] = ipVal
		reg = ops[ipVal].behaviour(reg, ops[ipVal].vals)
		ipVal = reg[boundReg]
		ipVal++
		// fmt.Printf("\t%v\n", reg)
	}

	return reg[0]
}

func parseOpInstruction(s string) opInstruction {
	opName := regexp.MustCompile(`[a-z]{4}`).FindStringSubmatch(s)[0]
	stringVals := regexp.MustCompile(`[0-9]+`).FindAllStringSubmatch(s, -1)
	vals := []int{}
	for _, str := range stringVals {
		v, _ := strconv.Atoi(str[0])
		vals = append(vals, v)
	}

	return opInstruction{opName, getOpFromName(opName), vals}
}

func getOpFromName(name string) opFunc {
	var op opFunc
	switch name {
	case "addi":
		op = addi

	case "addr":
		op = addr

	case "muli":
		op = muli

	case "mulr":
		op = mulr

	case "bani":
		op = bani

	case "banr":
		op = banr

	case "bori":
		op = bori

	case "borr":
		op = borr

	case "seti":
		op = seti

	case "setr":
		op = setr

	case "gtir":
		op = gtir

	case "gtri":
		op = gtri

	case "gtrr":
		op = gtrr

	case "eqir":
		op = eqir

	case "eqri":
		op = eqri

	case "eqrr":
		op = eqrr
	}
	return op
}
