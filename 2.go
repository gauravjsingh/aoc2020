package main

import (
	"aoc2020/reader"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var inputPath = flag.String("input_path", "input/2.txt", "path to the input data")

type sledPolicy struct {
	// number of times s must appear in a valid password.
	min, max int
	r        rune
}

func (p sledPolicy) valid(pass string) bool {
	cnt := 0
	for _, c := range pass {
		if c == p.r {
			cnt += 1
		}
	}
	return p.min <= cnt && cnt <= p.max
}

type tobogganPolicy struct {
	// 1-indexed positions in the string. exactly 1 must have r at the location.
	pos1, pos2 int
	b        byte
}

func (p tobogganPolicy) valid(pass string) bool {
	if p.pos1 > len(pass) || p.pos2 > len(pass) {
		log.Printf("invalid positions for pass %q: %v", pass, p)
		return false
	}
	// note that these are 1 indexed.
	return (pass[p.pos1-1] == p.b) != (pass[p.pos2-1] == p.b)
}

func parsePolicyParts(s string) (int, int, byte, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return 0, 0, 0,  fmt.Errorf("incorrect number of parts to create policy: %v", parts)
	}
	limits := strings.Split(parts[0], "-")
	if len(limits) != 2 {
		return 0, 0, 0,  fmt.Errorf("incorrect number of parts to create limits: %v", limits)
	}
	int1, err := strconv.Atoi(limits[0])
	if err != nil {
		return 0, 0, 0,  fmt.Errorf("error parsing int1 from %q: %v", limits[0], err)
	}
	int2, err := strconv.Atoi(limits[1])
	if err != nil {
		return 0, 0, 0,  fmt.Errorf("error parsing int2 from %q: %v", limits[1], err)
	}
	cs := strings.TrimSpace(parts[1])
	if len(cs) != 1 {
		return 0, 0, 0,  fmt.Errorf("expected only a rune as part of password policy, found: %q", cs)
}
	return int1, int2, cs[0], nil
}

func parseInput(ls []string) (map[tobogganPolicy][]string, error) {
	out := make(map[tobogganPolicy][]string)
	for _, l := range ls {
		ps := strings.Split(l, ":")
		if len(ps) != 2 {
			return nil, fmt.Errorf("unexpected number of parts in %v", ps)
		}
		int1, int2, b, err := parsePolicyParts(strings.TrimSpace(ps[0]))
		if err != nil {
			return nil, fmt.Errorf("error creating policy: %v", err)
		}
		p := tobogganPolicy{pos1: int1, pos2: int2, b: b}
		out[p] = append(out[p], strings.TrimSpace(ps[1]))
	}
	return out, nil
}

func validPasswords(in map[tobogganPolicy][]string) int {
	cnt := 0
	for pol, ps := range in {
		for _, p := range ps {
			if pol.valid(p) {
				cnt+=1
			}
		}
	}
	return cnt
}

func main() {
	flag.Parse()
	ls, err := reader.ReadInput(*inputPath)
	if err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	ps, err := parseInput(ls)
	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}
	cnt := validPasswords(ps)
	log.Printf("answer is: %d", cnt)
}
