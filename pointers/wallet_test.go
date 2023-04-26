package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	// test deposit logic
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	// test withdraw logic
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, wallet.balance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	if wallet.balance != want {
		t.Errorf("got %s want %s", wallet.balance, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("did not get an error")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
