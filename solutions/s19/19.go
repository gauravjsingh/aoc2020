package s19

import (
	"aoc2020/reader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type rule interface {
	toRegexString(rs *rules) string
}

type constant string

func (c constant) toRegexString(rs *rules) string {
	return string(c)
}

type list []int

func (l list) toRegexString(rs *rules) string {
	var res []string
	for _, i := range l {
		res = append(res, fmt.Sprintf("(%s)", rs.regexString(i)))
	}
	return strings.Join(res, "")
}

type or []rule

func (o or) toRegexString(rs *rules) string {
	var res []string
	for _, r := range o {
		res = append(res, fmt.Sprintf("(%s)", r.toRegexString(rs)))
	}
	return strings.Join(res, "|")
}

func parseList(s string) (rule, error) {
	s = strings.TrimSpace(s)
	rs := strings.Split(s, " ")
	var ns []int
	for _, r := range rs {
		n, err := strconv.Atoi(r)
		if err != nil {
			return nil, fmt.Errorf("error parsing rule number (from %q): %v", s, err)
		}
		ns = append(ns, n)
	}
	return list(ns), nil
}

type rules struct {
	rs    map[int]rule
	cache map[int]string
}

func (rs *rules) regexString(i int) string {
	if s, ok := rs.cache[i]; ok {
		return s
	}
	return rs.rs[i].toRegexString(rs)
}

func newRules() *rules {
	return &rules{rs: make(map[int]rule), cache: make(map[int]string)}
}

var ruleRE = regexp.MustCompile(`(?:"(\w+)"|([\d\s|]+))$`)

func parseRule(s string) (rule, error) {
	matches := ruleRE.FindStringSubmatch(s)
	if len(matches) != 3 {
		return nil, fmt.Errorf("error parsing rule: %v", matches)
	}
	if matches[1] != "" {
		if matches[2] != "" {
			return nil, fmt.Errorf("unexpected match result for %q: %v", s, matches)
		}
		return constant(matches[1]), nil
	}
	ps := strings.Split(matches[2], "|")
	if len(ps) == 1 {
		return parseList(matches[2])
	}
	var rs []rule
	for _, p := range ps {
		r, err := parseList(p)
		if err != nil {
			return nil, fmt.Errorf("error parsing sublist %v: %v", p, err)
		}
		rs = append(rs, r)
	}
	return or(rs), nil
}

func parseRules(ls []string) (*rules, error) {
	rs := newRules()
	for _, l := range ls {
		ps := strings.Split(l, ":")
		if len(ps) != 2 {
			return nil, fmt.Errorf("error parsing rule: %q", l)
		}
		n, err := strconv.Atoi(ps[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing rule number: %v", err)
		}
		r, err := parseRule(strings.TrimSpace(ps[1]))
		if err != nil {
			return nil, fmt.Errorf("error parsing rule: %v", err)
		}
		rs.rs[n] = r
	}
	return rs, nil
}

// Hacky override of the values of 8 and 11 to "solve" part B.
func overrideB(rs *rules, iterations int) {
	l := []int{42}
	r := []int{31}
	var rs8 or
	var rs11 or
	for i := 0; i < iterations; i++ {
		rs8 = append(rs8, list(l))
		rs11 = append(rs11, list(append(l, r...)))
		l = append(l, 42)
		r = append(r, 31)
	}
	rs.rs[8] = rs8
	rs.rs[11] = rs11
}

func Solve(ls []string) (int, error) {
	gs := reader.GroupInput(ls)
	if len(gs) != 2 {
		return 0, fmt.Errorf("error grouping input, expected 2 groups, got %d", len(gs))
	}
	rs, err := parseRules(gs[0])
	if err != nil {
		return 0, fmt.Errorf("error parsing rules: %v", err)
	}
	overrideB(rs, 10)
	reStr := rs.rs[0].toRegexString(rs)
	re, err := regexp.Compile(fmt.Sprintf("^%s$", reStr))
	if err != nil {
		return 0, fmt.Errorf("error building regex from %q: %v", reStr, err)
	}
	cnt := 0
	for _, l := range gs[1] {
		if re.MatchString(l) {
			cnt += 1
		}
	}
	return cnt, nil
}
