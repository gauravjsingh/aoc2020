package s11

import (
	"log"
	"strings"
)

type seat rune

const (
	empty    seat = 'L'
	occupied seat = '#'
	floor    seat = '.'
)

type seatLayout struct {
	seats [][]seat
}

func (s seatLayout) occupiedNeighbors(row, col int) int {
	out := 0
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			// Check bounds.
			if r < 0 || r >= len(s.seats) || c < 0 || c >= len(s.seats[r]) {
				continue
			}
			// Ignore the square itself.
			if r == row && c == col {
				continue
			}
			if s.seats[r][c] == occupied {
				out += 1
			}
		}
	}
	return out
}

func (s seatLayout) inBounds(row, col int) bool {
	return row >= 0 && row < len(s.seats) && col >= 0 && col < len(s.seats[row])
}

func (s seatLayout) visibleOccupied(row, col int) int {
	visible := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			for r, c := row+x, col+y; s.inBounds(r, c); r, c = r+x, c+y {
				entry := s.seats[r][c]
				if entry == floor {
					continue
				}
				if entry == occupied {
					visible++
				}
				break
			}
		}
	}
	return visible
}

// returns the new seat layout as well as whether it changed.
func (s seatLayout) iterate() (seatLayout, bool) {
	out := seatLayout{}
	change := false
	for r := range s.seats {
		out.seats = append(out.seats, make([]seat, len(s.seats[r])))
		for c := range s.seats[r] {
			if s.seats[r][c] == floor {
				out.seats[r][c] = floor
				continue
			}
			if ns := s.visibleOccupied(r, c); ns == 0 {
				out.seats[r][c] = occupied
			} else if ns >= 5 {
				out.seats[r][c] = empty
			} else {
				out.seats[r][c] = s.seats[r][c]
			}
			if out.seats[r][c] != s.seats[r][c] {
				change = true
			}
		}
	}
	return out, change
}

func (s seatLayout) numOccupied() int {
	out := 0
	for _, row := range s.seats {
		for _, seat := range row {
			if seat == occupied {
				out++
			}
		}
	}
	return out
}

func (s seatLayout) String() string {
	var rows []string
	for _, r := range s.seats {
		var row []rune
		for _, seat := range r {
			row = append(row, rune(seat))
		}
		rows = append(rows, string(row))
	}
	return strings.Join(rows, "\n")
}

func parseInput(ls []string) seatLayout {
	out := seatLayout{}
	for _, l := range ls {
		var row []seat
		for _, s := range l {
			row = append(row, seat(s))
		}
		out.seats = append(out.seats, row)
	}
	return out
}

func Solve(ls []string) (int, error) {
	sl := parseInput(ls)
	change := true
	iterations := 0
	for change {
		sl, change = sl.iterate()
		iterations++
	}
	log.Printf("took %d iterations", iterations)
	return sl.numOccupied(), nil
}
