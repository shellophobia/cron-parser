package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	name     string
	args     []string
	expected [][2]string
	hasError bool
}

var tests = []testCase{
	{
		"generates for valid arguments",
		[]string{"*/15", "0", "1,15", "*", "1-5", "/usr/bin/find"},
		[][2]string{
			{"minute", "0 15 30 45"},
			{"hour", "0"},
			{"day of month", "1 15"},
			{"month", "1 2 3 4 5 6 7 8 9 10 11 12"},
			{"day of week", "1 2 3 4 5"},
			{"command", "/usr/bin/find"},
		},
		false,
	},
	{
		"throws error for invalid arguments",
		[]string{"*/*", "0", "1,15", "*", "1-5", "/usr/bin/find"},
		nil,
		true,
	},
}

func TestParseCronExpression(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			values, err := parseCronExpression(test.args)
			if test.hasError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, test.expected, values)
		})
	}
}
