package s3

import (
	"fmt"
	"log"
)

type grid struct {
	// trees represents the grid of trees. The first dimension is the rows. The columns repeat
	// indefinitely to the right.
	trees [][]bool
}

func (g grid) repWidth() int {
	if len(g.trees) == 0 {
		log.Fatal("invalid grid")
	}
	return len(g.trees[0])
}

func (g grid) rows() int {
	return len(g.trees)
}

func (g grid) isTree(r, c int) bool {
	return g.trees[r][c%g.repWidth()]
}

type runParams struct {
	rowRun, colRun int
}

func (g grid) countTrees(rp runParams, r, c int) int {
	cnt := 0
	r, c = r+rp.rowRun, c+rp.colRun
	for r < g.rows() {
		if g.isTree(r, c) {
			cnt += 1
		}
		r, c = r+rp.rowRun, c+rp.colRun
	}
	return cnt
}

func parseInput(ls []string) (grid, error) {
	var trees [][]bool
	for _, l := range ls {
		var rowTrees []bool
		for _, r := range l {
			if r != '.' && r != '#' {
				return grid{}, fmt.Errorf("unexpected rune in tree grid: %q", r)
			}
			rowTrees = append(rowTrees, r == '#')
		}
		if len(trees) > 0 && (len(trees[0]) != len(rowTrees)) {
			return grid{}, fmt.Errorf("row %q has %d entries, expected %d", l, len(rowTrees), len(trees[0]))
		}
		trees = append(trees, rowTrees)
	}
	return grid{trees: trees}, nil
}

var rps = []runParams{
	{rowRun: 1, colRun: 1},
	{rowRun: 1, colRun: 3},
	{rowRun: 1, colRun: 5},
	{rowRun: 1, colRun: 7},
	{rowRun: 2, colRun: 1},
}

func Solve(ls []string) (int, error) {
	g, err := parseInput(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing input as grid: %v", err)
	}
	prod := 1
	for _, rp := range rps {
		prod *= g.countTrees(rp /*r=*/, 0 /*c=*/, 0)
	}
	return prod, nil
}
