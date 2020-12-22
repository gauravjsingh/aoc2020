package s21

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type food struct {
	ingredients map[string]struct{}
	allergens   map[string]struct{}
}

var foodRE = regexp.MustCompile(`([\w\s]+) \(contains ([\w,\s]+)\)`)

func parseFood(s string) (food, error) {
	matches := foodRE.FindStringSubmatch(s)
	if len(matches) != 3 {
		return food{}, fmt.Errorf("error matching food from %q: %v", s, matches)
	}
	f := food{
		ingredients: make(map[string]struct{}),
		allergens:   make(map[string]struct{}),
	}
	for _, in := range strings.Split(strings.TrimSpace(matches[1]), " ") {
		f.ingredients[in] = struct{}{}
	}
	for _, a := range strings.Split(strings.TrimSpace(matches[2]), ", ") {
		f.allergens[a] = struct{}{}
	}
	return f, nil
}

func parseInput(ls []string) ([]food, error) {
	var fs []food
	for _, l := range ls {
		f, err := parseFood(l)
		if err != nil {
			return nil, fmt.Errorf("error parsing food: %v", err)
		}
		fs = append(fs, f)
	}
	return fs, nil
}

func SolveA(ls []string) (int, error) {
	fs, err := parseInput(ls)
	if err != nil {
		return 0, err
	}
	aToI := make(map[string]map[string]struct{})
	for _, f := range fs {
		for a := range f.allergens {
			if _, ok := aToI[a]; !ok {
				aToI[a] = make(map[string]struct{})
				for in := range f.ingredients {
					aToI[a][in] = struct{}{}
				}
				continue
			}
			for in := range aToI[a] {
				if _, ok := f.ingredients[in]; !ok {
					delete(aToI[a], in)
				}
			}
		}
	}
	possible := make(map[string]struct{})
	for _, ins := range aToI {
		for in := range ins {
			possible[in] = struct{}{}
		}
	}
	log.Printf("aToI: %v", aToI)
	impossibleCount := 0
	for _, f := range fs {
		for in := range f.ingredients {
			if _, ok := possible[in]; !ok {
				impossibleCount += 1
			}
		}
	}
	return impossibleCount, nil
}
