package validator

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	name string
	args []string
	err  error
}

var tests = []testCase{
	{
		"throws error for invalid cron args",
		[]string{"*/15", "0", "1,15"},
		errors.New("invalid number of cron args"),
	},
	{
		"throws error for invalid cron expression 1-*",
		[]string{"*/15", "0", "1,15", "*", "1-*"},
		errors.New("invalid cron expression for: 1-*"),
	},
	{
		"throws error for invalid cron expression */*",
		[]string{"*/*", "0", "1,15", "*", "1-5"},
		errors.New("invalid cron expression for: */*"),
	},
	{
		"throws error for invalid cron expression 5/*",
		[]string{"5/*", "0", "1,15", "*", "1-5"},
		errors.New("invalid cron expression for: 5/*"),
	},
	{
		"throws error for invalid cron expression **",
		[]string{"*/5", "**", "1,15", "*", "1-5"},
		errors.New("invalid cron expression for: **"),
	},
	{
		"throws error for invalid cron expression 1,*",
		[]string{"*/5", "0", "1,*", "*", "1-5"},
		errors.New("invalid cron expression for: 1,*"),
	},
	{
		"throws error for invalid cron expression *,1",
		[]string{"*/5", "0", "*,1", "*", "1-5"},
		errors.New("invalid cron expression for: *,1"),
	},
	{
		"throws error for values outside allowed range in expression 60/5",
		[]string{"60/5", "0", "1,15", "*", "1-5"},
		errors.New("invalid cron expression for: 60/5"),
	},
	{
		"throws error for values outside allowed range in expression 20",
		[]string{"*/5", "0", "1,15", "*", "20"},
		errors.New("invalid cron expression for: 20"),
	},
	{
		"throws error for values outside allowed range in expression 1xaf",
		[]string{"*/5", "0", "1,15", "1xaf", "1-5"},
		errors.New("invalid cron expression for: 1xaf"),
	},
	{
		"doesn't return error for valid expression",
		[]string{"*/5", "0", "1,15", "*", "1-5"},
		nil,
	},
}

func TestValidateExpression(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateExpression(test.args)
			assert.Equal(t, test.err, err)
		})
	}
}
