package main

import (
	"fmt"
	"os"
)

func main() {
	var bank *Bank
	var choice int

	for {
		fmt.Println("\n=== КОНСОЛЬНЕ МЕНЮ БАНКУ ===")
		fmt.Println("1. Створення банку")
		fmt.Println("2. Створення клієнта для роботи з кредитами (Бот)")
		fmt.Println("3. Створення клієнта для роботи з депозитами (Бот)")
		fmt.Println("4. Виведення інформації про клієнта за прізвищем")
		fmt.Println("5. Виведення інформації про клієнта за номером рахунку")
		fmt.Println("6. Завершення")
		fmt.Print("Оберіть пункт: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			var money float64
			fmt.Print("Введіть назву банку: ")
			fmt.Scan(&name)
			fmt.Print("Введіть початковий капітал банку: ")
			fmt.Scan(&money)
			bank = NewBank(name, money)
			fmt.Printf("Банк '%s' успішно створено з капіталом %.2f\n", bank.Name, bank.BankMoney)

		case 2, 3:
			if bank == nil {
				fmt.Println("Помилка! Спочатку потрібно створити банк (Пункт 1).")
				continue
			}
			var name, surname, accNum string
			var amount float64
			fmt.Print("Ім'я клієнта: ")
			fmt.Scan(&name)
			fmt.Print("Прізвище клієнта: ")
			fmt.Scan(&surname)
			fmt.Print("Номер рахунку: ")
			fmt.Scan(&accNum)

			if choice == 2 {
				fmt.Print("Сума кредиту: ")
				fmt.Scan(&amount)
				client := NewClient(name, surname, accNum, 0, amount)
				bank.AddClient(client)
				bank.Credit += amount
				bank.BankMoney -= amount
				bank.StartCreditBot(client)
				fmt.Println("Клієнта-кредитника створено. Бот запущено в горутині.")
			} else {
				fmt.Print("Сума депозиту: ")
				fmt.Scan(&amount)
				client := NewClient(name, surname, accNum, amount, 0)
				bank.AddClient(client)
				bank.Deposit += amount
				bank.BankMoney += amount
				bank.StartDepositBot(client)
				fmt.Println("Клієнта-депозитника створено. Бот запущено в горутині.")
			}

		case 4:
			if bank == nil {
				fmt.Println("Банк не створено.")
				continue
			}
			var surname string
			fmt.Print("Введіть прізвище: ")
			fmt.Scan(&surname)
			if c, found := bank.FindBySurname(surname); found {
				fmt.Printf("\nСтан клієнта %s %s:\nРахунок: %s | Депозит: %.2f | Кредит: %.2f\n",
					c.Name, c.Surname, c.AccountNumber, c.CDeposit, c.CCredit)
			} else {
				fmt.Println("Клієнта не знайдено.")
			}

		case 5:
			if bank == nil {
				fmt.Println("Банк не створено.")
				continue
			}
			var accNum string
			fmt.Print("Введіть номер рахунку: ")
			fmt.Scan(&accNum)
			if c, found := bank.FindByAccount(accNum); found {
				fmt.Printf("\nСтан клієнта %s %s:\nРахунок: %s | Депозит: %.2f | Кредит: %.2f\n",
					c.Name, c.Surname, c.AccountNumber, c.CDeposit, c.CCredit)
			} else {
				fmt.Println("Клієнта не знайдено.")
			}

		case 6:
			fmt.Println("Роботу завершено.")
			os.Exit(0)

		default:
			fmt.Println("Некоректний вибір, спробуйте ще раз.")
		}
	}
}
