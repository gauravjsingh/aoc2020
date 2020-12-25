package s23

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type node struct {
	n    int
	next *node
}

func (n *node) elide(num int) {
	for i := 0; i < num; i++ {
		n.next = n.next.next
	}
}

// for now assume there are size cups, 1-size
// current cup is head
type cups struct {
	head *node
	size int
	nMap map[int]*node
}

func (c cups) lower(n int) int {
	if n-1 < 1 {
		return c.size
	}
	return n - 1
}

func (c cups) dest(val int, move *node) int {
	m := make(map[int]bool)
	for i := 0; i < 3; i++ {
		m[move.n] = true
		move = move.next
	}
	for val = c.lower(val); m[val]; val = c.lower(val) {
	}
	return val
}

func (c *cups) splice(head, next *node, size int) {
	tail := head.next
	head.next = next
	// We only need to iterate size-1 times to find the last node to be moved.
	for i := 0; i < size-1; i++ {
		next = next.next
	}
	next.next = tail
}

func (c *cups) iterate() {
	moveHead := c.head.next
	c.head.elide(3)
	d := c.nMap[c.dest(c.head.n, moveHead)]
	c.splice(d, moveHead, 3)
	c.head = c.head.next
}

func (c cups) String() string {
	out := fmt.Sprint(c.head.n)
	for n := c.head.next; n != c.head; n = n.next {
		out = out + fmt.Sprint(n.n)
	}
	return out
}

// start from the cup after 1.
func (c cups) AnswerString() (string, error) {
	s := c.String()
	pcs := strings.Split(s, "1")
	if len(pcs) != 2 {
		return "", fmt.Errorf("invalid cup arrangement: %s", s)
	}
	return pcs[1] + pcs[0], nil
}

func parseCups(l string, size int) (cups, error) {
	c := cups{size: size, nMap: make(map[int]*node)}
	var head, prev *node
	for _, ch := range l {
		i, err := strconv.Atoi(string(ch))
		if err != nil {
			return cups{}, fmt.Errorf("error parsing cup: %v", err)
		}
		n := &node{n: i}
		c.nMap[i] = n
		if c.head == nil {
			c.head = n
			prev = c.head
			head = c.head
			continue
		}
		prev.next = n
		prev = prev.next
	}
	for i := len(l) + 1; i <= size; i++ {
		n := &node{n: i}
		c.nMap[i] = n
		prev.next = n
		prev = prev.next
	}
	prev.next = head
	return c, nil
}

func SolveA(ls []string) (string, error) {
	if len(ls) != 1 {
		return "", fmt.Errorf("invalid input: %v", ls)
	}
	cs, err := parseCups(ls[0], 9)
	if err != nil {
		return "", err
	}
	for i := 0; i < 100; i++ {
		log.Printf("cups: %s", cs)
		cs.iterate()
	}
	return cs.AnswerString()
}

func SolveB(ls []string) (int, error) {
	if len(ls) != 1 {
		return 0, fmt.Errorf("invalid input: %v", ls)
	}
	cs, err := parseCups(ls[0], 1000*1000)
	if err != nil {
		return 0, err
	}
	for i := 0; i < 10*1000*1000; i++ {
		cs.iterate()
	}
	n1 := cs.nMap[1]
	return n1.next.n * n1.next.next.n, nil
}
