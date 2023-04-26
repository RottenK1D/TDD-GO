package pointers

type Depositer interface {
	Deposit(int) int
}
type Balancer interface {
	Balance() int
}

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
