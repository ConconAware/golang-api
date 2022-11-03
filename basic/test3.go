package test

import (
	"fmt"
	"log"
)

type IBankAccount interface {
	GetBalance() (balance int)
	Deposit(amount int)
	Withdraw(amount int) (err error)
}

func (bitcoin BitcoinAccount) GetBalance() int {
	return bitcoin.balance
}

func (bitcoin *BitcoinAccount) Deposit(amount int) {
	fmt.Println("Bitcoin Deposit:\t", amount)
	bitcoin.balance = bitcoin.balance + amount
}

func (bitcoin *BitcoinAccount) Withdraw(amount int) error {
	fmt.Println("Bitcoin Withdraw:\t", amount)
	bitcoin.balance = bitcoin.balance - amount

	return nil
}

type BitcoinAccount struct {
	balance int
}

func NewBitcoinAccount(balance int) *BitcoinAccount {
	fmt.Println("NewBitcoinAccount balance:\t", balance)
	return &BitcoinAccount{
		balance: balance,
	}
}

func (bualuang BualuangAccount) GetBalance() int {
	return bualuang.balance
}

func (bualuang *BualuangAccount) Deposit(amount int) {
	fmt.Println("Bualuang Deposit:\t", amount)
	bualuang.balance = bualuang.balance + amount
}

func (bualuang *BualuangAccount) Withdraw(amount int) error {
	fmt.Println("Bualuang Withdraw:\t", amount)
	bualuang.balance = bualuang.balance - amount

	return nil
}

type BualuangAccount struct {
	balance int
}

func NewBualuangAccount(balance int) *BualuangAccount {
	fmt.Println("NewBualuangAccount balance:\t", balance)
	return &BualuangAccount{
		balance: balance,
	}
}

func Main3() {

	myAcc := []IBankAccount{
		NewBitcoinAccount(5000),
		NewBualuangAccount(1000)}

	for _, acc := range myAcc {
		fmt.Println()
		acc.Deposit(500)
		if err := acc.Withdraw(15); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Bank Account:\t", acc.GetBalance())
	}
}
