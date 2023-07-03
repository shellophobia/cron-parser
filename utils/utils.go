package utils

import (
	"cron/constants"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// ParseToInt checks the if the string is a valid integer or throws an exception.
// An exception is thrown instead of returning an error because the regex should ensure invalid values don't get parsed
func ParseToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Invalid value encountered as number, %s", s)
	}
	return val
}

// ExtractValues extracts the combination of var1, operator, var2 or just var1
// from the cron argument. If the value is invalid it returns an error
func ExtractValues(arg string) ([]string, error) {
	// Find all matching groups in the cron expression
	matchResult := constants.CronExprRegex.FindAllStringSubmatch(arg, 1)
	if len(matchResult) == 0 {
		return nil, errors.New(fmt.Sprintf("No match found for cron arg: %s", arg))
	}

	matches := matchResult[0]
	// Remove the empty matches
	matches = filter(matches, func(s string) bool {
		return strings.TrimSpace(s) != ""
	})
	if len(matches) < 1 {
		return nil, errors.New(fmt.Sprintf("No match found for cron arg: %s", arg))
	}

	// In the matched group result, 1st value is the complete match value that has the whole argument
	// So removing that from the concerned list
	matches = matches[1:]
	if len(matches) != 1 && len(matches) != 3 {
		return nil, errors.New(fmt.Sprintf("Parsing failure for expression: %s", arg))
	}

	return matches, nil
}

// PrintOutput is the utility method to print the output with adequate padding
func PrintOutput(keyName, output string) {
	fmt.Printf("%-*s %s\n", constants.RightSpacePadding, keyName, output)
}

func filter(args []string, callback func(s string) bool) []string {
	var filteredList []string
	for _, arg := range args {
		if callback(arg) {
			filteredList = append(filteredList, arg)
		}
	}
	return filteredList
}
