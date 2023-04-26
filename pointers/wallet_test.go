package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	// test deposit logic
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	// test withdraw logic
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(0)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
