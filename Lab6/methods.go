package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- Геттери та Сеттери для банку (із захистом мутекса) ---
func (b *Bank) AddClient(c *Client) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Clients = append(b.Clients, c)
}

func (b *Bank) FindBySurname(surname string) (*Client, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, c := range b.Clients {
		if c.Surname == surname {
			return c, true
		}
	}
	return nil, false
}

func (b *Bank) FindByAccount(accNum string) (*Client, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, c := range b.Clients {
		if c.AccountNumber == accNum {
			return c, true
		}
	}
	return nil, false
}

func (b *Bank) StartDepositBot(c *Client) {
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			time.Sleep(1 * time.Second)

			b.mu.Lock()

			if c.CDeposit <= 0 {
				fmt.Printf("\n[БОТ ДЕПОЗИТ]: Клієнт %s %s зняв усі кошти. Горутину завершено.\n", c.Name, c.Surname)
				b.mu.Unlock()
				break
			}

			withdraw := float64(r.Intn(100) + 10)
			if withdraw > c.CDeposit {
				withdraw = c.CDeposit
			}

			if b.BankMoney >= withdraw {
				c.CDeposit -= withdraw
				b.Deposit -= withdraw
				b.BankMoney -= withdraw
				fmt.Printf("\n[БОТ ДЕПОЗИТ]: %s %s зняв %.2f грн. Залишок на депозиті: %.2f грн\n", c.Name, c.Surname, withdraw, c.CDeposit)
			}
			b.mu.Unlock()
		}
	}()
}

func (b *Bank) StartCreditBot(c *Client) {
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			time.Sleep(1 * time.Second)

			b.mu.Lock()
			if c.CCredit <= 0 {
				fmt.Printf("\n[БОТ КРЕДИТ]: Клієнт %s %s повністю виплатив кредит. Горутину завершено.\n", c.Name, c.Surname)
				b.mu.Unlock()
				break
			}

			pay := float64(r.Intn(100) + 10)
			if pay > c.CCredit {
				pay = c.CCredit
			}

			c.CCredit -= pay
			b.Credit -= pay
			b.BankMoney += pay
			fmt.Printf("\n[БОТ КРЕДИТ]: %s %s сплатив %.2f грн боргу. Залишок кредиту: %.2f грн\n", c.Name, c.Surname, pay, c.CCredit)
			b.mu.Unlock()
		}
	}()
}
