package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertEqual(t, wallet.Balance(), Bitcoin(10))

		wallet.Deposit(Bitcoin(20))
		assertEqual(t, wallet.Balance(), Bitcoin(30))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		err := wallet.Withdraw(Bitcoin(2))
		assertNoErr(t, err)
		assertEqual(t, wallet.Balance(), Bitcoin(8))

		err = wallet.Withdraw(Bitcoin(4))
		assertNoErr(t, err)
		assertEqual(t, wallet.Balance(), Bitcoin(4))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(1))

		err := wallet.Withdraw(Bitcoin(1))
		assertNoErr(t, err)

		err = wallet.Withdraw(Bitcoin(1))
		assertErr(t, err, "insufficient funds, only 0.000000 BTC left")
		assertEqual(t, wallet.Balance(), Bitcoin(0))
	})
}

func assertEqual[V comparable](t testing.TB, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func assertErr(t testing.TB, err error, msg string) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted err but did not get one")
	}
	if err.Error() != msg {
		t.Errorf("got %v but wanted %v as err message", err.Error(), msg)
	}
}

func assertNoErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got err but did not want one: %v", err)
	}
}
