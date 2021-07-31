package problem0020

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one string
}

type ans struct {
	one bool
}

func Test_Problem0020(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{"()[]{}"},
			ans{true},
		},
		{
			para{"(]"},
			ans{false},
		},
		{
			para{"({[]})"},
			ans{true},
		},
		{
			para{"(){[({[]})]}"},
			ans{true},
		},
		{
			para{"((([[[{{{"},
			ans{false},
		},
		{
			para{"(())]]"},
			ans{false},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, isValid(p.one), "%v", p)
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 1; i < b.N; i++ {
		isValid("{{{{{[[[[[((((()))))]]]]]}}}}}")
	}
}
