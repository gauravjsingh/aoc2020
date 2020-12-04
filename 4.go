package main

import (
	"aoc2020/reader"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

var inputPath = flag.String("input_path", "input/4.txt", "path to the input data")

var reqFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

type passport struct {
	fields map[string]string
}

var hgtRE = regexp.MustCompile(`^(\d+)in|(\d+)cm$`)
var hclRE = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var eclRE = regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
var pidRE = regexp.MustCompile(`^[\d]{9}$`)

func (p passport) isValid() bool {
	for _, f := range reqFields {
		if _, ok := p.fields[f]; !ok {
			return false
		}
	}

	if byr, err := strconv.Atoi(p.fields["byr"]); err != nil || byr < 1920 || byr > 2002 {
		return false
	}
	if iyr, err := strconv.Atoi(p.fields["iyr"]); err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}
	if eyr, err := strconv.Atoi(p.fields["eyr"]); err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}

	matches := hgtRE.FindStringSubmatch(p.fields["hgt"])
	log.Print(matches)
	if len(matches) < 3 {
		return false
	}
	if matches[1] != "" {
		if in, err := strconv.Atoi(matches[1]); err != nil || in < 59 || in > 76 {
			return false
		}
	} else if matches[2] != "" {
		if cm, err := strconv.Atoi(matches[2]); err != nil || cm < 150 || cm > 193 {
			return false
		}
	} else {
		return false
	}

	if !hclRE.MatchString(p.fields["hcl"]) {
		return false
	}
	if !eclRE.MatchString(p.fields["ecl"]) {
		return false
	}
	if !pidRE.MatchString(p.fields["pid"]) {
		return false
	}

	return true
}

func newPassport() passport {
	return passport{fields: make(map[string]string)}
}

var re = regexp.MustCompile(`([\w#]+):([\w#]+)`)

func parseInput(ls []string) ([]passport, error) {
	var ps []passport
	p := newPassport()
	for _, l := range ls {
		if len(l) == 0 {
			ps = append(ps, p)
			p = newPassport()
			continue
		}
		matches := re.FindAllStringSubmatch(l, -1)
		for _, match := range matches {
			if len(match) != 3 {
				return nil, fmt.Errorf("unexpected return from re: %v", match)
			}
			p.fields[match[1]] = match[2]
		}
	}
	ps = append(ps, p)
	return ps, nil
}

func validPasses(ps []passport) int {
	valid := 0
	for _, p := range ps {
		if p.isValid() {
			valid += 1
		}
	}
	return valid
}

func main() {
	flag.Parse()
	ls, err := reader.ReadInput(*inputPath)
	if err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	ps, err := parseInput(ls)
	if err != nil {
		log.Fatalf("error parsing input as passports: %v", err)
	}
	ans := validPasses(ps)
	log.Printf("answer is: %d", ans)
}
