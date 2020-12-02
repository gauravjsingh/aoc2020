package input

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