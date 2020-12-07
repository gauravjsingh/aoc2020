package solutions

import (
	"aoc2020/solutions/s1"
	"aoc2020/solutions/s2"
	"aoc2020/solutions/s3"
	"aoc2020/solutions/s4"
	"aoc2020/solutions/s5"
	"aoc2020/solutions/s6"
	"aoc2020/solutions/s7"
	"fmt"
)

type Solver func([]string) (string, error)

func intWrap(f func([]string) (int, error)) Solver {
	return func(ls []string) (string, error) {
		i, err := f(ls)
		return fmt.Sprint(i), err
	}
}

var Solvers = map[int]Solver{
	1: intWrap(s1.Solve),
	2: intWrap(s2.Solve),
	3: intWrap(s3.Solve),
	4: intWrap(s4.Solve),
	5: intWrap(s5.Solve),
	6: intWrap(s6.Solve),
	7: intWrap(s7.Solve),
}
