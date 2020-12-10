package s10

import (
	"fmt"
	"sort"
	"strconv"
)

type adapters struct {
	as map[int]struct{}
	// Number of ways to get to a target value.
	waysCache map[int]int
}

func (a adapters) max() int {
	max := 0
	for adapter := range a.as {
		if adapter > max {
			max = adapter
		}
	}
	return max
}

func (a adapters) computeDiffs() (map[int]int, error) {
	var as []int
	for a := range a.as {
		as = append(as, a)
	}
	sort.Ints(as)

	var last int
	diffs := make(map[int]int)
	for _, a := range as {
		diffs[a-last]++
		last = a
	}
	return diffs, nil
}

func (a *adapters) ways(target int) int {
	if w, ok := a.waysCache[target]; ok {
		return w
	}
	if _, ok := a.as[target]; !ok {
		return 0
	}
	var out int
	for i := 1; i <= 3; i++ {
		out += a.ways(target - i)
	}
	a.waysCache[target] = out
	return out
}

func newAdapters() *adapters {
	return &adapters{
		as:        make(map[int]struct{}),
		waysCache: map[int]int{0: int(1)},
	}
}

func parseInput(ls []string) (*adapters, error) {
	as := newAdapters()
	for _, l := range ls {
		i, err := strconv.Atoi(l)
		if err != nil {
			return nil, fmt.Errorf("error parsing adapter from %q: %v", l, err)
		}
		as.as[i] = struct{}{}
	}
	return as, nil
}

func SolveA(ls []string) (int, error) {
	as, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input as adapters: %v", err)
	}

	// The phone is the max adapter + 3.
	diffs, err := as.computeDiffs()
	if err != nil {
		return 0, fmt.Errorf("error computing diffs for adapters %v: %v", as, err)
	}
	return diffs[1] * (diffs[3] + 1), nil
}

func Solve(ls []string) (int, error) {
	as, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input as adapters: %v", err)
	}

	// The number of ways to reach our charger is the number of ways to reach the max adapter since
	// they differ by 3.
	return as.ways(as.max()), nil
}
