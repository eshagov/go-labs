package hw_3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var text = `
the 
on on on on 
coo coo cood
you you you you you you you you
wau wau wau wau wau wau wau wau wau wau wau wau wau wau 
?`

type TestCaseType struct {
	s        string
	expected []string
}

var testCases = []TestCaseType {
	{
		s:        ",;-?",
		expected: []string{},
	},
	{
		s:        "",
		expected: []string{},
	},
	{
		s:        "a a a a b b b c c c c c y y y y y y y d d",
		expected: []string{"y", "c", "a", "b", "d"},
	},
	{
		s:        text,
		expected: []string{"wau", "you", "on", "coo", "the", "cood"},
	},
}

func Test_Top10(t *testing.T) {
	t.Parallel()
	for _, testCase := range testCases {
		actualValue := TopTenWords(testCase.s)
		//sort.Strings(data)
		require.Equal(t, testCase.expected, actualValue)
	}
}

