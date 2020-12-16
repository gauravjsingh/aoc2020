package s16

import (
	"aoc2020/reader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type intRange struct {
	min, max int
}

type constraint struct {
	name string
	rs   []intRange
}

func (c constraint) valid(n int) bool {
	for _, r := range c.rs {
		if r.min <= n && n <= r.max {
			return true
		}
	}
	return false
}

type constraints []constraint

func (cs constraints) valid(n int) bool {
	for _, c := range cs {
		if c.valid(n) {
			return true
		}
	}
	return false
}

var constraintRE = regexp.MustCompile(`^([ \w]+): (\d+)-(\d+) or (\d+)-(\d+)`)

func parseConstraint(l string) (constraint, error) {
	matches := constraintRE.FindStringSubmatch(l)
	if len(matches) != 6 {
		return constraint{}, fmt.Errorf("error matching constraintRE: %v", matches)
	}
	c := constraint{name: matches[1]}
	for i := 2; i+1 < len(matches); i += 2 {
		min, err := strconv.Atoi(matches[i])
		if err != nil {
			return constraint{}, fmt.Errorf("error parsing range min: %v", err)
		}
		max, err := strconv.Atoi(matches[i+1])
		if err != nil {
			return constraint{}, fmt.Errorf("error parsing range max: %v", err)
		}
		c.rs = append(c.rs, intRange{min: min, max: max})
	}
	return c, nil
}

type ticket []int

func (t ticket) valid(cs constraints) bool {
	for _, n := range t {
		if !cs.valid(n) {
			return false
		}
	}
	return true
}

func (t ticket) updatePossible(possible []map[string]constraint) error {
	if len(t) != len(possible) {
		return fmt.Errorf("ticket and possible have different lengths: %v, %v", t, possible)
	}
	for i, n := range t {
		var invalid []string
		for k, c := range possible[i] {
			if !c.valid(n) {
				invalid = append(invalid, k)
			}
		}
		for _, inv := range invalid {
			delete(possible[i], inv)
		}
	}
	return nil
}

func parseTicket(l string) (ticket, error) {
	var t ticket
	for _, s := range strings.Split(l, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("error parsing ticket value: %v", err)
		}
		t = append(t, i)
	}
	return t, nil
}

func parseInput(ls []string) (constraints, ticket, []ticket, error) {
	gs := reader.GroupInput(ls)
	if len(gs) != 3 {
		return nil, ticket{}, nil, fmt.Errorf("error grouping input: %v", gs)
	}
	var cs constraints
	for _, l := range gs[0] {
		c, err := parseConstraint(l)
		if err != nil {
			return nil, ticket{}, nil, err
		}
		cs = append(cs, c)
	}
	if len(gs[1]) != 2 {
		return nil, ticket{}, nil, fmt.Errorf("invalid in your ticket group: %v", gs[1])
	}
	yt, err := parseTicket(gs[1][1])
	if err != nil {
		return nil, ticket{}, nil, fmt.Errorf("error parsing your ticket: %v", err)
	}
	var ts []ticket
	for _, l := range gs[2][1:] {
		t, err := parseTicket(l)
		if err != nil {
			return nil, ticket{}, nil, fmt.Errorf("error parsing other ticket %q: %v", l, err)
		}
		ts = append(ts, t)
	}
	return cs, yt, ts, nil
}

func SolveA(ls []string) (int, error) {
	cs, _, ts, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("parsing error: %v", err)
	}
	invalidSum := 0
	for _, t := range ts {
		for _, n := range t {
			if !cs.valid(n) {
				invalidSum += n
			}
		}
	}
	return invalidSum, nil
}

func resolvePossible(possible []map[string]constraint) (map[string]int, error) {
	found := make(map[string]int)
	for len(found) < len(possible) {
		for i, m := range possible {
			for k := range m {
				if v, ok := found[k]; ok && v != i {
					delete(m, k)
				}
			}
			if len(m) == 0 {
				return nil, fmt.Errorf("error resolving possible for index %d\n%v", i, found)
			}
			if len(m) == 1 {
				for k := range m {
					found[k] = i
				}
			}
		}
	}
	return found, nil
}

func SolveB(ls []string) (int, error) {
	cs, yt, ts, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("parsing error: %v", err)
	}
	var valid []ticket
	for _, t := range ts {
		if t.valid(cs) {
			valid = append(valid, t)
		}
	}
	var possible []map[string]constraint
	for range yt {
		cMap := make(map[string]constraint)
		for _, c := range cs {
			cMap[c.name] = c
		}
		possible = append(possible, cMap)
	}
	for _, t := range valid {
		t.updatePossible(possible)
	}
	mapping, err := resolvePossible(possible)
	if err != nil {
		return 0, fmt.Errorf("error resoloving possible: %v", err)
	}
	prod := 1
	for c, i := range mapping {
		if strings.HasPrefix(c, "departure") {
			prod *= yt[i]
		}
	}
	return prod, nil
}
