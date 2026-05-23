package calc

import "testing"

func TestFindMin(t *testing.T) {
	expected := -3.0
	result := FindMin(1, 2, -3)
	if result != expected {
		t.Errorf("Тест не пройдено! Очікувалось %f, отримано %f", expected, result)
	}
}

func TestCalculateAverage(t *testing.T) {
	expected := 4.0
	result := CalculateAverage(2, 4, 6)
	if result != expected {
		t.Errorf("Тест не пройдено! Очікувалось %f, отримано %f", expected, result)
	}
}

func TestSolveEquation(t *testing.T) {

	expected := 4.0
	result, err := SolveEquation(2, -8)

	if err != nil {
		t.Errorf("Неочікувана помилка: %v", err)
	}
	if result != expected {
		t.Errorf("Тест не пройдено! Очікувалось %f, отримано %f", expected, result)
	}

	_, errZero := SolveEquation(0, 5)
	if errZero == nil {
		t.Errorf("Очікувалась помилка при a=0, але її не було")
	}
}

func BenchmarkFindMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMin(1, 2, -3)
	}
}
