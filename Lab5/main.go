package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Company struct {
	name     string
	position string
	salary   float64
}

type Worker struct {
	name      string
	year      int
	month     int
	workPlace *Company
}

func NewCompany(name, position string, salary float64) *Company {
	return &Company{
		name:     name,
		position: position,
		salary:   salary,
	}
}

func NewWorker(name string, year, month int, workPlace *Company) *Worker {
	return &Worker{
		name:      name,
		year:      year,
		month:     month,
		workPlace: workPlace,
	}
}

func (c *Company) SetName(name string)         { c.name = name }
func (c *Company) GetName() string             { return c.name }
func (c *Company) SetPosition(position string) { c.position = position }
func (c *Company) GetPosition() string         { return c.position }
func (c *Company) SetSalary(salary float64)    { c.salary = salary }
func (c *Company) GetSalary() float64          { return c.salary }

func (w *Worker) SetName(name string)             { w.name = name }
func (w *Worker) GetName() string                 { return w.name }
func (w *Worker) SetYear(year int)                { w.year = year }
func (w *Worker) GetYear() int                    { return w.year }
func (w *Worker) SetMonth(month int)              { w.month = month }
func (w *Worker) GetMonth() int                   { return w.month }
func (w *Worker) SetWorkPlace(workPlace *Company) { w.workPlace = workPlace }
func (w *Worker) GetWorkPlace() *Company          { return w.workPlace }

func (w *Worker) GetWorkerPosition() string {
	if w.workPlace != nil {
		return w.workPlace.GetPosition()
	}
	return "Невідомо"
}

func (w *Worker) GetWorkExperience() int {
	now := time.Now()
	currentYear := now.Year()
	currentMonth := int(now.Month())

	months := (currentYear-w.year)*12 + (currentMonth - w.month)
	if months < 0 {
		return 0
	}
	return months
}

func (w *Worker) GetTotalMoney() float64 {
	experience := w.GetWorkExperience()
	if w.workPlace != nil {
		return float64(experience) * w.workPlace.GetSalary()
	}
	return 0
}

func ReadWorkersArray() []*Worker {
	n := readInt("Введіть кількість працівників (n): ")
	var workers []*Worker

	for i := 0; i < n; i++ {
		fmt.Printf("\n--- Введення даних для працівника #%d ---\n", i+1)

		wName := readString("Прізвище та ініціали: ")
		wYear := readInt("Рік початку роботи: ")
		wMonth := readInt("Місяць початку роботи (1-12): ")

		fmt.Println("Дані про компанію:")
		cName := readString("  Назва компанії: ")
		cPosition := readString("  Посада: ")
		cSalary := readFloat("  Зарплата: ")

		company := NewCompany(cName, cPosition, cSalary)
		worker := NewWorker(wName, wYear, wMonth, company)

		workers = append(workers, worker)
	}

	return workers
}

func PrintWorker(w *Worker) {
	fmt.Printf("Працівник: %s | Стаж: %d міс. | Посада: %s | Зарплата: %.2f | Всього зароблено: %.2f\n",
		w.GetName(),
		w.GetWorkExperience(),
		w.GetWorkerPosition(),
		w.GetWorkPlace().GetSalary(),
		w.GetTotalMoney(),
	)
}

func PrintWorkers(workers []*Worker) {
	fmt.Println("\n=== Список усіх працівників ===")
	for _, w := range workers {
		PrintWorker(w)
	}
}

func GetWorkersInfo(workers []*Worker) (minSalary float64, maxSalary float64) {
	if len(workers) == 0 {
		return 0, 0
	}

	minSalary = workers[0].GetWorkPlace().GetSalary()
	maxSalary = workers[0].GetWorkPlace().GetSalary()

	for _, w := range workers {
		salary := w.GetWorkPlace().GetSalary()
		if salary < minSalary {
			minSalary = salary
		}
		if salary > maxSalary {
			maxSalary = salary
		}
	}
	return minSalary, maxSalary
}

func readString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text != "" {
			return text
		}
		fmt.Println("Помилка: рядок не може бути порожнім. Спробуйте ще раз.")
	}
}

func readInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		val, err := strconv.Atoi(text)
		if err == nil {
			return val
		}
		fmt.Println("Помилка: введіть коректне ціле число.")
	}
}

func readFloat(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		val, err := strconv.ParseFloat(text, 64)
		if err == nil {
			return val
		}
		fmt.Println("Помилка: введіть коректне число (можна з крапкою).")
	}
}

func main() {
	workers := ReadWorkersArray()
	PrintWorkers(workers)
	minSal, maxSal := GetWorkersInfo(workers)
	fmt.Printf("\n=== Статистика по зарплатах ===\n")
	fmt.Printf("Найменша зарплата: %.2f\n", minSal)
	fmt.Printf("Найбільша зарплата: %.2f\n", maxSal)
}
