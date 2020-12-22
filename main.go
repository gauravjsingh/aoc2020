package main

import (
	"aoc2020/reader"
	"aoc2020/solutions"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
)

var (
	inputPathDir = flag.String("input_dir", "input", "path to the input data directory")
	problem      = flag.Int("problem", -1, "Which problem to solve?")
	cpuProfile   = flag.String("cpuprofile", "", "write cpu profile to the given file path")
)

const fileFormat = "%d.txt"

func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if *problem == -1 {
		*problem = len(solutions.Solvers)
	}
	path := filepath.Join(*inputPathDir, fmt.Sprintf(fileFormat, *problem))
	ls, err := reader.ReadInput(path)
	if err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	s, ok := solutions.Solvers[*problem]
	if !ok {
		log.Fatalf("solver %d not found", *problem)
	}
	ans, err := s(ls)
	if err != nil {
		log.Fatalf("error solving problem: %v", err)
	}
	log.Printf("answer is: %v", ans)
}
