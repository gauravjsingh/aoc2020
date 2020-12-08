package s8

import (
	"fmt"
	"regexp"
	"strconv"
)

type operation int

const (
	nop operation = iota
	acc
	jmp
)

func parseOp(s string) (operation, error) {
	switch s {
	case "nop":
		return nop, nil
	case "acc":
		return acc, nil
	case "jmp":
		return jmp, nil
	}
	return 0, fmt.Errorf("unrecognized operation: %q", s)
}

type instruction struct {
	op  operation
	arg int
}

var re = regexp.MustCompile(`(\w+) ((?:-|\+)\d+)`)

func parseInput(ls []string) ([]instruction, error) {
	var is []instruction
	for _, l := range ls {
		matches := re.FindStringSubmatch(l)
		if len(matches) != 3 {
			return nil, fmt.Errorf("wrong number of matches: %v", matches)
		}
		op, err := parseOp(matches[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing operation: %v", err)
		}
		i, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("error converting %q to an int: %v", matches[2], err)
		}
		is = append(is, instruction{op: op, arg: i})
	}
	return is, nil
}

func lastAcc(is []instruction) (int, error) {
	sum := 0
	ptr := 0
	seen := make(map[int]struct{})
	var ok bool
	for !ok {
		if ptr >= len(is) || ptr < 0 {
			return 0, fmt.Errorf("index %d out of range", ptr)
		}
		seen[ptr] = struct{}{}
		ins := is[ptr]
		switch ins.op {
		case nop:
			ptr += 1
		case acc:
			sum += ins.arg
			ptr += 1
		case jmp:
			ptr += ins.arg
		}
		_, ok = seen[ptr]
	}
	return sum, nil
}

func Solve(ls []string) (int, error) {
	is, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input as seats: %v", err)
	}
	return lastAcc(is)
}
