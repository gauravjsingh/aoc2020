package s1

import (
	"aoc2020/reader"
	"fmt"
)

func sumProd(ns []int, tot int) (int, error) {
	inputs := make(map[int]bool)
	for _, n := range ns {
		if _, ok := inputs[tot-n]; ok {
			return n * (tot - n), nil
		}
		inputs[n] = true
	}
	return 0, fmt.Errorf("no integers summing to %d found", tot)
}

func sum3Prod(ns []int, tot int) (int, error) {
	inputs1 := make(map[int]int)
	inputs2 := make(map[int][]int)

	for _, n := range ns {
		if ps, ok := inputs2[tot-n]; ok {
			if len(ps) != 1 {
				return 0, fmt.Errorf("multiple products found: %v", ps)
			}
			return inputs2[tot-n][0] * n, nil
		}
		for k, v := range inputs1 {
			inputs2[k+n] = append(inputs2[k+n], v*n)
		}
		inputs1[n] = n
	}
	return 0, fmt.Errorf("no products found for 3 numbers with total %d", tot)
}

func Solve(ls []string) (int, error) {
	ls, err := reader.ReadInput("input/1.txt")
	if err != nil {
		return 0, fmt.Errorf("error reading input: %v", err)
	}
	ns, err := reader.ParseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input: %v", err)
	}
	ans, err := sum3Prod(ns, 2020)
	if err != nil {
		return 0, fmt.Errorf("error solving problem: %v", err)
	}
	return ans, nil
}
