package pointers

type Bitcoin int

type Depositer interface {
	Deposit(Bitcoin) Bitcoin
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

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
