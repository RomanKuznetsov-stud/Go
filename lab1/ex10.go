package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	// Завдання.
	// 1. Вивести українську літеру 'Ї'
	// 2. Пояснити призначення типу "rune"
	var ukrLetter rune = 'Ї'
	fmt.Printf("Code '%c' - %d\n", ukrLetter, ukrLetter)
	// Використовуємо тип rune, оскільки 'Ї' виходить за межі таблиці ASCII
}
