package main

import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	// Завдання.
	// 1. Створіть змінні різних типів, використовуючи короткий запис та ініціалізацію за замовчуванням. Результат вивести в консоль
	var defaultInt int
	var defaultBool bool
	var defaultString string

	shortFloat := 3.14
	shortString := "Go"
	shortBool := true

	fmt.Println("--- Ініціалізація за замовчуванням ---")
	fmt.Printf("defaultInt    = %d  (тип %T)\n", defaultInt, defaultInt)
	fmt.Printf("defaultBool   = %t  (тип %T)\n", defaultBool, defaultBool)
	fmt.Printf("defaultString = %q  (тип %T)\n\n", defaultString, defaultString)

	fmt.Println("--- Короткий запис ---")
	fmt.Printf("shortFloat  = %f (тип %T)\n", shortFloat, shortFloat)
	fmt.Printf("shortString = %q (тип %T)\n", shortString, shortString)
	fmt.Printf("shortBool   = %t (тип %T)\n", shortBool, shortBool)

}
