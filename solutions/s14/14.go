package s14

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type state struct {
	bm bitmask
	rs map[int]int
}

type command interface {
	Apply(*state)
}

type bitmask map[int]int

func (b bitmask) maskInt(n int) int {
	for i, mask := range b {
		if mask == 0 {
			n &^= 1 << i
		} else if mask == 1 {
			n |= 1 << i
		} else {
			log.Fatalf("unexpected mask value: %v", mask)
		}
	}
	return n
}

func (b bitmask) Apply(s *state) {
	s.bm = b
}

var maskRE = regexp.MustCompile(`[X01]{36}`)

func parseMask(l string) (bitmask, error) {
	b := make(bitmask)
	m := maskRE.FindString(l)
	if len(m) != 36 {
		return nil, fmt.Errorf("failed to extract mask from %q: %q", l, m)
	}
	for i, r := range m {
		if r == 'X' {
			continue
		}
		val := 0
		if r == '1' {
			val = 1
		}
		// 36 bit ints
		b[36-1-i] = val
	}
	return b, nil
}

type set struct {
	register, value int
}

func (s set) Apply(st *state) {
	st.rs[s.register] = st.bm.maskInt(s.value)
}

var cmdRE = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func parseCommand(l string) (set, error) {
	matches := cmdRE.FindStringSubmatch(l)
	if len(matches) != 3 {
		return set{}, fmt.Errorf("error matching set from %q: %v", l, matches)
	}
	r, err := strconv.Atoi(matches[1])
	if err != nil {
		return set{}, fmt.Errorf("error parsing register: %v", err)
	}
	v, err := strconv.Atoi(matches[2])
	if err != nil {
		return set{}, fmt.Errorf("error parsing value: %v", err)
	}
	return set{register: r, value: v}, nil
}

func parseInput(ls []string) ([]command, error) {
	var cs []command
	for _, l := range ls {
		bm, bmErr := parseMask(l)
		if bmErr == nil {
			cs = append(cs, bm)
			continue
		}
		c, cErr := parseCommand(l)
		if cErr != nil {
			return nil, fmt.Errorf("error parsing %q as command: %v, %v", l, bmErr, cErr)
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func Solve(ls []string) (int, error) {
	cs, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input: %v", err)
	}
	s := state{rs: make(map[int]int), bm: make(bitmask)}
	for _, c := range cs {
		c.Apply(&s)
	}
	sum := 0
	for _, v := range s.rs {
		sum += v
	}
	return sum, nil
}
