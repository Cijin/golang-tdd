package main

import "testing"

const startingBalance = Bitcoin(10)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{startingBalance}
		wallet.Withdraw(5)

		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("withdraw insufficent funds", func(t *testing.T) {
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(15)

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected an error, none returned")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
