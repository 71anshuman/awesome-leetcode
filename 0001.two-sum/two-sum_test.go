package problem0001

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: []int{3, 2, 4},
				two: 6,
			},
			a: ans{
				one: []int{1, 2},
			},
		},
		{
			p: para{
				one: []int{3, 2, 4},
				two: 8,
			},
			a: ans{
				one: nil,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		res := twoSum(p.one, p.two)
		sort.Ints(res)
		sort.Ints(a.one)
		ast.Equal(a.one, res, "%v", p)
	}
}
