package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assert := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		if wallet.balance != want {
			t.Errorf("got %s want %s", wallet.balance, want)
		}
	}
	// test deposit logic
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assert(t, wallet, Bitcoin(10))
	})
	// test withdraw logic
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(10))
		assert(t, wallet, Bitcoin(0))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(100))

		assert(t, wallet, Bitcoin(10))

		if err != nil {
			t.Error("no error to be found")
		}
	})
}
