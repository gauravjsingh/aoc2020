package main

import (
	"aoc2020/input"
	"fmt"
	"log"
)

func sumProd(ns []int, tot int) (int, error) {
	var inputs map[int]bool
	for _, n := range ns {
		if _, ok := input[tot-n]; ok {
			return n * (tot - n), nil
		}
		inputs[n] = true
	}
	return 0, fmt.Errorf("no integers summing to %d found", tot)
}

func main() {
	ls, err := input.ReadInput("input/1.txt")
	if err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	ns, err := input.ParseInput(ls)
	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}
	ans, err := sumProd(ns, 2020)
	if err != nil {
		log.Fatalf("error solving problem: %v", err)
	}
	log.Printf("answer is: %d", ans)
}
