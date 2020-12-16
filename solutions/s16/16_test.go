package s16

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		desc  string
		input string
		want  constraint
	}{
		{
			desc:  "2 ranges",
			input: "class: 1-3 or 5-7",
			want:  constraint{name: "class", rs: []intRange{{1, 3}, {5, 7}}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			c, err := parseConstraint(tc.input)
			if err != nil {
				t.Errorf("parseConstraint(%s) gave unexpected error: %v", tc.input, err)
			}
			if diff := cmp.Diff(c, tc.want, cmp.AllowUnexported(constraint{}, intRange{})); diff != "" {
				t.Errorf("parseConstraint(%s) = %v, want %v, diff\n:%s", tc.input, c, tc.want, diff)
			}
		})
	}
}

func TestValid(t *testing.T) {
	tests := []struct {
		desc string
		n    int
		cs   constraints
		want bool
	}{
		{
			desc: "one constraint, valid",
			n:    5,
			cs:   constraints{{rs: []intRange{{3, 7}}}},
			want: true,
		},
		{
			desc: "one constraint, invalid",
			n:    15,
			cs:   constraints{{rs: []intRange{{3, 7}}}},
		},
		{
			desc: "two constraint, valid",
			n:    15,
			cs:   constraints{{rs: []intRange{{3, 7}}}, {rs: []intRange{{13, 17}}}},
			want: true,
		},
		{
			desc: "two constraint, invalid",
			n:    1,
			cs:   constraints{{rs: []intRange{{3, 7}}}, {rs: []intRange{{13, 17}}}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.cs.valid(tc.n)
			if got != tc.want {
				t.Errorf("valid(%d) = %t, want %t", tc.n, got, tc.want)
			}
		})
	}

}
