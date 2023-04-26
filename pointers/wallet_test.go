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

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("did not get an error")
		}
		if got.Error() != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	// test deposit logic
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assert(t, wallet, Bitcoin(10))
	})

	// test withdraw logic
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "Cannot withdraw, insufficient funds")
		assert(t, wallet, wallet.balance)
	})
}
