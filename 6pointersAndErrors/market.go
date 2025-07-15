package pointersanderrors

import (
	"errors"
	"fmt"
)

type Wallet struct{
	balance Bitcoin
}

type Bitcoin int

type Stringer interface{
	String() string
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin{
	return w.balance
}
// In Go, errors are values, so we can refactor it out into a variable and have a 
// single source of truth for it.
// The var keyword allows us to define values global to the package.

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")


func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}