package pointers

import (
	"fmt"
)

type Bitcoin int

type Depositer interface {
	Deposit(Bitcoin) Bitcoin
}

type Withdrawal interface {
	Withdraw(Bitcoin)
}

type Balancer interface {
	Balance() Bitcoin
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
