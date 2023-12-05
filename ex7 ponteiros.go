package main

import (
	"fmt"
)

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionaSaldo(conta *Conta, valor float64) {
	conta.Saldo += valor
}

func main() {
	conta := Conta{
		Saldo:   100.0,
		Titular: "João",
	}
	fmt.Println("Saldo antes da operação:", conta.Saldo)
	adicionaSaldo(&conta, 50.0)
	fmt.Println("Saldo após adicionar 50.0:", conta.Saldo)
}
