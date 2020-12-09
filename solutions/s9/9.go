package s9

import (
	"fmt"
	"strconv"
)

type encoded struct {
	ns           []int
	preambleSize int
}

func parseInput(ls []string) (encoded, error) {
	var ns []int
	for _, l := range ls {
		n, err := strconv.Atoi(l)
		if err != nil {
			return encoded{}, fmt.Errorf("error parsing int from %q: %v", l, err)
		}
		ns = append(ns, n)
	}
	return encoded{ns: ns, preambleSize: 25}, nil
}

func firstInvalid(e encoded) (int, error) {
	if len(e.ns) < e.preambleSize {
		return 0, fmt.Errorf("e.ns too short, expected preamble of size %d: %v", e.preambleSize, e.ns)
	}
	activeSet := make(map[int]struct{})
	for i := 0; i < e.preambleSize; i++ {
		n := e.ns[i]
		if _, ok := activeSet[n]; ok {
			return 0, fmt.Errorf("unexpected duplicate: %d", n)
		}
		activeSet[n] = struct{}{}
	}
	for i := e.preambleSize; i < len(e.ns); i++ {
		n := e.ns[i]
		valid := false
		for k := range activeSet {
			if _, ok := activeSet[n-k]; ok {
				valid = true
				continue
			}
		}
		if !valid {
			return n, nil
		}

		// Update the active set
		if _, ok := activeSet[n]; ok {
			return 0, fmt.Errorf("unexpected duplicate: %d", n)
		}
		delete(activeSet, e.ns[i-e.preambleSize])
		activeSet[n] = struct{}{}
	}
	return 0, fmt.Errorf("no invalid number found in: %v", e)
}

func contiguous(e encoded, target int) (int, error) {

	return 0, nil
}

func Solve(ls []string) (int, error) {
	enc, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input as seats: %v", err)
	}
	fi, err := firstInvalid(enc)
	if err != nil {
		return 0, err
	}
	return contiguous(enc, fi)
}
