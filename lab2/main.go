package main

import (
	"fmt"
	"lab2/calc"
)

func main() {
	fmt.Println("Мінімум:", calc.FindMin(1, 2, -3))
	fmt.Println("Середнє:", calc.CalculateAverage(2, 4, 6))

	res, _ := calc.SolveEquation(2, -8)
	fmt.Println("Корінь рівняння:", res)
}
