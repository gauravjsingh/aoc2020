package s15

import (
	"fmt"
	"strconv"
	"strings"
)

type numbers struct {
	ns   []int
	prev map[int]int
}

func (n *numbers) add(i int) {
	if len(n.ns) > 0 {
		i := len(n.ns) - 1
		n.prev[n.ns[i]] = i
	}
	n.ns = append(n.ns, i)
}

func (n *numbers) next() int {
	i := len(n.ns) - 1
	val := n.ns[i]
	out := 0
	if p, ok := n.prev[val]; ok {
		out = i - p
	}
	n.add(out)
	return out
}

func (n *numbers) extend(goalLen int) int {
	var out int
	for len(n.ns) < goalLen {
		out = n.next()
		//log.Printf("%+v", n)
	}
	return out
}

func newNumbers(ns []int) *numbers {
	nums := &numbers{prev: make(map[int]int)}
	for _, n := range ns {
		nums.add(n)
	}
	return nums
}

func parseInput(ls []string) (*numbers, error) {
	if len(ls) != 1 {
		return nil, fmt.Errorf("expected 1 line of input, got: %v", ls)
	}
	var out []int
	for _, s := range strings.Split(ls[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("error parsing number: %v", err)
		}
		out = append(out, i)
	}
	return newNumbers(out), nil
}

const nth = 30000000

func Solve(ls []string) (int, error) {
	nums, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input: %v", err)
	}
	out := nums.extend(nth)
	return out, nil
}
