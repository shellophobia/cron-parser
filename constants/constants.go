package constants

import "regexp"

// CronVariable has the name of expression and the allowed value range
type CronVariable struct {
	Name       string
	MinAllowed int
	MaxAllowed int
}

// CronConfig stores the list of supported cron variables with allowed ranges
var CronConfig = []*CronVariable{
	{"minute", 0, 59},
	{"hour", 0, 23},
	{"day of month", 1, 31},
	{"month", 1, 12},
	{"day of week", 1, 7},
}

// RightSpacePadding for printing the output log
var RightSpacePadding = 13

// CommandKey key name for command in output log
var CommandKey = "command"

// CronExprRegex regex for cron expression for parsing
var CronExprRegex = regexp.MustCompile("^(\\*|[0-9]+)([/,-]?)([0-9]*)$")
