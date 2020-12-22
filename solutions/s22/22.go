package s22

import (
	"aoc2020/reader"
	"fmt"
	"log"
	"strconv"
)

type combat struct {
	d1, d2 []int
	// The strings in this map are output of hash().
	seen map[string]bool
}

// Hash of d1, d2 in combat.
func (c combat) hash() string {
	return fmt.Sprintf("d1: %v, d2: %v", c.d1, c.d2)
}

var globalCache = map[string]int{}

// returns 1 or 2 if their deck has all the cards, or 0 if there is no winner.
func (c *combat) winner() int {
	hash := c.hash()
	if w, ok := globalCache[hash]; ok {
		return w
	}
	w := 0
	for w == 0 {
		w = c.takeTurn()
	}
	globalCache[hash] = w
	return w
}

// assumes that there is no winner. returns 1 or 2 for the winner, 0 if there isn't one.
func (c *combat) takeTurn() int {
	if c.seen[c.hash()] {
		return 1
	}
	if len(c.d1) == 0 {
		return 2
	}
	if len(c.d2) == 0 {
		return 1
	}
	c.seen[c.hash()] = true
	a, b := c.d1[0], c.d2[0]
	c.d1 = c.d1[1:]
	c.d2 = c.d2[1:]
	roundWinner := 0
	if a <= len(c.d1) && b <= len(c.d2) {
		//build and use the result of the new game.
		subC := combat{seen: make(map[string]bool)}
		for _, card := range c.d1[:a] {
			subC.d1 = append(subC.d1, card)
		}
		for _, card := range c.d2[:b] {
			subC.d2 = append(subC.d2, card)
		}
		roundWinner = subC.winner()
	} else if a > b {
		roundWinner = 1
	} else if b > a {
		roundWinner = 2
	} else {
		log.Fatalf("there was a tie")
	}
	if roundWinner == 1 {
		c.d1 = append(c.d1, a, b)
	} else if roundWinner == 2 {
		c.d2 = append(c.d2, b, a)
	} else {
		log.Fatal("each round must have a winner")
	}
	return 0
}

func (c combat) String2() string {
	return fmt.Sprintf("d1: %v,\nd2: %v", c.d1, c.d2)
}

func parseInput(ls []string) (combat, error) {
	gs := reader.GroupInput(ls)
	if len(gs) != 2 {
		return combat{}, fmt.Errorf("unexpected number of decks to parse: %d", len(gs))
	}
	c := combat{seen: make(map[string]bool)}

	for _, l := range gs[0][1:] {
		i, err := strconv.Atoi(l)
		if err != nil {
			return combat{}, fmt.Errorf("error parsing card: %v", err)
		}
		c.d1 = append(c.d1, i)
	}
	for _, l := range gs[1][1:] {
		i, err := strconv.Atoi(l)
		if err != nil {
			return combat{}, fmt.Errorf("error parsing card: %v", err)
		}
		c.d2 = append(c.d2, i)
	}
	return c, nil
}

func Solve(ls []string) (int, error) {
	c, err := parseInput(ls)
	if err != nil {
		return 0, err
	}
	w := c.winner()
	d := c.d1
	if w == 2 {
		d = c.d2
	}
	log.Printf("winner: %d, with deck: %v. evaluated %d games", w, d, len(globalCache))
	score := 0
	for i, c := range d {
		score += c * (len(d) - i)
	}
	return score, nil
}
