package parser

import (
	"cron/constants"
	"cron/utils"
	"cron/validator"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ParseCronExpression parses the passed arguments and prints the output
func ParseCronExpression(args []string) error {
	outputLog, err := parseCronExpression(args)
	if err != nil {
		return err
	}
	generateOutput(outputLog)
	return nil
}

func parseCronExpression(args []string) ([][2]string, error) {
	cronArgs := args[:len(args)-1] // Last value has the command, so skip parsing it
	var outputLog [][2]string

	for idx, arg := range cronArgs {
		// Re-check validity of expression before parsing it to avoid exceptions
		valid, validErr := validator.IsValid(arg, constants.CronConfig[idx])
		if !valid || validErr != nil {
			return nil, errors.New(fmt.Sprintf("invalid cron expression for: %s", arg))
		}

		parsed, err := parseCronArg(arg, constants.CronConfig[idx])
		if err != nil {
			return nil, err
		}

		outputLog = append(outputLog, [2]string{constants.CronConfig[idx].Name, strings.Join(parsed, " ")})
	}

	// Add the command value to output log
	outputLog = append(outputLog, [2]string{constants.CommandKey, args[len(args)-1]})
	return outputLog, nil
}

func generateOutput(outputLog [][2]string) {
	for _, output := range outputLog {
		utils.PrintOutput(output[0], output[1])
	}
}

func parseCronArg(arg string, conf *constants.CronVariable) ([]string, error) {
	matches, err := utils.ExtractValues(arg)
	if err != nil {
		return nil, err
	}
	if len(matches) == 1 {
		return generateValuesSimple(matches, conf), nil
	}
	return generateValuesWithOperator(matches, conf), nil
}

// generateValuesSimple processes the case where there is one operand
func generateValuesSimple(parsedEls []string, conf *constants.CronVariable) []string {
	var result []string
	if parsedEls[0] == "*" {
		result = generateValues(conf.MinAllowed, conf.MaxAllowed, 1)
	} else {
		result = append(result, parsedEls[0])
	}
	return result
}

// generateValuesSimple processes the case where there are operands and operator
func generateValuesWithOperator(parsedEls []string, conf *constants.CronVariable) []string {
	var result []string
	if parsedEls[0] == "*" {
		parsedEls[0] = "0"
	}
	num1 := utils.ParseToInt(parsedEls[0])
	num2 := utils.ParseToInt(parsedEls[2])
	switch parsedEls[1] {
	case ",":
		result = []string{parsedEls[0], parsedEls[2]}
	case "-":
		result = generateValues(num1, num2, 1)
	case "/":
		result = generateValues(num1, conf.MaxAllowed, num2)
	}
	return result
}

func generateValues(min, max, step int) []string {
	var res []string
	for i := min; i <= max; i += step {
		res = append(res, strconv.Itoa(i))
	}
	return res
}
