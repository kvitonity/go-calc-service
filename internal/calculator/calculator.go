package calculator

import (
	"fmt"
)

func Calculate(op1 float64, operator string, op2 float64) (float64, error) {
	switch operator {
	case "+":
		return op1 + op2, nil
	case "-":
		return op1 - op2, nil
	case "*":
		return op1 * op2, nil
	case "/":
		if op2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return op1 / op2, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}
}
