package s13

import (
	"fmt"
	"strconv"
	"strings"
)

type buses struct {
	busFreq []int
}

func (b buses) nextBus(start int) (int, int) {
	var bestBus int
	earliest := start + b.busFreq[0]
	for _, f := range b.busFreq {
		if fe := start + (f - (start % f)); fe < earliest {
			bestBus = f
			earliest = fe
		}
	}
	return bestBus, earliest
}

func parseBuses(l string) (buses, error) {
	var freqs []int
	freqStrs := strings.Split(l, ",")
	for _, f := range freqStrs {
		if f == "x" {
			continue
		}
		i, err := strconv.Atoi(f)
		if err != nil {
			return buses{}, fmt.Errorf("error parsing frequency: %v", err)
		}
		freqs = append(freqs, i)
	}
	return buses{busFreq: freqs}, nil
}

func SolveA(ls []string) (int, error) {
	if len(ls) != 2 {
		return 0, fmt.Errorf("expected 2 lines of input, found: %d", len(ls))
	}
	start, err := strconv.Atoi(ls[0])
	if err != nil {
		return 0, fmt.Errorf("error parsing next time: %v", err)
	}
	bs, err := parseBuses(ls[1])
	if err != nil {
		return 0, fmt.Errorf("error parsing buses from %q: %v", ls[1], err)
	}
	bus, earliest := bs.nextBus(start)
	return bus * (earliest - start), nil
}

// Map from offset to frequency.
type busConstraint struct {
	offset, frequency int
}

func parseBusConstraints(l string) ([]busConstraint, error) {
	var bcs []busConstraint
	for offset, f := range strings.Split(l, ",") {
		if f == "x" {
			continue
		}
		freq, err := strconv.Atoi(f)
		if err != nil {
			return nil, fmt.Errorf("error parsing frequency: %v", err)
		}
		bcs = append(bcs, busConstraint{offset: offset, frequency: freq})
	}
	return bcs, nil
}

func earliest(bcs []busConstraint) (int, int) {
	if len(bcs) == 0 {
		return 0, 1
	}
	e, n := earliest(bcs[1:])
	for (e+bcs[0].offset)%bcs[0].frequency != 0 {
		e += n
	}
	// Assumes that these are all relatively prime.
	return e, n * bcs[0].frequency
}

func SolveB(ls []string) (int, error) {
	if len(ls) != 2 {
		return 0, fmt.Errorf("expected 2 lines of input, found: %d", len(ls))
	}
	bcs, err := parseBusConstraints(ls[1])
	if err != nil {
		return 0, fmt.Errorf("error parsing bus constraints from %q: %v", ls[1], err)
	}
	e, _ := earliest(bcs)
	return e, nil
}
