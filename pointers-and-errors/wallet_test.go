package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertEqual(t, Bitcoin(10), wallet.Balance())
		wallet.Deposit(Bitcoin(20))
		assertEqual(t, Bitcoin(30), wallet.Balance())
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.Withdraw(Bitcoin(2))
		assertEqual(t, Bitcoin(8), wallet.Balance())
		wallet.Withdraw(Bitcoin(4))
		assertEqual(t, Bitcoin(4), wallet.Balance())
	})
}

func assertEqual[V comparable](t testing.TB, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
