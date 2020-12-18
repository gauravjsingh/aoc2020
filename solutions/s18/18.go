package s18

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type expression interface {
	value() int
}

type intValue int

func (v intValue) value() int {
	return int(v)
}

func newIntValue(s string, existing expression) (intValue, error) {
	if existing != nil {
		return 0, fmt.Errorf("newIntValue got non-nil existing: %v", existing)
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("error parsing intValue from %q: %v", s, err)
	}
	return intValue(n), nil
}

type binaryOperator int

const (
	add binaryOperator = iota
	mul
)

func (b binaryOperator) apply(l, r int) int {
	switch b {
	case add:
		return l + r
	case mul:
		return l * r
	}
	log.Fatalf("unexpected operator: %v", b)
	return 0
}

type binaryExpression struct {
	left, right expression
	op          binaryOperator
}

func (b binaryExpression) value() int {
	return b.op.apply(b.left.value(), b.right.value())
}

// Assume that only 1 digit numbers are present
func parseExpression(l string, existing expression) (expression, error) {
	if len(l) == 0 {
		return nil, errors.New("error parsing empty string as expression")
	}
	e := existing
	var opStack []binaryOperator
	for i := 0; i < len(l); i++ {
		if l[i] == ' ' {
			continue
		}
		if l[i] == '(' {
			var right int
			parens := 1
			for right = i + 1; right < len(l); right++ {
				if l[right] == '(' {
					parens++
				}
				if l[right] == ')' {
					parens--
				}
				if parens == 0 {
					break
				}
			}
			if parens != 0 {
				return nil, fmt.Errorf("missing end paren at %d after %q", i, l[i:])
			}
			subE, err := parseExpression(l[i+1:right], nil)
			if err != nil {
				return nil, fmt.Errorf("error at %d parsing subexpression %q: %v", i, l[i+1:right], err)
			}
			i = right + 1
			if e == nil {
				e = subE
				continue
			}
			e = binaryExpression{left: e, right: subE, op: opStack[len(opStack)-1]}
			opStack = opStack[:len(opStack)-1]
			continue
		}
		if l[i] == '+' {
			opStack = append(opStack, add)
			continue
		}
		if l[i] == '*' {
			opStack = append(opStack, mul)
			continue
		}
		n, err := strconv.Atoi(l[i : i+1])
		if err != nil {
			return nil, fmt.Errorf("error parsing int at %d: %v", i, err)
		}
		if e == nil {
			e = intValue(n)
			continue
		}
		if len(opStack) == 0 {
			// is len ever != 1?
			return nil, fmt.Errorf("opStack is empty at %d, expression so far: %v", i, e)
		}
		e = binaryExpression{left: e, right: intValue(n), op: opStack[len(opStack)-1]}
		opStack = opStack[:len(opStack)-1]
	}
	return e, nil
}

func Solve(ls []string) (int, error) {
	sum := 0
	for _, l := range ls {
		e, err := parseExpression(l, nil)
		if err != nil {
			return 0, fmt.Errorf("error parsing %q as expression: %v", l, err)
		}
		log.Printf("value of %q: %d", l, e.value())
		sum += e.value()
	}
	return sum, nil
}
