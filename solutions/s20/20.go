package s20

import (
	"aoc2020/reader"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type edge [10]bool

func (e edge) hash() int {
	var n int
	for i := range e {
		n *= 2
		if e[i] {
			n++
		}
	}
	return n
}

func (e edge) String() string {
	return fmt.Sprint(e.hash())
}

func unorderedEdge(e edge) edge {
	var rev edge
	for i := range e {
		rev[i] = e[10-1-i]
	}
	if rev.hash() < e.hash() {
		return rev
	}
	return e
}

type tileEdges struct {
	left, top, right, bottom edge
}

func (te tileEdges) String() string {
	return fmt.Sprintf("left: %q, top: %q, right: %q, bottom: %q", te.left, te.top, te.right, te.bottom)
}

// # is true, . is false
type tile [10][10]bool

func (t tile) edges() tileEdges {
	var left, right, top, bottom edge
	top = t[0]
	bottom = t[9]
	for i, line := range t {
		left[i] = line[0]
		right[i] = line[9]
	}
	return tileEdges{left, top, right, bottom}
}

func (t tile) unorderedEdges() []edge {
	var es []edge
	te := t.edges()
	for _, e := range []edge{te.left, te.top, te.right, te.bottom} {
		es = append(es, unorderedEdge(e))
	}
	return es
}

func (t tile) verticalFlip() tile {
	var newT tile
	for i := 0; i < len(t); i++ {
		newT[i] = t[9-i]
	}
	return newT
}

func (t tile) rotateR() tile {
	var newT tile
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[0]); j++ {
			newT[i][j] = t[9-j][i]
		}
	}
	return newT
}

// 0-7 correspond to applying different elements of D4, where i is a^xb^y where x = i%2, and y = i/2.
func (t tile) positions() []tile {
	ts := []tile{t, t.verticalFlip()}

	for i := 0; i < 3; i++ {
		ts = append(ts, ts[2*i].rotateR(), ts[2*i+1].rotateR())
	}
	return ts
}

func (t tile) data() [8][8]bool {
	var d [8][8]bool
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[0]); j++ {
			d[i][j] = t[i+1][j+1]
		}
	}
	return d
}

func parseTile(ls []string) tile {
	t := tile{}
	for i, l := range ls {
		for j, ch := range l {
			t[i][j] = ch == '#'
		}
	}
	return t
}

type edgeMap map[edge][]int

func (em edgeMap) boundary(e edge) bool {
	return len(em[e]) == 1
}

type tileSet map[int]tile

// Return a map from edge to the tiles with that edge.
func (ts tileSet) edgeMapping() edgeMap {
	m := make(edgeMap)
	for id, t := range ts {
		for _, e := range t.unorderedEdges() {
			m[e] = append(m[e], id)
		}
	}
	return m
}

func (ts tileSet) corners() []int {
	em := ts.edgeMapping()
	var cs []int
	for id, t := range ts {
		cnt := 0
		for _, e := range t.unorderedEdges() {
			if len(em[e]) == 2 {
				cnt += 1
			}
		}
		if cnt == 2 {
			cs = append(cs, id)
		}
	}
	return cs
}

var tileIDRE = regexp.MustCompile(`\d+`)

func parseTileSet(ls []string) (tileSet, error) {
	gs := reader.GroupInput(ls)
	ts := make(tileSet)
	for index, g := range gs {
		if len(g) == 0 {
			return nil, fmt.Errorf("group %d is invalid tile", index)
		}
		id := tileIDRE.FindString(g[0])
		i, err := strconv.Atoi(id)
		if err != nil {
			return nil, fmt.Errorf("error parsing tile ID: %v", err)
		}
		ts[i] = parseTile(g[1:])
	}
	return ts, nil
}

type image [][]bool

func (im image) verticalFlip() image {
	var newIM image
	for i := 0; i < len(im); i++ {
		newIM = append(newIM, im[len(im)-1-i])
	}
	return newIM
}

func (im image) rotateSquareR() image {
	size := len(im)
	var newGrid image
	for i := 0; i < size; i++ {
		newGrid = append(newGrid, make([]bool, size))
		for j := 0; j < size; j++ {
			newGrid[i][j] = im[size-1-j][i]
		}
	}
	return newGrid
}

var seaMonsterStrings = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

func seaMonster() [][]bool {
	var out [][]bool
	for _, s := range seaMonsterStrings {
		var row []bool
		for _, ch := range s {
			row = append(row, ch == '#')
		}
		out = append(out, row)
	}
	return out
}

func (im image) seaMonsterAt(x, y int) bool {
	sm := seaMonster()
	for i := 0; i < len(sm); i++ {
		for j := 0; j < len(sm[0]); j++ {
			if x+i >= len(im) || y+j >= len(im[0]) {
				return false
			}
			if sm[i][j] && !im[x+i][y+j] {
				return false
			}
		}
	}
	return true
}

func (im image) seaMonsters() int {
	cnt := 0
	for i := 0; i < len(im); i++ {
		for j := 0; j < len(im[0]); j++ {
			if im.seaMonsterAt(i, j) {
				cnt++
			}
		}
	}
	return cnt
}

func (im image) String() string {
	var rows []string
	for _, r := range im {
		s := ""
		for _, b := range r {
			if b {
				s += "#"
			} else {
				s += "."
			}
		}
		rows = append(rows, s)
	}
	return strings.Join(rows, "\n")
}

func newImage(ts [][]tile) image {
	var im image
	for range ts {
		for j := 0; j < 8; j++ {
			im = append(im, make([]bool, len(ts[0])*8))
		}
	}
	for i, row := range ts {
		for j, t := range row {
			for di, dataRow := range t.data() {
				for dj, v := range dataRow {
					im[i*8+di][j*8+dj] = v
				}
			}
		}
	}
	return im
}

func findTile(left, top *int, em edgeMap, ts tileSet) (int, tile, error) {
	var id int
	if left == nil && top == nil {
		id = ts.corners()[0]
	} else if left != nil {
		l := ts[*left]
		for _, i := range em[unorderedEdge(l.edges().right)] {
			if i != *left {
				id = i
			}
		}
	} else {
		t := ts[*top]
		for _, i := range em[unorderedEdge(t.edges().bottom)] {
			if i != *top {
				id = i
			}
		}
	}

	for _, t := range ts[id].positions() {
		es := t.edges()
		if left == nil && len(em[unorderedEdge(es.left)]) != 1 {
			continue
		}
		if left != nil && es.left != ts[*left].edges().right {
			continue
		}
		if top == nil && len(em[unorderedEdge(es.top)]) != 1 {
			continue
		}
		if top != nil && es.top != ts[*top].edges().bottom {
			continue
		}
		return id, t, nil
	}
	return 0, tile{}, fmt.Errorf("did not find tile for left, top: %d, %d", left, top)
}

const numPositions = 8

func buildImage(ts tileSet, size int) (image, error) {
	var tiles [][]tile
	var tileIDs [][]int
	for i := 0; i < size; i++ {
		tiles = append(tiles, make([]tile, size))
		tileIDs = append(tileIDs, make([]int, size))
	}
	em := ts.edgeMapping()

	for i := range tiles {
		for j := range tiles[0] {
			var left, top *int
			if i > 0 {
				top = &tileIDs[i-1][j]
			}
			if j > 0 {
				left = &tileIDs[i][j-1]
			}
			id, t, err := findTile(left, top, em, ts)
			if err != nil {
				return image{}, err
			}
			ts[id] = t
			tiles[i][j] = t
			tileIDs[i][j] = id
		}
	}

	return newImage(tiles), nil
}

func SolveA(ls []string) (int, error) {
	ts, err := parseTileSet(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing tileset: %v", err)
	}
	cs := ts.corners()
	log.Printf("corners: %v", cs)
	prod := 1
	for _, c := range cs {
		prod *= c
	}
	return prod, nil
}

func SolveB(ls []string) (int, error) {
	ts, err := parseTileSet(ls)
	if err != nil {
		return 0, fmt.Errorf("error parsing tileset: %v", err)
	}
	size := int(math.Sqrt(float64(len(ts))))
	im, err := buildImage(ts, size)
	if err != nil {
		return 0, fmt.Errorf("error constructing image: %v", err)
	}
	log.Printf("image:\n%v", im)

	pos := []image{im, im.verticalFlip()}
	for i := 0; i < 3; i++ {
		pos = append(pos, pos[2*i].rotateSquareR(), pos[2*i+1].rotateSquareR())
	}

	for _, p := range pos {
		if p.seaMonsters() > 0 {
			im = p
			break
		}
	}

	cnt := 0
	for _, r := range im {
		for _, v := range r {
			if v {
				cnt++
			}
		}
	}
	sms := im.seaMonsters()
	ans := cnt - sms*15
	log.Printf("#s in grid: %d, sea monsters: %d", cnt, sms)
	return ans, nil
}
