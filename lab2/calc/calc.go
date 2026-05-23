package calc

import "fmt"

func FindMin(a, b, c float64) float64 {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func CalculateAverage(a, b, c float64) float64 {
	return (a + b + c) / 3.0
}

func SolveEquation(a, b float64) (float64, error) {
	if a == 0 {
		return 0, fmt.Errorf("рівняння не має єдиного розв'язку (a = 0)")
	}
	return -b / a, nil
}
