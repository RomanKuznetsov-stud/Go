package main

import (
	"sync"
)

type Client struct {
	Name          string
	Surname       string
	AccountNumber string
	CDeposit      float64
	CCredit       float64
}

type Bank struct {
	mu        sync.Mutex
	Name      string
	BankMoney float64
	Deposit   float64
	Credit    float64
	Clients   []*Client
}

func NewClient(name, surname, accNum string, dep, cred float64) *Client {
	return &Client{
		Name:          name,
		Surname:       surname,
		AccountNumber: accNum,
		CDeposit:      dep,
		CCredit:       cred,
	}
}

func NewBank(name string, initialMoney float64) *Bank {
	return &Bank{
		Name:      name,
		BankMoney: initialMoney,
		Clients:   make([]*Client, 0),
	}
}
