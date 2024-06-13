package main

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	calculator := Calculator{}
	arr := []float64{4, 4, 4}
	sum := calculator.Addition(4, 4, 4)

	if sum != 12 {
		t.Fatalf(fmt.Sprintf("Addition returned wrong result to Addition %v", arr))
	}
}

func TestSubtraction(t *testing.T) {
	calculator := Calculator{}
	arr := []float64{4, 5}
	sub := calculator.Subtraction(4, 5)

	if sub != -1 {
		t.Fatalf(fmt.Sprintf("Subtraction returned wrong result to Subtraction %v", arr))
	}
}

func TestMultiplication(t *testing.T) {
	calculator := Calculator{}
	arr := []float64{4, 5, 2}
	mult := calculator.Multiplication(4, 5, 2)

	if mult != 40 {
		t.Fatalf(fmt.Sprintf("Multiplication returned wrong result to Multiplication %v", arr))
	}
}
func TestDivision(t *testing.T) {
	calculator := Calculator{}
	arr := []float64{4, 2}
	div, _ := calculator.Division(4, 2)

	if div != 2 {
		t.Fatalf(fmt.Sprintf("Division returned wrong result to Division %v", arr))
	}
}
