package main

import (
	"fmt"
	"strings"
)

// Calculator provides basic arithmetic operations
type Calculator struct{}

// Add adds two numbers
func (c *Calculator) Add(a, b int) int {
	return a + b
}

// Subtract subtracts two numbers
func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two numbers
func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

// Divide divides two numbers
func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// StringProcessor provides string manipulation utilities
type StringProcessor struct{}

// ToUpper converts string to uppercase
func (s *StringProcessor) ToUpper(text string) string {
	return strings.ToUpper(text)
}

// Reverse reverses a string
func (s *StringProcessor) Reverse(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ProcessList processes a list of strings
func (s *StringProcessor) ProcessList(items []string) []string {
	var result []string
	for _, item := range items {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			result = append(result, strings.ToUpper(trimmed))
		}
	}
	return result
}

func main() {
	calc := Calculator{}
	processor := StringProcessor{}

	fmt.Println("Go Test Application")
	fmt.Printf("Add: %d\n", calc.Add(5, 3))
	fmt.Printf("Subtract: %d\n", calc.Subtract(10, 4))
	fmt.Printf("Multiply: %d\n", calc.Multiply(3, 4))
	
	result, err := calc.Divide(20, 4)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Divide: %d\n", result)
	}

	fmt.Printf("Uppercase: %s\n", processor.ToUpper("hello world"))
	fmt.Printf("Reverse: %s\n", processor.Reverse("golang"))
}

