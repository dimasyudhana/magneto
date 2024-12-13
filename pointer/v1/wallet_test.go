package main

import (
	"log"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	want := Bitcoin(10)

	log.Printf("%s", got)
	log.Printf("%s", want)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
