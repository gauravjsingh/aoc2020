package main

import (
	"aoc2020/reader"
	"aoc2020/registry"
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

var (
	inputPathDir = flag.String("input_dir", "input", "path to the input data directory")
	problem      = flag.Int("problem", -1, "Which problem to solve?")
)

const fileFormat = "%d.txt"

func main() {
	flag.Parse()
	if *problem == -1 {
		*problem = len(registry.Solvers)
	}
	path := filepath.Join(*inputPathDir, fmt.Sprintf(fileFormat, *problem))
	ls, err := reader.ReadInput(path)
	if err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	s, ok := registry.Solvers[*problem]
	if !ok {
		log.Fatalf("solver %d not found", *problem)
	}
	ans, err := s(ls)
	if err != nil {
		log.Fatalf("error solving problem: %v", err)
	}
	log.Printf("answer is: %v", ans)
}
