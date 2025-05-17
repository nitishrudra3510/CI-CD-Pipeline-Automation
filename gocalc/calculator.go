package main

import (
	"errors"
	"strconv"
	"strings"
)

func calculate(expression string) (float64, error) {

	// Split the expression by operators
	tokens := strings.Fields(expression)

	if len(tokens) < 3 || len(tokens)%2 == 0 {
		return 0, nil
	}

	// Convert the first token to a float64 number
	result, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, errors.New("invalid expression")
	}

	// Perform arithmetic operations on subsequent tokens
	for i := 1; i < len(tokens); i += 2 {
		operator := tokens[i]
		number, err := strconv.ParseFloat(tokens[i+2], 64)
		if err != nil {
			return 0, errors.New("invalid expression")
		}

		switch operator {
		case "+":
			result += number
		case "-":
			result -= number
		case "*":
			result *= number
		case "/":
			if number == 0 {
				return 0, errors.New("division by zero")
			}
			result /= number
		default:
			return 0, errors.New("invalid operator")
		}
	}

	return result, nil
}
