package problem0014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one []string
}

type ans struct {
	one string
}

func Test_Problem0014(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{
				[]string{"flower", "flight", "flow"},
			},
			ans{"fl"},
		},
		question{
			para{
				[]string{"dog", "racecar", "car"},
			},
			ans{""},
		},
		question{
			para{
				[]string{"japan", "janam", "jamun"},
			},
			ans{"ja"},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, longestCommonPrefix(p.one))
	}
}
