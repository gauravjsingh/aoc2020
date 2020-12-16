package reader

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInput(path string) ([]string, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(string(bs))
	return strings.Split(trimmed, "\n"), nil
}

func ParseInput(ls []string) ([]int, error) {
	var out []int
	for _, l := range ls {
		i, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		out = append(out, i)
	}

	return out, nil
}

// GroupInput groups input into groups split by blank lines.
func GroupInput(ls []string) [][]string {
	var out [][]string
	var grp []string
	for _, l := range ls {
		if len(l) == 0 {
			out = append(out, grp)
			grp = nil
			continue
		}
		grp = append(grp, l)
	}
	out = append(out, grp)
	return out
}
