package s24

import (
	"flag"
	"log"
	"math"
)

var iterations = flag.Int("iterations", 100, "number of iterations of game of life")

type direction string

const (
	nw direction = "nw"
	ne direction = "ne"
	e  direction = "e"
	se direction = "se"
	sw direction = "sw"
	w  direction = "w"
)

var directions = []direction{nw, ne, e, se, sw, w}

type point struct {
	x, y int
}

func (p point) size() int {
	s := 0
	if ax := int(math.Abs(float64(p.x))); ax > s {
		s = ax
	}
	if ay := int(math.Abs(float64(p.y))); ay > s {
		s = ay
	}
	return s
}

// y axis goes ne
func (p point) add(d direction) point {
	out := p
	switch d {
	case nw:
		out.x--
		out.y++
	case ne:
		out.y++
	case e:
		out.x++
	case se:
		out.x++
		out.y--
	case sw:
		out.y--
	case w:
		out.x--
	default:
		log.Fatalf("invalid direction: %v", d)
	}
	return out
}

func parseDirs(l string) []direction {
	var dirs []direction
	for i := 0; i < len(l); i++ {
		if l[i] == 'n' || l[i] == 's' {
			dirs = append(dirs, direction(l[i:i+2]))
			i++
			continue
		}
		dirs = append(dirs, direction(l[i:i+1]))
	}
	return dirs
}

// black is true, white is false
type grid map[point]bool

func (g grid) blackTiles() int {
	black := 0
	for _, v := range g {
		if v {
			black++
		}
	}
	return black
}

func (g grid) size() int {
	max := 0
	for k := range g {
		if size := k.size(); size > max {
			max = size
		}
	}
	return max
}

func (g grid) neighbors(p point) int {
	out := 0
	for _, d := range directions {
		if g[p.add(d)] {
			out += 1
		}
	}
	return out
}

func (g grid) update() grid {
	newSize := g.size() + 1
	newG := make(grid)
	for i := -newSize; i <= newSize; i++ {
		for j := -newSize; j <= newSize; j++ {
			p := point{i, j}
			if n := g.neighbors(p); n == 2 || (g[p] && n == 1) {
				newG[p] = true
			}
		}
	}
	return newG
}

func parseGrid(ls []string) grid {
	g := make(grid)
	for _, l := range ls {
		p := point{}
		for _, d := range parseDirs(l) {
			p = p.add(d)
		}
		g[p] = !g[p]
	}
	return g
}

func SolveA(ls []string) (int, error) {
	g := parseGrid(ls)
	return g.blackTiles(), nil
}

func SolveB(ls []string) (int, error) {
	g := parseGrid(ls)
	for i := 0; i < *iterations; i++ {
		g = g.update()
	}
	return g.blackTiles(), nil
}
