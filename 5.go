package main

import (
	"aoc2020/reader"
	"errors"
	"flag"
	"fmt"
	"log"
	"regexp"
)

var inputPath = flag.String("input_path", "input/5.txt", "path to the input data")

type seat struct {
	row, col int
}

func (s seat) id() int {
	return s.row*8 + s.col
}

func parseBinary(s string, zero, one rune) (int, error) {
	out := 0
	for _, r := range s {
		out *= 2
		if r == one {
			out += 1
		} else if r != zero {
			return 0, fmt.Errorf("unexpected rune %q in %s", r, s)
		}
	}
	return out, nil
}

var re = regexp.MustCompile(`([FB]{7})([RL]{3})`)

func parseSeat(rep string) (seat, error) {
	matches := re.FindStringSubmatch(rep)
	if len(matches) != 3 {
		return seat{}, fmt.Errorf("error matching regex: ", matches)
	}
	row, err := parseBinary(matches[1], 'F', 'B')
	if err != nil {
		return seat{}, fmt.Errorf("error parsing row: %v")
	}
	col, err := parseBinary(matches[2], 'L', 'R')
	if err != nil {
		return seat{}, fmt.Errorf("error parsing col: %v")
	}
	return seat{row: row, col: col}, nil
}

func parseInput(ls []string) ([]seat, error) {
	var ss []seat
	for _, l := range ls {
		s, err := parseSeat(l)
		if err != nil {
			return nil, fmt.Errorf("error parsing seat from: %q", l)
		}
		ss = append(ss, s)
	}
	return ss, nil
}

func maxID(ss []seat) int {
	max := -1
	for _, s := range ss {
		if id := s.id(); id > max {
			max = id
		}
	}
	return max
}

func findSeat(ss []seat) (int, error) {
	ids := make(map[int]bool)
	for _, s := range ss {
		ids[s.id()] = true
	}
	for i := maxID(ss); i > 0; i-- {
		// assume there is only one empty seat.
		if _, ok := ids[i]; !ok {
			return i, nil
		}
	}
	return 0, errors.New("failed to find empty seats")
}

func main() {
	flag.Parse()
	ls, err := reader.ReadInput(*inputPath)
	if err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	ss, err := parseInput(ls)
	if err != nil {
		log.Fatalf("error parsing input as seats: %v", err)
	}
	ans, err := findSeat(ss)
	if err != nil {
		log.Fatal("failed to find seat: %v")
	}
	log.Printf("answer is: %d", ans)
}
