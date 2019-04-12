package functions

import (
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name     string
		position int8
		result   int64
	}{
		{"-1", -1, 1},
		{"0", 0, 1},
		{"1", 1, 1},
		{"2", 2, 1},
		{"3", 3, 2},
		{"4", 4, 3},
		{"5", 5, 5},
		{"6", 6, 8},
		{"7", 7, 13},
		{"50", 50, 12586269025},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.position); got != tt.result {
				t.Errorf("Fib() = %v, want %v", got, tt.result)
			}
		})
	}
}

func TestCondition(t *testing.T) {
	if Condition(true) {
		t.Errorf("Condition() = %v, want %v", true, false)
	}
}
