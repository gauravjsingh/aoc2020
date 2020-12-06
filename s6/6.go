package s6

import (
	"fmt"
)

type form struct {
	qs map[rune]int
}

// Does not validate that only questions a-z are in the answers.
func newForm(s string) form {
	f := form{qs: make(map[rune]int)}
	for _, r := range s {
		f.qs[r] += 1
	}
	return f
}

type party struct {
	fs []form
}

func (p party) numYes() map[rune]int {
	m := make(map[rune]int)
	for _, f := range p.fs {
		for r := range f.qs {
			m[r] += 1
		}
	}
	return m
}

func (p party) anyYes() int {
	return len(p.numYes())
}

func (p party) allYes() int {
	m := p.numYes()
	cnt := 0
	for _, c := range m {
		if c == len(p.fs) {
			cnt += 1
		}
	}
	return cnt
}

func parseInput(ls []string) ([]party, error) {
	var p party
	var ps []party
	for _, l := range ls {
		if l == "" {
			ps = append(ps, p)
			p = party{}
			continue
		}
		f := newForm(l)
		p.fs = append(p.fs, f)
	}
	ps = append(ps, p)
	return ps, nil
}

func Solve(ls []string) (int, error) {
	ps, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input as seats: %v", err)
	}
	cnt := 0
	for _, p := range ps {
		cnt += p.allYes()
	}
	return cnt, nil
}
