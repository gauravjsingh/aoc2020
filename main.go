package main

import (
	"aoc2020/reader"
	"aoc2020/s6"
	"flag"
	"log"
	"path/filepath"
)

var (
	inputPathDir = flag.String("input_path_dir", "input", "path to the input data directory")
	inputFile    = flag.String("input", "", "name of the input file")
)

type solver interface {
	Solve(ls []string) (string, error)
}

var solvers = []solver{}

func main() {
	flag.Parse()
	path := filepath.Join(*inputPathDir, *inputFile)
	ls, err := reader.ReadInput(path)
	if err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	ans, err := s6.Solve(ls)
	if err != nil {
		log.Fatalf("error solving problem: %v", err)
	}
	log.Printf("answer is: %v", ans)
}
