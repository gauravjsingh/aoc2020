package s25

import (
	"fmt"
	"strconv"
)

func parseInput(ls []string) ([]int, error) {
	if len(ls) != 2 {
		return nil, fmt.Errorf("wrong number of public keys in: %v", ls)
	}
	var out []int
	for _, l := range ls {
		i, err := strconv.Atoi(l)
		if err != nil {
			return nil, fmt.Errorf("error parsing public key: %v", err)
		}
		out = append(out, i)
	}
	return out, nil
}

var modulus = 20201227

func iterate(v, base int) int {
	return (v * base) % modulus
}

func powMod(base, exp int) int {
	v := 1
	for i := 0; i < exp; i++ {
		v = iterate(v, base)
	}
	return v
}

func SolveA(ls []string) (int, error) {
	pks, err := parseInput(ls)
	if err != nil {
		return 0, err
	}
	v := 1
	for i := 1; i < modulus; i++ {
		v = iterate(v, 7)
		if v == pks[0] {
			return powMod(pks[1], i), nil
		}
	}
	return 0, fmt.Errorf("did not find exponent for first pk: %v", pks[0])
}
