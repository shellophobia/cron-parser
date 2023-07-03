package main

import (
	"cron/parser"
	"cron/validator"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("Please enter a valid cron expression with cron argument")
	}
	cronArgs := strings.Split(args[1], " ")

	err := validator.ValidateExpression(cronArgs[:len(cronArgs)-1])
	if err != nil {
		log.Fatalf("Please enter a valid cron expression. Failed with error: %#v", err.Error())
	}

	err = parser.ParseCronExpression(cronArgs)
	if err != nil {
		log.Fatalf("Please enter a valid cron expression. Failed with error: %#v", err.Error())
	}
}
