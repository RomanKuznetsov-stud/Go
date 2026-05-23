package main

import "fmt"

func main() {
	// Ініціалізація змінних
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4 // Такий варіант ініціалізації також можливий

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	// Короткий запис оголошення змінної
	// тільки для нових змінних
	intVar := 10

	//fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	// Завдання.
	// 1. Вивести типи всіх змінних
	fmt.Println("1---------------------")
	fmt.Printf("Type of userinit8: %T\n", userinit8)
	fmt.Printf("Type of userinit16: %T\n", userinit16)
	fmt.Printf("Type of userinit64: %T\n", userinit64)
	fmt.Printf("Type of userautoinit: %T\n", userautoinit)
	fmt.Printf("Type of intVar: %T\n", intVar)
	fmt.Println("2---------------------")
	// 2. Присвоїти змінній intVar змінні userinit16 і userautoinit. Результат вивести в консоль.
	intVar = int(userinit16)
	fmt.Println("intVar after userinit16 =", intVar)
	intVar = userautoinit
	fmt.Println("intVar after userautoinit =", intVar)

	intVar = int(userinit16) + userautoinit
	fmt.Println("Result (userinit16 + userautoinit) = ", intVar)
}
