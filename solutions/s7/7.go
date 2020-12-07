package s7

import (
	"fmt"
	"regexp"
	"strconv"
)

// eg. "bright white"
type color string

type bags struct {
	bs map[color]map[color]int

	// recursive contents of a bag: rc[c1][c2] has known values of how many times c1 recursively
	// contains c2. Missing values are unknown.
	rc map[color]map[color]int
}

// Note, could be more efficient by not redoing all computations.
// The map returned must not be modified.
func (bs *bags) recursiveContents(c color) map[color]int {
	rc, ok := bs.rc[c]
	if ok {
		return rc
	}
	rc = make(map[color]int)
	bs.rc[c] = rc

	for contentColor, i := range bs.bs[c] {
		rc[contentColor] += i
		for innerContentColor, n := range bs.recursiveContents(contentColor) {
			rc[innerContentColor] += n * i
		}
	}
	return rc
}

func (bs *bags) numContainingColor(target color) int {
	cnt := 0
	for c := range bs.bs {
		if bs.recursiveContents(c)[target] > 0 {
			cnt += 1
		}
	}
	return cnt
}

func newBags() bags {
	return bags{
		bs: make(map[color]map[color]int),
		rc: make(map[color]map[color]int),
	}
}

var contentsRE = regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)

func parseContents(s string) (map[color]int, error) {
	matches := contentsRE.FindAllStringSubmatch(s, -1)
	contents := make(map[color]int)
	for _, m := range matches {
		if len(m) != 3 {
			return nil, fmt.Errorf("error getting contents from %q: %v", s, matches)
		}
		i, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing number of bags from: %q", m[1])
		}
		contents[color(m[2])] = i
	}
	return contents, nil
}

var ruleRE = regexp.MustCompile(`(\w+ \w+) bags contain (?:(no other bags)|([\w\d ,]+)).`)

func newBag(s string) (color, map[color]int, error) {
	matches := ruleRE.FindStringSubmatch(s)
	if len(matches) != 4 {
		return "", nil, fmt.Errorf("matches had wrong length: %v", matches)
	}
	c := color(matches[1])
	if len(matches[2]) != 0 {
		return c, map[color]int{}, nil
	}
	b, err := parseContents(matches[3])
	return c, b, err
}

func parseInput(ls []string) (bags, error) {
	bs := newBags()
	for _, l := range ls {
		c, b, err := newBag(l)
		if err != nil {
			return bags{}, fmt.Errorf("error parsing bag from %q: %v", l, err)
		}
		bs.bs[c] = b
	}
	return bs, nil
}

const targetColor = "shiny gold"

func Solve(ls []string) (int, error) {
	bs, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input as bags: %v", err)
	}

	cnt := 0
	for _, i := range bs.recursiveContents(targetColor) {
		cnt += i
	}
	return cnt, nil
}
