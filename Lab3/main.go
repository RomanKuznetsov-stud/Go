package main

import (
	"fmt"
	"math"
)

const (
	a uint32 = 22695477
	c uint32 = 1
	K int    = 20000
	n int    = 150
)

// 1. Функція генерування цілочислової послідовності псевдовипадкових значень
func generateIntSequence(seed uint32) []int {
	seq := make([]int, K)
	x := seed
	for i := 0; i < K; i++ {
		x = a*x + c
		seq[i] = int(x % uint32(n))
	}
	return seq
}

// 2. Функція генерування послідовності дійсних значень
func generateFloatSequence(seed uint32) []float64 {
	seq := make([]float64, K)
	x := seed
	mFloat := float64(4294967296) // 2^32
	for i := 0; i < K; i++ {
		x = a*x + c
		seq[i] = (float64(x) / mFloat) * float64(n)
	}
	return seq
}

func main() {
	seed := uint32(12345)

	fmt.Println("--- 1. Аналіз цілочислової послідовності ---")
	intSeq := generateIntSequence(seed)

	frequencies := make([]int, n)
	for _, val := range intSeq {
		frequencies[val]++
	}

	probabilities := make([]float64, n)
	for i := 0; i < n; i++ {
		probabilities[i] = float64(frequencies[i]) / float64(K)
	}

	fmt.Println("Частота та ймовірність для значень (0-4):")
	for i := 0; i < 5; i++ {
		fmt.Printf("  Значення %d: Частота = %d, Ймовірність = %.4f\n", i, frequencies[i], probabilities[i])
	}

	var expectedValue float64
	for i, p := range probabilities {
		expectedValue += float64(i) * p
	}
	fmt.Printf("\nМатематичне сподівання: %.4f\n", expectedValue)

	var variance float64
	for i, p := range probabilities {
		diff := float64(i) - expectedValue
		variance += p * (diff * diff)
	}
	fmt.Printf("Дисперсія: %.4f\n", variance)

	stdDeviation := math.Sqrt(variance)
	fmt.Printf("Середньоквадратичне відхилення: %.4f\n\n", stdDeviation)

	fmt.Println("--- 2. Дійсна послідовність псевдовипадкових значень ---")
	floatSeq := generateFloatSequence(seed)

	fmt.Println("Перші 10 згенерованих дійсних значень:")
	for i := 0; i < 10; i++ {
		fmt.Printf("  %.4f\n", floatSeq[i])
	}
}
