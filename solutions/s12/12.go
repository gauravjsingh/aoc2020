package s12

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

func (d *direction) rotateR(deg int) error {
	turns := (deg / 90) % 4
	newD := direction((4 + int(*d) + turns) % 4)
	if newD < north || newD > west {
		return fmt.Errorf("unexpected direction: %d", newD)
	}
	*d = newD
	return nil
}

func (d *direction) rotateL(deg int) error { return d.rotateR(-deg) }

func (d direction) String() string {
	switch d {
	case north:
		return "N"
	case east:
		return "E"
	case south:
		return "S"
	case west:
		return "W"
	}
	return fmt.Sprintf("unrecognized direction %d", d)
}

type instruction struct {
	c string
	n int
}

var instructionRE = regexp.MustCompile(`([NSEWLRF])(\d+)`)

func parseInstruction(s string) (instruction, error) {
	matches := instructionRE.FindStringSubmatch(s)
	if len(matches) != 3 {
		return instruction{}, fmt.Errorf("error parsing instructions: wrong number of matches: %v", matches)
	}
	i, err := strconv.Atoi(matches[2])
	if err != nil {
		return instruction{}, fmt.Errorf("error parsing command argument: %v", err)
	}
	return instruction{c: matches[1], n: i}, nil
}

func parseInput(ls []string) ([]instruction, error) {
	var is []instruction
	for _, l := range ls {
		ins, err := parseInstruction(l)
		if err != nil {
			return nil, fmt.Errorf("error parsing instruction: %v", err)
		}
		is = append(is, ins)
	}
	return is, nil
}

type ship struct {
	x, y int
	dir  direction
}

func (s *ship) move(ins instruction) error {
	switch ins.c {
	case "N":
		s.y += ins.n
	case "E":
		s.x += ins.n
	case "S":
		s.y -= ins.n
	case "W":
		s.x -= ins.n
	case "R":
		s.dir.rotateR(ins.n)
	case "L":
		s.dir.rotateL(ins.n)
	case "F":
		return s.move(instruction{c: s.dir.String(), n: ins.n})
	default:
		return fmt.Errorf("unrecognized command: %q", ins.c)
	}
	return nil
}

func SolveA(ls []string) (int, error) {
	is, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input: %v", err)
	}
	s := ship{x: 0, y: 0, dir: east}
	for _, ins := range is {
		s.move(ins)
	}
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y))), nil
}

type point struct {
	x, y int
}

func (p *point) rotateR(deg int) {
	for i := 0; i < deg/90; i++ {
		p.x, p.y = p.y, -p.x
	}
}

func (p *point) rotateL(deg int) {
	p.rotateR(360 - deg)
}

func (p *point) move(ins instruction, pos *point) error {
	switch ins.c {
	case "N":
		p.y += ins.n
	case "E":
		p.x += ins.n
	case "S":
		p.y -= ins.n
	case "W":
		p.x -= ins.n
	case "R":
		p.rotateR(ins.n)
	case "L":
		p.rotateL(ins.n)
	case "F":
		pos.x += p.x * ins.n
		pos.y += p.y * ins.n
	default:
		return fmt.Errorf("unrecognized command: %q", ins.c)
	}
	return nil
}

func SolveB(ls []string) (int, error) {
	is, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input: %v", err)
	}
	waypoint := point{x: 10, y: 1}
	ship := point{}
	for _, ins := range is {
		waypoint.move(ins, &ship)
	}
	return int(math.Abs(float64(ship.x)) + math.Abs(float64(ship.y))), nil
}
