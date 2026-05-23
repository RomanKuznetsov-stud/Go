package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведення типів\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	// Завдання.
	// 1. Створіть 2 змінні різних типів. Виконайте арифметичні операції. Результат вивести в консоль
	var intVal int = 20
	var floatVal float64 = 2.7

	fmt.Println("\n--- Арифметичні операції з різними типами ---")
	fmt.Printf("intVal   = %d   (тип %T)\n", intVal, intVal)
	fmt.Printf("floatVal = %.1f (тип %T)\n", floatVal, floatVal)

	fmt.Println("Завдання-----------------")
	sumFloat := float64(intVal) + floatVal
	fmt.Printf("Сума (як float64): %.1f + %.1f = %.1f\n", float64(intVal), floatVal, sumFloat)

	diffInt := intVal - int(floatVal)
	fmt.Printf("Різниця (як int): %d - %d = %d\n", intVal, int(floatVal), diffInt)

	multFloat := float64(intVal) * floatVal
	fmt.Printf("Добуток (як float64): %.1f * %.1f = %.2f\n", float64(intVal), floatVal, multFloat)
}
