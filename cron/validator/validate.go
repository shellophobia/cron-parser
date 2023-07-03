package validator

import (
	"cron/constants"
	"cron/utils"
	"errors"
	"fmt"
)

// ValidateExpression accepts a list of cron arguments and checks if cron expressions
// are in a valid format. If not it returns an error
func ValidateExpression(cronArgs []string) error {
	if len(cronArgs) != len(constants.CronConfig) {
		return errors.New("invalid number of cron args")
	}
	for idx, cronVar := range cronArgs {
		valid, err := IsValid(cronVar, constants.CronConfig[idx])
		if !valid || err != nil {
			return errors.New(fmt.Sprintf("invalid cron expression for: %s", cronVar))
		}
	}
	return nil
}

// IsValid accepts a cron argument and checks if it's valid against the cron config value
// E.g. if it's the 1st argument, is it a valid minute argument
func IsValid(arg string, conf *constants.CronVariable) (bool, error) {
	if !constants.CronExprRegex.MatchString(arg) {
		return false, nil
	}
	// Extract the values from cron expression i.e. var1, operator, var2 or just var1
	values, err := utils.ExtractValues(arg)
	if err != nil {
		return false, err
	}
	if len(values) == 1 {
		if values[0] == "*" {
			return true, nil
		}
		val := utils.ParseToInt(values[0])
		return isWithinRange(val, conf), nil
	}
	isVal1Allowed := true
	if values[0] == "*" && values[1] != "/" || values[2] == "*" {
		return false, nil
	}
	if values[0] != "*" {
		isVal1Allowed = isWithinRange(utils.ParseToInt(values[0]), conf)
	}
	isVal2Allowed := isWithinRange(utils.ParseToInt(values[2]), conf)
	return isVal1Allowed && isVal2Allowed, nil
}

func isWithinRange(val int, conf *constants.CronVariable) bool {
	return conf.MinAllowed <= val && conf.MaxAllowed >= val
}
