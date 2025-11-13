package main

import (
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	calc := Calculator{}
	
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 5, 3, 8},
		{"zero", 0, 0, 0},
		{"negative numbers", -5, 5, 0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Subtract(t *testing.T) {
	calc := Calculator{}
	
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 10, 4, 6},
		{"zero result", 5, 5, 0},
		{"negative result", 3, 5, -2},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Multiply(t *testing.T) {
	calc := Calculator{}
	
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 3, 4, 12},
		{"zero", 0, 5, 0},
		{"negative numbers", -2, 3, -6},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	calc := Calculator{}
	
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
		hasError bool
	}{
		{"positive numbers", 20, 4, 5, false},
		{"divide by zero", 10, 0, 0, true},
		{"negative numbers", -10, 2, -5, false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Divide(tt.a, tt.b)
			if tt.hasError {
				if err == nil {
					t.Errorf("Divide(%d, %d) expected error, got nil", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.expected {
					t.Errorf("Divide(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

func TestStringProcessor_ToUpper(t *testing.T) {
	processor := StringProcessor{}
	
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercase", "hello", "HELLO"},
		{"empty string", "", ""},
		{"already uppercase", "WORLD", "WORLD"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processor.ToUpper(tt.input)
			if result != tt.expected {
				t.Errorf("ToUpper(%s) = %s, expected %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStringProcessor_Reverse(t *testing.T) {
	processor := StringProcessor{}
	
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"normal string", "golang", "gnalog"},
		{"empty string", "", ""},
		{"single character", "a", "a"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processor.Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%s) = %s, expected %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStringProcessor_ProcessList(t *testing.T) {
	processor := StringProcessor{}
	
	input := []string{"hello", "world", "", "  test  ", "go"}
	expected := []string{"HELLO", "WORLD", "TEST", "GO"}
	
	result := processor.ProcessList(input)
	
	if len(result) != len(expected) {
		t.Errorf("ProcessList length = %d, expected %d", len(result), len(expected))
	}
	
	for i, val := range expected {
		if i < len(result) && result[i] != val {
			t.Errorf("ProcessList[%d] = %s, expected %s", i, result[i], val)
		}
	}
}

