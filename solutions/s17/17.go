package s17

import (
	"flag"
	"log"
)

var rounds = flag.Int("rounds", 6, "number of evolution rounds")

type point struct {
	x, y, z, w int
}

func (p point) neighbors() []point {
	var ps []point
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					newP := point{p.x + i, p.y + j, p.z + k, p.w + l}
					if p == newP {
						continue
					}
					ps = append(ps, newP)
				}
			}
		}
	}
	return ps
}

type conway struct {
	cubes map[point]bool
	// Only values between (-size, -size, -size) and (size, size, size) (inclusive) can be active.
	size int
}

func (c conway) activeNeighbors(p point) int {
	active := 0
	for _, n := range p.neighbors() {
		// This is false if this is not in the current bounds.
		if c.cubes[n] {
			active++
		}
	}
	return active
}

func (c conway) allPoints(extraDims int) []point {
	min, max := -c.size-extraDims, c.size+extraDims
	var ps []point
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			for k := min; k <= max; k++ {
				for l := min; l <= max; l++ {
					ps = append(ps, point{i, j, k, l})
				}
			}
		}
	}
	return ps
}

func (c conway) evolve() conway {
	e := conway{cubes: make(map[point]bool), size: c.size + 1}
	for _, p := range c.allPoints(1) {
		ns := c.activeNeighbors(p)
		if (c.cubes[p] && (ns == 2 || ns == 3)) || (!c.cubes[p] && ns == 3) {
			e.cubes[p] = true
		} else {
			e.cubes[p] = false
		}
	}
	return e
}

func parseInput(ls []string) conway {
	c := conway{cubes: make(map[point]bool), size: len(ls)}
	for i, l := range ls {
		if len(l) > c.size {
			log.Printf("updating size from %d to %d", c.size, len(l))
			c.size = len(l)
		}
		for j, r := range l {
			c.cubes[point{x: i, y: j}] = r == '#'
		}
	}
	return c
}

func Solve(ls []string) (int, error) {
	c := parseInput(ls)
	for i := 0; i < *rounds; i++ {
		c = c.evolve()
	}
	active := 0
	for _, b := range c.cubes {
		if b {
			active++
		}
	}
	return active, nil
}
