package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to GoCalc")
	fmt.Println("--------------------------------------")

	for {
		fmt.Print("Enter an expression (or 'exit' to quit): ")
		expression, _ := reader.ReadString('\n')
		expression = strings.TrimSpace(expression)

		if expression == "exit" {
			break
		}

		result, err := calculate(expression)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			fmt.Printf("Results: %.2f\n", result)
		}
	}
}
