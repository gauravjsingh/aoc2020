package s18

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
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
	s = strings.TrimSpace(s)
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

func breakParens(s string) []string {
	i := strings.Index(s, "(")
	if i == -1 {
		return []string{s}
	}

	parens := 1
	for j := i + 1; j < len(s); j++ {
		if s[j] == '(' {
			parens++
		}
		if s[j] == ')' {
			parens--
		}
		if parens == 0 {
			var out []string
			out = append(out, s[:i])
			out = append(out, s[i+1:j])
			out = append(out, s[j+1:])
			return out
		}
	}
	return nil
}

func parseOperators(s string) (expression, error) {
	// add binds more tightly, so we apply it first.
	index := strings.Index(s, "*")
	op := mul
	if index == -1 {
		index = strings.Index(s, "+")
		op = add
	}
	//log.Printf("operator %d at index %d of %q", op, index, s)
	if index == -1 {
		return newIntValue(s, nil)
	}
	l, err := parseOperators(s[:index])
	if err != nil {
		return nil, fmt.Errorf("error parsing left subexpression at %d: %v", index, err)
	}
	r, err := parseOperators(s[index+1:])
	if err != nil {
		return nil, fmt.Errorf("error parsing right subexpression at %d: %v", index, err)
	}
	return binaryExpression{left: l, right: r, op: op}, nil
}

func simplifyExpr(s string) (int, error) {
	if ps := breakParens(s); len(ps) > 1 {
		nested, err := simplifyExpr(ps[1])
		if err != nil {
			return 0, fmt.Errorf("error evaluating nested %q in %q: %v", ps[1], s, err)
		}
		simplified := fmt.Sprintf("%s %d %s", ps[0], nested, ps[2])
		return simplifyExpr(simplified)
	}
	e, err := parseOperators(s)
	if err != nil {
		return 0, fmt.Errorf("error evaluating simple expression %q: %v", s, err)
	}
	return e.value(), nil
}

func SolveA(ls []string) (int, error) {
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

func SolveB(ls []string) (int, error) {
	sum := 0
	for _, l := range ls {
		v, err := simplifyExpr(l)
		if err != nil {
			return 0, fmt.Errorf("error simplifying %q: %v", l, err)
		}
		log.Printf("value of %q: %d", l, v)
		sum += v
	}
	return sum, nil
}
